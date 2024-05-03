package repo

import (
	"github.com/Malpizarr/dbproto/pkg/data"
	model "github.com/Malpizarr/testwprotodb/data"
)

type UserRepo interface {
	Create(user model.User) error
	GetAll() ([]model.User, error)
	Update(user model.User) error
	Delete(username string) error
}

type userRepo struct {
	db *data.Database
}

func NewUserRepo(db *data.Database) UserRepo {
	return &userRepo{db: db}
}

func (ur *userRepo) Create(user model.User) error {
	userRecord := data.Record{
		"username": user.Username,
		"email":    user.Email,
		"password": user.Password,
	}

	return ur.db.Tables["users"].Insert(userRecord)
}

func (ur *userRepo) GetAll() ([]model.User, error) {
	records, err := ur.db.Tables["users"].SelectAll()
	if err != nil {
		return nil, err
	}

	users := make([]model.User, 0)
	for _, record := range records {
		if record == nil {
			continue
		}

		username, ok1 := record.Fields["username"]
		email, ok2 := record.Fields["email"]
		password, ok3 := record.Fields["password"]
		if !ok1 || !ok2 || !ok3 {
			continue
		}

		user := model.User{
			Username: username,
			Email:    email,
			Password: password,
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *userRepo) Update(user model.User) error {
	key := user.Username

	userRecord := data.Record{
		"username": user.Username,
		"email":    user.Email,
		"password": user.Password,
	}

	return ur.db.Tables["users"].Update(key, userRecord)
}

func (ur *userRepo) Delete(username string) error {
	return ur.db.Tables["users"].Delete(username)
}
