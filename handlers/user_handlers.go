package handlers

import (
	"encoding/json"
	"net/http"
	
	"todolist/models"
	"todolist/services"
	

	
)

type UserHandler struct {
	Service *services.UserService
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	// collect the user details as a request body
	var signUp *models.User
	err := json.NewDecoder(r.Body).Decode(&signUp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// call service layer

	err = h.Service.RegisterUser(signUp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// send response back to client
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(signUp)
}

func (h *UserHandler)  Login(w http.ResponseWriter, r *http.Request) {
	var login models.User
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//call service layer
	token, err :=h.Service.Login(&login)
	if err !=nil{http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}



	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)

}