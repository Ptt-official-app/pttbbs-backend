package apitypes

import (
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

func ToURL(fbboardID FBoardID, farticleID FArticleID) string {
	if fbboardID == "" || farticleID == "" {
		return ""
	}
	return types.FRONTEND_PREFIX + "/" + "board" + "/" + string(fbboardID) + "/" + "article" + "/" + string(farticleID)
}
