package main

import (
	"fmt"
	"log"
	"net/http"
	"todolist/db"
	"todolist/handlers"
	"todolist/repository"

	"todolist/routes"
	"todolist/services"
)

func main() {

	db.InitDB()
	// initialising repo
	userREpo := &repository.UserRepo{}
	todoRepo := &repository.TodoRepo{}

	//initialise services
	userService := &services.UserService{Repo: userREpo}
	todoService := &services.TodoService{Repo: todoRepo}

	// initalise handlers
	userHandler := &handlers.UserHandler{Service: userService}
	todoHandler := &handlers.TodoHandler{Service: todoService}
	//routes
	r := routes.SetUpRouter(userHandler, todoHandler)
	

	//start server

	fmt.Println("server started on :8080")
	err := http.ListenAndServe(":8080",r)
	if err != nil {
		log.Fatal("failed to start server", err)
	}
}