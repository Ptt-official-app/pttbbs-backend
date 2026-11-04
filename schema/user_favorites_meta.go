package schema

import (
	"errors"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserFavoritesMeta_c *db.Collection

type UserFavoritesMeta struct {
	UserID          bbs.UUserID  `bson:"user_id"`
	DoubleBufferIdx int          `bson:"db_idx"`
	UpdateNanoTS    types.NanoTS `bson:"update_nano_ts"`
	MTime           types.NanoTS `bson:"mtime_nano_ts"`
	FolderMeta      FolderMeta   `bson:"folder_meta"`
}

var EMPTY_USER_FAVORITES_META = &UserFavoritesMeta{}

var (
	USER_FAVORITES_META_USER_ID_b           = getBSONName(EMPTY_USER_FAVORITES_META, "UserID")
	USER_FAVORITES_META_DOUBLE_BUFFER_IDX_b = getBSONName(EMPTY_USER_FAVORITES_META, "DoubleBufferIdx")
	USER_FAVORITES_META_UPDATE_NANO_TS_b    = getBSONName(EMPTY_USER_FAVORITES_META, "UpdateNanoTS")
	USER_FAVORITES_META_MTIME_b             = getBSONName(EMPTY_USER_FAVORITES_META, "MTime")
	USER_FAVORITES_META_FAV_NUM_b           = getBSONName(EMPTY_USER_FAVORITES_META, "FavNum")
)

func assertUserFavoritesMeta() error {
	if err := assertFields(EMPTY_USER_FAVORITES_META, EMPTY_USER_FAVORITES_META_SUMMARY); err != nil {
		return err
	}

	return nil
}

func GetUserFavoritesMeta(userID bbs.UUserID) (result *UserFavoritesMeta, err error) {
	query := bson.M{
		USER_FAVORITES_META_USER_ID_b: userID,
	}

	result = &UserFavoritesMeta{}
	err = UserFavoritesMeta_c.FindOne(query, result, nil)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateUserFavoritesMeta(meta *UserFavoritesMeta) (err error) {
	query := bson.M{
		USER_FAVORITES_META_USER_ID_b: meta.UserID,
	}

	r, err := UserFavoritesMeta_c.CreateOnly(query, meta)
	if err != nil {
		return err
	}
	if r.UpsertedCount > 0 {
		return nil
	}

	query = bson.M{
		USER_FAVORITES_META_USER_ID_b: meta.UserID,
		USER_FAVORITES_META_MTIME_b: bson.M{
			"$lt": meta.MTime,
		},
	}

	r, err = UserFavoritesMeta_c.UpdateOneOnly(query, meta)
	if err != nil {
		return err
	}
	if r.MatchedCount == 0 {
		return ErrNoMatch
	}
	return nil
}
