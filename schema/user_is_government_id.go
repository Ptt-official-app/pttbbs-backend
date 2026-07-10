package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
)

type UserIsGovernmentID struct {
	UserID bbs.UUserID `bson:"user_id"`

	IsGovernmentID   bool         `bson:"is_government_id"`
	IsGovernmentIDNS types.NanoTS `bson:"is_government_id_ns"`
}

func UpdateUserIsGovernmentID(userID bbs.UUserID, isGovernmentID bool, updateNS types.NanoTS) (err error) {
	query := bson.M{
		USER_USER_ID_b: userID,
	}

	userIsGovernmentID := &UserIsGovernmentID{
		UserID:           userID,
		IsGovernmentID:   isGovernmentID,
		IsGovernmentIDNS: updateNS,
	}
	_, err = User_c.UpdateOneOnly(query, userIsGovernmentID)
	if err != nil {
		return err
	}

	return nil
}
