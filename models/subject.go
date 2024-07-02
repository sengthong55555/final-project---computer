package models

import "time"

type Subject struct {
	ID          uint   `gorm:"primaryKey"`
	SubjectCode string `gorm:"unique;size:255"`
	SubjectName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
