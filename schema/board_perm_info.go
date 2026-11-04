package schema

import (
	"errors"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// BoardPermInfo
//
// Information related to Board permission
type BoardPermInfo struct {
	BBoardID bbs.BBoardID `bson:"bid"`

	Brdname string `bson:"brdname"`

	BrdAttr ptttype.BrdAttr `bson:"flag"`

	BMs []bbs.UUserID `bson:"bms"`

	Level ptttype.PERM `bson:"perm"`

	ParentID bbs.BBoardID `bson:"parent"`

	VoteLimitLogins  int `bson:"vote_limit_logins"`
	VoteLimitBadpost int `bson:"vote_limit_bad_post"`

	PostLimitLogins  int `bson:"post_limit_logins"`
	PostLimitBadpost int `bson:"post_limit_bad_post"`

	IsOver18 bool `bson:"is_over_18"` // 18歲板

	NUser int `bson:"nuser"` /* use aggregate to periodically get the data */
}

var (
	EMPTY_BOARD_PERM_INFO = &BoardPermInfo{}
	boardPermInfoFields   = getFields(EMPTY_BOARD, EMPTY_BOARD_PERM_INFO)
)

func GetBoardPermInfo(boardID bbs.BBoardID) (boardPermInfo *BoardPermInfo, err error) {
	query := bson.M{
		BOARD_BBOARD_ID_b: boardID,
	}

	err = Board_c.FindOne(query, &boardPermInfo, boardPermInfoFields)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return boardPermInfo, nil
}
