package service

import (
	model "github.com/Malpizarr/testwprotodb/data"
	"github.com/Malpizarr/testwprotodb/repo"
)

type FriendshipService interface {
	Create(friendship model.Friendship) error
	GetFriendship(id int) (*model.Friendship, error)
	AcceptFriendship(id int) error
	RejectFriendship(id int) error
	DeleteFriendship(id int) error
	GetAll() ([]model.Friendship, error)
}

type friendshipService struct {
	repo repo.FriendshipRepo
}

func NewFriendshipService(repo repo.FriendshipRepo) FriendshipService {
	return &friendshipService{repo: repo}
}

func (fs *friendshipService) Create(friendship model.Friendship) error {
	return fs.repo.Create(friendship)
}

func (fs *friendshipService) GetFriendship(id int) (*model.Friendship, error) {
	return fs.repo.GetFriendship(id)
}

func (fs *friendshipService) AcceptFriendship(id int) error {
	return fs.repo.AcceptFriendship(id)
}

func (fs *friendshipService) RejectFriendship(id int) error {
	return fs.repo.RejectFriendship(id)
}

func (fs *friendshipService) DeleteFriendship(id int) error {
	return fs.repo.DeleteFriendship(id)
}

func (fs *friendshipService) GetAll() ([]model.Friendship, error) {
	return fs.repo.GetAll()
}
