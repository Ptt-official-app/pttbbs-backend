package api

import (
	"strings"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/gin-gonic/gin"
)

const ADD_FAVORITE_LINE_R = "/user/:user_id/favorites/addline"

type AddFavoriteLineParams struct {
	LevelIdx schema.LevelIdx `json:"level_idx,omitempty" form:"level_idx,omitempty" url:"level_idx,omitempty"`
}

type AddFavoriteLinePath struct {
	UserID bbs.UUserID `uri:"user_id"`
}

func AddFavoriteLineWrapper(c *gin.Context) {
	params := &AddFavoriteLineParams{}
	path := &AddFavoriteLinePath{}
	LoginRequiredPathJSON(AddFavoriteLine, params, path, c)
}

func AddFavoriteLine(remoteAddr string, user *UserInfo, params any, path any, c *gin.Context) (result any, statusCode int, err error) {
	theParams, ok := params.(*AddFavoriteLineParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*AddFavoriteLinePath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	userID := user.UserID

	if userID != thePath.UserID {
		return nil, 403, ErrInvalidUser
	}

	userFavoritesMeta, userFavorites, err := tryGetAllUserFavorites(userID, c)
	if err != nil {
		return nil, 500, err
	}

	rootFav, err := schema.UserFavoritesToFav(&userFavoritesMeta.FolderMeta, userFavorites, 0)
	if err != nil {
		return nil, 500, err
	}

	levelIdxList := strings.Split(string(theParams.LevelIdx), ":")
	theFav, err := rootFav.LocateFav(levelIdxList)
	if err != nil {
		return nil, 500, err
	}

	theIdx, _, err := theFav.AddLine()
	if err != nil {
		return nil, 500, err
	}

	return postAddFavorite(userID, rootFav, remoteAddr, theParams.LevelIdx, theIdx, c, false)
}
