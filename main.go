package main

import (
	"log"
	"net/http"

	"github.com/Malpizarr/dbproto/pkg/data"
	"github.com/Malpizarr/testwprotodb/handler"
	"github.com/Malpizarr/testwprotodb/repo"
	"github.com/Malpizarr/testwprotodb/service"
)

func main() {
	server := data.NewServer()

	err := server.Initialize()
	if err != nil {
		log.Fatalf("error initializing server: %v", err)
	}

	err = server.CreateDatabase("users")

	db := server.Databases["users"]
	db.CreateTable("users", "Username")
	db.CreateTable("posts", "ID")
	db.CreateTable("friendships", "ID")
	userrepo := repo.NewUserRepo(db)
	userservice := service.NewUserService(userrepo)
	userhandler := handler.NewUserHandler(userservice)

	postrepo := repo.NewPostRepo(db)
	postservice := service.NewPostService(postrepo)
	posthandler := handler.NewPostHandler(postservice)

	friendshiprepo := repo.NewFriendshipRepo(db, userrepo)
	friendshipservice := service.NewFriendshipService(friendshiprepo)
	friendshiphandler := handler.NewFriendshipHandler(friendshipservice)

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

	mux.HandleFunc("POST /friendships", friendshiphandler.Create)
	mux.HandleFunc("GET /friendships/{id}", friendshiphandler.GetFriendship)
	mux.HandleFunc("GET /friendships", friendshiphandler.GetAll)
	mux.HandleFunc("PUT /friendships/accept/{id}", friendshiphandler.AcceptFriendship)
	mux.HandleFunc("PUT /friendships/reject/{id}", friendshiphandler.RejectFriendship)
	mux.HandleFunc("DELETE /friendships/{id}", friendshiphandler.DeleteFriendship)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
