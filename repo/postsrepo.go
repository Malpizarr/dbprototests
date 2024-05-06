package repo

import (
	"fmt"
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
	db *data.Database
}

func NewPostRepo(db *data.Database) PostRepo {
	return &postRepo{db: db}
}

func (pr *postRepo) Create(post model.Post) error {
	s := strconv.Itoa(post.ID)
	postRecord := data.Record{
		"ID":       s,
		"Username": post.Username,
		"Title":    post.Title,
		"Content":  post.Content,
	}
	err := pr.db.Tables["posts"].Insert(postRecord)
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
	posts := make([]model.Post, 0)
	for _, record := range records {
		if record == nil {
			continue
		}
		idStr, ok1 := record.Fields["ID"]
		username, ok2 := record.Fields["Username"]
		title, ok3 := record.Fields["Title"]
		content, ok4 := record.Fields["Content"]
		if !ok1 || !ok2 || !ok3 || !ok4 {
			continue
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			continue
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
		return nil, err
	}

	var posts []model.Post
	for _, record := range joinedRecords {
		idStr, ok := record["t1.ID"].(string)
		if !ok {
			continue
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			continue
		}

		post := model.Post{
			ID:       id,
			Username: record["t1.Username"].(string),
			Title:    record["t1.Title"].(string),
			Content:  record["t1.Content"].(string),
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
	idStr, ok1 := record.Fields["ID"]
	username, ok2 := record.Fields["Username"]
	title, ok3 := record.Fields["Title"]
	content, ok4 := record.Fields["Content"]
	if !ok1 || !ok2 || !ok3 || !ok4 {
		return model.Post{}, fmt.Errorf("error getting post by id")
	}
	idS, err := strconv.Atoi(idStr)
	if err != nil {
		return model.Post{}, fmt.Errorf("error converting id to int")
	}
	post := model.Post{
		ID:       idS,
		Username: username,
		Title:    title,
		Content:  content,
	}
	return post, nil
}

func (pr *postRepo) Update(post model.Post) error {
	key := strconv.Itoa(post.ID)
	postRecord := data.Record{
		"ID":       post.ID,
		"Username": post.Username,
		"Title":    post.Title,
		"Content":  post.Content,
	}
	return pr.db.Tables["posts"].Update(key, postRecord)
}

func (pr *postRepo) Delete(id int) error {
	key := strconv.Itoa(id)
	return pr.db.Tables["posts"].Delete(key)
}
