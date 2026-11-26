package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
)

type UserReadArticle struct {
	UserID           bbs.UUserID          `bson:"user_id"`
	BoardID          bbs.BBoardID         `bson:"bid"`
	ArticleID        bbs.ArticleID        `bson:"aid"`
	BoardArticleID   types.BoardArticleID `bson:"baid"`
	ReadUpdateNanoTS types.NanoTS         `bson:"update_nano_ts"`
}

var EMPTY_USER_READ_ARTICLE = &UserReadArticle{}

type UserReadArticleQuery struct {
	UserID    bbs.UUserID   `bson:"user_id"`
	BoardID   bbs.BBoardID  `bson:"bid"`
	ArticleID bbs.ArticleID `bson:"aid"`
}

var EMPTY_USER_READ_ARTICLE_QUERY = &UserReadArticleQuery{}

func UpdateUserReadArticle(userReadArticle *UserReadArticle) (err error) {
	// ensure board-article-id
	userReadArticle.BoardArticleID = types.ToBoardArticleID(userReadArticle.BoardID, userReadArticle.ArticleID)

	query := bson.M{
		USER_ARTICLE_USER_ID_b:    userReadArticle.UserID,
		USER_ARTICLE_BOARD_ID_b:   userReadArticle.BoardID,
		USER_ARTICLE_ARTICLE_ID_b: userReadArticle.ArticleID,
	}

	r, err := UserArticle_c.CreateOnly(query, userReadArticle)
	if err != nil {
		return err
	}
	if r.UpsertedCount == 1 { // created, no need to update.
		return nil
	}

	query[USER_ARTICLE_READ_UPDATE_NANO_TS_b] = bson.M{
		"$lt": userReadArticle.ReadUpdateNanoTS,
	}

	_, err = UserArticle_c.UpdateOneOnly(query, userReadArticle)

	return err
}

func UpdateUserReadArticles(userReadArticles []*UserReadArticle, updateNanoTS types.NanoTS) (err error) {
	if len(userReadArticles) == 0 {
		return nil
	}

	// ensure board-article-id
	for _, each := range userReadArticles {
		each.BoardArticleID = types.ToBoardArticleID(each.BoardID, each.ArticleID)
	}

	theList := make([]*db.UpdatePair, len(userReadArticles))
	for idx, each := range userReadArticles {
		query := &UserReadArticleQuery{
			UserID:    each.UserID,
			BoardID:   each.BoardID,
			ArticleID: each.ArticleID,
		}

		theList[idx] = &db.UpdatePair{
			Filter: query,
			Update: each,
		}
	}

	r, err := UserArticle_c.BulkCreateOnly(theList)
	if err != nil {
		return err
	}
	if r.UpsertedCount == int64(len(userReadArticles)) { // all are created
		return nil
	}

	upsertedIDs := r.UpsertedIDs
	updateUserReadArticles := make([]*db.UpdatePair, 0, len(userReadArticles))
	for idx, each := range theList {
		_, ok := upsertedIDs[int64(idx)]
		if ok {
			continue
		}

		origFilter := each.Filter.(*UserReadArticleQuery)
		filter := bson.M{
			USER_ARTICLE_USER_ID_b:    origFilter.UserID,
			USER_ARTICLE_BOARD_ID_b:   origFilter.BoardID,
			USER_ARTICLE_ARTICLE_ID_b: origFilter.ArticleID,
			USER_ARTICLE_READ_UPDATE_NANO_TS_b: bson.M{
				"$lt": updateNanoTS,
			},
		}
		each.Filter = filter
		updateUserReadArticles = append(updateUserReadArticles, each)
	}

	_, err = UserArticle_c.BulkUpdateOneOnly(updateUserReadArticles)

	return err
}

func FindUserReadArticles(userID bbs.UUserID, boardID bbs.BBoardID, articleIDs []bbs.ArticleID) ([]*UserReadArticle, error) {
	// query
	query := bson.M{
		USER_ARTICLE_USER_ID_b:  userID,
		USER_ARTICLE_BOARD_ID_b: boardID,
		USER_ARTICLE_ARTICLE_ID_b: bson.M{
			"$in": articleIDs,
		},
	}

	var dbResults []*UserReadArticle
	err := UserArticle_c.Find(query, 0, &dbResults, nil, nil)
	if err != nil {
		return nil, err
	}

	return dbResults, nil
}

func FindUserReadArticlesByBoardArticleIDs(userID bbs.UUserID, boardArticleIDs []types.BoardArticleID) ([]*UserReadArticle, error) {
	// query

	query := bson.M{
		USER_ARTICLE_USER_ID_b: userID,
		USER_ARTICLE_BOARD_ARTICLE_ID_b: bson.M{
			"$in": boardArticleIDs,
		},
	}

	var dbResults []*UserReadArticle
	err := UserArticle_c.Find(query, 0, &dbResults, nil, nil)
	if err != nil {
		return nil, err
	}

	return dbResults, nil
}
