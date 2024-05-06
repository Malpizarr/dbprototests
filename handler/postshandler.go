package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	model "github.com/Malpizarr/testwprotodb/data"
	"github.com/Malpizarr/testwprotodb/service"
)

type PostHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByUsername(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type postHandler struct {
	service service.PostService
}

func NewPostHandler(service service.PostService) PostHandler {
	return &postHandler{service: service}
}

func (ph *postHandler) Create(w http.ResponseWriter, r *http.Request) {
	var post model.Post

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	post.ID = rand.Intn(1000)

	if err := ph.service.Create(post); err != nil {
		http.Error(w, "Failed to create post: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (ph *postHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	posts, err := ph.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(posts)
}

func (ph *postHandler) GetByUsername(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	posts, err := ph.service.GetByUsername(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ph *postHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	post, err := ph.service.GetByID(ids)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(post)
}

func (ph *postHandler) Update(w http.ResponseWriter, r *http.Request) {
	var post model.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = ph.service.Update(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}

func (ph *postHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = ph.service.Delete(ids)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusNoContent)
}
