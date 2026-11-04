package schema

import (
	"errors"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ManArticleContentInfo struct {
	ContentMD5          string          `bson:"content_md5"`
	ContentID           types.ContentID `bson:"content_id"`
	Content             [][]*types.Rune `bson:"content"`
	ContentUpdateNanoTS types.NanoTS    `bson:"content_update_nano_ts"`
}

var (
	EMPTY_MAN_ARTICLE_CONTENT_INFO = &ManArticleContentInfo{}
	manArticleContentInfoFields    = getFields(EMPTY_MAN_ARTICLE, EMPTY_MAN_ARTICLE_CONTENT_INFO)
)

func GetManArticleContentInfo(bboardID bbs.BBoardID, articleID types.ManArticleID, isContent bool) (contentInfo *ManArticleContentInfo, err error) {
	query := bson.M{
		MAN_ARTICLE_BBOARD_ID_b:  bboardID,
		MAN_ARTICLE_ARTICLE_ID_b: articleID,
	}

	contentInfo = &ManArticleContentInfo{}
	err = ManArticle_c.FindOne(query, &contentInfo, manArticleContentInfoFields)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if contentInfo.ContentID == "" || !isContent {
		return contentInfo, nil
	}

	contentBlocks, err := GetAllManContentBlocks(bboardID, articleID, contentInfo.ContentID)
	if err != nil {
		return nil, err
	}

	content := contentBlocksToContent(contentBlocks)
	contentInfo.Content = content

	return contentInfo, nil
}

func UpdateManArticleContentInfo(bboardID bbs.BBoardID, articleID types.ManArticleID, contentInfo *ManArticleContentInfo) (err error) {
	query := bson.M{
		MAN_ARTICLE_BBOARD_ID_b:  bboardID,
		MAN_ARTICLE_ARTICLE_ID_b: articleID,
	}

	r, err := ManArticle_c.CreateOnly(query, contentInfo)
	if err != nil {
		return err
	}
	if r.UpsertedCount > 0 {
		return nil
	}

	query = bson.M{
		"$or": bson.A{
			bson.M{
				MAN_ARTICLE_BBOARD_ID_b:  bboardID,
				MAN_ARTICLE_ARTICLE_ID_b: articleID,
				MAN_ARTICLE_CONTENT_UPDATE_NANO_TS_b: bson.M{
					"$exists": false,
				},

				MAN_ARTICLE_IS_DELETED_b: bson.M{"$exists": false},
			},
			bson.M{
				MAN_ARTICLE_BBOARD_ID_b:  bboardID,
				MAN_ARTICLE_ARTICLE_ID_b: articleID,
				MAN_ARTICLE_CONTENT_UPDATE_NANO_TS_b: bson.M{
					"$lt": contentInfo.ContentUpdateNanoTS,
				},

				MAN_ARTICLE_IS_DELETED_b: bson.M{"$exists": false},
			},
		},
	}

	r, err = ManArticle_c.UpdateOneOnly(query, contentInfo)
	if err != nil {
		return err
	}
	if r.MatchedCount == 0 {
		return ErrNoMatch
	}
	return nil
}
