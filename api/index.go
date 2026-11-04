package api

import (
	"github.com/gin-gonic/gin"
)

const INDEX_R = "/"

type IndexParams struct {
	In int `json:"in,omitzero" form:"in,omitempty" url:"in,omitempty"`
}

func NewIndexParams() *IndexParams {
	return &IndexParams{}
}

type IndexResult struct {
	Data any `json:"data,omitempty"`
}

func IndexWrapper(c *gin.Context) {
	params := NewIndexParams()
	Query(Index, params, c)
}

func Index(remoteAddr string, user *UserInfo, params any, c *gin.Context) (result any, statusCode int, err error) {
	result = &IndexResult{Data: params}
	return result, 200, nil
}
