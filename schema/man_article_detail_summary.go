package schema

import (
	"errors"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ManArticleDetailSummary struct {
	BBoardID   bbs.BBoardID       `bson:"bid"`                 //
	ArticleID  types.ManArticleID `bson:"aid"`                 //
	IsDeleted  bool               `bson:"deleted,omitempty"`   //
	LevelIdx   types.ManArticleID `bson:"level_idx"`           //
	CreateTime types.NanoTS       `bson:"create_time_nano_ts"` //
	MTime      types.NanoTS       `bson:"mtime_nano_ts"`

	Title string `bson:"title"` //
	IsDir bool   `bson:"is_dir"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"` // used by article-summary

	ContentMTime types.NanoTS    `bson:"content_mtime_nano_ts"` //
	ContentMD5   string          `bson:"content_md5"`
	ContentID    types.ContentID `bson:"content_id"`

	ContentUpdateNanoTS types.NanoTS `bson:"content_update_nano_ts"`

	Idx int `bson:"pttidx"`
}

var (
	EMPTY_MAN_ARTICLE_DETAIL_SUMMARY = &ManArticleDetailSummary{}
	manArticleDetailSummaryFields    = getFields(EMPTY_MAN_ARTICLE, EMPTY_MAN_ARTICLE_DETAIL_SUMMARY)
)

func GetManArticleDetailSummary(bboardID bbs.BBoardID, articleID types.ManArticleID) (result *ManArticleDetailSummary, err error) {
	query := bson.M{
		MAN_ARTICLE_BBOARD_ID_b:  bboardID,
		MAN_ARTICLE_ARTICLE_ID_b: articleID,
	}

	result = &ManArticleDetailSummary{}
	err = ManArticle_c.FindOne(query, &result, manArticleDetailSummaryFields)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetManArticleDetailSummaries(boardID bbs.BBoardID, levelIdx types.ManArticleID) (result []*ManArticleDetailSummary, err error) {
	// setup query
	query := bson.M{
		MAN_ARTICLE_BBOARD_ID_b: boardID,
		MAN_ARTICLE_LEVEL_IDX_b: levelIdx,
	}

	// sort opts
	sortOpts := bson.D{
		{Key: MAN_ARTICLE_IDX_b, Value: 1},
	}

	// find
	err = ManArticle_c.Find(query, int64(0), &result, manArticleDetailSummaryFields, sortOpts)
	if err != nil {
		return nil, err
	}

	return result, nil
}
