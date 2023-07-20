package database

import (
	"log"

	"github.com/elliot14A/practice/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init() error {
	log.Println("connecting to Database.....")
	dsn := "host=localhost user=postgres password=postgres dbname=practice port=5432"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}
	log.Println("Database connected successfully.....")
	return nil
}
