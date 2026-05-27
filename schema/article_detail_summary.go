package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ArticleDetailSummary
type ArticleDetailSummary struct {
	BBoardID       bbs.BBoardID         `bson:"bid"`
	ArticleID      bbs.ArticleID        `bson:"aid"`
	BoardArticleID types.BoardArticleID `bson:"baid"`

	IsDeleted  bool         `bson:"deleted,omitempty"`
	CreateTime types.NanoTS `bson:"create_time_nano_ts"`
	MTime      types.NanoTS `bson:"mtime_nano_ts"`

	Recommend int              `bson:"recommend"`
	Owner     bbs.UUserID      `bson:"owner"`
	Title     string           `bson:"title"`
	Money     int              `bson:"money"`
	Class     string           `bson:"class"`
	Filemode  ptttype.FileMode `bson:"mode"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`

	ContentMTime        types.NanoTS    `bson:"content_mtime_nano_ts"`
	ContentMD5          string          `bson:"content_md5"`
	ContentID           types.ContentID `bson:"content_id"`
	ContentUpdateNanoTS types.NanoTS    `bson:"content_update_nano_ts"`

	FirstCommentsMD5          string       `bson:"first_comments_md5"`
	FirstCommentsLastTime     types.NanoTS `bson:"first_comments_last_time_nano_ts"`
	FirstCommentsUpdateNanoTS types.NanoTS `bson:"first_comments_update_nano_ts"`

	NComments            int          `bson:"n_comments"`
	CommentsUpdateNanoTS types.NanoTS `bson:"comments_update_nano_ts"`

	IP   string `bson:"ip"`   //
	Host string `bson:"host"` // ip 的中文呈現, 外國則為國家.
	BBS  string `bson:"bbs"`  //

	Rank             int          `bson:"rank"` // 評價
	RankUpdateNanoTS types.NanoTS `bson:"rank_update_nano_ts"`

	Idx string `bson:"pttidx"`

	SubjectType ptttype.SubjectType `bson:"subject_type"`
}

var (
	EMPTY_ARTICLE_DETAIL_SUMMARY = &ArticleDetailSummary{}
	articleDetailSummaryFields   = getFields(EMPTY_ARTICLE, EMPTY_ARTICLE_DETAIL_SUMMARY)
)

func GetArticleDetailSummary(bboardID bbs.BBoardID, articleID bbs.ArticleID) (result *ArticleDetailSummary, err error) {
	query := &ArticleQuery{
		BBoardID:  bboardID,
		ArticleID: articleID,
	}

	result = &ArticleDetailSummary{}
	err = Article_c.FindOne(query, &result, articleDetailSummaryFields)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetArticleDetailSummaries(boardID bbs.BBoardID, startIdx string, descending bool, limit int, withDeleted bool) (result []*ArticleDetailSummary, err error) {
	// setup query
	var query bson.M
	if startIdx == "" {
		query = bson.M{
			ARTICLE_BBOARD_ID_b: boardID,
		}
	} else {
		theDir := "$gte"
		if descending {
			theDir = "$lte"
		}
		query = bson.M{
			ARTICLE_BBOARD_ID_b: boardID,
			ARTICLE_IDX_b: bson.M{
				theDir: startIdx,
			},
		}
	}
	if !withDeleted {
		query[ARTICLE_IS_DELETED_b] = bson.M{
			"$exists": false,
		}
	}

	// sort opts
	var sortOpts bson.D
	if descending {
		sortOpts = bson.D{
			{Key: ARTICLE_IDX_b, Value: -1},
			{Key: ARTICLE_ARTICLE_ID_b, Value: -1},
		}
	} else {
		sortOpts = bson.D{
			{Key: ARTICLE_IDX_b, Value: 1},
			{Key: ARTICLE_ARTICLE_ID_b, Value: 1},
		}
	}

	// find
	err = Article_c.Find(query, int64(limit), &result, articleDetailSummaryFields, sortOpts)
	if err != nil {
		return nil, err
	}

	return result, nil
}
