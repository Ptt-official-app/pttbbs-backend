package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
)

const GET_MAN_ARTICLE_R = "/board/:bid/manual/*aid"

type GetManArticleDetailPath struct {
	FBoardID   apitypes.FBoardID   `uri:"bid"`
	FArticleID apitypes.FArticleID `uri:"aid"`
}

type GetManArticleDetailResult struct {
	BBoardID   apitypes.FBoardID   `json:"bid"` //
	ArticleID  apitypes.FArticleID `json:"aid"` //
	LevelIdx   apitypes.FArticleID `json:"level_idx"`
	IsDeleted  bool                `json:"deleted"`     //
	CreateTime types.Time8         `json:"create_time"` //
	MTime      types.Time8         `json:"modified"`    //
	Title      string              `json:"title"`       //
	IsDir      bool                `json:"is_dir"`

	Content [][]*types.Rune `json:"content"`

	TokenUser bbs.UUserID `json:"tokenuser"`
}

func GetManArticleDetailWrapper(c *gin.Context) {
	path := &GetManArticleDetailPath{}
	LoginRequiredPathQuery(GetManArticleDetail, nil, path, c)
}

func GetManArticleDetail(remoteAddr string, user *UserInfo, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	thePath, ok := path.(*GetManArticleDetailPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	userID := user.UserID
	boardID, err := toBoardID(thePath.FBoardID, remoteAddr, userID, c)
	if err != nil {
		return nil, 500, err
	}

	farticleID := thePath.FArticleID
	if farticleID[0] == '/' {
		farticleID = farticleID[1:]
	}
	articleID := farticleID.ToManArticleID()

	// validate user
	_, statusCode, err = isBoardValidUser(boardID, userID)
	if err != nil {
		return nil, statusCode, err
	}

	// ensure that we do have the article.
	content, _, articleDetailSummary, err := TryGetManArticleContentInfo(userID, boardID, articleID, c, false, true)
	if err != nil {
		return nil, statusCode, err
	}

	result = &GetManArticleDetailResult{
		BBoardID:   thePath.FBoardID,
		ArticleID:  farticleID,
		LevelIdx:   apitypes.ToFArticleIDFromManArticleID(articleDetailSummary.LevelIdx),
		IsDeleted:  articleDetailSummary.IsDeleted,
		CreateTime: articleDetailSummary.CreateTime.ToTime8(),
		MTime:      articleDetailSummary.CreateTime.ToTime8(),
		Title:      articleDetailSummary.Title,
		IsDir:      articleDetailSummary.IsDir,

		Content: content,

		TokenUser: userID,
	}

	return result, 200, nil
}
