package handlers

import (
	"encoding/json"
	"net/http"
	"todolist/models"
	"todolist/services"
)

type TodoHandler struct {

	Service *services.TodoService
}

func (h *TodoHandler)CreateToDo(w http.ResponseWriter, r *http.Request) {
	//get todo from request body
	var todo models.ToDoList
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//call service layer
	h.Service.CreateToDo(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	
	}

	//response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}