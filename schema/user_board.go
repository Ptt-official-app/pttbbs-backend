package schema

import (
	"errors"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserBoard_c *db.Collection

type UserBoard struct {
	// 已讀板紀錄

	UserID           bbs.UUserID  `bson:"user_id"`
	BBoardID         bbs.BBoardID `bson:"bid"`
	ReadUpdateNanoTS types.NanoTS `bson:"update_nano_ts"`

	BoardFriend             bool         `bson:"board_friend"`
	BoardFriendUpdateNanoTS types.NanoTS `bson:"board_friend_update_nano_ts"`

	BoardBucketExpireNanoTS types.NanoTS `bson:"board_bucket_expire_nano_ts"` // 水桶
	BoardBucketReason       string       `bson:"board_bucket_reason"`
	BoardBucketUpdateNanoTS types.NanoTS `bson:"board_bucket_update_nano_ts"`

	BoardBlocked             bool         `bson:"board_blocked"` // 不看這個板
	BoardBlockedReason       string       `bson:"board_blocked_reason"`
	BoardBlockedUpdateNanoTS types.NanoTS `bson:"board_blocked_update_nano_ts"`

	BoardReported             bool         `bson:"board_reported"` // 檢舉這個板
	BoardReportedReason       string       `bson:"board_reported_reason"`
	BoardReportedUpdateNanoTS types.NanoTS `bson:"board_reported_update_nano_ts"`
}

var EMPTY_USER_BOARD = &UserBoard{}

var (
	USER_BOARD_USER_ID_b             = getBSONName(EMPTY_USER_BOARD, "UserID")
	USER_BOARD_BBOARD_ID_b           = getBSONName(EMPTY_USER_BOARD, "BBoardID")
	USER_BOARD_READ_UPDATE_NANO_TS_b = getBSONName(EMPTY_USER_BOARD, "ReadUpdateNanoTS")

	USER_BOARD_BOARD_FRIEND_b                = getBSONName(EMPTY_USER_BOARD, "BoardFriend")
	USER_BOARD_BOARD_FRIEND_UPDATE_NANO_TS_b = getBSONName(EMPTY_USER_BOARD, "BoardFriendUpdateNanoTS")

	USER_BOARD_BOARD_BUCKET_b                = getBSONName(EMPTY_USER_BOARD, "BoardBucket")
	USER_BOARD_BOARD_BUCKET_REASON_b         = getBSONName(EMPTY_USER_BOARD, "BoardBucketReason")
	USER_BOARD_BOARD_BUCKET_UPDATE_NANO_TS_b = getBSONName(EMPTY_USER_BOARD, "BoardBucketUpdateNanoTS")

	USER_BOARD_BOARD_BLOCKED_b                = getBSONName(EMPTY_USER_BOARD, "BoardBlocked")
	USER_BOARD_BOARD_BLOCKED_REASON_b         = getBSONName(EMPTY_USER_BOARD, "BoardBlockedReason")
	USER_BOARD_BOARD_BLOCKED_UPDATE_NANO_TS_b = getBSONName(EMPTY_USER_BOARD, "BoardBlockedUpdateNanoTS")

	USER_BOARD_BOARD_REPORTED_b                = getBSONName(EMPTY_USER_BOARD, "BoardReported")
	USER_BOARD_BOARD_REPORTED_REASON_b         = getBSONName(EMPTY_USER_BOARD, "BoardReportedReason")
	USER_BOARD_BOARD_REPORTED_UPDATE_NANO_TS_b = getBSONName(EMPTY_USER_BOARD, "BoardReportedUpdateNanoTS")
)

func assertUserBoardFields() error {
	if err := assertFields(EMPTY_USER_BOARD, EMPTY_USER_READ_BOARD); err != nil {
		return err
	}

	if err := assertFields(EMPTY_USER_BOARD, EMPTY_USER_READ_BOARD_QUERY); err != nil {
		return err
	}

	return nil
}

func FindUserBoard(userID bbs.UUserID, boardID bbs.BBoardID) (result *UserBoard, err error) {
	// query
	query := bson.M{
		USER_BOARD_USER_ID_b:   userID,
		USER_BOARD_BBOARD_ID_b: boardID,
	}

	err = UserBoard_c.FindOne(query, &result, nil)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return &UserBoard{UserID: userID, BBoardID: boardID}, nil
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}

func FindUserBoards(userID bbs.UUserID, boardIDs []bbs.BBoardID) (result []*UserBoard, err error) {
	// query
	query := bson.M{
		USER_BOARD_USER_ID_b: userID,
		USER_BOARD_BBOARD_ID_b: bson.M{
			"$in": boardIDs,
		},
	}

	err = UserBoard_c.Find(query, 0, &result, nil, nil)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return []*UserBoard{}, nil
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}
