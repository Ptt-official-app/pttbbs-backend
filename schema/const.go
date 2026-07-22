package schema

import (
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/db"
	redis "github.com/go-redis/redis/v8"
)

const (
	TITLE_REGEX_N_GRAM              = 5
	TIME_CALC_ALL_USER_VISIT_COUNTS = -10 * time.Minute

	MAX_CONTENT_BLOCK = 5

	MAX_ALL_CONTENT_BLOCK = 2000

	STR_REPLY       = "Re:"
	STR_REPLY_LOWER = "re:"

	STR_FORWARD       = "Fw:"
	STR_FORWARD_LOWER = "fw:"

	STR_LEGACY_FORWARD = "[轉錄]"

	MAX_COMMENT_BYTES = 90

	EMAIL_VERIFICATION_TOKEN_LEN = 32

	ARGON2_SALT_LEN = 32
	ARGON2_TIME     = 1
	ARGON2_MEMORY   = 64 * 1024
	ARGON2_THREADS  = 4
	ARGON2_KEYLEN   = 32

	RDB_PREFIX_SEPARATOR                             = ":"
	RDB_PREFIX_2FA                                   = "2fa" + RDB_PREFIX_SEPARATOR
	RDB_PREFIX_EMAIL                                 = "email" + RDB_PREFIX_SEPARATOR
	RDB_PREFIX_LOCK                                  = "lock" + RDB_PREFIX_SEPARATOR
	RDB_PREFIX_OIDC_OP_CODE                          = "opcode" + RDB_PREFIX_SEPARATOR
	RDB_PREFIX_OIDC_OP_REQUEST                       = "opreq" + RDB_PREFIX_SEPARATOR
	RDB_PREFIX_OIDC_OP_REQUEST_CHALLENGE             = "opreqchal" + RDB_PREFIX_SEPARATOR
	RDB_PREFIX_OIDC_OP_REQUEST_IS_AUTH               = "opreqauth" + RDB_PREFIX_SEPARATOR
	RDB_OIDC_OP_REQUEST_TRUE                         = "1"
	RDB_PREFIX_OIDC_OP_CHALLENGE                     = "opchal" + RDB_PREFIX_SEPARATOR
	RDB_PREFIX_OIDC_OP_ACCESS_TOKEN                  = "opaccess" + RDB_PREFIX_SEPARATOR
	RDB_PREFIX_OIDC_OP_ACCESS_TOKEN_SUBJECT_CLIENTID = "opaccesssubcli" + RDB_PREFIX_SEPARATOR
	RDB_PREFIX_OIDC_OP_REFRESH_TOKEN                 = "oprefresh" + RDB_PREFIX_SEPARATOR
)

var (
	client *db.Client

	rdb *redis.Client

	DEFAULT_POST_TYPE = []string{"問題", "建議", "討論", "心得", "閒聊", "請益", "情報", "公告"}
)
