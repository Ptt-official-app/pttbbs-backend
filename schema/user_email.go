package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserEmail_c *db.Collection

type UserEmail struct {
	UserID bbs.UUserID `bson:"user_id"`
	Email  string      `bson:"email"`

	IsDefault bool `bson:"is_default,omitempty"`

	CreateNanoTS types.NanoTS `bson:"create_nano_ts"`
	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}

var EMPTY_USER_EMAIL = &UserEmail{}

var (
	USER_EMAIL_USER_ID_b        = getBSONName(EMPTY_USER_EMAIL, "UserID")
	USER_EMAIL_EMAIL_b          = getBSONName(EMPTY_USER_EMAIL, "Email")
	USER_EMAIL_IS_DEFAULT_b     = getBSONName(EMPTY_USER_EMAIL, "IsDefault")
	USER_EMAIL_CREATE_NANO_TS_b = getBSONName(EMPTY_USER_EMAIL, "CreateNanoTS")
	USER_EMAIL_UPDATE_NANO_TS_b = getBSONName(EMPTY_USER_EMAIL, "UpdateNanoTS")
)

func GetUserEmailByUserID(userID bbs.UUserID) (userEmail *UserEmail, err error) {
	query := bson.M{
		USER_EMAIL_USER_ID_b: userID,
	}

	return getUserEmailCore(query)
}

func GetUserEmailByEmail(email string) (userEmail *UserEmail, err error) {
	query := bson.M{
		USER_EMAIL_EMAIL_b: email,
	}

	return getUserEmailCore(query)
}

func getUserEmailCore(query bson.M) (userEmail *UserEmail, err error) {
	userEmail = &UserEmail{}
	err = UserEmail_c.FindOne(query, userEmail, nil)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = nil
		}
		return nil, err
	}

	return userEmail, nil
}

func CreateUserEmail(userID bbs.UUserID, email string, isDefault bool, updateNanoTS types.NanoTS) (err error) {
	// only 1 record for each email, but potentially multiple records for some user-ids.
	query := bson.M{
		USER_EMAIL_EMAIL_b: email,
	}

	userEmail := &UserEmail{
		UserID:       userID,
		Email:        email,
		IsDefault:    isDefault,
		CreateNanoTS: updateNanoTS,
		UpdateNanoTS: updateNanoTS,
	}

	ret, err := UserEmail_c.CreateOnly(query, userEmail)
	if err != nil {
		return err
	}
	if ret.UpsertedCount != 1 {
		return ErrNoCreate
	}

	return nil
}

func UpdateUserEmailIsDefault(userID bbs.UUserID, email string, isDefault bool, updateNanoTS types.NanoTS) (err error) {
	query := bson.M{
		USER_EMAIL_USER_ID_b: userID,
		USER_EMAIL_EMAIL_b:   email,
		USER_EMAIL_UPDATE_NANO_TS_b: bson.M{
			"$lt": updateNanoTS,
		},
	}

	toUpdate := bson.M{
		USER_EMAIL_IS_DEFAULT_b:     isDefault,
		USER_EMAIL_UPDATE_NANO_TS_b: updateNanoTS,
	}

	r, err := UserEmail_c.UpdateOneOnly(query, toUpdate)
	if err != nil {
		return err
	}

	if r.MatchedCount == 0 {
		return ErrNoMatch
	}

	return nil
}
