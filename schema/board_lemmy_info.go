package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type BoardLemmyInfo struct {
	BBoardID bbs.BBoardID `bson:"bid"`

	LemmyID       int          `bson:"lemmy_id"` // lemmy id.
	LemmyUpdateNS types.NanoTS `bson:"lemmy_update_ns"`
}

var (
	EMPTY_BOARD_LEMMY_INFO = &BoardLemmyInfo{}
	boardLemmyInfoFields   = getFields(EMPTY_BOARD, EMPTY_BOARD_LEMMY_INFO)
)

func GetBoardLemmyInfo(boardID bbs.BBoardID) (ret *BoardLemmyInfo, err error) {
	query := &BoardQuery{
		BBoardID: boardID,
	}

	ret = &BoardLemmyInfo{}
	err = Board_c.FindOne(query, ret, boardLemmyInfoFields)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return ret, nil
}
