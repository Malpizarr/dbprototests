package repo

import (
	"github.com/Malpizarr/dbproto/pkg/data"
	model "github.com/Malpizarr/testwprotodb/data"
)

type UserRepo interface {
	Create(user model.User) error
	GetAll() ([]model.User, error)
}

type userRepo struct {
	db *data.Database
}

func NewUserRepo(db *data.Database) UserRepo {
	return &userRepo{db: db}
}

func (ur *userRepo) Create(user model.User) error {
	// Convertir model.User a data.Record
	userRecord := data.Record{
		"username": user.Username, // Asumiendo que ID es int
		"email":    user.Email,    // Asumiendo que Name es string
		"password": user.Password, // Asumiendo que Email es string
	}

	// Insertar el record en la tabla de usuarios
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
