package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

var Board_c *db.Collection

type Board struct {
	Version   int             `bson:"version"`
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

var EMPTY_BOARD = &Board{}

var (
	BOARD_BBOARD_ID_b      = getBSONName(EMPTY_BOARD, "BBoardID")
	BOARD_BRDNAME_b        = getBSONName(EMPTY_BOARD, "Brdname")
	BOARD_TITLE_b          = getBSONName(EMPTY_BOARD, "Title")
	BOARD_BRD_ATTR_b       = getBSONName(EMPTY_BOARD, "BrdAttr")
	BOARD_BOARD_TYPE_b     = getBSONName(EMPTY_BOARD, "BoardType")
	BOARD_CATEGORY_b       = getBSONName(EMPTY_BOARD, "Category")
	BOARD_NUSER_b          = getBSONName(EMPTY_BOARD, "NUser")
	BOARD_BMS_b            = getBSONName(EMPTY_BOARD, "BMs")
	BOARD_TOTAL_b          = getBSONName(EMPTY_BOARD, "Total")
	BOARD_LAST_POST_TIME_b = getBSONName(EMPTY_BOARD, "LastPostTime")
	BOARD_UDPATE_TIME_b    = getBSONName(EMPTY_BOARD, "UpdateTime")

	BOARD_VOTE_LIMIT_LOGINS_b  = getBSONName(EMPTY_BOARD, "VoteLimitLogins")
	BOARD_POST_LIMIT_LOGINS_b  = getBSONName(EMPTY_BOARD, "PostLimitLogins")
	BOARD_VOTE_LIMIT_BADPOST_b = getBSONName(EMPTY_BOARD, "VoteLimitBadpost")
	BOARD_POST_LIMIT_BADPOST_b = getBSONName(EMPTY_BOARD, "PostLimitBadpost")

	BOARD_PARENT_b = getBSONName(EMPTY_BOARD, "Parent")

	BOARD_NVOTE_b             = getBSONName(EMPTY_BOARD, "NVote")
	BOARD_VOTE_CLOSING_TIME_b = getBSONName(EMPTY_BOARD, "VoteClosingTime")

	BOARD_LEVEL_b                = getBSONName(EMPTY_BOARD, "Level")
	BOARD_LAST_SET_TIME_b        = getBSONName(EMPTY_BOARD, "LastSetTime")
	BOARD_POST_EXPIRE_b          = getBSONName(EMPTY_BOARD, "PostExpire")
	BOARD_END_GAMBLE_NANO_TS_b   = getBSONName(EMPTY_BOARD, "EndGambleNanoTS")
	BOARD_POST_TYPE_b            = getBSONName(EMPTY_BOARD, "PostType")
	BOARD_POST_TEMPLATE_b        = getBSONName(EMPTY_BOARD, "PostTemplate")
	BOARD_FAST_RECOMMEND_PAUSE_b = getBSONName(EMPTY_BOARD, "FastRecommendPause")
	BOARD_IS_DELETED_b           = getBSONName(EMPTY_BOARD, "IsDeleted")
	BOARD_UPDATE_NANO_TS_b       = getBSONName(EMPTY_BOARD, "UpdateNanoTS")

	BOARD_GID_b          = getBSONName(EMPTY_BOARD, "Gid")
	BOARD_BID_b          = getBSONName(EMPTY_BOARD, "Bid")
	BOARD_IDX_BY_NAME_b  = getBSONName(EMPTY_BOARD, "IdxByName")
	BOARD_IDX_BY_CLASS_b = getBSONName(EMPTY_BOARD, "IdxByClass")

	BOARD_CHESS_COUNTRY_b = getBSONName(EMPTY_BOARD, "ChessCountry")

	BOARD_IS_POPULAR_b = getBSONName(EMPTY_BOARD, "IsPopular")
	BOARD_IS_OVER18_b  = getBSONName(EMPTY_BOARD, "IsOver18")

	BOARD_IS_NO_STATS_b               = getBSONName(EMPTY_BOARD, "IsNoStats")
	BOARD_IS_GROUP_BOARD_b            = getBSONName(EMPTY_BOARD, "IsGroupBoard")
	BOARD_IS_HIDE_b                   = getBSONName(EMPTY_BOARD, "IsHide")
	BOARD_IS_POST_MASK_b              = getBSONName(EMPTY_BOARD, "IsPostMask")
	BOARD_IS_ANONYMOUS_b              = getBSONName(EMPTY_BOARD, "IsAnonymous")
	BOARD_IS_DEFAULT_ANONYMOUS_b      = getBSONName(EMPTY_BOARD, "IsDefaultAnonymous")
	BOARD_IS_NO_CREDIT_b              = getBSONName(EMPTY_BOARD, "IsNoCredit")
	BOARD_IS_VOTE_BOARD_b             = getBSONName(EMPTY_BOARD, "IsVoteBoard")
	BOARD_IS_WARN_EOL_b               = getBSONName(EMPTY_BOARD, "IsWarnEOL")
	BOARD_IS_NO_COMMENT_b             = getBSONName(EMPTY_BOARD, "IsNoComment")
	BOARD_IS_ANGEL_ANONYMOUS_b        = getBSONName(EMPTY_BOARD, "IsAngelAnonymous")
	BOARD_IS_SYM_LINK_b               = getBSONName(EMPTY_BOARD, "IsSymLink")
	BOARD_IS_NO_BOO_b                 = getBSONName(EMPTY_BOARD, "IsNoBoo")
	BOARD_IS_BOARD_MEMBER_ONLY_POST_b = getBSONName(EMPTY_BOARD, "IsBoardMemberOnlyPost")
	BOARD_IS_GUEST_POST_b             = getBSONName(EMPTY_BOARD, "IsGuestPost")
	BOARD_IS_COOLDOWN_b               = getBSONName(EMPTY_BOARD, "IsCooldown")
	BOARD_IS_IP_LOG_COMMENT_b         = getBSONName(EMPTY_BOARD, "IsIPLogComment")
	BOARD_IS_NO_REPLY_b               = getBSONName(EMPTY_BOARD, "IsNoReply")
	BOARD_IS_ALIGNED_COMMENT_b        = getBSONName(EMPTY_BOARD, "IsAlignedComment")
	BOARD_IS_NO_SELF_DEL_POST_b       = getBSONName(EMPTY_BOARD, "IsNoSelfDelPost")
	BOARD_IS_BM_MASK_CONTENT_b        = getBSONName(EMPTY_BOARD, "IsBMMaskContent")
)

func assertBoardFields() error {
	if err := assertFields(EMPTY_BOARD, EMPTY_BOARD_QUERY); err != nil {
		return err
	}

	if err := assertFields(EMPTY_BOARD, EMPTY_BOARD_SUMMARY); err != nil {
		return err
	}

	if err := assertFields(EMPTY_BOARD, EMPTY_BOARD_DETAIL); err != nil {
		return err
	}

	if err := assertFields(EMPTY_BOARD, EMPTY_BOARD_ID); err != nil {
		return err
	}

	if err := assertFields(EMPTY_BOARD, EMPTY_BOARD_BID); err != nil {
		return err
	}

	if err := assertFields(EMPTY_BOARD, EMPTY_BOARD_PERM_INFO); err != nil {
		return err
	}

	if err := assertFields(EMPTY_BOARD, EMPTY_BOARD_IS_POPULAR); err != nil {
		return err
	}

	return nil
}

type BoardQuery struct {
	BBoardID bbs.BBoardID `bson:"bid"`

	IsDeleted interface{} `bson:"deleted,omitempty"`
}

var EMPTY_BOARD_QUERY = &BoardQuery{}
