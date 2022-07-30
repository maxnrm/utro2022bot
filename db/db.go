package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// DBHandler singleton instance of db
var DBHandler Handler = New()

// Init is init
func Init() *gorm.DB {

	var DatabaseURL string = os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(DatabaseURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Timetable{})
	db.AutoMigrate(&User{})

	return db
}

// Handler is handler
type Handler struct {
	DB *gorm.DB
}

// New is new
func New() Handler {
	db := Init()
	return Handler{db}
}

// AddUser adds user. 'columns' regulates what columns will be updated
func (h Handler) AddUser(user *User, columns []string) {

	if result := h.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns(columns),
	}).Create(&user); result.Error != nil {
		fmt.Println(result.Error)
	}

	return
}

// AddTimetable saves timetable in case it needs to be loaded on server startup
func (h Handler) AddTimetable(ttString string) {

	var timetable Timetable = Timetable{TimetableString: ttString}

	if result := h.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"timetable_string"}),
	}).Create(&timetable); result.Error != nil {
		fmt.Println(result.Error)
	}

	return
}

// GetTimetable gets latest timetable
func (h Handler) GetTimetable() (string, error) {

	var timetable Timetable
	result := h.DB.First(&timetable)

	return timetable.TimetableString, result.Error
}

// GetUser gets user by id
func (h Handler) GetUser(tgUserID int64) User {

	var user User
	if result := h.DB.First(&user, tgUserID); result.Error != nil {
		fmt.Printf("User with id %v not found\n", tgUserID)
	}

	return user
}
