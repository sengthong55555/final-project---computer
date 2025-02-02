package repositories

import (
	"fmt"
	"go_starter/models"
	"gorm.io/gorm"
)

type SubjectRepository interface {
	FilterSubjectBySubjectCodeRepository(subjectCode string) (*models.Subject, error)
	CreateSubjectRepository(model models.Subject) error
	GetSubjectRepository() ([]models.Subject, error)
	//not yet
	DeleteSubjectRepository(model models.Subject) error
	UpdateSubjectRepository(model models.Subject) error
}

type subjectRepository struct {
	db *gorm.DB
}

func (s *subjectRepository) FilterSubjectBySubjectCodeRepository(subjectCode string) (*models.Subject, error) {
	var model models.Subject
	if err := s.db.Where("subject_code=?", subjectCode).First(&model).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (s *subjectRepository) DeleteSubjectRepository(model models.Subject) error {
	//TODO implement me
	panic("implement me")
}

func (s *subjectRepository) UpdateSubjectRepository(model models.Subject) error {
	//TODO implement me
	panic("implement me")
}

func (s *subjectRepository) checkSubjectRepository(subjectCode string) (bool, error) {
	var count int64
	err := s.db.Model(&models.Subject{}).Where("subject_code = ?", subjectCode).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (s *subjectRepository) CreateSubjectRepository(model models.Subject) error {
	checkSubjectCode, err := s.checkSubjectRepository(model.SubjectCode)
	if err != nil {
		return err
	}

	if checkSubjectCode == true {
		return fmt.Errorf("subject code already had been created")
	}

	if err = s.db.Create(&model).Error; err != nil {
		return err
	}
	return nil
}

func (s *subjectRepository) GetSubjectRepository() ([]models.Subject, error) {
	var model []models.Subject
	query := s.db.Find(&model).Error
	if query != nil {
		return nil, query
	}
	return model, nil
}

func NewSubjectRepository(
	db *gorm.DB,
) SubjectRepository {
	return &subjectRepository{
		db: db,
	}
}
