package dbcs

import (
	"bytes"

	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

var (
	MATCH_COMMENT_RECOMMEND_BYTES = []byte{ // 推
		//\x1b[1;37m\xb1\xc0 \x1b[33mabcd   \x1b[m\x1b[33m: 77777777                                         \x1b[m 03/18 12:07
		0x1b, 0x5b, 0x31, 0x3b, 0x33, 0x37, 0x6d,
		0xb1, 0xc0, 0x20, 0x1b, 0x5b, 0x33, 0x33, 0x6d,
	}

	MATCH_COMMENT_BOO_BYTES = []byte{ // 噓
		//\x1b[1;31m\xbcN \x1b[33mabcd  \x1b[m\x1b[33m: \xb0\xa3\xabD\xb9\xef\xa4\xe2\xa4]....          \x1b[m 03/18 18:13
		0x1b, 0x5b, 0x31, 0x3b, 0x33, 0x31, 0x6d,
		0xbc, 0x4e, 0x20, 0x1b, 0x5b, 0x33, 0x33, 0x6d,
	}

	MATCH_COMMENT_ARROW_BYTES = []byte{ //→
		//\x1b[1;31m\xa1\xf7 \x1b[33mabcd     \x1b[m\x1b[33m: 7777777777                                       \x1b[m 03/18 12:07
		0x1b, 0x5b, 0x31, 0x3b, 0x33, 0x31, 0x6d,
		0xa1, 0xf7, 0x20, 0x1b, 0x5b, 0x33, 0x33, 0x6d,
	}

	MATCH_COMMENT_INFIX = []byte("\x1b[m\x1b[33m:")

	//※ 編輯: abcd (1.2.3.4 臺灣), 03/21/2021 03:04:47
	//\xa1\xb0 \xbds\xbf\xe8: abcd (1.2.3.4 \xbbO\xc6W), 03/18/2021 12:07:22
	MATCH_COMMENT_EDIT_BYTES = []byte("\xa1\xb0 \xbds\xbf\xe8: ")

	MATCH_COMMENT_EDIT_FROM_BYTES = []byte("\xa8\xd3\xa6\xdb: ")

	//※ abcde:轉錄至看板 SYSOP
	//\xa1\xb0 \x1b[1;32mabcd\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xac\xdd\xaaO Mavericks\x1b[m                               03/18 12:07
	//※ jasome:轉錄至某隱形看板
	//\xa1\xb0 \x1b[1;32mjasome\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc\xacY\xc1\xf4\xa7\xce\xac\xdd\xaaO\x1b[m                                         01/29 02:39
	MATCH_COMMENT_FORWARD_BYTES       = []byte("\x1b[0;32m:\xc2\xe0\xbf\xfd\xa6\xdc") //\x1b[0;32m:\xc2
	MATCH_COMMENT_FORWARD_BOARD_BYTES = []byte("\xac\xdd\xaaO ")

	MATCH_COMMENT_FORWARD_PREFIX = []byte("\xa1\xb0 \x1b[1;32m")

	MATCH_COMMENT_FORWARD_HIDDEN_BYTES = []byte("\xacY\xc1\xf4\xa7\xce\xac\xdd\xaaO")

	//(teemocogs 刪除 teemocogs 的推文: 誤植)
	//\x1b[1;30m(teemocogs \xa7R\xb0\xa3 teemocogs \xaa\xba\xb1\xc0\xa4\xe5: \xbb~\xb4\xd3)\x1b[m
	MATCH_COMMENT_DELETED_PREFIX  = []byte("\x1b[1;30m(")
	MATCH_COMMENT_DELETED_INFIX0  = []byte(" \xa7R\xb0\xa3 ")
	MATCH_COMMENT_DELETED_INFIX1  = []byte(" \xaa\xba\xb1\xc0\xa4\xe5: ")
	MATCH_COMMENT_DELETED_POSTFIX = []byte(")\x1b[m")

	MATCH_COMMENT_GREEN_PREFIX = []byte("\xa1\xb0 ") //※
)

// MatchComment
//
// TODO: record the idxes of each condition, rematch only the condition with the smallest idx.
func MatchComment(content []byte) int {
	theIdx := -1

	idxRecommend := bytes.Index(content, MATCH_COMMENT_RECOMMEND_BYTES)
	if idxRecommend != -1 && (idxRecommend == 0 || content[idxRecommend-1] == '\n') {
		theIdx = matchCommentIntegratedIdx(theIdx, idxRecommend)
	}

	idxBoo := bytes.Index(content, MATCH_COMMENT_BOO_BYTES)
	if idxBoo != -1 && (idxBoo == 0 || content[idxBoo-1] == '\n') {
		theIdx = matchCommentIntegratedIdx(theIdx, idxBoo)
	}
	idxArrow := bytes.Index(content, MATCH_COMMENT_ARROW_BYTES)
	if idxArrow != -1 && (idxArrow == 0 || content[idxArrow-1] == '\n') {
		theIdx = matchCommentIntegratedIdx(theIdx, idxArrow)
	}
	idxEdit := bytes.Index(content, MATCH_COMMENT_EDIT_BYTES)
	if idxEdit != -1 && (idxEdit == 0 || content[idxEdit-1] == '\n') {
		theIdx = matchCommentIntegratedIdx(theIdx, idxEdit)
	}
	idxForward := matchCommentForward(content)
	if idxForward != -1 {
		theIdx = matchCommentIntegratedIdx(theIdx, idxForward)
	}
	idxDeleted := matchCommentDeleted(content)
	if idxDeleted != -1 {
		theIdx = matchCommentIntegratedIdx(theIdx, idxDeleted)
	}

	if theIdx == -1 {
		return theIdx
	}

	return theIdx // do not include the leading \n
}

func MatchCommentType(commentDBCS []byte) (theType ptttype.CommentType, nextCommentDBCS []byte) {
	if bytes.HasPrefix(commentDBCS, MATCH_COMMENT_GREEN_PREFIX) {
		if bytes.HasPrefix(commentDBCS, MATCH_COMMENT_EDIT_BYTES) {
			return ptttype.COMMENT_TYPE_EDIT, commentDBCS[len(MATCH_COMMENT_EDIT_BYTES):]
		}

		isForward, newCommentDBCS := hasPrefixCommentForward(commentDBCS)
		if isForward {
			return ptttype.COMMENT_TYPE_FORWARD, newCommentDBCS
		}
	} else if bytes.HasPrefix(commentDBCS, MATCH_COMMENT_RECOMMEND_BYTES) {
		return ptttype.COMMENT_TYPE_RECOMMEND, commentDBCS[len(MATCH_COMMENT_RECOMMEND_BYTES):]
	} else if bytes.HasPrefix(commentDBCS, MATCH_COMMENT_BOO_BYTES) {
		return ptttype.COMMENT_TYPE_BOO, commentDBCS[len(MATCH_COMMENT_BOO_BYTES):]
	} else if bytes.HasPrefix(commentDBCS, MATCH_COMMENT_ARROW_BYTES) {
		return ptttype.COMMENT_TYPE_COMMENT, commentDBCS[len(MATCH_COMMENT_ARROW_BYTES):]
	} else {
		isDeleted, newCommentDBCS := hasPrefixCommentDeleted(commentDBCS)
		if isDeleted {
			return ptttype.COMMENT_TYPE_DELETED, newCommentDBCS
		}
	}

	return ptttype.COMMENT_TYPE_UNKNOWN, commentDBCS
}

func matchCommentIntegratedIdx(theIdx int, idx int) int {
	if theIdx == -1 {
		return idx
	}
	if theIdx < idx {
		return theIdx
	}
	return idx
}

// matchCommentForward returns the offset of the first forward comment in content, or -1.
//
// base tracks how much of content has already been skipped, so that the returned
// offset is always relative to the original content.
func matchCommentForward(content []byte) (idx int) {
	for base := 0; base < len(content); {
		rel := bytes.Index(content[base:], MATCH_COMMENT_FORWARD_BYTES)
		if rel == -1 {
			return -1
		}

		pos := base + rel
		next := pos + len(MATCH_COMMENT_FORWARD_BYTES)
		if rel < 3 {
			base = next
			continue
		}

		idx = pos - 1
		//nolint:revive
		for ; idx >= base+2 && content[idx] != ' '; idx-- {
		}
		start := idx - 2
		if start >= base && content[idx] == ' ' && content[idx-1] == '\xb0' && content[idx-2] == '\xa1' &&
			(start == 0 || content[start-1] == '\n') {
			return start
		}

		base = next
	}

	return -1
}

func hasPrefixCommentForward(content []byte) (isForward bool, nextCommentDBCS []byte) {
	idx := bytes.Index(content, MATCH_COMMENT_FORWARD_BYTES)
	if idx == -1 {
		return false, nil
	}

	if idx < 3 {
		return false, nil
	}
	idx--
	//nolint:revive
	for ; idx >= 2 && content[idx] != ' '; idx-- {
	}
	if idx == 2 && content[idx-2] == '\xa1' && content[idx-1] == '\xb0' && content[idx] == ' ' {
		return true, content[len(MATCH_COMMENT_FORWARD_PREFIX):]
	}

	return false, nil
}

// matchCommentDeleted returns the offset of the first deleted comment in content, or -1.
//
// base tracks how much of content has already been skipped, so that the returned
// offset is always relative to the original content.
func matchCommentDeleted(content []byte) (idx int) {
	for base := 0; base < len(content); {
		rest := content[base:]
		idx = bytes.Index(rest, MATCH_COMMENT_DELETED_PREFIX)
		if idx == -1 {
			return -1
		}
		startIdx := idx
		//nolint:revive
		for ; idx < len(rest) && rest[idx] != ' '; idx++ {
		}
		if !bytes.HasPrefix(rest[idx:], MATCH_COMMENT_DELETED_INFIX0) {
			base += idx
			continue
		}
		idx += len(MATCH_COMMENT_DELETED_INFIX0)
		//nolint:revive
		for ; idx < len(rest) && rest[idx] != ' '; idx++ {
		}
		if !bytes.HasPrefix(rest[idx:], MATCH_COMMENT_DELETED_INFIX1) {
			base += idx
			continue
		}

		return base + startIdx
	}

	return -1
}

func hasPrefixCommentDeleted(content []byte) (isDeleted bool, nextCommentDBCS []byte) {
	if !bytes.HasPrefix(content, MATCH_COMMENT_DELETED_PREFIX) {
		return false, nil
	}
	idx := len(MATCH_COMMENT_DELETED_PREFIX)
	//nolint:revive
	for ; idx < len(content) && content[idx] != ' '; idx++ {
	}
	if !bytes.HasPrefix(content[idx:], MATCH_COMMENT_DELETED_INFIX0) {
		return false, nil
	}
	idx += len(MATCH_COMMENT_DELETED_INFIX0)
	//nolint:revive
	for ; idx < len(content) && content[idx] != ' '; idx++ {
	}
	if !bytes.HasPrefix(content[idx:], MATCH_COMMENT_DELETED_INFIX1) {
		return false, nil
	}

	return true, content[len(MATCH_COMMENT_DELETED_PREFIX):]
}
