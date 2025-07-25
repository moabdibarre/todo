package routes

import (
	"todolist/handlers"

	"github.com/gorilla/mux"
)


func SetUpRouter(userHandler *handlers.UserHandler, todoHandler *handlers.TodoHandler)*mux.Router {
	r := mux.NewRouter()
	// public routes
	r.HandleFunc("/register", userHandler.RegisterUser)
	r.HandleFunc("/login", userHandler.Login)



	// //sub router for protect routes
	// protected := r.PathPrefix("/").Subrouter()
	// protected.Use(auth.AuthMiddleware)

	// // authenticated routes
	// protected.HandleFunc("/protect", handlers.ProtectedHandler).Methods("GET")
	// protected.HandleFunc("/todos", handlers.CreateTask).Methods("POST")
	// protected.HandleFunc("/todos", handlers.GetTask).Methods("GET")
	// protected.HandleFunc("/todos/{id}", handlers.UpdateTask).Methods("PUT")
	// protected.HandleFunc("/todos/{id}", handlers.DeleteTask).Methods("DELETE")

	return r


}