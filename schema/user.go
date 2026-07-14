package schema

import (
	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

var User_c *db.Collection

type User struct {
	Version    int          `bson:"version"`
	UserID     bbs.UUserID  `bson:"user_id"`
	Username   string       `bson:"username"`
	UsernameNS types.NanoTS `bson:"username_ns"`

	FamilyName string `bson:"familyname"`
	GivenName  string `bson:"givenname"`
	Locale     string `bson:"locale"`

	Realname string `bson:"realname"`
	Nickname string `bson:"nickname"`

	IsGovernmentID   bool         `bson:"is_government_id"`
	IsGovernmentIDNS types.NanoTS `bson:"is_government_id_ns"`

	Uflag        ptttype.UFlag `bson:"flag"`
	Userlevel    ptttype.PERM  `bson:"perm"`
	Numlogindays int           `bson:"login_days"` /* 考慮透過 db-count, 但是可能無法跟以前版本相容 */
	Numposts     int           `bson:"posts"`      /* 考慮透過 db-count, 因為 post 可能會被消失, 但是這個欄位不應該因為被消失的而減少? (poster 主動消失的需要減少)? */
	Firstlogin   types.NanoTS  `bson:"first_login_nano_ts"`
	Lastlogin    types.NanoTS  `bson:"last_login_nano_ts"` /* 考慮透過 db-max, 但是可能在拉 user-detail 時會花很多時間. */
	LastIP       string        `bson:"last_ip"`
	LastHost     string        `bson:"last_host"` // last-ip 的中文呈現, 外國則為國家.

	Money    int    `bson:"money"`
	PttEmail string `bson:"pttemail"`
	Justify  string `bson:"justify"`
	Over18   bool   `bson:"over18"`

	PagerUIType uint8             `bson:"pager_ui"` /* 呼叫器界面類別 (was: WATER_*) */
	Pager       ptttype.PagerMode `bson:"pager"`    /* 呼叫器狀態 */
	Invisible   bool              `bson:"hide"`
	Exmailbox   uint32            `bson:"exmail"`

	Career        string       `bson:"career"`
	Role          uint32       `bson:"role"`
	LastSeen      types.NanoTS `bson:"last_seen_nano_ts"`
	TimeSetAngel  types.NanoTS `bson:"time_set_angel_nano_ts"`
	TimePlayAngel types.NanoTS `bson:"time_play_angel_nano_ts"`

	LastSong  types.NanoTS `bson:"last_song_nano_ts"`
	LoginView uint32       `bson:"login_view"`

	Vlcount   int   `bson:"violation"`
	FiveWin   int   `bson:"five_win"`
	FiveLose  int   `bson:"five_lose"`
	FiveTie   int   `bson:"five_tie"`
	ChcWin    int   `bson:"chc_win"`
	ChcLose   int   `bson:"chc_lose"`
	ChcTie    int   `bson:"chc_tie"`
	Conn6Win  int   `bson:"conn6_win"`
	Conn6Lose int   `bson:"conn6_lose"`
	Conn6Tie  int   `bson:"conn6_tie"`
	GoWin     int   `bson:"go_win"`
	GoLose    int   `bson:"go_lose"`
	GoTie     int   `bson:"go_tie"`
	DarkWin   int   `bson:"dark_win"`
	DarkLose  int   `bson:"dark_lose"`
	DarkTie   int   `bson:"dark_tie"` /* 暗棋戰績 和 */
	UaVersion uint8 `bson:"ua_version"`

	Signature uint8       `bson:"signaure"` /* 慣用簽名檔 */
	BadPost   int         `bson:"bad_post"` /* 評價為壞文章數 */
	MyAngel   bbs.UUserID `bson:"angel"`    /* 我的小天使 */

	ChessEloRating int `bson:"chess_rank"` /* 象棋等級 */

	TimeRemoveBadPost types.NanoTS `bson:"time_remove_bad_post_nano_ts"`
	TimeViolateLaw    types.NanoTS `bson:"time_violate_law_nano_ts"`

	IsDeleted    bool         `bson:"deleted,omitempty"`
	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`

	UserLevel2    ptttype.PERM2 `bson:"perm2"`
	UpdateNanoTS2 types.NanoTS  `bson:"update_nano_ts2"`

	// NFriend int `bson:"n_friend"` /* 需要透過 db-count */
	Avatar       []byte       `bson:"avatar"`
	AvatarNanoTS types.NanoTS `bson:"avatar_nano_ts"`

	CooldownNanoTS       types.NanoTS `bson:"cooldown_nano_ts"`
	CooldownUpdateNanoTS types.NanoTS `bson:"cooldown_update_nano_ts"`

	Posttime             int          `bson:"postitme"`
	PosttimeUpdateNanoTS types.NanoTS `bson:"posttime_update_nano_ts"`
}

var EMPTY_USER = &User{}

var (
	USER_USER_ID_b     = getBSONName(EMPTY_USER, "UserID")
	USER_USERNAME_b    = getBSONName(EMPTY_USER, "Username")
	USER_USERNAME_NS_b = getBSONName(EMPTY_USER, "UsernameNS")

	USER_REALNAME_b = getBSONName(EMPTY_USER, "Realname")
	USER_NICKNAME_b = getBSONName(EMPTY_USER, "Nickname")

	USER_IS_GOVERNMENT_ID_b    = getBSONName(EMPTY_USER, "IsGovernmentID")
	USER_IS_GOVERNMENT_ID_NS_b = getBSONName(EMPTY_USER, "IsGovernmentIDNS")

	USER_UFLAG_b           = getBSONName(EMPTY_USER, "Uflag")
	USER_USER_LEVEL_b      = getBSONName(EMPTY_USER, "Userlevel")
	USER_NUMLOGINDAYS_b    = getBSONName(EMPTY_USER, "Numlogindays")
	USER_NUMPOSTS_b        = getBSONName(EMPTY_USER, "Numposts")
	USER_FIRSTLOGIN_b      = getBSONName(EMPTY_USER, "Firstlogin")
	USER_LASTLOGIN_b       = getBSONName(EMPTY_USER, "Lastlogin")
	USER_LAST_IP_b         = getBSONName(EMPTY_USER, "LastIP")
	USER_LAST_HOST_b       = getBSONName(EMPTY_USER, "LastHost")
	USER_MONEY_b           = getBSONName(EMPTY_USER, "Money")
	USER_PTTEMAIL_b        = getBSONName(EMPTY_USER, "PttEmail")
	USER_JUSTIFY_b         = getBSONName(EMPTY_USER, "Justify")
	USER_OVER18_b          = getBSONName(EMPTY_USER, "Over18")
	USER_PAGER_UI_TYPE_b   = getBSONName(EMPTY_USER, "PagerUIType")
	USER_PAGER_b           = getBSONName(EMPTY_USER, "Pager")
	USER_INVISIBLE_b       = getBSONName(EMPTY_USER, "Invisible")
	USER_EXMAILBOX_b       = getBSONName(EMPTY_USER, "Exmailbox")
	USER_CAREER_b          = getBSONName(EMPTY_USER, "Career")
	USER_ROLE_b            = getBSONName(EMPTY_USER, "Role")
	USER_LAST_SEEN_b       = getBSONName(EMPTY_USER, "LastSeen")
	USER_TIME_SET_ANGEL_b  = getBSONName(EMPTY_USER, "TimeSetAngel")
	USER_TIME_PLAY_ANGEL_b = getBSONName(EMPTY_USER, "TimePlayAngel")
	USER_LAST_SONG_b       = getBSONName(EMPTY_USER, "LastSong")
	USER_LOGIN_VIEW_b      = getBSONName(EMPTY_USER, "LoginView")

	USER_VLCOUNT_b    = getBSONName(EMPTY_USER, "Vlcount")
	USER_FIVE_WIN_b   = getBSONName(EMPTY_USER, "FiveWin")
	USER_FIVE_LOSE_b  = getBSONName(EMPTY_USER, "FiveLose")
	USER_FIVE_TIE_b   = getBSONName(EMPTY_USER, "FiveTie")
	USER_CHC_WIN_b    = getBSONName(EMPTY_USER, "ChcWin")
	USER_CHC_LOSE_b   = getBSONName(EMPTY_USER, "ChcLose")
	USER_CHC_TIE_b    = getBSONName(EMPTY_USER, "ChcTie")
	USER_CONN6_WIN_b  = getBSONName(EMPTY_USER, "Conn6Win")
	USER_CONN6_LOSE_b = getBSONName(EMPTY_USER, "Conn6Lose")
	USER_CONN6_TIE_b  = getBSONName(EMPTY_USER, "Conn6Tie")
	USER_GO_WIN_b     = getBSONName(EMPTY_USER, "GoWin")
	USER_GO_LOSE_b    = getBSONName(EMPTY_USER, "GoLose")
	USER_GO_TIE_b     = getBSONName(EMPTY_USER, "GoTie")
	USER_DARK_WIN_b   = getBSONName(EMPTY_USER, "DarkWin")
	USER_DARK_LOSE_b  = getBSONName(EMPTY_USER, "DarkLose")
	USER_DARK_TIE_b   = getBSONName(EMPTY_USER, "DarkTie")
	USER_UA_VERSION_b = getBSONName(EMPTY_USER, "UaVersion")

	USER_SIGNATURE_b = getBSONName(EMPTY_USER, "Signature")
	USER_BAD_POST_b  = getBSONName(EMPTY_USER, "BadPost")
	USER_MY_ANGEL_b  = getBSONName(EMPTY_USER, "MyAngel")

	USER_CHESS_ELO_RATING_b = getBSONName(EMPTY_USER, "ChessEloRating")

	USER_TIME_REMOVE_BAD_POST_b = getBSONName(EMPTY_USER, "TimeRemoveBadPost")
	USER_TIME_VIOLATE_LAW_b     = getBSONName(EMPTY_USER, "TimeViolateLaw")

	USER_IS_DELETED_b     = getBSONName(EMPTY_USER, "IsDeleted")
	USER_UPDATE_NANO_TS_b = getBSONName(EMPTY_USER, "UpdateNanoTS")

	USER_USER_LEVEL2_b     = getBSONName(EMPTY_USER, "UserLevel2")
	USER_UPDATE_NANO_TS2_b = getBSONName(EMPTY_USER, "UpdateNanoTS2")

	USER_AVATAR_b         = getBSONName(EMPTY_USER, "Avatar")
	USER_AVATAR_NANO_TS_b = getBSONName(EMPTY_USER, "AvatorNanoTS")

	USER_EMAIL_b         = getBSONName(EMPTY_USER, "Email")
	USER_EMAIL_NANO_TS_b = getBSONName(EMPTY_USER, "EmailNanoTS")

	USER_TWO_FACTOR_ENABLED_b         = getBSONName(EMPTY_USER, "TwoFactorEnabled")
	USER_TWO_FACTOR_ENABLED_NANO_TS_b = getBSONName(EMPTY_USER, "TwoFactorEnabledNanoTS")

	USER_COOLDOWN_NANO_TS_b        = getBSONName(EMPTY_USER, "CooldownNanoTS")
	USER_COOLDOWN_UPDATE_NANO_TS_b = getBSONName(EMPTY_USER, "CooldownUpdateNanoTS")

	USER_POSTTIME_b                = getBSONName(EMPTY_USER, "Posttime")
	USER_POSTTIME_UPDATE_NANO_TS_b = getBSONName(EMPTY_USER, "PosttimeUpdateNanoTS")
)

func assertUserFields() error {
	if err := assertFields(EMPTY_USER, EMPTY_USER_QUERY); err != nil {
		return err
	}

	if err := assertFields(EMPTY_USER, EMPTY_USER_INFO_SUMMARY); err != nil {
		return err
	}

	if err := assertFields(EMPTY_USER, EMPTY_USER_DETAIL); err != nil {
		return err
	}

	if err := assertFields(EMPTY_USER, EMPTY_USER_NEW_INFO); err != nil {
		return err
	}

	if err := assertFields(EMPTY_USER, EMPTY_USER_NAME_INFO); err != nil {
		return err
	}

	if err := assertFields(EMPTY_USER, EMPTY_USER_PERM_INFO); err != nil {
		return err
	}

	if err := assertFields(EMPTY_USER, EMPTY_USER_POSTTIME); err != nil {
		return err
	}

	if err := assertFields(EMPTY_USER, EMPTY_USER_USERNAME); err != nil {
		return err
	}

	if err := assertFields(EMPTY_USER, EMPTY_USER_IS_GOVERNMENT_ID); err != nil {
		return err
	}

	if err := assertFields(EMPTY_USER, EMPTY_USER_IS_GOVERNMENT_ID_OVER18); err != nil {
		return err
	}

	if err := assertFields(EMPTY_USER, EMPTY_USER_PROFILE); err != nil {
		return err
	}

	return nil
}

type UserQuery struct {
	UserID    bbs.UUserID `bson:"user_id"`
	IsDeleted interface{} `bson:"deleted,omitempty"`
}

var EMPTY_USER_QUERY = &UserQuery{}
