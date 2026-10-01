package apitypes

import (
	"strconv"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	pttbbsfav "github.com/Ptt-official-app/go-pttbbs/ptt/fav"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

type BoardSummary struct {
	FBoardID     FBoardID        `json:"bid"`
	Brdname      string          `json:"brdname"`
	Title        string          `json:"title"`
	BrdAttr      ptttype.BrdAttr `json:"flag"`
	BoardType    string          `json:"type"`
	Category     string          `json:"class"`
	NUser        int             `json:"nuser"`
	BMs          []bbs.UUserID   `json:"moderators"`
	Reason       string          `json:"reason"`
	Read         bool            `json:"read"`
	Fav          bool            `json:"fav"`
	Total        int             `json:"total"`
	LastPostTime types.Time8     `json:"last_post_time"`

	StatAttr ptttype.BoardStatAttr `json:"stat_attr,omitempty"`
	LevelIdx schema.LevelIdx       `json:"level_idx,omitempty"` // sub-level-idx for folder, "" if type is line or board.

	URL string `json:"url,omitempty"`

	Gid ptttype.Bid `json:"gid"`
	Bid ptttype.Bid `json:"pttbid"`

	Idx string `json:"idx,omitempty"`

	TokenUser bbs.UUserID `json:"tokenuser,omitempty"`

	IsPopular bool `json:"is_popular,omitempty"` // 熱門板
	IsOver18  bool `json:"is_over_18"`           // 18歲板

	IsNoStats             bool `json:"is_no_stats,omitempty"`               // 不列入統計
	IsGroupBoard          bool `json:"is_group_board,omitempty"`            // 群組板
	IsHide                bool `json:"is_hide,omitempty"`                   // 隱板
	IsPostMask            bool `json:"is_post_mask,omitempty"`              // 限制發表或是閱讀
	IsAnonymous           bool `json:"is_anony,omitempty"`                  // 匿名板
	IsDefaultAnonymous    bool `json:"is_default_anony,omitempty"`          // 預設匿名板
	IsNoCredit            bool `json:"is_no_credit,omitempty"`              // 發文無獎勵板
	IsVoteBoard           bool `json:"is_vote_board,omitempty"`             // 連署機看板
	IsWarnEOL             bool `json:"is_warn_eol,omitempty"`               // 已警告要廢除
	IsNoComment           bool `json:"is_no_comment,omitempty"`             // 不可推文
	IsAngelAnonymous      bool `json:"is_angel_anony,omitempty"`            // 小天使可匿名
	IsSymLink             bool `json:"is_sym_link,omitempty"`               // 文章是 sym-link (AllPost/HiddenAllPost)
	IsNoBoo               bool `json:"is_no_boo,omitempty"`                 // 不可噓文
	IsBoardMemberOnlyPost bool `json:"is_board_member_only_post,omitempty"` // 板友才可以發文
	IsGuestPost           bool `json:"is_guest_post,omitempty"`             // guest 可發文
	IsCooldown            bool `json:"is_cooldown,omitempty"`               // 靜
	IsIPLogComment        bool `json:"is_ip_log_comment,omitempty"`         // 推文記錄 IP
	IsNoReply             bool `json:"is_no_reply,omitempty"`               // 不可回文
	IsAlignedComment      bool `json:"is_aligned_comment,omitempty"`        // 對齊式推文
	IsNoSelfDelPost       bool `json:"is_no_self_del_post,omitempty"`       // 不可以自己刪文
	IsBMMaskContent       bool `json:"is_bm_mask_content,omitempty"`        // 允許板主刪除特定文字
}

func NewBoardSummary(b_db *schema.BoardSummary, idx string, userBoardInfo *UserBoardInfo, userID bbs.UUserID) *BoardSummary {
	if b_db == nil {
		return nil
	}

	fboardID := ToFBoardID(b_db.BBoardID)

	url := ""
	if b_db.BrdAttr.HasPerm(ptttype.BRD_GROUPBOARD) {
		// XXX api.LOAD_CLASS_BOARDS_R
		bidStr := strconv.Itoa(int(b_db.Bid))
		url = "/cls/" + bidStr
	} else {
		// XXX api.LOAD_GENERAL_ARTICLES_R
		url = "/board/" + string(fboardID) + "/articles"
	}

	bms := b_db.BMs
	if bms == nil {
		bms = []bbs.UUserID{}
	}
	return &BoardSummary{
		FBoardID:     fboardID,
		Brdname:      b_db.Brdname,
		Title:        b_db.Title,
		BrdAttr:      b_db.BrdAttr,
		BoardType:    b_db.BoardType,
		Category:     b_db.Category,
		BMs:          bms,
		Total:        b_db.Total,
		LastPostTime: b_db.LastPostTime.ToTime8(),
		NUser:        b_db.NUser,

		Gid: b_db.Gid,
		Bid: b_db.Bid,

		Idx: idx,

		StatAttr: userBoardInfo.Stat,
		Read:     userBoardInfo.Read,
		Fav:      userBoardInfo.Fav,

		URL: url,

		TokenUser: userID,

		IsPopular: b_db.IsPopular,
		IsOver18:  b_db.IsOver18,

		IsNoStats:             b_db.IsNoStats,
		IsGroupBoard:          b_db.IsGroupBoard,
		IsHide:                b_db.IsHide,
		IsPostMask:            b_db.IsPostMask,
		IsAnonymous:           b_db.IsAnonymous,
		IsDefaultAnonymous:    b_db.IsDefaultAnonymous,
		IsNoCredit:            b_db.IsNoCredit,
		IsVoteBoard:           b_db.IsVoteBoard,
		IsWarnEOL:             b_db.IsWarnEOL,
		IsNoComment:           b_db.IsNoComment,
		IsAngelAnonymous:      b_db.IsAngelAnonymous,
		IsSymLink:             b_db.IsSymLink,
		IsNoBoo:               b_db.IsNoBoo,
		IsBoardMemberOnlyPost: b_db.IsBoardMemberOnlyPost,
		IsGuestPost:           b_db.IsGuestPost,
		IsCooldown:            b_db.IsCooldown,
		IsIPLogComment:        b_db.IsIPLogComment,
		IsNoReply:             b_db.IsNoReply,
		IsAlignedComment:      b_db.IsAlignedComment,
		IsNoSelfDelPost:       b_db.IsNoSelfDelPost,
		IsBMMaskContent:       b_db.IsBMMaskContent,
	}
}

func NewBoardSummaryFromUserFavorites(userID bbs.UUserID, uf_db *schema.UserFavorites, b_db *schema.BoardSummary, userBoardInfo *UserBoardInfo) *BoardSummary {
	idxStr := strconv.Itoa(uf_db.Idx)

	switch uf_db.TheType {
	case pttbbsfav.FAVT_LINE:
		return &BoardSummary{
			StatAttr: ptttype.NBRD_LINE,
			Idx:      idxStr,
			BMs:      []bbs.UUserID{},

			TokenUser: userID,
		}
	case pttbbsfav.FAVT_FOLDER:
		// XXX api.LOAD_FAVORITE_BOARDS_R
		subLevelIdx := schema.SetLevelIdx(uf_db.LevelIdx, uf_db.Idx)
		url := "/user/" + string(userID) + "/favorites?level_idx=" + string(subLevelIdx)
		return &BoardSummary{
			Title: uf_db.FolderTitle,

			StatAttr: ptttype.NBRD_FOLDER,
			LevelIdx: subLevelIdx,
			Idx:      idxStr,
			URL:      url,
			BMs:      []bbs.UUserID{},

			TokenUser: userID,
		}

	case pttbbsfav.FAVT_BOARD:
		return NewBoardSummary(b_db, idxStr, userBoardInfo, userID)
	}

	return nil
}
