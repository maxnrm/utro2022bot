package db

import (
	"gorm.io/gorm"
)

// User is telegram user
type User struct {
	ID            int64  `json:"id" gorm:"primaryKey"`
	Group         string `json:"group"`
	JoinedChannel bool   `json:"joinedChannel"`
}

// Timetable for storing timetables
type Timetable struct {
	gorm.Model
	TimetableString string `json:"timetableString"`
}
