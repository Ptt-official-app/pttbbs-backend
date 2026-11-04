package schema

import (
	"errors"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ArticlePermInfo struct {
	BBoardID       bbs.BBoardID         `bson:"bid"` //
	ArticleID      bbs.ArticleID        `bson:"aid"` //
	BoardArticleID types.BoardArticleID `bson:"baid"`

	Owner bbs.UUserID `bson:"owner"` //

	IsDeleted bool `bson:"deleted,omitempty"` //
}

var (
	EMPTY_ARTICLE_PERM_INFO = &ArticlePermInfo{}
	articlePermInfoFields   = getFields(EMPTY_ARTICLE, EMPTY_ARTICLE_PERM_INFO)
)

func GetArticlePermInfo(boardID bbs.BBoardID, articleID bbs.ArticleID) (articlePermInfo *ArticlePermInfo, err error) {
	query := bson.M{
		ARTICLE_BBOARD_ID_b:  boardID,
		ARTICLE_ARTICLE_ID_b: articleID,
	}

	err = Article_c.FindOne(query, &articlePermInfo, articlePermInfoFields)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return articlePermInfo, nil
}

func GetArticlesPermInfo(boardID bbs.BBoardID, articleIDs []bbs.ArticleID) (articlesPermInfo []*ArticlePermInfo, err error) {
	query := bson.M{
		ARTICLE_BBOARD_ID_b: boardID,
		ARTICLE_ARTICLE_ID_b: bson.M{
			"$in": articleIDs,
		},
	}

	err = Article_c.Find(query, 0, &articlesPermInfo, articlePermInfoFields, nil)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return articlesPermInfo, nil
}
