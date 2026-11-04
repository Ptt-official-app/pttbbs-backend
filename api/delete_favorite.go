package api

import (
	"strconv"
	"strings"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/gin-gonic/gin"
)

const DELETE_FAVORITE_R = "/user/:user_id/favorites/delete"

type DeleteFavoriteParams struct {
	LevelIdx schema.LevelIdx `json:"level_idx,omitempty" form:"level_idx,omitempty" url:"level_idx,omitempty"`
	Idx      string          `json:"idx,omitempty" form:"idx,omitempty" url:"idx,omitempty"`
}

type DeleteFavoritePath struct {
	UserID bbs.UUserID `uri:"user_id"`
}

type DeleteFavoriteResult struct {
	Success   bool        `json:"success"`
	TokenUser bbs.UUserID `json:"tokenuser"`
}

func DeleteFavoriteWrapper(c *gin.Context) {
	params := &DeleteFavoriteParams{}
	path := &DeleteFavoritePath{}
	LoginRequiredPathJSON(DeleteFavorite, params, path, c)
}

func DeleteFavorite(remoteAddr string, user *UserInfo, params any, path any, c *gin.Context) (result any, statusCode int, err error) {
	theParams, ok := params.(*DeleteFavoriteParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*DeleteFavoritePath)
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

	idx, err := strconv.Atoi(theParams.Idx)
	if err != nil {
		return nil, 500, err
	}

	err = theFav.DeleteIdx(idx)
	if err != nil {
		return nil, 500, err
	}

	statusCode, err = tryWriteFav(rootFav, remoteAddr, userID, c)
	if err != nil {
		return nil, statusCode, err
	}

	_, _, statusCode, err = tryGetUserFavorites(userID, theParams.LevelIdx, "", true, 1, c)
	if err != nil {
		return nil, statusCode, err
	}

	return &DeleteFavoriteResult{Success: true, TokenUser: userID}, 200, nil
}
