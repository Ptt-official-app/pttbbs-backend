package schema

import (
	"errors"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserPermInfo struct {
	UserID bbs.UUserID `bson:"user_id"`

	Userlevel ptttype.PERM `bson:"perm"`

	Over18 bool `bson:"over18"`

	Numlogindays int `bson:"login_days"`

	BadPost int `bson:"bad_post"` /* 評價為壞文章數 */

	CooldownNanoTS types.NanoTS `bson:"cooldown_nano_ts"`
	Posttime       int          `bson:"postitme"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}

var (
	EMPTY_USER_PERM_INFO = &UserPermInfo{}
	userPermInfoFields   = getFields(EMPTY_USER, EMPTY_USER_PERM_INFO)
)

func GetUserPermInfo(userID bbs.UUserID) (userPermInfo *UserPermInfo, err error) {
	// XXX for guest
	if userID == bbs.UUserID(pttbbsapi.GUEST) {
		return &UserPermInfo{
			UserID: userID,

			Over18: true,
		}, nil
	}

	query := bson.M{
		USER_USER_ID_b: userID,
	}

	err = User_c.FindOne(query, &userPermInfo, userPermInfoFields)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return userPermInfo, nil
}
