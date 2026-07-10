package schema

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
)

type UserDetail struct {
	UserID   bbs.UUserID `bson:"user_id"`
	Username string      `bson:"username"`
	Realname string      `bson:"realname"`
	Nickname string      `bson:"nickname"`

	IsGovernmentID bool `bson:"is_government_id"`

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
}

var (
	EMPTY_USER_DETAIL = &UserDetail{}
	userDetailFields  = getFields(EMPTY_USER, EMPTY_USER_DETAIL) //nolint // consistent with programming pattern
)

func NewUserDetailGuest(updateNanoTS types.NanoTS) (user *UserDetail) {
	return &UserDetail{
		UserID:   bbs.UUserID(pttbbsapi.GUEST),
		Username: pttbbsapi.GUEST,
		Realname: pttbbsapi.GUEST,
		Nickname: pttbbsapi.GUEST,

		Userlevel: ptttype.PERM_BASIC,

		Justify: pttbbsapi.GUEST,
		Over18:  true,

		UpdateNanoTS: updateNanoTS,
	}
}

func GetUserDetailByUsername(username string) (userDetail *UserDetail, err error) {
	query := bson.M{
		USER_USERNAME_b: username,
	}

	userDetail = &UserDetail{}
	err = User_c.FindOne(query, userDetail, userDetailFields)
	if err != nil {
		return nil, err
	}

	return userDetail, nil
}

func NewUserDetail(user_b pttbbsapi.GetUserResult, updateNanoTS types.NanoTS) (user *UserDetail) {
	return &UserDetail{
		UserID:   user_b.UUserID,
		Username: user_b.Username,
		Realname: types.Big5ToUtf8(user_b.Realname),
		Nickname: types.Big5ToUtf8(user_b.Nickname),

		Uflag:        user_b.Uflag,
		Userlevel:    user_b.Userlevel,
		Numlogindays: int(user_b.Numlogindays),
		Numposts:     int(user_b.Numposts),
		Firstlogin:   types.Time4ToNanoTS(user_b.Firstlogin),
		Lastlogin:    types.Time4ToNanoTS(user_b.Lastlogin),
		LastIP:       user_b.Lasthost,
		LastHost:     user_b.Lasthost,

		Money:    int(user_b.Money),
		PttEmail: user_b.Email,
		Justify:  types.Big5ToUtf8(user_b.Justify),
		Over18:   user_b.Over18,

		PagerUIType: user_b.PagerUIType,
		Pager:       user_b.Pager,
		Invisible:   user_b.Invisible,
		Exmailbox:   user_b.Exmailbox,

		Career:        types.Big5ToUtf8(user_b.Career),
		Role:          user_b.Role,
		LastSeen:      types.Time4ToNanoTS(user_b.LastSeen),
		TimeSetAngel:  types.Time4ToNanoTS(user_b.TimeSetAngel),
		TimePlayAngel: types.Time4ToNanoTS(user_b.TimePlayAngel),

		LastSong:  types.Time4ToNanoTS(user_b.LastSong),
		LoginView: user_b.LoginView,

		Vlcount:   int(user_b.Vlcount),
		FiveWin:   int(user_b.FiveWin),
		FiveLose:  int(user_b.FiveLose),
		FiveTie:   int(user_b.FiveTie),
		ChcWin:    int(user_b.ChcWin),
		ChcLose:   int(user_b.ChcLose),
		ChcTie:    int(user_b.ChcTie),
		Conn6Win:  int(user_b.Conn6Win),
		Conn6Lose: int(user_b.Conn6Lose),
		Conn6Tie:  int(user_b.Conn6Tie),
		GoWin:     int(user_b.GoWin),
		GoLose:    int(user_b.GoLose),
		GoTie:     int(user_b.GoTie),
		DarkWin:   int(user_b.DarkWin),
		DarkLose:  int(user_b.DarkLose),
		DarkTie:   int(user_b.DarkTie),
		UaVersion: user_b.UaVersion,

		Signature: user_b.Signature,
		BadPost:   int(user_b.BadPost),
		MyAngel:   user_b.MyAngel,

		ChessEloRating: int(user_b.ChessEloRating),

		TimeRemoveBadPost: types.Time4ToNanoTS(user_b.TimeRemoveBadPost),
		TimeViolateLaw:    types.Time4ToNanoTS(user_b.TimeViolateLaw),

		UpdateNanoTS: updateNanoTS,

		UserLevel2:    user_b.UserLevel2,
		UpdateNanoTS2: types.Time4ToNanoTS(user_b.UpdateTS2),
	}
}

func UpdateUserDetail(user *UserDetail) (err error) {
	userID := user.UserID
	query := bson.M{
		USER_USER_ID_b: userID,
	}

	r, err := User_c.CreateOnly(query, user)
	if err != nil {
		return err
	}
	if r.UpsertedCount > 0 {
		return nil
	}

	query = bson.M{
		"$or": bson.A{
			bson.M{
				USER_USER_ID_b: userID,
				USER_UPDATE_NANO_TS_b: bson.M{
					"$exists": false,
				},

				USER_IS_DELETED_b: bson.M{"$exists": false},
			},
			bson.M{
				USER_USER_ID_b: userID,
				USER_UPDATE_NANO_TS_b: bson.M{
					"$lt": user.UpdateNanoTS,
				},
				USER_IS_DELETED_b: bson.M{"$exists": false},
			},
		},
	}

	r, err = User_c.UpdateOneOnly(query, user)
	if err != nil {
		return err
	}
	if r.MatchedCount == 0 {
		return ErrNoMatch
	}

	return nil
}
