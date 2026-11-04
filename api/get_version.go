package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"
)

const GET_VERSION_R = "/version"

type GetVersionResult struct {
	Version    string `json:"version"`
	GitVersion string `json:"commit"`

	PttVersion    string `json:"pttversion"`
	PttGitVersion string `json:"pttcommit"`
}

func GetVersionWrapper(c *gin.Context) {
	Query(GetVersion, nil, c)
}

func GetVersion(remoteAddr string, user *UserInfo, params any, c *gin.Context) (result any, statusCode int, err error) {
	url := pttbbsapi.GET_VERSION_R

	var result_b *pttbbsapi.GetVersionResult

	statusCode, err = utils.BackendGet(c, url, nil, nil, &result_b)
	if err != nil {
		return nil, statusCode, err
	}

	return &GetVersionResult{
		Version:       types.VERSION,
		GitVersion:    types.GIT_VERSION,
		PttVersion:    result_b.Version,
		PttGitVersion: result_b.GitVersion,
	}, 200, nil
}
