package service

import (
	model "github.com/Malpizarr/testwprotodb/data"
	"github.com/Malpizarr/testwprotodb/repo"
)

type UserService interface {
	CreateUser(user model.User) error
	UpdateUser(user model.User) error
	GetUser(username string) (model.User, error)
	GetUsers() ([]model.User, error)
	DeleteUser(username string) error
}

type userService struct {
	repo repo.UserRepo
}

func NewUserService(repo repo.UserRepo) UserService {
	return &userService{repo: repo}
}

func (us *userService) CreateUser(user model.User) error {
	return us.repo.Create(user)
}

func (us *userService) UpdateUser(user model.User) error {
	return us.repo.Update(user)
}

func (us *userService) GetUser(username string) (model.User, error) {
	return us.repo.Get(username)
}

func (us *userService) GetUsers() ([]model.User, error) {
	return us.repo.GetAll()
}

func (us *userService) DeleteUser(username string) error {
	return us.repo.Delete(username)
}
