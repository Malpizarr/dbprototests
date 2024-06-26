package repo

import (
	"fmt"

	"github.com/Malpizarr/dbproto/pkg/data"
	model "github.com/Malpizarr/testwprotodb/data"
)

type UserRepo interface {
	Create(user model.User) error
	GetAll() ([]model.User, error)
	Get(username string) (*model.User, error)
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
		"Username": user.Username,
		"Email":    user.Email,
		"Password": user.Password,
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

		username, ok1 := record["Username"].(string)
		email, ok2 := record["Email"].(string)
		password, ok3 := record["Password"].(string)
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
		"Username": user.Username,
		"Email":    user.Email,
		"Password": user.Password,
	}

	return ur.db.Tables["users"].Update(key, userRecord)
}

func (ur *userRepo) Delete(username string) error {
	return ur.db.Tables["users"].Delete(username)
}

func (ur *userRepo) Get(username string) (*model.User, error) {
	record, err := ur.db.Tables["users"].Select(username)
	if err != nil {
		return nil, err
	}

	usernameValue, ok1 := record["Username"]
	emailValue, ok2 := record["Email"]
	passwordValue, ok3 := record["Password"]

	if !ok1 || !ok2 || !ok3 {
		return &model.User{}, fmt.Errorf("one or more required fields are missing in the record")
	}

	var user model.User

	if usernameValue != nil {
		user.Username = usernameValue.(string)
	}

	if emailValue != nil {
		user.Email = emailValue.(string)
	}

	if passwordValue != nil {
		user.Password = passwordValue.(string)
	}

	return &user, nil
}
