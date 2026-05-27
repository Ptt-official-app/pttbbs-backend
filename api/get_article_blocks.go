package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
)

const GET_ARTICLE_BLOCKS_R = "/board/:bid/articleblocks/:aid"

type GetArticleBlocksParams struct {
	StartIdx string `json:"start_idx,omitempty" form:"start_idx,omitempty" url:"start_idx,omitempty"`
	Max      int    `json:"limit,omitempty" form:"limit,omitempty" url:"limit,omitempty"`
}

type GetArticleBlocksPath struct {
	FBoardID   apitypes.FBoardID   `uri:"bid"`
	FArticleID apitypes.FArticleID `uri:"aid"`
}

type GetArticleBlocksResult struct {
	Content [][]*types.Rune `json:"content,omitempty"`

	IsDeleted   bool                `json:"deleted,omitempty"`     //
	CreateTime  types.Time8         `json:"create_time,omitempty"` //
	MTime       types.Time8         `json:"modified,omitempty"`    //
	Recommend   int                 `json:"recommend,omitempty"`   //
	NComments   int                 `json:"n_comments,omitempty"`  //
	Owner       bbs.UUserID         `json:"owner,omitempty"`       //
	Nickname    string              `json:"nickname,omitempty"`
	Title       string              `json:"title,omitempty"` //
	Money       int                 `json:"money,omitempty"` //
	Class       string              `json:"class,omitempty"` // can be: R: 轉, [class]
	Filemode    ptttype.FileMode    `json:"mode,omitempty"`  //
	SubjectType ptttype.SubjectType `json:"subject_type"`

	IP   string `json:"ip,omitempty"`
	Host string `json:"host,omitempty"` // ip 的中文呈現, 外國則為國家.
	BBS  string `json:"bbs,omitempty"`

	Rank int `json:"rank,omitempty"` // 評價

	NextIdx string `json:"next_idx"`

	TokenUser bbs.UUserID `json:"tokenuser"`
}

func NewGetArticleBlocksParams() *GetArticleBlocksParams {
	return &GetArticleBlocksParams{
		Max: DEFAULT_MAX_ARTICLE_BLOCK_LIST,
	}
}

func GetArticleBlocksWrapper(c *gin.Context) {
	params := NewGetArticleBlocksParams()
	path := &GetArticleBlocksPath{}
	LoginRequiredPathQuery(GetArticleBlocks, params, path, c)
}

func GetArticleBlocks(remoteAddr string, user *UserInfo, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*GetArticleBlocksParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*GetArticleBlocksPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	userID := user.UserID
	boardID, err := toBoardID(thePath.FBoardID, remoteAddr, userID, c)
	if err != nil {
		return nil, 500, err
	}
	articleID := thePath.FArticleID.ToArticleID()

	// check permission
	err = CheckUserArticlePermReadable(user, boardID, articleID, true, c)
	if err != nil {
		return nil, 403, err
	}

	if theParams.StartIdx == "" {
		_, _, _, _, _, _, _, _, _, _, _, statusCode, err = TryGetArticleContentInfo(userID, boardID, articleID, c, false, false, false)
		if err != nil {
			return nil, statusCode, err
		}
	}

	contentID, contentIdx, articleDetailSummary, err := getArticleBlocksContentBlocksDeserializeContentIdx(boardID, articleID, theParams.StartIdx)
	if err != nil {
		return nil, 400, err
	}

	contentBlocks, err := schema.GetContentBlocks(boardID, articleID, contentID, contentIdx, int64(theParams.Max+1))
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

	ret := &GetArticleBlocksResult{
		Content: content,
		NextIdx: nextIdx,

		TokenUser: userID,
	}
	if articleDetailSummary != nil {
		ret.IsDeleted = articleDetailSummary.IsDeleted
		ret.CreateTime = articleDetailSummary.CreateTime.ToTime8()
		ret.MTime = articleDetailSummary.MTime.ToTime8()
		ret.Recommend = articleDetailSummary.Recommend
		ret.NComments = articleDetailSummary.NComments
		ret.Owner = articleDetailSummary.Owner
		ret.Title = articleDetailSummary.Title
		ret.Money = articleDetailSummary.Money
		ret.Class = articleDetailSummary.Class
		ret.Filemode = articleDetailSummary.Filemode
		ret.IP = articleDetailSummary.IP
		ret.Host = articleDetailSummary.Host
		ret.BBS = articleDetailSummary.BBS
		ret.Rank = articleDetailSummary.Rank
		ret.SubjectType = articleDetailSummary.SubjectType
	}

	return ret, 200, nil
}

func getArticleBlocksContentBlocksDeserializeContentIdx(boardID bbs.BBoardID, articleID bbs.ArticleID, theIdx string) (contentID types.ContentID, contentIdx int, articleDetailSummary *schema.ArticleDetailSummary, err error) {
	if theIdx != "" {
		contentID, contentIdx, err = apitypes.DeserializeContentIdx(theIdx)
		return contentID, contentIdx, nil, err
	}

	articleDetailSummary, err = schema.GetArticleDetailSummary(boardID, articleID)
	if err != nil {
		return "", 0, nil, err
	}

	return articleDetailSummary.ContentID, 0, articleDetailSummary, nil
}

func getArticleBlocksContentBlocksToContent(contentBlocks []*schema.ContentBlock) (content [][]*types.Rune) {
	lenContent := 0
	for _, each := range contentBlocks {
		lenContent += len(each.Content)
	}

	if lenContent == 0 {
		return nil
	}

	content = make([][]*types.Rune, 0, lenContent)
	for _, each := range contentBlocks {
		content = append(content, each.Content...)
	}

	return content
}
