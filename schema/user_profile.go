package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/text/language"
)

type UserProfile struct {
	UserID   bbs.UUserID `bson:"user_id"`
	Username string      `bson:"username"`

	FamilyName string       `bson:"familyname"`
	GivenName  string       `bson:"givenname"`
	Locale     language.Tag `bson:"locale"`
}

var EMPTY_USER_PROFILE = &UserProfile{}

func GetUserProfile(userID bbs.UUserID) (userProfile *UserProfile, err error) {
	query := bson.M{
		USER_USER_ID_b: userID,
	}
	userProfile = &UserProfile{}

	err = User_c.FindOne(query, userProfile, nil)
	if err != nil {
		return nil, err
	}

	return userProfile, nil
}
