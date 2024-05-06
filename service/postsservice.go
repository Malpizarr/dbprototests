package service

import (
	model "github.com/Malpizarr/testwprotodb/data"
	"github.com/Malpizarr/testwprotodb/repo"
)

type PostService interface {
	Create(post model.Post) error
	GetAll() ([]model.Post, error)
	GetByUsername(username string) ([]model.Post, error)
	GetByID(id int) (model.Post, error)
	Update(post model.Post) error
	Delete(id int) error
}

type postService struct {
	repo repo.PostRepo
}

func NewPostService(repo repo.PostRepo) PostService {
	return &postService{repo: repo}
}

func (ps *postService) Create(post model.Post) error {
	return ps.repo.Create(post)
}

func (ps *postService) GetAll() ([]model.Post, error) {
	return ps.repo.GetAll()
}

func (ps *postService) GetByUsername(username string) ([]model.Post, error) {
	return ps.repo.GetByUsername(username)
}

func (ps *postService) GetByID(id int) (model.Post, error) {
	return ps.repo.GetByID(id)
}

func (ps *postService) Update(post model.Post) error {
	return ps.repo.Update(post)
}

func (ps *postService) Delete(id int) error {
	return ps.repo.Delete(id)
}
