package models

import "time"

type ClassroomSubject struct {
	ID          uint `gorm:"primaryKey"`
	ClassroomID uint
	SubjectID   uint
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Classroom Classroom `gorm:"foreignKey:ClassroomID"`
	Subject   Subject   `gorm:"foreignKey:SubjectID"`
}
