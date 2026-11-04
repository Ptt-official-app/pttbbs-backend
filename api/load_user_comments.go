package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
)

const LOAD_USER_COMMENTS_R = "/user/:user_id/comments"

type LoadUserCommentsParams struct {
	StartIdx   string `json:"start_idx,omitempty" form:"start_idx,omitempty" url:"start_idx,omitempty"`
	Descending bool   `json:"desc,omitzero"  form:"desc,omitempty" url:"desc,omitempty"`
	Max        int    `json:"limit,omitzero" form:"limit,omitempty" url:"limit,omitempty"`
}

type LoadUserCommentsPath struct {
	UserID bbs.UUserID `uri:"user_id"`
}

type LoadUserCommentsResult struct {
	List    []*apitypes.ArticleComment `json:"list"`
	NextIdx string                     `json:"next_idx"`

	TokenUser bbs.UUserID `json:"tokenuser,omitempty"`
}

func NewLoadUserCommentsParams() *LoadUserCommentsParams {
	return &LoadUserCommentsParams{
		Descending: DEFAULT_DESCENDING,
		Max:        DEFAULT_MAX_LIST,
	}
}

func LoadUserCommentsWrapper(c *gin.Context) {
	params := NewLoadUserCommentsParams()
	path := &LoadUserCommentsPath{}
	LoginRequiredPathQuery(LoadUserComments, params, path, c)
}

func LoadUserComments(remoteAddr string, user *UserInfo, params any, path any, c *gin.Context) (result any, statusCode int, err error) {
	theParams, ok := params.(*LoadUserCommentsParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*LoadUserCommentsPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	commentSummaries_db, nextIdx, err := loadUserComments(thePath.UserID, theParams.StartIdx, theParams.Descending, theParams.Max)
	if err != nil {
		return nil, 500, err
	}

	userID := user.UserID
	articleSummaryMap, userReadBoardArticleMap, err := getArticleSummaryMapFromCommentSummaries(userID, commentSummaries_db)
	if err != nil {
		return nil, 500, err
	}

	r := NewLoadUserCommentsResult(commentSummaries_db, articleSummaryMap, userReadBoardArticleMap, nextIdx, userID)

	return r, 200, nil
}

func loadUserComments(ownerID bbs.UUserID, startIdx string, descending bool, theMax int) (commentSummaries_db []*schema.CommentSummary, nextIdx string, err error) {
	commentSummaries_db = make([]*schema.CommentSummary, 0, theMax+1)

	var nextSortTime types.NanoTS
	if startIdx != "" {
		_, startIdx, err = apitypes.DeserializeArticleCommentIdx(startIdx)
		if err != nil {
			return nil, "", err
		}
		nextSortTime, _ = apitypes.DeserializeCommentIdx(startIdx)
	}

	isEndLoop := false
	remaining := theMax
	for !isEndLoop && remaining > 0 {
		eachCommentSummaries_db, err := schema.GetBasicCommentSummariesByOwnerID(ownerID, nextSortTime, descending, theMax+1)
		if err != nil {
			return nil, "", err
		}

		// check is-last query
		if len(eachCommentSummaries_db) < theMax+1 {
			isEndLoop = true
			nextSortTime = 0
		} else {
			// setup next
			nextCommentSummary := eachCommentSummaries_db[len(eachCommentSummaries_db)-1]
			eachCommentSummaries_db = eachCommentSummaries_db[:len(eachCommentSummaries_db)-1]

			nextSortTime = nextCommentSummary.SortTime
			nextIdx = apitypes.SerializeCommentIdx(nextSortTime, nextCommentSummary.CommentID)
		}
		if len(eachCommentSummaries_db) == 0 {
			break
		}

		// is-valid
		validCommentSummaries_db, err := isValidCommentSummaries(eachCommentSummaries_db)
		if err != nil {
			return nil, "", err
		}

		// append
		if len(validCommentSummaries_db) > remaining {
			nextCommentSummary := validCommentSummaries_db[remaining]
			validCommentSummaries_db = validCommentSummaries_db[:remaining]

			nextSortTime = nextCommentSummary.SortTime
			nextIdx = apitypes.SerializeCommentIdx(nextSortTime, nextCommentSummary.CommentID)
		}

		commentSummaries_db = append(commentSummaries_db, validCommentSummaries_db...)
		remaining -= len(validCommentSummaries_db)
	}

	return commentSummaries_db, nextIdx, nil
}

// isValidCommentSummaries
// XXX TODO
func isValidCommentSummaries(commentSummaries_db []*schema.CommentSummary) ([]*schema.CommentSummary, error) {
	return commentSummaries_db, nil
}

func getArticleSummaryMapFromCommentSummaries(userID bbs.UUserID, commentSummaries_db []*schema.CommentSummary) (articleSummaryMap map[types.BoardArticleID]*schema.ArticleSummary, userReadBoardArticleMap map[types.BoardArticleID]types.NanoTS, err error) {
	boardArticleIDs := make([]types.BoardArticleID, 0, len(commentSummaries_db))
	boardArticleIDMap := make(map[types.BoardArticleID]bool)

	var eachBoardArticleID types.BoardArticleID
	for _, each := range commentSummaries_db {
		eachBoardArticleID = types.ToBoardArticleID(each.BBoardID, each.ArticleID)
		_, ok := boardArticleIDMap[eachBoardArticleID]
		if ok {
			continue
		}

		boardArticleIDMap[eachBoardArticleID] = true
		boardArticleIDs = append(boardArticleIDs, eachBoardArticleID)
	}

	// article summaries
	articleSummaries, err := schema.GetArticleSummariesByBoardArticleIDs(boardArticleIDs)
	if err != nil {
		return nil, nil, err
	}

	articleSummaryMap = make(map[types.BoardArticleID]*schema.ArticleSummary)
	for _, each := range articleSummaries {
		articleSummaryMap[each.BoardArticleID] = each
	}

	// user read articles
	userReadArticles, err := schema.FindUserReadArticlesByBoardArticleIDs(userID, boardArticleIDs)
	if err != nil {
		return nil, nil, err
	}

	userReadBoardArticleMap = make(map[types.BoardArticleID]types.NanoTS)
	for _, each := range userReadArticles {
		userReadBoardArticleMap[each.BoardArticleID] = each.ReadUpdateNanoTS
	}

	return articleSummaryMap, userReadBoardArticleMap, nil
}

func NewLoadUserCommentsResult(
	commentSummaries_db []*schema.CommentSummary,
	articleSummaryMap map[types.BoardArticleID]*schema.ArticleSummary,
	userReadBoardArticleMap map[types.BoardArticleID]types.NanoTS,
	nextIdx string,
	userID bbs.UUserID,
) (result *LoadUserCommentsResult) {
	comments := make([]*apitypes.ArticleComment, 0, len(commentSummaries_db))
	for _, each := range commentSummaries_db {
		boardArticleID := types.ToBoardArticleID(each.BBoardID, each.ArticleID)
		articleSummary, ok := articleSummaryMap[boardArticleID]
		if !ok {
			continue
		}

		eachComment := apitypes.NewArticleCommentFromComment(articleSummary, each)

		// read
		readNanoTS, ok := userReadBoardArticleMap[boardArticleID]
		if !ok {
			comments = append(comments, eachComment)
			continue
		}

		if readNanoTS > articleSummary.MTime {
			eachComment.Read = types.READ_STATUS_MTIME
		} else if readNanoTS > each.SortTime {
			eachComment.Read = types.READ_STATUS_COMMENT_TIME
		} else if readNanoTS > each.CreateTime {
			eachComment.Read = types.READ_STATUS_CREATE_TIME
		} else {
			eachComment.Read = types.READ_STATUS_UNREAD
		}

		comments = append(comments, eachComment)
	}

	nextIdx = apitypes.SerializeArticleCommentIdx(apitypes.ARTICLE_COMMENT_TYPE_COMMENT, nextIdx)

	result = &LoadUserCommentsResult{
		List:    comments,
		NextIdx: nextIdx,

		TokenUser: userID,
	}
	return result
}
