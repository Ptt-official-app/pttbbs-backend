package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
)

type ArticleIsDeleted struct {
	IsDeleted    bool         `bson:"deleted,omitempty"`
	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}

// DeleteArticles deletes articles in Database
func DeleteArticles(boardID bbs.BBoardID, articleIDs []bbs.ArticleID, updateNanoTS types.NanoTS) (err error) {
	query := bson.M{
		ARTICLE_BBOARD_ID_b:      boardID,
		ARTICLE_ARTICLE_ID_b:     bson.M{"$in": articleIDs},
		ARTICLE_UPDATE_NANO_TS_b: bson.M{"$lt": updateNanoTS},
	}
	update := &ArticleIsDeleted{
		IsDeleted:    true,
		UpdateNanoTS: updateNanoTS,
	}
	_, err = Article_c.UpdateManyOnly(query, update)
	return err
}

// DeleteArticles deletes articles in Database
func DeleteArticlesByBoardID(boardID bbs.BBoardID, updateNanoTS types.NanoTS) (err error) {
	query := bson.M{
		ARTICLE_BBOARD_ID_b:      boardID,
		ARTICLE_UPDATE_NANO_TS_b: bson.M{"$lt": updateNanoTS},
	}
	update := &ArticleIsDeleted{
		IsDeleted:    true,
		UpdateNanoTS: updateNanoTS,
	}
	_, err = Article_c.UpdateManyOnly(query, update)
	return err
}
