package handler

import (
	"encoding/json"
	"net/http"

	model "github.com/Malpizarr/testwprotodb/data"
	"github.com/Malpizarr/testwprotodb/service"
)

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return &userHandler{service: service}
}

func (uh *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "error decoding user", http.StatusBadRequest)
		return
	}
	err = uh.service.CreateUser(user)
	if err != nil {
		http.Error(w, "error creating user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "error encoding user", http.StatusInternalServerError)
		return
	}
}

func (uh *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "error decoding user", http.StatusBadRequest)
		return
	}
	err = uh.service.UpdateUser(user)
	if err != nil {
		http.Error(w, "error updating user", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "error encoding user", http.StatusInternalServerError)
		return
	}
}

func (uh *userHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	user, err := uh.service.GetUser(username)
	if err != nil {
		http.Error(w, "error getting user", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "error encoding user", http.StatusInternalServerError)
		return
	}
}

func (uh *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uh.service.GetUsers()
	if err != nil {
		http.Error(w, "error getting users", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "error encoding users", http.StatusInternalServerError)
		return
	}
}

func (uh *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	err := uh.service.DeleteUser(username)
	if err != nil {
		http.Error(w, "error deleting user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
