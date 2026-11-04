package schema

import (
	"errors"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserIDEmail_c *db.Collection

type UserIDEmail struct {
	UserID  bbs.UUserID `bson:"user_id"`
	IDEmail string      `bson:"idemail"`

	IsSet bool `bson:"is_set,omitempty"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}

var EMPTY_USER_ID_EMAIL = &UserIDEmail{}

var (
	USER_ID_EMAIL_USER_ID_b        = getBSONName(EMPTY_USER_ID_EMAIL, "UserID")
	USER_ID_EMAIL_ID_EMAIL_b       = getBSONName(EMPTY_USER_ID_EMAIL, "IDEmail")
	USER_ID_EMAIL_IS_SET_b         = getBSONName(EMPTY_USER_ID_EMAIL, "IsSet")
	USER_ID_EMAIL_UPDATE_NANO_TS_b = getBSONName(EMPTY_USER_ID_EMAIL, "UpdateNanoTS")
)

func GetUserIDEmailByUserID(userID bbs.UUserID, updateNanoTS types.NanoTS) (userIDEmail *UserIDEmail, err error) {
	query := bson.M{
		USER_ID_EMAIL_USER_ID_b: userID,
	}

	return getUserIDEmailCore(query, updateNanoTS)
}

func GetUserIDEmailByEmail(email string, updateNanoTS types.NanoTS) (userIDEmail *UserIDEmail, err error) {
	query := bson.M{
		USER_ID_EMAIL_ID_EMAIL_b: email,
	}

	return getUserIDEmailCore(query, updateNanoTS)
}

func getUserIDEmailCore(query bson.M, updateNanoTS types.NanoTS) (userIDEmail *UserIDEmail, err error) {
	userIDEmail = &UserIDEmail{}
	err = UserIDEmail_c.FindOne(query, userIDEmail, nil)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = nil
		}
		return nil, err
	}

	if userIDEmail.IsSet {
		return userIDEmail, nil
	}

	if updateNanoTS-userIDEmail.UpdateNanoTS < types.EXPIRE_USER_ID_EMAIL_IS_NOT_SET_NANO_TS {
		return userIDEmail, nil
	}

	// to remove query.
	err = UserIDEmail_c.Remove(query)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func CreateUserIDEmail(userID bbs.UUserID, email string, updateNanoTS types.NanoTS) (err error) {
	query := bson.M{
		USER_ID_EMAIL_USER_ID_b: userID,
	}

	userIDEmail := &UserIDEmail{
		UserID:       userID,
		IDEmail:      email,
		UpdateNanoTS: updateNanoTS,
	}

	r, createErr := UserIDEmail_c.CreateOnly(query, userIDEmail)
	if createErr == nil && r.UpsertedCount > 0 {
		return nil
	}

	// check duplicate
	if createErr != nil {
		return createErr
	}

	gotUserIDEmail := &UserIDEmail{}
	err = UserIDEmail_c.FindOne(query, gotUserIDEmail, nil)
	if err != nil {
		return err
	}

	if gotUserIDEmail.IsSet && updateNanoTS-gotUserIDEmail.UpdateNanoTS < types.EXPIRE_USER_ID_EMAIL_IS_SET_NANO_TS {
		return ErrNoCreate
	}

	if updateNanoTS-gotUserIDEmail.UpdateNanoTS < types.EXPIRE_USER_ID_EMAIL_IS_NOT_SET_NANO_TS {
		return ErrNoCreate
	}

	query[USER_ID_EMAIL_UPDATE_NANO_TS_b] = bson.M{
		"$lt": updateNanoTS,
	}

	_, err = UserIDEmail_c.UpdateOneOnly(query, userIDEmail)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUserIDEmailIsSet(userID bbs.UUserID, email string, isSet bool, updateNanoTS types.NanoTS) (err error) {
	query := bson.M{
		USER_ID_EMAIL_USER_ID_b:  userID,
		USER_ID_EMAIL_ID_EMAIL_b: email,
		USER_ID_EMAIL_UPDATE_NANO_TS_b: bson.M{
			"$lt": updateNanoTS,
		},
	}

	toUpdate := bson.M{
		USER_ID_EMAIL_IS_SET_b:         isSet,
		USER_ID_EMAIL_UPDATE_NANO_TS_b: updateNanoTS,
	}

	r, err := UserIDEmail_c.UpdateOneOnly(query, toUpdate)
	if err != nil {
		return err
	}

	if r.MatchedCount == 0 {
		return ErrNoMatch
	}

	return nil
}
