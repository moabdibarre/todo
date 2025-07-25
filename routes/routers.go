package routes

import (
	"todolist/handlers"
	"todolist/middleware"

	"github.com/gorilla/mux"
)


func SetUpRouter(userHandler *handlers.UserHandler, todoHandler *handlers.TodoHandler)*mux.Router {
	r := mux.NewRouter()
	// public routes
	r.HandleFunc("/register", userHandler.RegisterUser)
	r.HandleFunc("/login", userHandler.Login)



	// //sub router for protect routes
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	// // authenticated routes
	// protected.HandleFunc("/protect", handlers.ProtectedHandler).Methods("GET")
	protected.HandleFunc("/todos", todoHandler.CreateToDo).Methods("POST")
	protected.HandleFunc("/todos", todoHandler.GetToDo).Methods("GET")
	// protected.HandleFunc("/todos/{id}", handlers.UpdateTask).Methods("PUT")
	// protected.HandleFunc("/todos/{id}", handlers.DeleteTask).Methods("DELETE")

	return r


}