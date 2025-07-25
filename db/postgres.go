package db

import (
	"fmt"
	"log"
	"os"
	"todolist/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB

func InitDB() {
	// create connection str
	// connStr := "user=postgres password=password dbname=tasks port=5432 sslmode=disable"
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// open db connection
	var err error
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//migrate the schema
	err = DB.AutoMigrate(&models.User{}, &models.ToDoList{})
	if err != nil {
		log.Fatal("failed to migrate the schema", err)
	}

	fmt.Println("connected to db successfully!")
}