package dbcs

import (
	"strings"

	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

var (
	MATCH_COMMENT_RECOMMEND_STR = "\x1b[1;37m推 \x1b[33m" // 推
	//\x1b[1;37m推 \x1b[33mfennyccc\x1b[m\x1b[33m: 讚                                    \x1b[m 115.82.149.161 01/28 12:34

	MATCH_COMMENT_BOO_STR = "\x1b[1;31m噓 \x1b[33m" // 噓
	//\x1b[1;31m噓 \x1b[33m

	MATCH_COMMENT_ARROW_STR = "\x1b[1;31m→ \x1b[33m" //→
	//\x1b[1;31m→ \x1b[33mSYSOP\x1b[m\x1b[33m:推推                                                     \x1b[m 12/13 03:51

	MATCH_DEFAULT_INFIX_STR = "\x1b[m\x1b[33m"

	//※ 編輯: peter50505      來自: 163.27.69.176        (10/03 14:42)
	//\xa1\xb0 \xbds\xbf\xe8: abcd (1.2.3.4 \xbbO\xc6W), 03/18/2021 12:07:22
	MATCH_COMMENT_EDIT_STR = "※ 編輯: "

	MATCH_COMMENT_EDIT_FROM_STR = "來自: "

	//※ abcde:轉錄至看板 SYSOP
	//※ \x1b[1;32mPttACT\x1b[0;32m:轉錄至看板 OriginalSong\x1b[m                                  01/26 17:19
	//
	//※ jasome:轉錄至某隱形看板
	//\xa1\xb0 \x1b[1;32mjasome\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xacY\xc1\xf4\xa7\xce\xac\xdd\xaaO\x1b[m                                         01/29 02:39
	MATCH_COMMENT_FORWARD_STR        = "\x1b[0;32m:轉錄至"
	MATCH_COMMENT_FORWARD_BOARD_STR  = "看板 "
	MATCH_COMMENT_FORWARD_PREFIX_STR = "※ \x1b[1;32m"

	//(teemocogs 刪除 teemocogs 的推文: 誤植)
	//\x1b[1;30m(teemocogs 刪除 teemocogs 的推文: 誤植)\x1b[m
	MATCH_COMMENT_DELETED_PREFIX_STR  = "\x1b[1;30m("
	MATCH_COMMENT_DELETED_INFIX0_STR  = " 刪除 "
	MATCH_COMMENT_DELETED_INFIX1_STR  = " 的推文: "
	MATCH_COMMENT_DELETED_POSTFIX_STR = ")\x1b[m"

	MATCH_COMMENT_GREEN_PREFIX_STR = "※ " //※
)

// MatchCommentStr is the utf-8 version of MatchComment.
func MatchCommentStr(content string) int {
	for lineStart := 0; lineStart < len(content); {
		line := content[lineStart:]
		idxNewLine := strings.IndexByte(line, '\n')
		if idxNewLine != -1 {
			line = line[:idxNewLine]
		}

		theType, _ := MatchCommentTypeStr(line)
		if theType != ptttype.COMMENT_TYPE_UNKNOWN {
			return lineStart
		}

		if idxNewLine == -1 {
			return -1
		}
		lineStart += idxNewLine + 1
	}

	return -1
}

func MatchCommentTypeStr(commentDBCS string) (theType ptttype.CommentType, nextCommentDBCS string) {
	if strings.HasPrefix(commentDBCS, MATCH_COMMENT_GREEN_PREFIX_STR) {
		if strings.HasPrefix(commentDBCS, MATCH_COMMENT_EDIT_STR) {
			return ptttype.COMMENT_TYPE_EDIT, commentDBCS[len(MATCH_COMMENT_EDIT_STR):]
		}

		isForward, newCommentDBCS := hasPrefixCommentForwardStr(commentDBCS)
		if isForward {
			return ptttype.COMMENT_TYPE_FORWARD, newCommentDBCS
		}
	} else if strings.HasPrefix(commentDBCS, MATCH_COMMENT_RECOMMEND_STR) {
		return ptttype.COMMENT_TYPE_RECOMMEND, commentDBCS[len(MATCH_COMMENT_RECOMMEND_STR):]
	} else if strings.HasPrefix(commentDBCS, MATCH_COMMENT_BOO_STR) {
		return ptttype.COMMENT_TYPE_BOO, commentDBCS[len(MATCH_COMMENT_BOO_STR):]
	} else if strings.HasPrefix(commentDBCS, MATCH_COMMENT_ARROW_STR) {
		return ptttype.COMMENT_TYPE_COMMENT, commentDBCS[len(MATCH_COMMENT_ARROW_STR):]
	} else {
		isDeleted, newCommentDBCS := hasPrefixCommentDeletedStr(commentDBCS)
		if isDeleted {
			return ptttype.COMMENT_TYPE_DELETED, newCommentDBCS
		}
	}

	return ptttype.COMMENT_TYPE_UNKNOWN, commentDBCS
}

func hasPrefixCommentForwardStr(content string) (isForward bool, nextCommentDBCS string) {
	idx := strings.Index(content, MATCH_COMMENT_FORWARD_STR)
	if idx == -1 {
		return false, ""
	}

	lenPrefix := len("※")
	if idx < lenPrefix+1 {
		return false, ""
	}
	idx--
	//nolint:revive
	for ; idx >= lenPrefix && content[idx] != ' '; idx-- {
	}
	if idx == lenPrefix && content[0:idx] == "※" && content[idx] == ' ' {
		return true, content[len(MATCH_COMMENT_FORWARD_PREFIX_STR):]
	}

	return false, ""
}

func hasPrefixCommentDeletedStr(content string) (isDeleted bool, nextCommentDBCS string) {
	if !strings.HasPrefix(content, MATCH_COMMENT_DELETED_PREFIX_STR) {
		return false, ""
	}
	idx := len(MATCH_COMMENT_DELETED_PREFIX_STR)
	//nolint:revive
	for ; idx < len(content) && content[idx] != ' '; idx++ {
	}
	if !strings.HasPrefix(content[idx:], MATCH_COMMENT_DELETED_INFIX0_STR) {
		return false, ""
	}
	idx += len(MATCH_COMMENT_DELETED_INFIX0_STR)
	//nolint:revive
	for ; idx < len(content) && content[idx] != ' '; idx++ {
	}
	if !strings.HasPrefix(content[idx:], MATCH_COMMENT_DELETED_INFIX1_STR) {
		return false, ""
	}

	return true, content[len(MATCH_COMMENT_DELETED_PREFIX_STR):]
}
