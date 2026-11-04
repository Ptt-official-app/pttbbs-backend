package dbcs

import (
	"strings"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

func ParseCommentsStr(ownerID bbs.UUserID, commentsDBCS string, allCommentsDBCS string) (comments []*schema.Comment) {
	if len(commentsDBCS) == 0 {
		return nil
	}

	// 1. estimate nComments
	nEstimatedComments := strings.Count(commentsDBCS, "\n")

	comments = make([]*schema.Comment, 0, nEstimatedComments)

	nextCommentIdx := MatchCommentStr(commentsDBCS)

	// 2. reply
	if nextCommentIdx > 0 || nextCommentIdx == -1 {
		if nextCommentIdx == -1 {
			nextCommentIdx = len(commentsDBCS)
		}
		replyDBCS := commentsDBCS[:nextCommentIdx]

		reply := parseReplyStr(replyDBCS, allCommentsDBCS, ownerID)
		if reply != nil {
			comments = append(comments, reply)
		}

		allCommentsDBCS = allCommentsDBCS[len(replyDBCS):]
		commentsDBCS = commentsDBCS[len(replyDBCS):]
	}

	var comment *schema.Comment
	// 3. for each comment-leading
	for idxNewLine := strings.Index(commentsDBCS, "\n"); len(commentsDBCS) > 0 && idxNewLine != -1; idxNewLine = strings.Index(commentsDBCS, "\n") {

		// 3.0 parse comment
		eachCommentDBCS := commentsDBCS[:idxNewLine]
		comment = parseCommentStr(eachCommentDBCS)
		comments = append(comments, comment)

		commentsDBCS = commentsDBCS[idxNewLine:]
		allCommentsDBCS = allCommentsDBCS[idxNewLine:]

		// 3.1 find next comment
		nextCommentIdx = MatchCommentStr(commentsDBCS)

		if nextCommentIdx == -1 { // no more comments
			// 3.1.1 assume the rest are reply
			commentsDBCS = commentsDBCS[1:] // step forward "\n"
			allCommentsDBCS = allCommentsDBCS[1:]
			if len(commentsDBCS) > 0 {
				replyDBCS := commentsDBCS
				reply := parseReplyStr(replyDBCS, allCommentsDBCS, ownerID)
				if reply != nil {
					comments = append(comments, reply)
				}

				allCommentsDBCS = allCommentsDBCS[len(commentsDBCS):] //nolint // consistent with programming pattern
				commentsDBCS = commentsDBCS[len(commentsDBCS):]
			}
			break
		}

		// 3.2 assume the context before the next comment is reply
		if nextCommentIdx > 1 { // commentsDBCS[0] == "\n", get reply from commentsDBCS[1:]
			replyDBCS := commentsDBCS[1:nextCommentIdx]
			reply := parseReplyStr(replyDBCS, allCommentsDBCS[1:], ownerID)
			if reply != nil {
				comments = append(comments, reply)
			}
		}

		commentsDBCS = commentsDBCS[nextCommentIdx:]
		allCommentsDBCS = allCommentsDBCS[nextCommentIdx:]
	}

	// 4. last comment without '\n'

	if len(commentsDBCS) > 0 {
		comment := parseCommentStr(commentsDBCS)
		comments = append(comments, comment)
	}

	if len(comments) == 0 {
		comments = nil
	}

	return comments
}

func parseCommentStr(commentDBCS string) (comment *schema.Comment) {
	theType, p_commentDBCS := MatchCommentTypeStr(commentDBCS)

	switch theType {
	case ptttype.COMMENT_TYPE_EDIT:
		return parseCommentEditStr(p_commentDBCS, commentDBCS)
	case ptttype.COMMENT_TYPE_FORWARD:
		return parseCommentForwardStr(p_commentDBCS, commentDBCS)
	case ptttype.COMMENT_TYPE_DELETED:
		return parseCommentDeletedStr(p_commentDBCS, commentDBCS)

	default:
		return parseCommentDefaultStr(theType, p_commentDBCS, commentDBCS)
	}
}

func parseCommentEditStr(commentDBCS string, origCommentDBCS string) (comment *schema.Comment) {
	ownerID, commentDBCS := parseCommentEditOwnerIDStr(commentDBCS)
	ip, host, theDateStr := parseCommentEditIPCreateTimeStr(commentDBCS)
	commentMD5 := Md5sum([]byte(origCommentDBCS))

	createNanoTS := types.NanoTS(0)
	commentID := types.CommentID("")
	theTime, err := types.DateYearTimeStrToTime(theDateStr)
	if err == nil {
		createNanoTS = types.TimeToNanoTS(theTime)
		commentID = types.ToCommentID(createNanoTS, commentMD5)
	}

	comment = &schema.Comment{
		TheType: ptttype.COMMENT_TYPE_EDIT,
		Owner:   ownerID,
		IP:      ip,
		Host:    host,
		MD5:     commentMD5,
		DBCSStr: origCommentDBCS,
		TheDate: theDateStr,

		CreateTime:         createNanoTS,
		InferredCreateTime: createNanoTS,
		SortTime:           createNanoTS,

		CommentID: commentID,
	}

	return comment
}

// parseCommentEditOwnerIDStr
//
// peter50505      來自: 163.27.69.176        (10/03 14:42)
// mapinkie (58.152.169.221 臺灣), 03/07/2020 15:49:09
func parseCommentEditOwnerIDStr(commentDBCS string) (ownerID bbs.UUserID, newCommentDBCS string) {
	p_commentDBCS := commentDBCS

	before, after, ok := strings.Cut(p_commentDBCS, " ")
	if !ok {
		return "", commentDBCS
	}

	ownerID = bbs.UUserID(string(before))

	return ownerID, after
}

// parseCommentEditIPCreateTime
//
// 來自: 163.27.69.176        (10/03 14:42)
//
// (58.152.169.221 臺灣), 03/07/2020 15:49:09
func parseCommentEditIPCreateTimeStr(commentDBCS string) (ip string, host string, theDate string) {
	if commentDBCS[0] != '(' { // old
		theIdx := strings.Index(commentDBCS, MATCH_COMMENT_EDIT_FROM_STR)
		if theIdx == -1 {
			return "", "", ""
		}

		commentDBCS = commentDBCS[theIdx+len(MATCH_COMMENT_EDIT_FROM_STR):]

		// ip
		for idx, each := range commentDBCS {
			if each == ' ' || each == '(' {
				ip = string(commentDBCS[:idx])
				commentDBCS = commentDBCS[idx:]
				break
			}
		}

		// create-time
		theIdx = strings.Index(commentDBCS, "(")
		if theIdx == -1 {
			return ip, "", ""
		}

		commentDBCS = commentDBCS[theIdx+1:]

		theIdx = strings.Index(commentDBCS, ")")
		theDate = string(commentDBCS[:theIdx])

		return ip, "", theDate
	}

	// ip
	theIdx := strings.Index(commentDBCS, ")")
	ipList := strings.Split(commentDBCS[1:theIdx], " ")
	ip = string(ipList[0])
	if len(ipList) == 2 {
		host = ipList[1]
	}

	theDate = strings.TrimSpace(string(commentDBCS[theIdx+3:])) //"), "

	return ip, host, theDate
}

// parseCommentForwardStr
//
// ※ \x1b[1;32mPttACT\x1b[0;32m:轉錄至看板 OriginalSong\x1b[m                                  01/26 17:19
func parseCommentForwardStr(commentDBCS string, origCommentDBCS string) (comment *schema.Comment) {
	ownerID, commentDBCS := parseCommentForwardOwnerIDStr(commentDBCS)
	contentDBCS, commentDBCS := parseCommentForwardContentStr(commentDBCS)
	contentUtf8 := dbcsToUtf8(contentDBCS)
	ip, theDateStr := parseCommentForwardIPCreateTimeStr(commentDBCS)
	commentMD5 := Md5sum([]byte(origCommentDBCS))

	comment = &schema.Comment{
		TheType: ptttype.COMMENT_TYPE_FORWARD,
		Owner:   ownerID,
		Content: contentUtf8,
		IP:      ip,
		MD5:     commentMD5,
		DBCSStr: origCommentDBCS,
		TheDate: theDateStr,
	}
	comment.CleanComment()

	return comment
}

// parseCommentForwardOwnerIDStr
//
// PttACT\x1b[0;32m:轉錄至看板 OriginalSong\x1b[m                                  01/26 17:19
func parseCommentForwardOwnerIDStr(commentDBCS string) (ownerID bbs.UUserID, newCommentDBCS string) {
	idx := strings.Index(commentDBCS, "\x1b")
	if idx == -1 {
		return "", ""
	}

	ownerID = bbs.UUserID(strings.TrimSpace(commentDBCS[:idx]))

	return ownerID, commentDBCS[idx:]
}

// parseCommentForwardContentStr
//
// \x1b[0;32m:轉錄至看板 OriginalSong\x1b[m                                  01/26 17:19
func parseCommentForwardContentStr(commentDBCS string) (contentDBCS string, newCommentDBCS string) {
	if !strings.HasPrefix(commentDBCS, MATCH_COMMENT_FORWARD_STR) {
		return "", ""
	}

	commentDBCS = commentDBCS[len(MATCH_COMMENT_FORWARD_STR):]
	if commentDBCS[:len(MATCH_COMMENT_FORWARD_BOARD_STR)] == MATCH_COMMENT_FORWARD_BOARD_STR {
		commentDBCS = commentDBCS[len(MATCH_COMMENT_FORWARD_BOARD_STR):]
	}

	idx := strings.Index(commentDBCS, "\x1b")
	if idx == -1 {
		return "", ""
	}

	contentDBCS = commentDBCS[:idx]
	newCommentDBCS = commentDBCS[idx:]

	return contentDBCS, newCommentDBCS
}

// parseCommentForwardIPCreateTimeStr
//
// \x1b[m                                  01/26 17:19
func parseCommentForwardIPCreateTimeStr(commentDBCS string) (ip string, dateStr string) {
	theIdx := -1
	for idx, each := range commentDBCS {
		if each >= '0' && each <= '9' {
			theIdx = idx
			break
		}
	}
	if theIdx == -1 {
		return "", ""
	}

	return "", strings.TrimSpace(string(commentDBCS[theIdx:]))
}

// parseCommentDeletedStr
//
// (teemocogs 刪除 teemocogs 的推文: 誤植)
// \x1b[1;30m(teemocogs \xa7R\xb0\xa3 teemocogs \xaa\xba\xb1\xc0\xa4\xe5: \xbb~\xb4\xd3)\x1b[m
func parseCommentDeletedStr(commentDBCS string, origCommentDBCS string) (comment *schema.Comment) {
	ownerID, commentDBCS := parseCommentDeletedOwnerIDStr(commentDBCS)
	contentDBCS, _ := parseCommentDeletedContentStr(commentDBCS)
	contentUtf8 := dbcsToUtf8(contentDBCS)
	commentMD5 := Md5sum([]byte(origCommentDBCS))

	comment = &schema.Comment{
		TheType: ptttype.COMMENT_TYPE_DELETED,
		Owner:   ownerID,
		Content: contentUtf8,
		MD5:     commentMD5,
		DBCSStr: origCommentDBCS,
	}
	comment.CleanComment()

	return comment
}

// parseCommentDeletedOwnerIDStr
//
// teemocogs 刪除 teemocogs 的推文: 誤植)
// teemocogs \xa7R\xb0\xa3 teemocogs \xaa\xba\xb1\xc0\xa4\xe5: \xbb~\xb4\xd3)\x1b[m
func parseCommentDeletedOwnerIDStr(commentDBCS string) (ownerID bbs.UUserID, newCommentDBCS string) {
	origCommentDBCS := commentDBCS
	startIdx := strings.Index(commentDBCS, MATCH_COMMENT_DELETED_INFIX0_STR)
	if startIdx == -1 {
		return "", commentDBCS
	}
	commentDBCS = commentDBCS[startIdx+len(MATCH_COMMENT_DELETED_INFIX0_STR):]
	before, _, ok := strings.Cut(commentDBCS, " ")
	if !ok {
		return "", origCommentDBCS
	}

	ownerID = bbs.UUserID(string(before))

	return ownerID, origCommentDBCS
}

// parseCommentDeletedContentStr
//
// teemocogs 刪除 teemocogs 的文章: 誤植)\x1b[m
func parseCommentDeletedContentStr(commentDBCS string) (contentDBCS string, newCommentDBCS string) {
	idx := strings.Index(commentDBCS, MATCH_COMMENT_DELETED_POSTFIX_STR)
	if idx == -1 {
		return commentDBCS, ""
	}

	return commentDBCS[:idx], commentDBCS[idx:]
}

// parseCommentDefaultStr
//
// fennyccc\x1b[m\x1b[33m: 讚                                    \x1b[m 115.82.149.161 01/28 12:34
//
// SYSOP\x1b[m\x1b[33m:推推                                                     \x1b[m 12/13 03:51
func parseCommentDefaultStr(theType ptttype.CommentType, commentDBCS string, origCommentDBCS string) (comment *schema.Comment) {
	ownerID, commentDBCS := parseCommentDefaultOwnerIDStr(commentDBCS)

	contentDBCS, commentDBCS := parseCommentContentStr(commentDBCS)
	contentUtf8 := dbcsToUtf8(contentDBCS)
	ip, theDateStr := parseCommentDefaultIPCreateTimeStr(commentDBCS)
	origCommentDBCSBytes := []byte(origCommentDBCS)
	commentMD5 := Md5sum(origCommentDBCSBytes)

	comment = &schema.Comment{
		TheType: theType,
		Owner:   ownerID,
		Content: contentUtf8,
		IP:      ip,
		MD5:     commentMD5,
		DBCSStr: origCommentDBCS,
		TheDate: theDateStr,
	}
	comment.CleanComment()

	return comment
}

// parseCommentDefaultOwnerIDStr
//
// fennyccc\x1b[m\x1b[33m: 讚                                    \x1b[m 115.82.149.161 01/28 12:34
//
// SYSOP\x1b[m\x1b[33m:推推                                                     \x1b[m 12/13 03:51
func parseCommentDefaultOwnerIDStr(commentDBCS string) (ownerID bbs.UUserID, nextCommentDBCS string) {
	if len(commentDBCS) == 0 {
		return "", ""
	}
	theIdx := strings.Index(commentDBCS, "\x1b")
	if theIdx == -1 {
		return "", ""
	}

	ownerID = bbs.UUserID(strings.TrimSpace(commentDBCS[:theIdx]))
	if len(commentDBCS) <= theIdx+len(MATCH_DEFAULT_INFIX_STR) {
		return ownerID, ""
	}
	nextCommentDBCS = commentDBCS[theIdx+len(MATCH_DEFAULT_INFIX_STR):]

	return ownerID, nextCommentDBCS
}

// parseCommentContentStr
//
// : 讚                                    \x1b[m 115.82.149.161 01/28 12:34
// :推推                                                     \x1b[m 12/13 03:51
func parseCommentContentStr(commmentDBCS string) (contentDBCS string, nextCommentDBCS string) {
	if len(commmentDBCS) == 0 {
		return "", ""
	}

	idx := strings.Index(commmentDBCS, "\x1b")
	if idx == -1 {
		return strings.TrimSpace(commmentDBCS[1:]), ""
	}

	contentDBCS, nextCommentDBCS = commmentDBCS[1:idx], commmentDBCS[idx:]
	contentDBCS = strings.TrimSpace(contentDBCS)
	if len(contentDBCS) == 0 {
		contentDBCS = ""
	}
	if len(nextCommentDBCS) == 0 {
		nextCommentDBCS = ""
	}
	idx = strings.Index(nextCommentDBCS, "m")
	if idx == -1 {
		nextCommentDBCS = ""
	}
	nextCommentDBCS = nextCommentDBCS[idx+1:]
	if len(nextCommentDBCS) == 0 {
		nextCommentDBCS = ""
	}

	return contentDBCS, nextCommentDBCS
}

// parseCommentDefaultIPCreateTimeStr
//
// Already separate the data by color.
// There are only ip/create-time information in p_commentDBCS.
// \x1b[m                        推  140.112.28.31 11/19
// \x1b[m                        推  				  11/19
// \x1b[m 12/13
// \x1b[m 118.160.112.18 03/20 09:44
func parseCommentDefaultIPCreateTimeStr(p_commentDBCS string) (ip string, dateStr string) {
	if len(p_commentDBCS) == 0 {
		return "", ""
	}
	_, after, ok := strings.Cut(p_commentDBCS, "推") // 推
	if ok {                                         // old
		postfix := strings.TrimSpace(after)
		postfixList := strings.Split(postfix, " ")
		if len(postfixList) != 2 { // assuming 推 11/19
			dateStr = strings.TrimSpace(postfix)
			dateStr = dateStr[:LEN_OLD_RECOMMEND_DATE]
			return "", dateStr
		}
		ip = postfixList[0]
		dateStr = postfixList[1]

		ip = strings.TrimSpace(ip)
		dateStr = strings.TrimSpace(dateStr)
		dateStr = dateStr[:LEN_OLD_RECOMMEND_DATE]

		return ip, dateStr
	}

	// new: MM/DD HH:mm
	ip = ""
	dateStr = strings.TrimSpace(dbcsToUtf8PurifyColor(p_commentDBCS))
	if len(dateStr) > LEN_RECOMMEND_DATE { // with ip
		theIdx := strings.Index(dateStr, " ")
		if theIdx != -1 {
			ip = dateStr[:theIdx]
			dateStr = dateStr[(theIdx + 1):]
		}
	}
	return ip, dateStr
}

// parseReply
//
// 只考慮parse
func parseReplyStr(replyDBCS string, editDBCS string, ownerID bbs.UUserID) (reply *schema.Comment) {
	if len(replyDBCS) == 0 {
		return nil
	}

	replyDBCS = strings.TrimSpace(replyDBCS)
	if len(replyDBCS) == 0 {
		return nil
	}

	// origReplyDBCS should exclude the last '\n'
	origReplyDBCSBytes := []byte(replyDBCS)
	replyMD5 := Md5sum(origReplyDBCSBytes)
	replyUtf8 := dbcsToUtf8(replyDBCS)

	editUserID, editNanoTS, editDateTimeStr, editIP, editHost := parseReplyUserIPHostStr(editDBCS)

	if editUserID == "" {
		editUserID = ownerID
	}

	reply = &schema.Comment{
		TheType: ptttype.COMMENT_TYPE_REPLY,
		Owner:   editUserID,
		Content: replyUtf8,
		IP:      editIP,
		Host:    editHost,
		MD5:     replyMD5,

		EditNanoTS: editNanoTS,

		DBCSStr: replyDBCS,
		TheDate: editDateTimeStr,
	}
	reply.CreateTime = editNanoTS

	reply.CleanReply()
	if len(reply.Content) == 0 {
		return nil
	}

	return reply
}

// parseReplyUserIPHost
// ※ 編輯: abcd (1.2.3.4 臺灣), 03/21/2021 03:04:47
//
// 1. 找到 EDIT_PREFIX
// 2. 找到 (, 設定 userID
// 3. 找到 ), 設定 ip/host
// 4. 設定時間.
func parseReplyUserIPHostStr(editDBCS string) (editUserID bbs.UUserID, editNanoTS types.NanoTS, editDateTimeStr string, editIP string, editHost string) {
	// 1.  get EDIT_PREFIX
	p_editDBCS := editDBCS
	theIdx := strings.Index(p_editDBCS, MATCH_COMMENT_EDIT_STR)
	if theIdx == -1 {
		return "", 0, "", "", ""
	}

	// 2. get (
	p_editDBCS = p_editDBCS[theIdx+len(MATCH_COMMENT_EDIT_STR):]

	theIdx = strings.Index(p_editDBCS, "(")
	if theIdx == -1 {
		return "", 0, "", "", ""
	}
	editUserID = bbs.UUserID(strings.TrimSpace(p_editDBCS[:theIdx]))

	// 3. get )
	p_editDBCS = p_editDBCS[theIdx+1:]

	theIdx = strings.Index(p_editDBCS, ")")
	if theIdx == -1 {
		return "", 0, "", "", ""
	}
	ipHost := p_editDBCS[:theIdx]

	ipHostList := strings.Split(ipHost, " ")
	if len(ipHostList) == 1 {
		editIP = ipHostList[0]
	} else {
		editIP = ipHostList[0]
		editHost = ipHostList[1]
	}

	// 4. get time.
	p_editDBCS = p_editDBCS[theIdx:]

	theIdx = strings.Index(p_editDBCS, ", ")
	p_editDBCS = p_editDBCS[theIdx+2:]

	editDateTimeStr = p_editDBCS[:types.LEN_DATE_YEAR_TIME_STR]
	editNanoTS = types.NanoTS(0)
	theTime, err := types.DateYearTimeStrToTime(editDateTimeStr)
	if err == nil {
		editNanoTS = types.TimeToNanoTS(theTime)
	}

	return editUserID, editNanoTS, editDateTimeStr, editIP, editHost
}
