package schema

import (
	"errors"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserInfoSummary struct {
	UserID       bbs.UUserID  `bson:"user_id"`
	IsDeleted    bool         `bson:"deleted,omitempty"`
	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}

var (
	EMPTY_USER_INFO_SUMMARY = &UserInfoSummary{}
	userInfoSummaryFields   = getFields(EMPTY_USER, EMPTY_USER_INFO_SUMMARY)
)

func GetUserInfoSummary(userID bbs.UUserID) (result *UserInfoSummary, err error) {
	query := bson.M{
		USER_USER_ID_b: userID,
	}

	err = User_c.FindOne(query, &result, userInfoSummaryFields)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}
