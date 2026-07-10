package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserNameInfo struct {
	Username string `bson:"username"`
	Nickname string `bson:"nickname"`
}

var (
	EMPTY_USER_NAME_INFO = &UserNameInfo{}
	userNameInfoFields   = getFields(EMPTY_USER, EMPTY_USER_NAME_INFO)
)

func GetUserNameInfo(userID bbs.UUserID) (userNameInfo *UserNameInfo, err error) {
	query := bson.M{
		USER_USER_ID_b: userID,
	}

	err = User_c.FindOne(query, &userNameInfo, userNameInfoFields)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return userNameInfo, nil
}
