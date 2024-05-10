package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	model "github.com/Malpizarr/testwprotodb/data"
	"github.com/Malpizarr/testwprotodb/service"
)

type FriendshipHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetFriendship(w http.ResponseWriter, r *http.Request)
	AcceptFriendship(w http.ResponseWriter, r *http.Request)
	RejectFriendship(w http.ResponseWriter, r *http.Request)
	DeleteFriendship(w http.ResponseWriter, r *http.Request)
}

type friendshipHandler struct {
	service service.FriendshipService
}

func NewFriendshipHandler(service service.FriendshipService) FriendshipHandler {
	return &friendshipHandler{service: service}
}

func (fh *friendshipHandler) Create(w http.ResponseWriter, r *http.Request) {
	var friendship model.Friendship
	err := json.NewDecoder(r.Body).Decode(&friendship)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = fh.service.Create(friendship)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (fh *friendshipHandler) GetFriendship(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "error: id parameter is required", http.StatusBadRequest)
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("error: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	friendship, err := fh.service.GetFriendship(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(friendship)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (fh *friendshipHandler) AcceptFriendship(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "error: id parameter is required", http.StatusBadRequest)
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = fh.service.AcceptFriendship(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (fh *friendshipHandler) RejectFriendship(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "error: id parameter is required", http.StatusBadRequest)
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = fh.service.RejectFriendship(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (fh *friendshipHandler) DeleteFriendship(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "error: id parameter is required", http.StatusBadRequest)
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = fh.service.DeleteFriendship(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
