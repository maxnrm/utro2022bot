package db

import (
	"github.com/maxnrm/utro2022bot/timetable"
	"gorm.io/gorm"
)

// User is telegram user
type User struct {
	ID            int64  `json:"id" gorm:"primaryKey"`
	Group         string `json:"group"`
	JoinedChannel bool   `json:"joinedChannel"`
}

// TimetableWrapper for storing timetables
type TimetableWrapper struct {
	gorm.Model
	timetable.Wrapper
}
