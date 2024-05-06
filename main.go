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
	db.CreateTable("posts", "ID")
	userrepo := repo.NewUserRepo(db)
	userservice := service.NewUserService(userrepo)
	userhandler := handler.NewUserHandler(userservice)

	postrepo := repo.NewPostRepo(db)
	postservice := service.NewPostService(postrepo)
	posthandler := handler.NewPostHandler(postservice)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", userhandler.CreateUser)
	mux.HandleFunc("PUT /users", userhandler.UpdateUser)
	mux.HandleFunc("GET /users", userhandler.GetUsers)
	mux.HandleFunc("GET /users/{username}", userhandler.GetUser)
	mux.HandleFunc("DELETE /users/{username}", userhandler.DeleteUser)

	mux.HandleFunc("POST /posts", posthandler.Create)
	mux.HandleFunc("GET /posts", posthandler.GetAll)
	mux.HandleFunc("GET /posts/{username}", posthandler.GetByUsername)

	mux.HandleFunc("PUT /posts", posthandler.Update)
	mux.HandleFunc("DELETE /posts/{id}", posthandler.Delete)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
