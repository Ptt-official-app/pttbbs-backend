package dbcs

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

func ParseContentBlocks(boardID bbs.BBoardID, articleID bbs.ArticleID, content [][]*types.Rune, contentMD5 string, updateNanoTS types.NanoTS) (contentID types.ContentID, contentBlocks []*schema.ContentBlock) {
	if len(content) == 0 {
		return "", nil
	}

	nBlocks := len(content)/N_LINES_PER_CONTENT_BLOCK + 1

	contentBlocks = make([]*schema.ContentBlock, 0, nBlocks)

	contentID = types.ToContentID(updateNanoTS, contentMD5)

	var eachContent [][]*types.Rune
	for idx := 0; idx < nBlocks && len(content) > 0; idx++ {
		theLen := min(len(content), N_LINES_PER_CONTENT_BLOCK)

		eachContent, content = content[:theLen], content[theLen:]
		eachContentBlock := &schema.ContentBlock{
			BBoardID:     boardID,
			ArticleID:    articleID,
			ContentID:    contentID,
			Idx:          idx,
			Content:      eachContent,
			UpdateNanoTS: updateNanoTS,
		}

		contentBlocks = append(contentBlocks, eachContentBlock)
	}

	return contentID, contentBlocks
}
