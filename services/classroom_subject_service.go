package services

import (
	"go_starter/models"
	"go_starter/repositories"
	"go_starter/requests"
	"go_starter/responses"
	"strconv"
)

type ClassroomSubjectService interface {
	GetClassroomSubjectService() ([]responses.ClassroomSubjectResponse, error)
	CreateClassroomSubjectService(request requests.CreateClassroomSubjectRequest) (*responses.CreateClassRoomSubjectResponse, error)
}

type classroomSubjectService struct {
	repositoryClassroomSubject repositories.ClassroomSubjectRepository
}

func (c *classroomSubjectService) GetClassroomSubjectService() ([]responses.ClassroomSubjectResponse, error) {
	classroomSubjectData, err := c.repositoryClassroomSubject.GetClassroomSubjectRepository()
	if err != nil {
		return nil, err
	}
	var response []responses.ClassroomSubjectResponse
	for _, data := range classroomSubjectData {
		classYearString := strconv.Itoa(data.Classroom.ClassYear)
		classNameString := strconv.Itoa(data.Classroom.ClassName)
		className := classYearString + data.Classroom.Major + classNameString
		classroomSubjectResponse := responses.ClassroomSubjectResponse{
			ID:          data.ID,
			ClassName:   className,
			SubjectCode: data.Subject.SubjectCode,
			SubjectName: data.Subject.SubjectName,
		}
		response = append(response, classroomSubjectResponse)
	}
	return response, nil
}

func (c *classroomSubjectService) CreateClassroomSubjectService(request requests.CreateClassroomSubjectRequest) (*responses.CreateClassRoomSubjectResponse, error) {
	classroomSubject := models.ClassroomSubject{
		ClassroomID: request.ClassroomID,
		SubjectID:   request.SubjectID,
	}
	err := c.repositoryClassroomSubject.InsertClassroomSubjectRepository(classroomSubject)
	if err != nil {
		return &responses.CreateClassRoomSubjectResponse{
			Message: "Failed to create classroom subject",
			Status:  false,
		}, err
	}

	// If successful, return a success message response
	response := &responses.CreateClassRoomSubjectResponse{
		Message: "Classroom subject created successfully",
		Status:  true,
	}
	return response, nil
}

func NewClassroomServices(
	repositoryClassroomSubject repositories.ClassroomSubjectRepository,
) ClassroomSubjectService {
	return &classroomSubjectService{
		repositoryClassroomSubject: repositoryClassroomSubject,
	}
}
