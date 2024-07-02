package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	CodeID    string `gorm:"unique;size:255"`
	Firstname string
	Lastname  string
	Password  string
	Phone     string `gorm:"unique;size:255"`
	Gender    string
	Image     string
	Degree    string
	Skill     string
	UserType  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
