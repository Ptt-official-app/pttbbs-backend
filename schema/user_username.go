package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserUsername struct {
	UserID     bbs.UUserID  `bson:"user_id"`
	Username   string       `bson:"username"`
	UsernameNS types.NanoTS `bson:"username_ns"`
}

var (
	EMPTY_USER_USERNAME = &UserUsername{}
	userUsernameFields  = getFields(EMPTY_USER, EMPTY_USER_USERNAME)
)

func CreateUsername(userID bbs.UUserID, username string, updateNS types.NanoTS) (err error) {
	query := bson.M{
		USER_USER_ID_b: userID,
	}

	userUsername := &UserUsername{
		UserID:     userID,
		Username:   username,
		UsernameNS: updateNS,
	}

	_, err = User_c.CreateOnly(query, userUsername)
	if err != nil {
		return err
	}

	return nil
}

func GetUserUsernameByUsername(username string) (userUsername *UserUsername, err error) {
	query := bson.M{
		USER_USERNAME_b: username,
	}

	return getUserUsernameCore(query)
}

func GetUserUsernameByUserID(userID bbs.UUserID) (userUsername *UserUsername, err error) {
	query := bson.M{
		USER_USER_ID_b: userID,
	}

	return getUserUsernameCore(query)
}

func getUserUsernameCore(query bson.M) (userUsername *UserUsername, err error) {
	userUsername = &UserUsername{}
	err = User_c.FindOne(query, userUsername, userUsernameFields)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = nil
		}
		return nil, err
	}

	return userUsername, nil
}
