package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type BoardID struct {
	BBoardID bbs.BBoardID `bson:"bid"`
	Brdname  string       `bson:"brdname"`
}

var (
	EMPTY_BOARD_ID = &BoardID{}
	boardIDFields  = getFields(EMPTY_BOARD, EMPTY_BOARD_ID)
)

func GetBoardID(brdname string) (boardID bbs.BBoardID, err error) {
	query := bson.M{
		BOARD_BRDNAME_b: brdname,
	}

	result := &BoardID{}
	err = Board_c.FindOne(query, result, boardIDFields)
	if err != nil {
		logrus.Warnf("schema.GetBoardID: unable to FindOne: query: %v e: %v", query, err)
		return "", err
	}

	return result.BBoardID, nil
}

func GetBoardIDByPttbid(bid ptttype.Bid) (boardID bbs.BBoardID, err error) {
	if bid == 0 {
		return "", ErrInvalidPTTBID
	}

	query := bson.M{
		BOARD_BID_b: bid,
	}

	result := &BoardID{}
	err = Board_c.FindOne(query, result, boardIDFields)
	if err != nil {
		logrus.Warnf("schema.GetBoardIDByPttbid: unable to FindOne: query: %v e: %v", query, err)
		return "", err
	}

	return result.BBoardID, nil
}

func GetBoardIDs(startBrdname string, descending bool, limit int, withDeleted bool) (result []*BoardID, err error) {
	// setup query
	var query bson.M
	if startBrdname == "" {
		query = bson.M{}
	} else {
		theDir := "$gte"
		if descending {
			theDir = "$lte"
		}
		query = bson.M{
			BOARD_BRDNAME_b: bson.M{
				theDir: startBrdname,
			},
		}

	}

	if !withDeleted {
		query[BOARD_IS_DELETED_b] = bson.M{
			"$exists": false,
		}
	}
	// sort opts
	var sortOpts bson.D
	if descending {
		sortOpts = bson.D{
			{Key: BOARD_BRDNAME_b, Value: -1},
		}
	} else {
		sortOpts = bson.D{
			{Key: BOARD_BRDNAME_b, Value: 1},
		}
	}

	// find
	err = Board_c.Find(query, int64(limit), &result, boardIDFields, sortOpts)
	if err != nil {
		return nil, err
	}

	return result, nil
}
