package db

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Init is init
func Init() *gorm.DB {

	var DatabaseURL string = os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(DatabaseURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

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

// GetUser adds user
func (h Handler) GetUser(tgUserID string) {
	id, _ := strconv.Atoi(tgUserID)

	var user User
	if result := h.DB.First(&user, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	return
}
