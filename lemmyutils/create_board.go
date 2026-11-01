package lemmyutils

import (
	"net/http"

	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"
)

type CreateBoardParams struct {
	Name  string `json:"name"`
	Title string `json:"title"`
	NSFW  bool   `json:"nsfw"`
}

type CreateBoardResult struct {
	DiscussionLanguages []int         `json:"discussion_languages"`
	CommunityView       CommunityView `json:"community_view"`
}

func TryCreateBoard(name string, title string, nsfw bool, c *gin.Context) (ret *CreateBoardResult, err error) {
	err = TryLemmyLogin(c)
	if err != nil {
		return nil, err
	}

	ret, err = CreateBoard(name, title, nsfw, c)
	if err == nil {
		return ret, nil
	}

	if err != ErrInvalidBackendStatusCode {
		return nil, err
	}

	err = LemmyLogin(c)
	if err != nil {
		return nil, err
	}

	return CreateBoard(name, title, nsfw, c)
}

func CreateBoard(name string, title string, nsfw bool, c *gin.Context) (ret *CreateBoardResult, err error) {
	// https://join-lemmy.org/api/main#tag/Community/operation/CreateCommunity
	url := types.LEMMY_PREFIX + "/community"

	params := &CreateBoardParams{
		Name:  name,
		Title: title,
		NSFW:  nsfw,
	}

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + LEMMY_ADMIN_TOKEN,
	}

	statusCode, err := utils.BackendPost(c, url, params, headers, &ret)
	if err != nil {
		return nil, err
	}
	if statusCode != http.StatusOK {
		return nil, ErrInvalidBackendStatusCode
	}

	return ret, nil
}
