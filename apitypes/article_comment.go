package apitypes

import (
	"strconv"
	"strings"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

type ArticleCommentType string

const (
	ARTICLE_COMMENT_TYPE_ARTICLE ArticleCommentType = "a"
	ARTICLE_COMMENT_TYPE_COMMENT ArticleCommentType = "c"
)

type ArticleCommentIdx string

type ArticleComment struct {
	FBoardID   FBoardID    `json:"bid"`         //
	FArticleID FArticleID  `json:"aid"`         //
	IsDeleted  bool        `json:"deleted"`     //
	CreateTime types.Time8 `json:"create_time"` //
	MTime      types.Time8 `json:"modified"`    //
	Recommend  int         `json:"recommend"`   //
	NComments  int         `json:"n_comments"`  //
	Owner      bbs.UUserID `json:"owner"`       //
	Title      string      `json:"title"`       //
	Money      int         `json:"money"`       //
	Class      string      `json:"class"`       // can be: R: 轉, [class]

	URL  string           `json:"url"`  //
	Read types.ReadStatus `json:"read"` //

	Idx string `json:"idx"`

	Rank int `json:"rank"`

	TheType ArticleCommentType `json:"type"`

	CommentID         types.CommentID     `json:"cid,omitempty"`
	CommentType       ptttype.CommentType `json:"ctype,omitzero"`
	CommentCreateTime types.Time8         `json:"ctime,omitzero"`
	Comment           [][]*types.Rune     `json:"comment,omitempty"`

	TokenUser bbs.UUserID `json:"tokenuser"`
}

func NewArticleCommentFromArticle(a_db *schema.ArticleSummary) *ArticleComment {
	fboardID := ToFBoardID(a_db.BBoardID)
	farticleID := ToFArticleID(a_db.ArticleID)

	url := ToURL(fboardID, farticleID)

	idx := SerializeArticleCommentIdx(ARTICLE_COMMENT_TYPE_ARTICLE, a_db.Idx)

	return &ArticleComment{
		FBoardID:   fboardID,
		FArticleID: farticleID,
		IsDeleted:  a_db.IsDeleted,
		CreateTime: a_db.CreateTime.ToTime8(),
		MTime:      a_db.MTime.ToTime8(),
		Recommend:  a_db.Recommend,
		NComments:  a_db.NComments,
		Owner:      a_db.Owner,
		Title:      a_db.Title,
		Money:      a_db.Money,
		Class:      a_db.Class,
		URL:        url,
		Idx:        idx,
		Rank:       a_db.Rank,

		TheType: ARTICLE_COMMENT_TYPE_ARTICLE,
	}
}

func NewArticleCommentFromComment(a_db *schema.ArticleSummary, c_db *schema.CommentSummary) *ArticleComment {
	fboardID := ToFBoardID(c_db.BBoardID)
	farticleID := ToFArticleID(c_db.ArticleID)

	url := ToURL(fboardID, farticleID)

	commentIdx := SerializeCommentIdx(c_db.SortTime, c_db.CommentID)
	idx := SerializeArticleCommentIdx(ARTICLE_COMMENT_TYPE_COMMENT, commentIdx)

	return &ArticleComment{
		FBoardID:   fboardID,
		FArticleID: farticleID,
		IsDeleted:  c_db.IsDeleted,
		CreateTime: a_db.CreateTime.ToTime8(),
		MTime:      a_db.MTime.ToTime8(),
		Recommend:  a_db.Recommend,
		NComments:  a_db.NComments,
		Owner:      a_db.Owner,
		Title:      a_db.Title,
		Money:      a_db.Money,
		Class:      a_db.Class,
		URL:        url,
		Idx:        idx,
		Rank:       a_db.Rank,

		TheType: ARTICLE_COMMENT_TYPE_COMMENT,

		CommentID:         c_db.CommentID,
		CommentType:       c_db.TheType,
		CommentCreateTime: c_db.CreateTime.ToTime8(),
		Comment:           c_db.Content,
	}
}

func SerializeArticleCommentIdx(theType ArticleCommentType, idx string) string {
	if idx == "" {
		return ""
	}

	return string(theType) + "#" + idx
}

func DeserializeArticleCommentIdx(theIdx string) (theType ArticleCommentType, subIdx string, err error) {
	theList := strings.Split(theIdx, "#")
	if len(theList) < 2 {
		return "", "", ErrInvalidIdx
	}

	theType = ArticleCommentType(theList[0])
	subIdx = strings.Join(theList[1:], "#")

	return theType, subIdx, nil
}

func SerializeContentIdx(contentID types.ContentID, contentIdx int) (theIdx string) {
	return string(contentID) + "^" + strconv.Itoa(contentIdx)
}

func DeserializeContentIdx(theIdx string) (contentID types.ContentID, contentIdx int, err error) {
	theList := strings.Split(theIdx, "^")
	if len(theList) != 2 {
		return "", 0, ErrInvalidIdx
	}

	contentIdx, err = strconv.Atoi(theList[1])
	if err != nil {
		return "", 0, ErrInvalidIdx
	}

	return types.ContentID(theList[0]), contentIdx, nil
}
