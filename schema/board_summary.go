package schema

import (
	"strings"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/boardd"
	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// BoardSummary
type BoardSummary struct {
	BBoardID  bbs.BBoardID    `bson:"bid"`
	Brdname   string          `bson:"brdname"`
	Title     string          `bson:"title"`
	BrdAttr   ptttype.BrdAttr `bson:"flag"`
	BoardType string          `bson:"the_type"`
	Category  string          `bson:"class"`
	NUser     int             `bson:"nuser"`
	BMs       []bbs.UUserID   `bson:"bms"`
	Total     int             `bson:"total"` /* total articles, 需要即時知道. 因為 read 頻率高. 並且跟 last-post-time-ts 一樣 write 頻率 << read 頻率 */

	LastPostTime types.NanoTS `bson:"last_post_time_nano_ts"` /* 需要即時知道來做板的已讀 */

	IsDeleted bool `bson:"deleted,omitempty"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`

	ParentID bbs.BBoardID `bson:"parent"`

	Gid        ptttype.Bid `bson:"pttgid"`
	Bid        ptttype.Bid `bson:"pttbid"`
	IdxByName  string      `bson:"pttidxname"`
	IdxByClass string      `bson:"pttidxclass"`

	IsPopular bool `bson:"is_popular"`
	IsOver18  bool `bson:"is_over_18"` // 18歲板

	IsOver18UpdateNanoTS types.NanoTS `bson:"is_over_18_update_nano_ts"`

	IsNoStats             bool `bson:"is_no_stats"`               // 不列入統計
	IsGroupBoard          bool `bson:"is_group_board"`            // 群組板
	IsHide                bool `bson:"is_hide"`                   // 隱板
	IsPostMask            bool `bson:"is_post_mask"`              // 限制發表或是閱讀
	IsAnonymous           bool `bson:"is_anony"`                  // 匿名板
	IsDefaultAnonymous    bool `bson:"is_default_anony"`          // 預設匿名板
	IsNoCredit            bool `bson:"is_no_credit"`              // 發文無獎勵板
	IsVoteBoard           bool `bson:"is_vote_board"`             // 連署機看板
	IsWarnEOL             bool `bson:"is_warn_eol"`               // 已警告要廢除
	IsNoComment           bool `bson:"is_no_comment"`             // 不可推文
	IsAngelAnonymous      bool `bson:"is_angel_anony"`            // 小天使可匿名
	IsSymLink             bool `bson:"is_sym_link"`               // 文章是 sym-link (AllPost/HiddenAllPost)
	IsNoBoo               bool `bson:"is_no_boo"`                 // 不可噓文
	IsBoardMemberOnlyPost bool `bson:"is_board_member_only_post"` // 板友才可以發文
	IsGuestPost           bool `bson:"is_guest_post"`             // guest 可發文
	IsCooldown            bool `bson:"is_cooldown"`               // 靜
	IsIPLogComment        bool `bson:"is_ip_log_comment"`         // 推文記錄 IP
	IsNoReply             bool `bson:"is_no_reply"`               // 不可回文
	IsAlignedComment      bool `bson:"is_aligned_comment"`        // 對齊式推文
	IsNoSelfDelPost       bool `bson:"is_no_self_del_post"`       // 不可以自己刪文
	IsBMMaskContent       bool `bson:"is_bm_mask_content"`        // 允許板主刪除特定文字
}

var (
	EMPTY_BOARD_SUMMARY = &BoardSummary{}
	boardSummaryFields  = getFields(EMPTY_BOARD, EMPTY_BOARD_SUMMARY)
)

func NewBoardSummary(b_b *bbs.BoardSummary, updateNanoTS types.NanoTS) *BoardSummary {
	parentID, _ := GetBoardIDByPttbid(b_b.Gid)

	return &BoardSummary{
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

		UpdateNanoTS: updateNanoTS,

		ParentID: parentID,

		Gid:        b_b.Gid,
		Bid:        b_b.Bid,
		IdxByName:  b_b.IdxByName,
		IdxByClass: b_b.IdxByClass,

		// IsPopular: b_b.BrdAttr.HasPerm(ptttype.BRD_TOP),
		IsOver18:             b_b.BrdAttr.HasPerm(ptttype.BRD_OVER18),
		IsOver18UpdateNanoTS: updateNanoTS,

		IsNoStats:             b_b.BrdAttr.HasPerm(ptttype.BRD_NOCOUNT),
		IsGroupBoard:          b_b.BrdAttr.HasPerm(ptttype.BRD_GROUPBOARD),
		IsHide:                b_b.BrdAttr.HasPerm(ptttype.BRD_HIDE),
		IsPostMask:            b_b.BrdAttr.HasPerm(ptttype.BRD_POSTMASK),
		IsAnonymous:           b_b.BrdAttr.HasPerm(ptttype.BRD_ANONYMOUS),
		IsDefaultAnonymous:    b_b.BrdAttr.HasPerm(ptttype.BRD_DEFAULTANONYMOUS),
		IsNoCredit:            b_b.BrdAttr.HasPerm(ptttype.BRD_NOCREDIT),
		IsVoteBoard:           b_b.BrdAttr.HasPerm(ptttype.BRD_VOTEBOARD),
		IsWarnEOL:             b_b.BrdAttr.HasPerm(ptttype.BRD_WARNEL),
		IsNoComment:           b_b.BrdAttr.HasPerm(ptttype.BRD_NORECOMMEND),
		IsAngelAnonymous:      b_b.BrdAttr.HasPerm(ptttype.BRD_ANGELANONYMOUS),
		IsSymLink:             b_b.BrdAttr.HasPerm(ptttype.BRD_SYMBOLIC),
		IsNoBoo:               b_b.BrdAttr.HasPerm(ptttype.BRD_NOBOO),
		IsBoardMemberOnlyPost: b_b.BrdAttr.HasPerm(ptttype.BRD_RESTRICTEDPOST),
		IsGuestPost:           b_b.BrdAttr.HasPerm(ptttype.BRD_GUESTPOST),
		IsCooldown:            b_b.BrdAttr.HasPerm(ptttype.BRD_COOLDOWN),
		IsIPLogComment:        b_b.BrdAttr.HasPerm(ptttype.BRD_IPLOGRECMD),
		IsNoReply:             b_b.BrdAttr.HasPerm(ptttype.BRD_NOREPLY),
		IsAlignedComment:      b_b.BrdAttr.HasPerm(ptttype.BRD_ALIGNEDCMT),
		IsNoSelfDelPost:       b_b.BrdAttr.HasPerm(ptttype.BRD_NOSELFDELPOST),
		IsBMMaskContent:       b_b.BrdAttr.HasPerm(ptttype.BRD_BM_MASK_CONTENT),
	}
}

func NewBoardSummaryFromPBBoard(b_b *boardd.Board, updateNanoTS types.NanoTS) *BoardSummary {
	rawModerators := strings.Split(b_b.RawModerators, "/")
	bms := make([]bbs.UUserID, 0, len(rawModerators))
	for _, each := range rawModerators {
		each = strings.TrimSpace(each)
		if each == "" {
			continue
		}
		bms = append(bms, bbs.UUserID(each))
	}

	idxByName := bbs.SerializeBoardIdxByNameStr(b_b.Name)
	clsBig5 := types.Utf8ToBig5(b_b.Bclass)
	idxByClass := bbs.SerializeBoardIdxByClassStr(clsBig5, idxByName)

	boardIDRaw := &ptttype.BoardID_t{}
	copy(boardIDRaw[:], []byte(b_b.Name))
	bboardID := bbs.ToBBoardID(ptttype.Bid(b_b.Bid), boardIDRaw)

	brdAttr := ptttype.BrdAttr(b_b.Attributes)
	boardType := "◎"
	if brdAttr.HasPerm(ptttype.BRD_GROUPBOARD) {
		boardType = "Σ"
	}

	isOver18 := brdAttr.HasPerm(ptttype.BRD_OVER18)

	parentID, _ := GetBoardIDByPttbid(ptttype.Bid(b_b.Parent))

	return &BoardSummary{
		BBoardID: bboardID,
		Brdname:  b_b.Name,
		Title:    b_b.Title,
		BrdAttr:  brdAttr,
		Category: b_b.Bclass,
		BMs:      bms,
		Total:    int(b_b.NumPosts),
		NUser:    int(b_b.NumUsers),

		IsOver18:             isOver18,
		IsOver18UpdateNanoTS: updateNanoTS,

		UpdateNanoTS: updateNanoTS,

		ParentID:   parentID,
		Gid:        ptttype.Bid(b_b.Parent),
		Bid:        ptttype.Bid(b_b.Bid),
		IdxByName:  idxByName,
		IdxByClass: idxByClass,

		BoardType: boardType,
	}
}

func NewBoardSummaryFromPttWeb(boardName string, boardTitle string, boardType string, boardClass string, nUser int, total int, lastPostTime types.NanoTS, isOver18 bool, updateNanoTS types.NanoTS) (boardSummary *BoardSummary, err error) {
	return &BoardSummary{
		BBoardID:  bbs.BBoardID(boardName),
		Brdname:   boardName,
		Title:     boardTitle,
		BoardType: boardType,
		Category:  boardClass,

		NUser:        nUser,
		Total:        total,
		LastPostTime: lastPostTime,

		IsPopular:            true,
		IsOver18:             isOver18,
		IsOver18UpdateNanoTS: updateNanoTS,

		UpdateNanoTS: updateNanoTS,
	}, nil
}

func UpdateBoardSummaries(boardSummaries []*BoardSummary, updateNanoTS types.NanoTS) (err error) {
	if len(boardSummaries) == 0 {
		return nil
	}

	// create items which do not exists yet.
	theList := make([]*db.UpdatePair, len(boardSummaries))
	for idx, each := range boardSummaries {
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
	if r.UpsertedCount == int64(len(boardSummaries)) { // all are created
		return nil
	}

	// update items with comparing update-nano-ts
	upsertedIDs := r.UpsertedIDs
	updateBoardSummaries := make([]*db.UpdatePair, 0, len(theList))
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
		updateBoardSummaries = append(updateBoardSummaries, each)
	}

	_, err = Board_c.BulkUpdateOneOnly(updateBoardSummaries)

	return err
}

func GetBoardSummary(bboardID bbs.BBoardID) (result *BoardSummary, err error) {
	query := &BoardQuery{
		BBoardID: bboardID,
	}

	result = &BoardSummary{}
	err = Board_c.FindOne(query, result, boardSummaryFields)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetBoardSummariesByClsID(clsID ptttype.Bid, startIdx string, isAsc bool, limit int, sortBy ptttype.BSortBy) (boardSummaries []*BoardSummary, err error) {
	idx := ""
	switch sortBy {
	case ptttype.BSORT_BY_NAME:
		idx = BOARD_IDX_BY_NAME_b
	case ptttype.BSORT_BY_CLASS:
		idx = BOARD_IDX_BY_CLASS_b
	}

	var query bson.M
	if startIdx == "" {
		query = bson.M{
			BOARD_GID_b: clsID,
			BOARD_IS_DELETED_b: bson.M{
				"$exists": false,
			},
		}
	} else {
		theDIR := "$lte"
		if isAsc {
			theDIR = "$gte"
		}

		query = bson.M{
			BOARD_GID_b: clsID,
			idx: bson.M{
				theDIR: startIdx,
			},
			BOARD_IS_DELETED_b: bson.M{
				"$exists": false,
			},
		}
	}

	// sort opts
	var sortOpts bson.D
	if isAsc {
		sortOpts = bson.D{
			{Key: idx, Value: 1},
		}
	} else {
		sortOpts = bson.D{
			{Key: idx, Value: -1},
		}
	}

	err = Board_c.Find(query, int64(limit), &boardSummaries, boardSummaryFields, sortOpts)
	if err != nil {
		return nil, err
	}

	return boardSummaries, nil
}
