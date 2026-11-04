package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	ptttypes "github.com/Ptt-official-app/go-pttbbs/types"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
)

func LoadGeneralArticlesByKeyword(remoteAddr string, user *UserInfo, params any, path any, c *gin.Context) (result any, statusCode int, err error) {
	theParams, ok := params.(*LoadGeneralArticlesParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*LoadGeneralArticlesPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	userID := user.UserID
	boardID, err := toBoardID(thePath.FBoardID, remoteAddr, userID, c)
	if err != nil {
		return nil, 500, err
	}

	// check board permission
	userBoardPerm, err := CheckUserBoardPermReadable(user, boardID, c)
	if err != nil {
		return nil, 403, err
	}

	// get article-summaries
	keywordList := types.StringsSplitAsRune(theParams.Keyword, " ")

	var createNanoTS types.NanoTS
	var articleID bbs.ArticleID
	var createTSTime4 ptttypes.Time4
	if theParams.StartIdx != "" {
		createTSTime4, articleID, err = bbs.DeserializeArticleIdxStr(theParams.StartIdx)
		if err != nil {
			return nil, 400, err
		}
		createNanoTS = types.Time4ToNanoTS(createTSTime4)
	}

	articleSummaries_db, err := schema.GetArticleSummariesByRegex(boardID, keywordList, createNanoTS, articleID, theParams.Descending, theParams.Max+1)
	if err != nil {
		return nil, 500, err
	}

	articleIDs := make([]bbs.ArticleID, len(articleSummaries_db))
	for idx, each := range articleSummaries_db {
		articleIDs[idx] = each.ArticleID
	}

	// check article permission
	articlePermEditableMap, articlePermDeletableMap, err := CheckUserArticlesPermEditableDeletable(user, boardID, articleIDs, userBoardPerm, c)
	if err != nil {
		return nil, 500, err
	}

	// nextIdx
	var nextIdx string
	if len(articleSummaries_db) > theParams.Max {
		nextIdx = articleSummaries_db[theParams.Max].Idx
		articleSummaries_db = articleSummaries_db[:theParams.Max]
	}

	userReadArticleMap := make(map[bbs.ArticleID]bool)
	userReadArticleMap, err = checkReadArticles(userID, boardID, userReadArticleMap, articleSummaries_db)
	if err != nil {
		return nil, 500, err
	}

	return NewLoadGeneralArticlesResult(articleSummaries_db, userReadArticleMap, articlePermEditableMap, articlePermDeletableMap, nextIdx, userID), 200, nil
}
