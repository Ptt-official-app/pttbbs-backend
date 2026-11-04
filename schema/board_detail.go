package schema

import (
	"errors"
	"strings"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	ptttypes "github.com/Ptt-official-app/go-pttbbs/types"
	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BoardDetail struct {
	BBoardID  bbs.BBoardID    `bson:"bid"`
	Brdname   string          `bson:"brdname"`
	Title     string          `bson:"title"`
	BrdAttr   ptttype.BrdAttr `bson:"flag"`
	BoardType string          `bson:"the_type"`
	Category  string          `bson:"class"`
	NUser     int             `bson:"nuser"` /* use aggregate to periodically get the data */
	BMs       []bbs.UUserID   `bson:"bms"`
	Total     int             `bson:"total"` /* total articles, 需要即時知道. 因為 read 頻率高. 並且跟 last-post-time-ts 一樣 write 頻率 << read 頻率 */

	LastPostTime types.NanoTS `bson:"last_post_time_nano_ts"` /* 需要即時知道來做板的已讀 */

	UpdateTime types.NanoTS `bson:"update_time_nano_ts"` /* show 進板畫面, 目前只有 INT_MAX - 1 或 0 */

	VoteLimitLogins  int `bson:"vote_limit_logins"`
	PostLimitLogins  int `bson:"post_limit_logins"`
	VoteLimitBadpost int `bson:"vote_limit_bad_post"`
	PostLimitBadpost int `bson:"post_limit_bad_post"`

	ParentID bbs.BBoardID `bson:"parent"`

	NVote           int          `bson:"vote"` /* use db-count to get current #vote */
	VoteClosingTime types.NanoTS `bson:"vtime_nano_ts"`

	Level              ptttype.PERM `bson:"perm"`
	LastSetTime        types.NanoTS `bson:"last_set_time_nano_ts"` /* perm-reload */
	PostExpire         ptttype.Bid  `bson:"post_expire"`           /* 看板連結的 bid */
	PostType           []string     `bson:"post_type"`
	PostTemplate       []bool       `bson:"post_tmpl"`
	EndGambleNanoTS    types.NanoTS `bson:"end_gamble_nano_ts"`
	FastRecommendPause types.NanoTS `bson:"fast_recommend_pause_nano_ts"`

	IsDeleted bool `bson:"deleted,omitempty"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`

	Gid        ptttype.Bid `bson:"pttgid"`
	Bid        ptttype.Bid `bson:"pttbid"`
	IdxByName  string      `bson:"pttidxname"`
	IdxByClass string      `bson:"pttidxclass"`

	ChessCountry ptttype.ChessCode `bson:"chesscountry"`
}

var (
	EMPTY_BOARD_DETAIL = &BoardDetail{}
	boardDetailFields  = getFields(EMPTY_BOARD, EMPTY_BOARD_DETAIL)
)

func NewBoardDetail(b_b *bbs.BoardDetail, updateNanoTS types.NanoTS) *BoardDetail {
	postType := make([]string, 0, len(b_b.PostType))
	nIsValid := 0
	for _, each := range b_b.PostType {
		eachPostType := strings.TrimSpace(types.Big5ToUtf8(ptttypes.CstrToBytes(each)))
		if len(eachPostType) > 0 {
			nIsValid++
		}
		postType = append(postType, eachPostType)
	}

	if nIsValid == 0 {
		postType = DEFAULT_POST_TYPE
	}

	var parentID bbs.BBoardID

	// XXX hack for Gid == 1 (1 should be SYSOP)
	if b_b.Gid != 1 {
		parentID, _ = GetBoardIDByPttbid(b_b.Gid)
	}

	return &BoardDetail{
		BBoardID:  b_b.BBoardID,
		Brdname:   b_b.Brdname,
		Title:     types.Big5ToUtf8(b_b.RealTitle),
		BrdAttr:   b_b.BrdAttr,
		BoardType: types.Big5ToUtf8(b_b.BoardType),
		Category:  types.Big5ToUtf8(b_b.BoardClass),
		BMs:       b_b.BM,
		Total:     int(b_b.Total),
		NUser:     int(b_b.NUser),

		LastPostTime: types.Time4ToNanoTS(b_b.LastPostTime),

		UpdateTime: types.Time4ToNanoTS(b_b.BUpdate),

		VoteLimitLogins:  int(b_b.VoteLimitLogins),
		PostLimitLogins:  int(b_b.PostLimitLogins),
		VoteLimitBadpost: int(b_b.VoteLimitBadpost),
		PostLimitBadpost: int(b_b.PostLimitBadPost),

		NVote:           int(b_b.BVote),
		VoteClosingTime: types.Time4ToNanoTS(b_b.VTime),

		Level:              b_b.Level,
		LastSetTime:        types.Time4ToNanoTS(b_b.PermReload),
		PostExpire:         b_b.PostExpire,
		PostType:           postType,
		PostTemplate:       b_b.PostTypeTemplate,
		EndGambleNanoTS:    types.Time4ToNanoTS(b_b.EndGamble),
		FastRecommendPause: types.Time4ToNanoTS(ptttypes.Time4(b_b.FastRecommendPause)),

		UpdateNanoTS: updateNanoTS,

		ParentID:   parentID,
		Gid:        b_b.Gid,
		Bid:        b_b.Bid,
		IdxByName:  b_b.IdxByName,
		IdxByClass: b_b.IdxByClass,

		ChessCountry: b_b.ChessCountry,
	}
}

func UpdateBoardDetails(boardDetails []*BoardDetail, updateNanoTS types.NanoTS) (err error) {
	if len(boardDetails) == 0 {
		return nil
	}

	// create items which do not exists yet.
	theList := make([]*db.UpdatePair, len(boardDetails))
	for idx, each := range boardDetails {
		query := &BoardQuery{
			BBoardID: each.BBoardID,
		}

		theList[idx] = &db.UpdatePair{
			Filter: query,
			Update: each,
		}
	}

	r, err := Board_c.BulkCreateOnly(theList)
	if err != nil {
		return err
	}
	if r.UpsertedCount == int64(len(boardDetails)) { // all are created
		return nil
	}

	// update items with comparing update-nano-ts
	upsertedIDs := r.UpsertedIDs
	updateBoardDetails := make([]*db.UpdatePair, 0, len(theList))
	for idx, each := range theList {
		_, ok := upsertedIDs[int64(idx)]
		if ok {
			continue
		}

		origFilter := each.Filter.(*BoardQuery)
		filter := bson.M{
			"$or": bson.A{
				bson.M{
					BOARD_BBOARD_ID_b: origFilter.BBoardID,
					BOARD_UPDATE_NANO_TS_b: bson.M{
						"$exists": false,
					},

					BOARD_IS_DELETED_b: bson.M{"$exists": false},
				},
				bson.M{
					BOARD_BBOARD_ID_b: origFilter.BBoardID,
					BOARD_UPDATE_NANO_TS_b: bson.M{
						"$lt": updateNanoTS,
					},

					BOARD_IS_DELETED_b: bson.M{"$exists": false},
				},
			},
		}
		each.Filter = filter
		updateBoardDetails = append(updateBoardDetails, each)
	}

	_, err = Board_c.BulkUpdateOneOnly(updateBoardDetails)

	return err
}

func GetBoardDetail(bboardID bbs.BBoardID, fields map[string]bool) (result *BoardDetail, err error) {
	query := &BoardQuery{
		BBoardID: bboardID,
	}

	result = &BoardDetail{}
	if fields == nil {
		fields = boardDetailFields
	}
	err = Board_c.FindOne(query, result, fields)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}
