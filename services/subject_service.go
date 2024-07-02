package services

import (
	"go_starter/models"
	"go_starter/repositories"
	"go_starter/requests"
	"go_starter/responses"
)

type SubjectService interface {
	CreateSubjectService(request requests.InsertSubjectRequest) (responses.SubjectMessageResponse, error)
	FilterSubjectBySubjectCodeService(request requests.SubjectRequest) (*responses.SubjectResponse, error)
	GetSubjectService() ([]responses.SubjectResponse, error)
}

type subjectService struct {
	repositorySubject repositories.SubjectRepository
}

func (s *subjectService) CreateSubjectService(request requests.InsertSubjectRequest) (responses.SubjectMessageResponse, error) {
	subjectModel := models.Subject{
		SubjectCode: request.SubjectCode,
		SubjectName: request.SubjectName,
	}
	err := s.repositorySubject.CreateSubjectRepository(subjectModel)
	if err != nil {
		return responses.SubjectMessageResponse{
			Message: "failed to create subject ",
			Status:  false,
		}, err
	}
	return responses.SubjectMessageResponse{
		Message: "success",
		Status:  true,
	}, nil
}

func (s *subjectService) FilterSubjectBySubjectCodeService(request requests.SubjectRequest) (*responses.SubjectResponse, error) {
	getSubjectData, err := s.repositorySubject.FilterSubjectBySubjectCodeRepository(request.SubjectCode)
	if err != nil {
		return nil, err
	}
	response := responses.SubjectResponse{
		ID:          getSubjectData.ID,
		SubjectCode: getSubjectData.SubjectCode,
		SubjectName: getSubjectData.SubjectName,
	}
	return &response, err
}

func (s *subjectService) GetSubjectService() ([]responses.SubjectResponse, error) {
	getSubjectData, err := s.repositorySubject.GetSubjectRepository()
	if err != nil {
		return nil, err
	}
	if getSubjectData == nil {
		return []responses.SubjectResponse{}, nil
	}
	var response []responses.SubjectResponse
	for _, data := range getSubjectData {
		response = append(response, responses.SubjectResponse{
			ID:          data.ID,
			SubjectCode: data.SubjectCode,
			SubjectName: data.SubjectName,
		})
	}
	return response, err
}

func NewSubjectService(
	repositorySubject repositories.SubjectRepository,
) SubjectService {
	return &subjectService{
		repositorySubject: repositorySubject,
	}
}
