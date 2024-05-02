package main

import (
	"log"

	"github.com/Malpizarr/dbproto/pkg/data"
	model "github.com/Malpizarr/testwprotodb/data"
	"github.com/Malpizarr/testwprotodb/repo"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	server := data.NewServer()

	err := server.Initialize()
	if err != nil {
		log.Fatalf("error initializing server: %v", err)
	}

	server.CreateDatabase("users")
	db := server.Databases["users"]
	db.CreateTable("users", "username")

	repo := repo.NewUserRepo(db)
	err = repo.Create(model.User{
		Username: "malpizarr",
		Email:    "mau",
		Password: "123",
	})
	if err != nil {
		log.Fatalf("error creating user: %v", err)
	}

	users, err := repo.GetAll()
	if err != nil {
		log.Fatalf("error getting users: %v", err)
	}
	for _, user := range users {
		log.Println(user)
	}
}
