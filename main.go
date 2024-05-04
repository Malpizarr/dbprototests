package main

import (
	"log"
	"net/http"

	"github.com/Malpizarr/dbproto/pkg/data"
	"github.com/Malpizarr/testwprotodb/handler"
	"github.com/Malpizarr/testwprotodb/repo"
	"github.com/Malpizarr/testwprotodb/service"
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
	db.CreateTable("users", "Username")
	repo := repo.NewUserRepo(db)
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", handler.CreateUser)
	mux.HandleFunc("PUT /users", handler.UpdateUser)
	mux.HandleFunc("GET /users", handler.GetUsers)
	mux.HandleFunc("GET /users/{username}", handler.GetUser)
	mux.HandleFunc("DELETE /users/{username}", handler.DeleteUser)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
