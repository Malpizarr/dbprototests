package repo

import (
	"fmt"

	"github.com/Malpizarr/dbproto/pkg/data"
	model "github.com/Malpizarr/testwprotodb/data"
)

type FriendshipRepo interface {
	Create(friendship model.Friendship) error
	GetFriendship(id int) (model.Friendship, error)
	AcceptFriendship(id int) error
	RejectFriendship(id int) error
	DeleteFriendship(id int) error
}

type friendshipRepo struct {
	db *data.Database
}

func NewFriendshipRepo(db *data.Database) FriendshipRepo {
	return &friendshipRepo{db: db}
}

func (fr *friendshipRepo) Create(friendship model.Friendship) error {
	friendshipRecord := data.Record{
		"ID":     friendship.ID,
		"User1":  friendship.User1,
		"User2":  friendship.User2,
		"Status": friendship.Status,
	}
	err := fr.db.Tables["friendships"].Insert(friendshipRecord)
	if err != nil {
		return err
	}
	return nil
}

func (fr *friendshipRepo) GetFriendship(id int) (model.Friendship, error) {
	friendshipRecord, err := fr.db.Tables["friendships"].Select(id)
	if err != nil {
		return model.Friendship{}, err
	}
	idStr, ok1 := friendshipRecord.Fields["ID"]
	user1Str, ok2 := friendshipRecord.Fields["User1"]
	user2Str, ok3 := friendshipRecord.Fields["User2"]
	statusStr, ok4 := friendshipRecord.Fields["Status"]
	if !ok1 || !ok2 || !ok3 || !ok4 {
		return model.Friendship{}, fmt.Errorf("error: friendship record not found")
	}

	iD := idStr.GetNumberValue()
	idS := int(iD)
	fmt.Println(idS)

	friendship := model.Friendship{
		ID:     idS,
		User1:  user1Str.GetStringValue(),
		User2:  user2Str.GetStringValue(),
		Status: statusStr.GetStringValue(),
	}
	return friendship, nil
}

func (fr *friendshipRepo) AcceptFriendship(id int) error {
	friendshipRecord, err := fr.db.Tables["friendships"].Select(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	statusValue, ok := friendshipRecord.Fields["Status"]
	if !ok || statusValue == nil {
		return fmt.Errorf("error: friendship status field is missing or not a proper *structpb.Value")
	}

	currentStatus := statusValue.GetStringValue()
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
	statusValue, ok := friendshipRecord.Fields["Status"]
	if !ok || statusValue == nil {
		return fmt.Errorf("error: friendship status field is missing or not a proper *structpb.Value")
	}
	currentStatus := statusValue.GetStringValue()
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
