package api

import (
	"github.com/gin-gonic/gin"
)

const INDEX_R = "/"

type IndexParams struct {
	In int `json:"in,omitempty" form:"in,omitempty" url:"in,omitempty"`
}

func NewIndexParams() *IndexParams {
	return &IndexParams{}
}

type IndexResult struct {
	Data interface{} `json:"data,omitempty"`
}

func IndexWrapper(c *gin.Context) {
	params := NewIndexParams()
	Query(Index, params, c)
}

func Index(remoteAddr string, user *UserInfo, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	result = &IndexResult{Data: params}
	return result, 200, nil
}
