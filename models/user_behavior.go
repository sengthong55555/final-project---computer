package models

import "time"

type UserBehavior struct {
	ID               uint `gorm:"primaryKey"`
	UserID           uint
	ClassroomID      uint
	StudentCheck     bool
	StudentAbsent    bool
	StudentVacation  bool
	StudentBreakRule bool
	CountCheck       int
	CountAbsent      int
	CountVacation    int
	CountBreakRule   int
	CreatedAt        time.Time
	UpdatedAt        time.Time

	User      User      `gorm:"foreignKey:UserID"`
	Classroom Classroom `gorm:"foreignKey:ClassroomID"`
}
