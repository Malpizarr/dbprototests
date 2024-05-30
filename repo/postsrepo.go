package repo

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Malpizarr/dbproto/pkg/data"
	model "github.com/Malpizarr/testwprotodb/data"
)

type PostRepo interface {
	Create(post model.Post) error
	GetAll() ([]model.Post, error)
	GetByUsername(username string) ([]model.Post, error)
	GetByID(id int) (model.Post, error)
	Update(post model.Post) error
	Delete(id int) error
}

type postRepo struct {
	db       *data.Database
	userRepo UserRepo
}

func NewPostRepo(db *data.Database, ur UserRepo) PostRepo {
	return &postRepo{db: db, userRepo: ur}
}

func (pr *postRepo) Create(post model.Post) error {
	postRecord := data.Record{
		"ID":       post.ID,
		"Username": post.Username,
		"Title":    post.Title,
		"Content":  post.Content,
	}
	userExists, err := pr.userRepo.Get(post.Username)
	if userExists == nil {
		return fmt.Errorf("error: user does not exist")

	}
	err = pr.db.Tables["posts"].Insert(postRecord)
	if err != nil {
		return err
	}
	return nil
}

func (pr *postRepo) GetAll() ([]model.Post, error) {
	records, err := pr.db.Tables["posts"].SelectAll()
	if err != nil {
		return nil, err
	}
	var posts []model.Post
	for _, record := range records {
		if record == nil {
			continue
		}
		idValue, ok1 := record["ID"]
		usernameValue, ok2 := record["Username"]
		titleValue, ok3 := record["Title"]
		contentValue, ok4 := record["Content"]
		if !ok1 || !ok2 || !ok3 || !ok4 {
			continue
		}

		var id int64
		if idValue != nil {
			id = idValue.(int64)

		}

		username := ""
		if usernameValue != nil {
			username = usernameValue.(string)
		}

		title := ""
		if titleValue != nil {
			title = titleValue.(string)
		}

		content := ""
		if contentValue != nil {
			content = contentValue.(string)
		}

		post := model.Post{
			ID:       id,
			Username: username,
			Title:    title,
			Content:  content,
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (pr *postRepo) GetByUsername(username string) ([]model.Post, error) {
	usersTable := pr.db.Tables["users"]
	postsTable := pr.db.Tables["posts"]

	joinedRecords, err := data.JoinTables(postsTable, usersTable, "Username", "Username", data.InnerJoin)
	if err != nil {
		return nil, fmt.Errorf("failed to join tables: %v", err)
	}

	var posts []model.Post
	for _, record := range joinedRecords {
		var id int64
		if idVal, ok := record["t1.ID"].(int64); ok {
			id = idVal
		} else {
			log.Printf("ID value is missing or not a float64 for username: %s", username)
			continue
		}
		usernameValue := record["t1.Username"].(string)
		titleValue := record["t1.Title"].(string)
		contentValue := record["t1.Content"].(string)

		post := model.Post{
			ID:       id,
			Username: usernameValue,
			Title:    titleValue,
			Content:  contentValue,
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (pr *postRepo) GetByID(id int) (model.Post, error) {
	ids := strconv.Itoa(id)
	record, err := pr.db.Tables["posts"].Select(ids)
	if err != nil {
		return model.Post{}, err
	}
	if record == nil {
		return model.Post{}, nil
	}
	idStr, ok1 := record["ID"]
	username, ok2 := record["Username"]
	title, ok3 := record["Title"]
	content, ok4 := record["Content"]
	if !ok1 || !ok2 || !ok3 || !ok4 {
		return model.Post{}, fmt.Errorf("error getting post by id")
	}

	post := model.Post{
		ID:       idStr.(int64),
		Username: username.(string),
		Title:    title.(string),
		Content:  content.(string),
	}
	return post, nil
}

func (pr *postRepo) Update(post model.Post) error {
	postRecord := data.Record{
		"ID":       post.ID,
		"Username": post.Username,
		"Title":    post.Title,
		"Content":  post.Content,
	}
	return pr.db.Tables["posts"].Update(post.ID, postRecord)
}

func (pr *postRepo) Delete(id int) error {
	key := strconv.Itoa(id)
	return pr.db.Tables["posts"].Delete(key)
}
