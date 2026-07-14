package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
)

type UserIsGovernmentIDOver18 struct {
	UserID bbs.UUserID `bson:"user_id"`

	IsGovernmentID   bool         `bson:"is_government_id"`
	IsGovernmentIDNS types.NanoTS `bson:"is_government_id_ns"`
	Over18           bool         `bson:"over18"`
}

var EMPTY_USER_IS_GOVERNMENT_ID_OVER18 = &UserIsGovernmentIDOver18{}

func SetUserIsGovernmentIDOver18(userID bbs.UUserID, updateNS types.NanoTS) (err error) {
	query := bson.M{
		USER_USER_ID_b: userID,
	}

	userIsGovernmentIDOver18 := &UserIsGovernmentIDOver18{
		UserID:           userID,
		IsGovernmentID:   true,
		IsGovernmentIDNS: updateNS,
		Over18:           true,
	}
	_, err = User_c.UpdateOneOnly(query, userIsGovernmentIDOver18)
	if err != nil {
		return err
	}

	return nil
}
