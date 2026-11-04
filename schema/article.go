package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
)

var Article_c *db.Collection

type Article struct {
	Version        int                  `bson:"version"`
	BBoardID       bbs.BBoardID         `bson:"bid"` //
	ArticleID      bbs.ArticleID        `bson:"aid"` //
	BoardArticleID types.BoardArticleID `bson:"baid"`

	IsDeleted  bool             `bson:"deleted,omitempty"`   //
	Filename   string           `bson:"filename"`            //
	CreateTime types.NanoTS     `bson:"create_time_nano_ts"` //
	MTime      types.NanoTS     `bson:"mtime_nano_ts"`       //
	Recommend  int              `bson:"recommend"`           //
	Owner      bbs.UUserID      `bson:"owner"`               //
	FullTitle  string           `bson:"full_title"`
	Title      string           `bson:"title"` //
	Money      int              `bson:"money"` //
	Class      string           `bson:"class"` //
	Filemode   ptttype.FileMode `bson:"mode"`  //

	TitleRegex []string `bson:"title_regex"`

	Idx string `bson:"pttidx"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"` // used by article-summary

	ContentMTime  types.NanoTS    `bson:"content_mtime_nano_ts"` //
	ContentMD5    string          `bson:"content_md5"`
	ContentID     types.ContentID `bson:"content_id"`
	Content       [][]*types.Rune `bson:"content"`        //
	ContentPrefix [][]*types.Rune `bson:"content_prefix"` //

	IP                  string       `bson:"ip"`   //
	Host                string       `bson:"host"` // ip 的中文呈現, 外國則為國家.
	BBS                 string       `bson:"bbs"`  //
	ContentUpdateNanoTS types.NanoTS `bson:"content_update_nano_ts"`

	SignatureDBCS []byte `bson:"signature_dbcs"`
	SignatureMD5  string `bson:"signature_md5"`

	FirstCommentsMD5          string       `bson:"first_comments_md5"`
	FirstCommentsLastTime     types.NanoTS `bson:"first_comments_last_time_nano_ts"`
	FirstCommentsUpdateNanoTS types.NanoTS `bson:"first_comments_update_nano_ts"`

	NComments            int          `bson:"n_comments"`
	CommentsUpdateNanoTS types.NanoTS `bson:"comments_update_nano_ts"`

	Rank               int          `bson:"rank"` // 評價
	RankToUpdateNanoTS types.NanoTS `bson:"rank_to_update_nano_ts"`
	RankUpdateNanoTS   types.NanoTS `bson:"rank_update_nano_ts"`

	SubjectType ptttype.SubjectType `bson:"subject_type"`

	IsBottom bool `bson:"is_bottom"`
}

var EMPTY_ARTICLE = &Article{}

var ( // bson-name
	ARTICLE_BBOARD_ID_b        = getBSONName(EMPTY_ARTICLE, "BBoardID")
	ARTICLE_ARTICLE_ID_b       = getBSONName(EMPTY_ARTICLE, "ArticleID")
	ARTICLE_BOARD_ARTICLE_ID_b = getBSONName(EMPTY_ARTICLE, "BoardArticleID")

	ARTICLE_IS_DELETED_b  = getBSONName(EMPTY_ARTICLE, "IsDeleted")
	ARTICLE_FILENAME_b    = getBSONName(EMPTY_ARTICLE, "Filename")
	ARTICLE_CREATE_TIME_b = getBSONName(EMPTY_ARTICLE, "CreateTime")
	ARTICLE_MTIME_b       = getBSONName(EMPTY_ARTICLE, "MTime")
	ARTICLE_RECOMMEND_b   = getBSONName(EMPTY_ARTICLE, "Recommend")
	ARTICLE_OWNER_b       = getBSONName(EMPTY_ARTICLE, "Owner")
	ARTICLE_FULL_TITLE_b  = getBSONName(EMPTY_ARTICLE, "FullTitle")
	ARTICLE_TITLE_b       = getBSONName(EMPTY_ARTICLE, "Title")
	ARTICLE_MONEY_b       = getBSONName(EMPTY_ARTICLE, "Money")
	ARTICLE_CLASS_b       = getBSONName(EMPTY_ARTICLE, "Class")
	ARTICLE_FILEMODE_b    = getBSONName(EMPTY_ARTICLE, "Filemode")
	ARTICLE_TITLE_REGEX_b = getBSONName(EMPTY_ARTICLE, "TitleRegex")

	ARTICLE_IDX_b = getBSONName(EMPTY_ARTICLE, "Idx")

	ARTICLE_UPDATE_NANO_TS_b = getBSONName(EMPTY_ARTICLE, "UpdateNanoTS")

	ARTICLE_CONTENT_MTIME_b  = getBSONName(EMPTY_ARTICLE, "ContentMTime")
	ARTICLE_CONTENT_MD5_B    = getBSONName(EMPTY_ARTICLE, "ContentMD5")
	ARTICLE_CONTENT_b        = getBSONName(EMPTY_ARTICLE, "Content")
	ARTICLE_CONTENT_PREFIX_b = getBSONName(EMPTY_ARTICLE, "ContentPrefix")

	ARTICLE_IP_b   = getBSONName(EMPTY_ARTICLE, "IP")
	ARTICLE_HOST_b = getBSONName(EMPTY_ARTICLE, "Host")
	ARTICLE_BBS_b  = getBSONName(EMPTY_ARTICLE, "BBS")

	ARTICLE_CONTENT_UPDATE_NANO_TS_b = getBSONName(EMPTY_ARTICLE, "ContentUpdateNanoTS")

	ARTICLE_SIGNATURE_DBCS_b = getBSONName(EMPTY_ARTICLE, "SignatureDBCS")
	ARTICLE_SIGNATURE_MD5_b  = getBSONName(EMPTY_ARTICLE, "SignatureMD5")

	ARTICLE_FIRST_COMMENTS_MD5_b       = getBSONName(EMPTY_ARTICLE, "FirstCommentsMD5")
	ARTICLE_FIRST_COMMENTS_LAST_TIME_b = getBSONName(EMPTY_ARTICLE, "FirstCommentsLastTime")

	ARTICLE_FIRST_COMMENTS_UPDATE_NANO_TS_b = getBSONName(EMPTY_ARTICLE, "FirstCommentsUpdateNanoTS")

	ARTICLE_N_COMMENTS_b              = getBSONName(EMPTY_ARTICLE, "NComments")
	ARTICLE_COMMENTS_UPDATE_NANO_TS_b = getBSONName(EMPTY_ARTICLE, "CommentsUpdateNanoTS")

	ARTICLE_RANK_b                   = getBSONName(EMPTY_ARTICLE, "Rank")
	ARTICLE_RANK_TO_UPDATE_NANO_TS_b = getBSONName(EMPTY_ARTICLE, "RankToUpdateNanoTS")
	ARTICLE_RANK_UPDATE_NANO_TS_b    = getBSONName(EMPTY_ARTICLE, "RankUpdateNanoTS")

	ARTICLE_SUBJECT_TYPE_b = getBSONName(EMPTY_ARTICLE, "SubjectType")

	ARTICLE_IS_BOTTOM_b = getBSONName(EMPTY_ARTICLE, "IsBottom")
)

func assertArticleFields() error {
	if err := assertFields(EMPTY_ARTICLE, EMPTY_ARTICLE_QUERY); err != nil {
		return err
	}

	// article_content_info
	if err := assertFields(EMPTY_ARTICLE, EMPTY_ARTICLE_CONTENT_INFO); err != nil {
		return err
	}

	// article_content_mtime
	if err := assertFields(EMPTY_ARTICLE, EMPTY_ARTICLE_CONTENT_MTIME); err != nil {
		return err
	}

	// article_detail_summary
	if err := assertFields(EMPTY_ARTICLE, EMPTY_ARTICLE_DETAIL_SUMMARY); err != nil {
		return err
	}

	// article_first_comments
	if err := assertFields(EMPTY_ARTICLE, EMPTY_ARTICLE_FIRST_COMMENTS); err != nil {
		return err
	}

	// article_summary
	if err := assertFields(EMPTY_ARTICLE, EMPTY_ARTICLE_SUMMARY); err != nil {
		return err
	}

	// article_summary_with_regex
	if err := assertFields(EMPTY_ARTICLE, EMPTY_ARTICLE_SUMMARY_WITH_REGEX); err != nil {
		return err
	}

	// article-n-comments
	if err := assertFields(EMPTY_ARTICLE, EMPTY_ARTICLE_N_COMMENTS); err != nil {
		return err
	}

	// article-rank

	// article-perm-info
	if err := assertFields(EMPTY_ARTICLE, EMPTY_ARTICLE_PERM_INFO); err != nil {
		return err
	}

	return nil
}

type ArticleQuery struct {
	BBoardID  bbs.BBoardID  `bson:"bid"`
	ArticleID bbs.ArticleID `bson:"aid"`
	IsDeleted any           `bson:"deleted,omitempty"` //
}

var EMPTY_ARTICLE_QUERY = &ArticleQuery{}

type ArticleCreateQuery struct {
	BBoardID       bbs.BBoardID         `bson:"bid"`
	ArticleID      bbs.ArticleID        `bson:"aid"`
	BoardArticleID types.BoardArticleID `bson:"baid"`

	IsDeleted any `bson:"deleted,omitempty"` //
}

var EMPTY_ARTICLE_CREATE_QUERY = &ArticleCreateQuery{}

func ResetArticleIsBottom(boardID bbs.BBoardID) (err error) {
	query := bson.M{
		ARTICLE_BBOARD_ID_b: boardID,
		ARTICLE_IS_BOTTOM_b: true,
	}
	update := bson.M{
		"$unset": bson.M{
			ARTICLE_IS_BOTTOM_b: true,
		},
	}

	_, err = Article_c.UpdateManyOnlyNoSet(query, update)

	return err
}
