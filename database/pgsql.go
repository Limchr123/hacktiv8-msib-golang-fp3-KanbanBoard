package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"kanban_board/entity"
	"log"
	"os"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	if env := godotenv.Load(); env != nil {
		log.Panic("Error occurred while trying to read environment file", env)
	}

	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		dbname   = os.Getenv("DB_NAME")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Error occurred while trying to connect to database:", err)
	}

	if err := db.AutoMigrate(&entity.User{}, &entity.Category{}, &entity.Task{}); err != nil {
		log.Panic("Error occurred while trying to perform database migrations:", err)
	}

}

func GetDataBaseInstance() *gorm.DB {
	return db
}
