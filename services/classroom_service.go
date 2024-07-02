package services

import (
	"go_starter/models"
	"go_starter/repositories"
	"go_starter/requests"
	"go_starter/responses"
	"strconv"

	"github.com/pkg/errors"
)

type ClassRoomService interface {
	GetAllClassRoomServices() ([]responses.ClassroomResponse, error)
	GetClassRoomByIdService(id uint) (*responses.ClassroomResponse, error)
	CreateClassRoomService(request requests.CreateClassRoomRequest) (*responses.MessageClassRoomResponse, error)

	//UpdateClassRoomService(request requests.ClassRoomRequest) (*responses.MessageClassRoomResponse, error)
	//DeleteClassRoomService(request requests.ClassRoomCodeRequest) (*responses.MessageClassRoomResponse, error)
}

type classroomService struct {
	repositoryClassRoom repositories.ClassRoomRepository
}

func (c *classroomService) GetAllClassRoomServices() ([]responses.ClassroomResponse, error) {

	data, err := c.repositoryClassRoom.GetAllClassRoomRepositories()
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.New("data slice is nil")
	}
	var response []responses.ClassroomResponse

	for _, request := range data {
		classYearString := strconv.Itoa(request.ClassYear)
		classNameString := strconv.Itoa(request.ClassName)
		model := responses.ClassroomResponse{
			ID:        request.ID,
			ClassName: classYearString + request.Major + classNameString,
		}
		response = append(response, model)
	}
	return response, err
}
func (c *classroomService) GetClassRoomByIdService(id uint) (*responses.ClassroomResponse, error) {
	data, err := c.repositoryClassRoom.GetClassRoomByIdRepository(uint(id))
	if err != nil {
		return nil, err
	}
	classYearString := strconv.Itoa(data.ClassYear)
	classNameString := strconv.Itoa(data.ClassName)
	model := &responses.ClassroomResponse{
		ID:        data.ID,
		ClassName: classYearString + data.Major + classNameString,
	}
	return model, nil
}

func (c *classroomService) CreateClassRoomService(request requests.CreateClassRoomRequest) (*responses.MessageClassRoomResponse, error) {

	model := models.Classroom{
		Major:     request.Major,
		ClassYear: request.ClassYear,
		ClassName: request.ClassName,
	}
	if err := c.repositoryClassRoom.CreateClassRoomRepository(&model); err != nil {
		return nil, err
	}

	// If successful, return a success message response
	response := &responses.MessageClassRoomResponse{Message: "success"}
	return response, nil

}

//func (c *classroomService) DeleteClassRoomService(request requests.ClassRoomCodeRequest) (*responses.MessageClassRoomResponse, error) {
//	if request.Id == 0 {
//		return nil, errors.New("Class Year cant be empty")
//	}
//	if err := c.repositoryClassRoom.DeleteClassRoomRepository(request.Id); err != nil {
//		return nil, err
//	}
//	response := &responses.MessageClassRoomResponse{Message: "success"}
//	return response, nil
//}

//func (c *classroomService) UpdateClassRoomService(request requests.ClassRoomRequest) (*responses.MessageClassRoomResponse, error) {
//
//	model := models.Classroom{
//		ClassName: request.ClassName,
//		Major:     request.Major,
//		ClassYear: request.ClassYear,
//		//SubjectName:   request.SubjectName,
//		//ClassRoomCode: request.ClassRoomCode,
//	}
//	if err := c.repositoryClassRoom.UpdateClassRoomRepository(&model); err != nil {
//		return nil, err
//	}
//	response := &responses.MessageClassRoomResponse{Message: "success"}
//	return response, nil
//}

func NewRoomServices(repositoryClassRoom repositories.ClassRoomRepository) ClassRoomService {
	return &classroomService{
		repositoryClassRoom: repositoryClassRoom,
	}
}
