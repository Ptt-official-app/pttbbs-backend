package schema

import (
	"errors"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserUsernameOver18 struct {
	UserID   bbs.UUserID `bson:"user_id"`
	Username string      `bson:"username"`
	Over18   bool        `bson:"over18"`
}

var (
	EMPTY_USER_USERNAME_OVER18 = &UserUsernameOver18{}
	// userUsernameOver18Fields must be its own projection: userUsernameFields does not include over18.
	userUsernameOver18Fields = getFields(EMPTY_USER, EMPTY_USER_USERNAME_OVER18)
)

func GetUserUsernameOver18ByUsername(username string) (ret *UserUsernameOver18, err error) {
	query := bson.M{
		USER_USERNAME_b: username,
	}

	ret = &UserUsernameOver18{}
	err = User_c.FindOne(query, ret, userUsernameOver18Fields)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = nil
		}
		return nil, err
	}

	return ret, nil
}
