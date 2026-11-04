package schema

import (
	"errors"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ManArticleContentMTime struct {
	ContentMTime types.NanoTS `bson:"content_mtime_nano_ts"`

	IsDeleted bool `bson:"deleted,omitempty"` //
}

var EMPTY_MAN_ARTICLE_CONTENT_MTIME = &ManArticleContentMTime{}

var manArticleContentMTimeFields = getFields(EMPTY_MAN_ARTICLE, EMPTY_MAN_ARTICLE_CONTENT_MTIME)

func GetManArticleContentMTime(bboardID bbs.BBoardID, articleID types.ManArticleID) (ret *ManArticleContentMTime, err error) {
	query := bson.M{
		MAN_ARTICLE_BBOARD_ID_b:  bboardID,
		MAN_ARTICLE_ARTICLE_ID_b: articleID,
	}

	ret = &ManArticleContentMTime{}
	err = ManArticle_c.FindOne(query, &ret, manArticleContentMTimeFields)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func UpdateManArticleContentMTime(bboardID bbs.BBoardID, articleID types.ManArticleID, contentMTime types.NanoTS) (err error) {
	query := bson.M{
		"$or": bson.A{
			bson.M{
				MAN_ARTICLE_BBOARD_ID_b:  bboardID,
				MAN_ARTICLE_ARTICLE_ID_b: articleID,
				MAN_ARTICLE_CONTENT_MTIME_b: bson.M{
					"$exists": false,
				},
			},
			bson.M{
				MAN_ARTICLE_BBOARD_ID_b:  bboardID,
				MAN_ARTICLE_ARTICLE_ID_b: articleID,
				MAN_ARTICLE_CONTENT_MTIME_b: bson.M{
					"$lt": contentMTime,
				},
			},
		},
	}

	update := &ManArticleContentMTime{
		ContentMTime: contentMTime,
	}

	r, err := ManArticle_c.UpdateOneOnly(query, update)
	if err != nil {
		return err
	}
	if r.MatchedCount == 0 {
		return ErrNoMatch
	}

	return err
}
