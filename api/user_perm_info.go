package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/gin-gonic/gin"
)

func getUserPermInfo(userID bbs.UUserID, c *gin.Context) (userPermInfo *schema.UserPermInfo, err error) {
	return schema.GetUserPermInfo(userID)
}
