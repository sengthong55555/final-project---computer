package repositories

import (
	"go_starter/models"
	"gorm.io/gorm"
)

type ClassroomSubjectRepository interface {
	GetClassroomSubjectRepository() ([]models.ClassroomSubject, error)
	DeleteClassroomSubjectRepository(request models.ClassroomSubject) error
	UpdateClassroomSubjectRepository(request models.ClassroomSubject) error
	InsertClassroomSubjectRepository(request models.ClassroomSubject) error
}

type classroomSubjectRepository struct {
	db *gorm.DB
}

//func (c *classroomSubjectRepository) GetClassroomSubjectRepository() ([]models.ClassroomSubject, error) {
//	var classroomSubjects []models.ClassroomSubject
//
//	// Inner join related Classroom and Subject data
//	err := c.db.
//		Table("classroom_subjects").
//		Select("classroom_subjects.id, classroom_subjects.classroom_id, classroom_subjects.subject_id, classrooms.id as classroom_id, classrooms.major, classrooms.class_year, classrooms.class_name, subjects.id as subject_id, subjects.subject_code, subjects.subject_name").
//		Joins("INNER JOIN classrooms ON classroom_subjects.classroom_id = classrooms.id").
//		Joins("INNER JOIN subjects ON classroom_subjects.subject_id = subjects.id").
//		Find(&classroomSubjects).Error
//	if err != nil {
//		return nil, err
//	}
//
//	return classroomSubjects, nil
//}

func (c *classroomSubjectRepository) GetClassroomSubjectRepository() ([]models.ClassroomSubject, error) {
	var classroomSubjects []models.ClassroomSubject

	// Preload related Classroom and Subject data
	err := c.db.Preload("Classroom").Preload("Subject").Find(&classroomSubjects).Error
	if err != nil {
		return nil, err
	}

	return classroomSubjects, nil
}

func (c *classroomSubjectRepository) DeleteClassroomSubjectRepository(request models.ClassroomSubject) error {
	//TODO implement me
	panic("implement me")
}

func (c *classroomSubjectRepository) UpdateClassroomSubjectRepository(request models.ClassroomSubject) error {
	//TODO implement me
	panic("implement me")
}

func (c *classroomSubjectRepository) InsertClassroomSubjectRepository(request models.ClassroomSubject) error {
	if err := c.db.Create(&request).Error; err != nil {
		return err
	}
	return nil
}

func NewClassroomSubjectRepository(
	db *gorm.DB,
) ClassroomSubjectRepository {
	return &classroomSubjectRepository{
		db: db,
	}
}
