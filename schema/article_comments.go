package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
)

type ArticleComments struct {
	BBoardID  bbs.BBoardID  `bson:"bid"` //
	ArticleID bbs.ArticleID `bson:"aid"` //

	NComments int `bson:"n_comments"`

	CommentsUpdateNanoTS types.NanoTS `bson:"comments_update_nano_ts"`

	IsDeleted bool `bson:"deleted,omitempty"` //
}

var (
	EMPTY_ARTICLE_COMMENTS = &ArticleComments{}
	articleCommentsFields  = getFields(EMPTY_ARTICLE, EMPTY_ARTICLE_COMMENTS) //nolint // consistent with programming pattern
)

func UpdateArticleCommentsByArticleID(boardID bbs.BBoardID, articleID bbs.ArticleID, updateNanoTS types.NanoTS) (err error) {
	nComments, err := CountComments(boardID, articleID)
	if err != nil {
		// do not persist a bogus n_comments with a newer timestamp.
		return err
	}

	articleComments := &ArticleComments{
		BBoardID:             boardID,
		ArticleID:            articleID,
		NComments:            nComments,
		CommentsUpdateNanoTS: updateNanoTS,
	}

	return UpdateArticleComments(articleComments)
}

// UpdateArticleComments
//
// We've already have article-content-info, we don't do create UpdateArticleComments
func UpdateArticleComments(articleComments *ArticleComments) (err error) {
	query := bson.M{
		"$or": bson.A{
			bson.M{
				ARTICLE_BBOARD_ID_b:  articleComments.BBoardID,
				ARTICLE_ARTICLE_ID_b: articleComments.ArticleID,
				ARTICLE_COMMENTS_UPDATE_NANO_TS_b: bson.M{
					"$exists": false,
				},

				ARTICLE_IS_DELETED_b: bson.M{"$exists": false},
			},
			bson.M{
				ARTICLE_BBOARD_ID_b:  articleComments.BBoardID,
				ARTICLE_ARTICLE_ID_b: articleComments.ArticleID,
				ARTICLE_COMMENTS_UPDATE_NANO_TS_b: bson.M{
					"$lt": articleComments.CommentsUpdateNanoTS,
				},

				ARTICLE_IS_DELETED_b: bson.M{"$exists": false},
			},
		},
	}

	r, err := Article_c.UpdateOneOnly(query, articleComments)
	if err != nil {
		return err
	}
	if r.MatchedCount == 0 {
		return ErrNoMatch
	}

	return nil
}
