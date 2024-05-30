package repo

import (
	"fmt"

	"github.com/Malpizarr/dbproto/pkg/data"
	model "github.com/Malpizarr/testwprotodb/data"
)

type FriendshipRepo interface {
	Create(friendship model.Friendship) error
	GetFriendship(id int) (*model.Friendship, error)
	AcceptFriendship(id int) error
	RejectFriendship(id int) error
	DeleteFriendship(id int) error
	GetAll() ([]model.Friendship, error)
}

type friendshipRepo struct {
	db       *data.Database
	userRepo UserRepo
}

func NewFriendshipRepo(db *data.Database, ur UserRepo) FriendshipRepo {
	return &friendshipRepo{db: db, userRepo: ur}
}

func (fr *friendshipRepo) Create(friendship model.Friendship) error {
	query := data.Query{
		Filters: map[string]interface{}{
			"User1": friendship.User1,
			"User2": friendship.User2,
		},
		Limit: 1,
	}
	friendshipExists, err := fr.db.Tables["friendships"].Query(query)
	if err != nil {
		return err
	}

	if len(friendshipExists) > 0 {
		return fmt.Errorf("error: friendship record already exists")
	}
	user1exists, err := fr.userRepo.Get(friendship.User1)
	user2exists, err := fr.userRepo.Get(friendship.User2)
	if user1exists == nil || user2exists == nil {
		return fmt.Errorf("error: user does not exist")
	}

	friendshipRecord := data.Record{
		"ID":     friendship.ID,
		"User1":  friendship.User1,
		"User2":  friendship.User2,
		"Status": friendship.Status,
	}
	err = fr.db.Tables["friendships"].Insert(friendshipRecord)
	fmt.Print(fr.db.Tables["friendships"].PrimaryKey)
	if err != nil {
		return err
	}
	return nil
}

func (fr *friendshipRepo) GetFriendship(id int) (*model.Friendship, error) {
	friendshipRecord, err := fr.db.Tables["friendships"].Select(id)
	if err != nil {
		return nil, err
	}
	idS, ok1 := friendshipRecord["ID"]
	user1Str, ok2 := friendshipRecord["User1"]
	user2Str, ok3 := friendshipRecord["User2"]
	statusStr, ok4 := friendshipRecord["Status"]
	if !ok1 || !ok2 || !ok3 || !ok4 {
		return nil, fmt.Errorf("error: friendship record not found")
	}

	friendship := model.Friendship{
		ID:     idS.(int64),
		User1:  user1Str.(string),
		User2:  user2Str.(string),
		Status: statusStr.(string),
	}
	return &friendship, nil
}

func (fr *friendshipRepo) AcceptFriendship(id int) error {
	friendshipRecord, err := fr.db.Tables["friendships"].Select(id)
	if err != nil {
		return err
	}

	statusValue, ok := friendshipRecord["Status"]
	if !ok || statusValue == nil {
		return fmt.Errorf("error: friendship status field is missing or not a proper *structpb.Value")
	}

	currentStatus := statusValue.(string)
	if currentStatus == "accepted" {
		return nil
	}

	updates := data.Record{"Status": "accepted"}

	err = fr.db.Tables["friendships"].Update(id, updates)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (fr *friendshipRepo) RejectFriendship(id int) error {
	friendshipRecord, err := fr.db.Tables["friendships"].Select(id)
	if err != nil {
		return err
	}
	statusValue, ok := friendshipRecord["Status"]
	if !ok || statusValue == nil {
		return fmt.Errorf("error: friendship status field is missing or not a proper *structpb.Value")
	}
	currentStatus := statusValue.(string)
	if currentStatus == "rejected" {
		return nil
	}
	updates := data.Record{"Status": "rejected"}
	err = fr.db.Tables["friendships"].Update(id, updates)
	if err != nil {
		return err
	}
	return nil
}

func (fr *friendshipRepo) DeleteFriendship(id int) error {
	err := fr.db.Tables["friendships"].Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (fr *friendshipRepo) GetAll() ([]model.Friendship, error) {
	records, err := fr.db.Tables["friendships"].SelectAll()
	if err != nil {
		return nil, err
	}
	var friendships []model.Friendship
	for _, record := range records {
		idStr, ok1 := record["ID"]
		user1Str, ok2 := record["User1"]
		user2Str, ok3 := record["User2"]
		statusStr, ok4 := record["Status"]
		if !ok1 || !ok2 || !ok3 || !ok4 {
			return nil, fmt.Errorf("error: error getting friendship records")
		}

		friendship := model.Friendship{
			ID:     idStr.(int64),
			User1:  user1Str.(string),
			User2:  user2Str.(string),
			Status: statusStr.(string),
		}
		friendships = append(friendships, friendship)
	}
	return friendships, nil
}
