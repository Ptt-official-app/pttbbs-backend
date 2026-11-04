package api

import (
	"strings"

	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/gin-gonic/gin"
)

const GET_BOARD_DETAIL_R = "/board/:bid"

type GetBoardDetailParams struct {
	Fields string `json:"fields,omitempty" form:"fields,omitempty" uri:"fields,omitempty"`
}

type GetBoardDetailPath struct {
	FBoardID apitypes.FBoardID `uri:"bid"`
}

type GetBoardDetailResult *apitypes.BoardDetail

func GetBoardDetailWrapper(c *gin.Context) {
	params := &GetBoardDetailParams{}
	path := &GetBoardDetailPath{}
	LoginRequiredPathQuery(GetBoardDetail, params, path, c)
}

func GetBoardDetail(remoteAddr string, user *UserInfo, params any, path any, c *gin.Context) (result any, statusCode int, err error) {
	theParams, ok := params.(*GetBoardDetailParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*GetBoardDetailPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	userID := user.UserID
	boardID, err := toBoardID(thePath.FBoardID, remoteAddr, userID, c)
	if err != nil {
		return nil, 400, err
	}

	// is board-valid-user
	_, err = CheckUserBoardPermReadable(user, boardID, c)
	if err != nil {
		return nil, 403, err
	}

	// fields
	var fieldMap map[string]bool
	if len(theParams.Fields) > 0 {
		fieldMap = make(map[string]bool)
		for each := range strings.SplitSeq(theParams.Fields, ",") {
			each_db, ok := apitypes.BOARD_DETAIL_FIELD_MAP[each]
			if !ok {
				each_db = each
			}
			fieldMap[each_db] = true
		}
		fieldMap[schema.BOARD_BBOARD_ID_b] = true
	}

	boardDetail_db, err := schema.GetBoardDetail(boardID, fieldMap)
	if err != nil {
		return nil, 500, err
	}

	boardDetail := apitypes.NewBoardDetail(boardDetail_db, "", userID)

	result = GetBoardDetailResult(boardDetail)

	return result, 200, nil
}
