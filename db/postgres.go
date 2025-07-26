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

func Initdb() {
	// fmt.Println("\n--- DEBUG: Values from os.Getenv() ---")
	// fmt.Printf("DB_HOST: %q\n", os.Getenv("DB_HOST"))
	// fmt.Printf("DB_PORT: %q\n", os.Getenv("DB_PORT"))
	// fmt.Printf("DB_USER: %q\n", os.Getenv("DB_USER"))
	// fmt.Printf("DB_PASSWORD: %q\n", os.Getenv("DB_PASSWORD"))
	// fmt.Printf("DB_NAME: %q\n", os.Getenv("DB_NAME"))
	// fmt.Printf("APP_PORT: %q\n", os.Getenv("APP_PORT"))
	// fmt.Println("------------------------------------")
	// create connection str
	// connStr := "user=postgres password=password dbname=tasks port=5432 sslmode=disable"
	connStr := fmt.Sprintf(
		"host=%s, user=%s, password+%s, dbname=%s port=%s, sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	fmt.Println(os.Getenv("DB_PORT"))
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
