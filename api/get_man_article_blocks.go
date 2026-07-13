package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
)

const GET_MAN_ARTICLE_BLOCKS_R = "/board/:bid/manualblocks/*aid"

type GetManArticleBlocksParams struct {
	StartIdx string `json:"start_idx,omitempty" form:"start_idx,omitempty" url:"start_idx,omitempty"`
	Max      int    `json:"limit,omitempty" form:"limit,omitempty" url:"limit,omitempty"`
}

type GetManArticleBlocksPath struct {
	FBoardID   apitypes.FBoardID   `uri:"bid"`
	FArticleID apitypes.FArticleID `uri:"aid"`
}

type GetManArticleBlocksResult struct {
	Content [][]*types.Rune `json:"content,omitempty"`

	IsDeleted  bool        `json:"deleted,omitempty"`     //
	CreateTime types.Time8 `json:"create_time,omitempty"` //
	MTime      types.Time8 `json:"modified,omitempty"`    //
	Title      string      `json:"title,omitempty"`       //

	NextIdx string `json:"next_idx"`

	TokenUser bbs.UUserID `json:"tokenuser"`
}

func NewGetManArticleBlocksParams() *GetManArticleBlocksParams {
	return &GetManArticleBlocksParams{
		Max: DEFAULT_MAX_ARTICLE_BLOCK_LIST,
	}
}

func GetManArticleBlocksWrapper(c *gin.Context) {
	params := NewGetManArticleBlocksParams()
	path := &GetManArticleBlocksPath{}
	LoginRequiredPathQuery(GetManArticleBlocks, params, path, c)
}

func GetManArticleBlocks(remoteAddr string, user *UserInfo, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*GetManArticleBlocksParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*GetManArticleBlocksPath)
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

	if theParams.StartIdx == "" {
		_, _, _, err = TryGetManArticleContentInfo(userID, boardID, articleID, c, false, false)
		if err != nil {
			return nil, 500, err
		}
	}

	contentID, contentIdx, articleDetailSummary, err := getManArticleBlocksContentBlocksDeserializeContentIdx(boardID, articleID, theParams.StartIdx)
	if err != nil {
		return nil, 400, err
	}

	contentBlocks, err := schema.GetManContentBlocks(boardID, articleID, contentID, contentIdx, int64(theParams.Max+1))
	if err != nil {
		return nil, 500, err
	}

	nextIdx := ""
	if len(contentBlocks) > theParams.Max {
		nextBlock := contentBlocks[theParams.Max]
		contentBlocks = contentBlocks[:theParams.Max]
		nextIdx = apitypes.SerializeContentIdx(contentID, nextBlock.Idx)
	}

	content := getArticleBlocksContentBlocksToContent(contentBlocks)

	ret := &GetManArticleBlocksResult{
		Content: content,
		NextIdx: nextIdx,

		TokenUser: userID,
	}
	if articleDetailSummary != nil {
		ret.IsDeleted = articleDetailSummary.IsDeleted
		ret.CreateTime = articleDetailSummary.CreateTime.ToTime8()
		ret.MTime = articleDetailSummary.MTime.ToTime8()
		ret.Title = articleDetailSummary.Title
	}

	return ret, 200, nil
}

func getManArticleBlocksContentBlocksDeserializeContentIdx(boardID bbs.BBoardID, articleID types.ManArticleID, theIdx string) (contentID types.ContentID, contentIdx int, articleDetailSummary *schema.ManArticleDetailSummary, err error) {
	if theIdx != "" {
		contentID, contentIdx, err = apitypes.DeserializeContentIdx(theIdx)
		return contentID, contentIdx, nil, err
	}

	articleDetailSummary, err = schema.GetManArticleDetailSummary(boardID, articleID)
	if err != nil {
		return "", 0, nil, err
	}

	return articleDetailSummary.ContentID, 0, articleDetailSummary, nil
}
