package schema

import (
	"errors"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ContentMTime
type ArticleContentMTime struct {
	ContentMTime types.NanoTS `bson:"content_mtime_nano_ts"`

	IsDeleted bool `bson:"deleted,omitempty"` //
}

var EMPTY_ARTICLE_CONTENT_MTIME = &ArticleContentMTime{}

var articleContentMTimeFields = getFields(EMPTY_ARTICLE, EMPTY_ARTICLE_CONTENT_MTIME)

func GetArticleContentMTime(bboardID bbs.BBoardID, articleID bbs.ArticleID) (ret *ArticleContentMTime, err error) {
	query := &ArticleQuery{
		BBoardID:  bboardID,
		ArticleID: articleID,
	}

	ret = &ArticleContentMTime{}
	err = Article_c.FindOne(query, &ret, articleContentMTimeFields)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func UpdateArticleContentMTime(bboardID bbs.BBoardID, articleID bbs.ArticleID, contentMTime types.NanoTS) (err error) {
	query := bson.M{
		"$or": bson.A{
			bson.M{
				ARTICLE_BBOARD_ID_b:  bboardID,
				ARTICLE_ARTICLE_ID_b: articleID,
				ARTICLE_CONTENT_MTIME_b: bson.M{
					"$exists": false,
				},

				ARTICLE_IS_DELETED_b: bson.M{"$exists": false},
			},
			bson.M{
				ARTICLE_BBOARD_ID_b:  bboardID,
				ARTICLE_ARTICLE_ID_b: articleID,
				ARTICLE_CONTENT_MTIME_b: bson.M{
					"$lt": contentMTime,
				},

				ARTICLE_IS_DELETED_b: bson.M{"$exists": false},
			},
		},
	}

	update := &ArticleContentMTime{
		ContentMTime: contentMTime,
	}

	r, err := Article_c.UpdateOneOnly(query, update)
	if err != nil {
		return err
	}
	if r.MatchedCount == 0 {
		return ErrNoMatch
	}

	return err
}
