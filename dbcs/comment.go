package dbcs

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

// ParseComments
//
// 有可能 reply-edit-info (編輯) 不在 commentsDBCS 裡
// 但是會在 allCommentsDBCS 裡 (firstComments)
// 只考慮:
//  1. appropriately split comments.
//  2. 對於每個 comment 裡的 DBCS Parse 成 Utf8.
//  3. type / IP / Host / MD5 / TheDate
//
// 不考慮:
//  1. boardID / articleID / commentID.
//  2. createTime / firstCreateTime / InferredCreateTime / AddCreateTime (除了編輯以外)
//
// steps:
//  1. 根據 '\n' 估計 nComments
//  2. 找出 pre-comment reply.
//  3. 對於每個 comment-leading newline for-loop:
//     3.0. parse comment
//     3.1. 找下一個 comment
//     3.1.1. 如果沒有更多 comment: 假設剩下 text 的都是 reply.
//     3.2. 假設下一個 comment 之前的 text 都是 reply.
//  4. (outside for-loop): 處理最後一個沒有 '\n' 的 comment.
func ParseComments(
	ownerID bbs.UUserID,
	commentsDBCS []byte,
	allCommentsDBCS []byte,
) (comments []*schema.Comment) {
	if len(commentsDBCS) == 0 {
		return nil
	}

	// 1. estimate nComments
	nEstimatedComments := bytes.Count(commentsDBCS, []byte{'\n'})

	comments = make([]*schema.Comment, 0, nEstimatedComments)

	p_commentsDBCS := commentsDBCS
	p_allCommentsDBCS := allCommentsDBCS

	var comment *schema.Comment

	// 2. reply
	nextCommentIdx := MatchComment(p_commentsDBCS)
	if nextCommentIdx > 0 || nextCommentIdx == -1 {
		replyDBCS := p_commentsDBCS
		if nextCommentIdx > 0 {
			replyDBCS = p_commentsDBCS[:nextCommentIdx]
		}
		reply := parseReply(replyDBCS, p_allCommentsDBCS, ownerID)
		if reply != nil {
			comments = append(comments, reply)
		}

		p_allCommentsDBCS = p_allCommentsDBCS[len(replyDBCS):]
		p_commentsDBCS = p_commentsDBCS[len(replyDBCS):]
	}

	// 3. for each comment-leading
	for idxNewLine := bytes.Index(p_commentsDBCS, []byte{'\n'}); len(p_commentsDBCS) > 0 && idxNewLine != -1; idxNewLine = bytes.Index(p_commentsDBCS, []byte{'\n'}) {

		// 3.0 parse comment
		commentDBCS := p_commentsDBCS[:idxNewLine]
		comment = parseComment(commentDBCS)
		comments = append(comments, comment)

		p_commentsDBCS = p_commentsDBCS[idxNewLine:] // with '\n'
		p_allCommentsDBCS = p_allCommentsDBCS[idxNewLine:]

		// 3.1 find next comment
		nextCommentIdx := MatchComment(p_commentsDBCS)

		if nextCommentIdx == -1 { // no more comments
			// 3.1.1 assume the rest are reply.
			p_commentsDBCS = p_commentsDBCS[1:] // step forward '\n'
			p_allCommentsDBCS = p_allCommentsDBCS[1:]
			if len(p_commentsDBCS) > 0 {
				replyDBCS := p_commentsDBCS
				reply := parseReply(replyDBCS, p_allCommentsDBCS, ownerID)
				if reply != nil {
					comments = append(comments, reply)
				}

				p_allCommentsDBCS = p_allCommentsDBCS[len(p_commentsDBCS):] //nolint // consistent with programming pattern
				p_commentsDBCS = nil
			}
			break
		}

		// 3.2 assume the context before the next comment is reply.
		if nextCommentIdx > 1 { // p_commentsDBCS[0] is '\n', get reply from p_commentsDBCS[1:]
			replyDBCS := p_commentsDBCS[1:nextCommentIdx]

			reply := parseReply(replyDBCS, p_allCommentsDBCS[1:], ownerID)
			if reply != nil {
				comments = append(comments, reply)
			}
		}

		p_commentsDBCS = p_commentsDBCS[nextCommentIdx:]
		p_allCommentsDBCS = p_allCommentsDBCS[nextCommentIdx:]

	}

	// 4. last comment without '\n'
	if len(p_commentsDBCS) > 0 {
		comment = parseComment(p_commentsDBCS)
		comments = append(comments, comment)
	}

	if len(comments) == 0 {
		comments = nil
	}

	return comments
}

func parseComment(commentDBCS []byte) (comment *schema.Comment) {
	theType, p_commentDBCS := MatchCommentType(commentDBCS)

	switch theType {
	case ptttype.COMMENT_TYPE_EDIT:
		return parseCommentEdit(p_commentDBCS, commentDBCS)
	case ptttype.COMMENT_TYPE_FORWARD:
		return parseCommentForward(p_commentDBCS, commentDBCS)
	case ptttype.COMMENT_TYPE_DELETED:
		return parseCommentDeleted(p_commentDBCS, commentDBCS)
	default:
		return parseCommentDefault(theType, p_commentDBCS, commentDBCS)
	}
}

func parseCommentEdit(commentDBCS []byte, origCommentDBCS []byte) (comment *schema.Comment) {
	ownerID, commentDBCS := parseCommentEditOwnerID(commentDBCS)
	ip, host, theDateStr := parseCommentEditIPCreateTime(commentDBCS)
	commentMD5 := Md5sum(origCommentDBCS)

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
		DBCS:    origCommentDBCS,
		TheDate: theDateStr,

		CreateTime:         createNanoTS,
		InferredCreateTime: createNanoTS,
		SortTime:           createNanoTS,

		CommentID: commentID,
	}

	return comment
}

func parseCommentEditOwnerID(commentDBCS []byte) (ownerID bbs.UUserID, newCommentDBCS []byte) {
	p_commentDBCS := commentDBCS

	before, after, ok := bytes.Cut(p_commentDBCS, []byte(" "))
	if !ok {
		return "", commentDBCS
	}

	ownerID = bbs.UUserID(string(before))

	return ownerID, after
}

func parseCommentEditIPCreateTime(commentDBCS []byte) (ip string, host string, theDate string) {
	if commentDBCS[0] != '(' { // old
		theIdx := bytes.Index(commentDBCS, MATCH_COMMENT_EDIT_FROM_BYTES)
		if theIdx == -1 {
			return "", "", ""
		}

		commentDBCS = commentDBCS[theIdx+len(MATCH_COMMENT_EDIT_FROM_BYTES):]

		// ip
		for idx, each := range commentDBCS {
			if each == ' ' || each == '(' {
				ip = string(commentDBCS[:idx])
				commentDBCS = commentDBCS[idx:]
				break
			}
		}

		// create-time
		theIdx = bytes.Index(commentDBCS, []byte("("))
		if theIdx == -1 {
			return ip, "", ""
		}

		commentDBCS = commentDBCS[theIdx+1:]

		theIdx = bytes.Index(commentDBCS, []byte(")"))
		theDate = string(commentDBCS[:theIdx])

		return ip, "", theDate
	}

	// ip
	theIdx := bytes.Index(commentDBCS, []byte(")"))
	ipList := bytes.Split(commentDBCS[1:theIdx], []byte(" "))
	ip = string(ipList[0])
	if len(ipList) == 2 {
		host = types.Big5ToUtf8(ipList[1])
	}

	theDate = strings.TrimSpace(string(commentDBCS[theIdx+3:])) //"), "

	return ip, host, theDate
}

func parseCommentForward(commentDBCS []byte, origCommentDBCS []byte) (comment *schema.Comment) {
	ownerID, commentDBCS := parseCommentForwardOwnerID(commentDBCS)
	contentDBCS, commentDBCS := parseCommentForwardContent(commentDBCS)
	contentBig5 := dbcsToBig5(contentDBCS) // the last 11 chars are the dates
	contentUtf8 := big5ToUtf8(contentBig5)
	ip, theDateStr := parseCommentForwardIPCreateTime(commentDBCS)
	commentMD5 := Md5sum(origCommentDBCS)

	comment = &schema.Comment{
		TheType: ptttype.COMMENT_TYPE_FORWARD,
		Owner:   ownerID,
		Content: contentUtf8,
		IP:      ip,
		MD5:     commentMD5,
		DBCS:    origCommentDBCS,
		TheDate: theDateStr,
	}
	comment.CleanComment()

	return comment
}

func parseCommentForwardOwnerID(commentDBCS []byte) (ownerID bbs.UUserID, newCommentDBCS []byte) {
	p_commentDBCS := commentDBCS
	idx := bytes.Index(p_commentDBCS, []byte{'\x1b'})
	if idx == -1 {
		return "", nil
	}

	ownerID = bbs.UUserID(strings.TrimSpace(string(p_commentDBCS[:idx])))

	return ownerID, p_commentDBCS[idx:]
}

func parseCommentForwardContent(commentDBCS []byte) (contentDBCS []byte, newCommentDBCS []byte) {
	if !bytes.HasPrefix(commentDBCS, MATCH_COMMENT_FORWARD_BYTES) {
		return nil, nil
	}

	p_commentDBCS := commentDBCS[len(MATCH_COMMENT_FORWARD_BYTES):]
	if bytes.Equal(p_commentDBCS[:len(MATCH_COMMENT_FORWARD_BOARD_BYTES)], MATCH_COMMENT_FORWARD_BOARD_BYTES) {
		p_commentDBCS = p_commentDBCS[len(MATCH_COMMENT_FORWARD_BOARD_BYTES):]
	}

	idx := bytes.Index(p_commentDBCS, []byte{'\x1b'})
	if idx == -1 {
		return nil, nil
	}

	contentDBCS = p_commentDBCS[:idx]
	newCommentDBCS = p_commentDBCS[idx:]

	return contentDBCS, newCommentDBCS
}

func parseCommentForwardIPCreateTime(commentDBCS []byte) (ip string, dateStr string) {
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

// parseCommentDeleted
//
// (teemocogs 刪除 teemocogs 的推文: 誤植)
// \x1b[1;30m(teemocogs \xa7R\xb0\xa3 teemocogs \xaa\xba\xb1\xc0\xa4\xe5: \xbb~\xb4\xd3)\x1b[m
func parseCommentDeleted(commentDBCS []byte, origCommentDBCS []byte) (comment *schema.Comment) {
	ownerID, commentDBCS := parseCommentDeletedOwnerID(commentDBCS)
	contentDBCS, _ := parseCommentDeletedContent(commentDBCS)
	contentBig5 := dbcsToBig5(contentDBCS) // the last 11 chars are the dates
	contentUtf8 := big5ToUtf8(contentBig5)
	commentMD5 := Md5sum(origCommentDBCS)

	comment = &schema.Comment{
		TheType: ptttype.COMMENT_TYPE_DELETED,
		Owner:   ownerID,
		Content: contentUtf8,
		MD5:     commentMD5,
		DBCS:    origCommentDBCS,
	}
	comment.CleanComment()

	return comment
}

func parseCommentDeletedOwnerID(commentDBCS []byte) (ownerID bbs.UUserID, newCommentDBCS []byte) {
	origCommentDBCS := commentDBCS
	startIdx := bytes.Index(commentDBCS, MATCH_COMMENT_DELETED_INFIX0)
	if startIdx == -1 {
		return "", commentDBCS
	}
	commentDBCS = commentDBCS[startIdx+len(MATCH_COMMENT_DELETED_INFIX0):]
	idx := bytes.Index(commentDBCS, []byte{' '})
	if idx == -1 {
		return "", origCommentDBCS
	}

	ownerID = bbs.UUserID(string(commentDBCS[:idx]))

	return ownerID, origCommentDBCS
}

func parseCommentDeletedContent(commentDBCS []byte) (contentDBCS []byte, newCommentDBCS []byte) {
	idx := bytes.Index(commentDBCS, MATCH_COMMENT_DELETED_POSTFIX)
	if idx == -1 {
		return commentDBCS, nil
	}

	return commentDBCS[:idx], commentDBCS[idx:]
}

func parseCommentDefault(theType ptttype.CommentType, commentDBCS []byte, origCommentDBCS []byte) (comment *schema.Comment) {
	ownerID, commentDBCS := parseCommentDefaultOwnerID(commentDBCS)

	contentDBCS, commentDBCS := parseCommentContent(commentDBCS)
	contentBig5 := dbcsToBig5(contentDBCS) // the last 11 chars are the dates
	contentUtf8 := big5ToUtf8(contentBig5)
	ip, theDateStr := parseCommentDefaultIPCreateTime(commentDBCS)
	commentMD5 := Md5sum(origCommentDBCS)

	comment = &schema.Comment{
		TheType: theType,
		Owner:   ownerID,
		Content: contentUtf8,
		IP:      ip,
		MD5:     commentMD5,
		DBCS:    origCommentDBCS,
		TheDate: theDateStr,
	}
	comment.CleanComment()

	return comment
}

func parseCommentDefaultOwnerID(p_commmentDBCS []byte) (ownerID bbs.UUserID, nextCommentDBCS []byte) {
	if len(p_commmentDBCS) == 0 {
		return "", nil
	}
	theIdx := bytes.Index(p_commmentDBCS, []byte{'\x1b'})
	if theIdx == -1 {
		return "", nil
	}

	ownerID = bbs.UUserID(strings.TrimSpace(string(p_commmentDBCS[:theIdx])))
	if len(p_commmentDBCS) <= theIdx+8 {
		return ownerID, nil
	}
	nextCommentDBCS = p_commmentDBCS[theIdx+8:]

	return ownerID, nextCommentDBCS
}

func parseCommentContent(p_commmentDBCS []byte) (contentDBCS []byte, nextCommentDBCS []byte) {
	if len(p_commmentDBCS) == 0 {
		return nil, nil
	}

	idx := bytes.Index(p_commmentDBCS, []byte{'\x1b'})
	if idx == -1 {
		return p_commmentDBCS[1:], nil
	}

	contentDBCS, nextCommentDBCS = p_commmentDBCS[1:idx], p_commmentDBCS[idx:]
	if len(contentDBCS) > 0 && contentDBCS[0] == ' ' {
		contentDBCS = contentDBCS[1:]
	}
	if len(contentDBCS) == 0 {
		contentDBCS = nil
	}
	if len(nextCommentDBCS) == 0 {
		nextCommentDBCS = nil
	}
	idx = bytes.Index(nextCommentDBCS, []byte{'m'})
	if idx == -1 {
		nextCommentDBCS = nil
	}
	nextCommentDBCS = nextCommentDBCS[idx+1:]
	if len(nextCommentDBCS) == 0 {
		nextCommentDBCS = nil
	}

	return contentDBCS, nextCommentDBCS
}

// parseCommentDefaultIPCreateTime
//
// Already separate the data by color.
// There are only ip/create-time information in p_commentDBCS.
func parseCommentDefaultIPCreateTime(p_commentDBCS []byte) (ip string, dateStr string) {
	if len(p_commentDBCS) == 0 {
		return "", ""
	}
	_, after, ok := bytes.Cut(p_commentDBCS, []byte("\xb1\xc0")) // 推
	if ok {                                                      // old
		postfix := strings.TrimSpace(types.Big5ToUtf8(after))
		postfixList := strings.Split(postfix, " ")
		if len(postfixList) != 2 { // unable to parse. return createTime + 10-millisecond
			return "", ""
		}
		ip = postfixList[0]
		dateStr = postfixList[1]
		if len(dateStr) > LEN_OLD_RECOMMEND_DATE {
			dateStr = dateStr[:LEN_OLD_RECOMMEND_DATE]
		}

		return ip, dateStr
	}

	// new: MM/DD HH:mm
	ip = ""
	dateStr = strings.TrimSpace(string(dbcsToBig5PurifyColor(p_commentDBCS)))
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
func parseReply(replyDBCS []byte, editDBCS []byte, ownerID bbs.UUserID) (reply *schema.Comment) {
	if len(replyDBCS) == 0 {
		return nil
	}

	if replyDBCS[len(replyDBCS)-1] == '\n' {
		replyDBCS = replyDBCS[:len(replyDBCS)-1]
	}
	if len(replyDBCS) == 0 {
		return nil
	}

	// origReplyDBCS should exclude the last '\n'
	origReplyDBCS := replyDBCS

	// remove '\r'
	if replyDBCS[len(replyDBCS)-1] == '\r' {
		replyDBCS = replyDBCS[:len(replyDBCS)-1]
	}
	if len(replyDBCS) == 0 {
		return nil
	}

	replyMD5 := Md5sum(origReplyDBCS)
	replyBig5 := dbcsToBig5(replyDBCS)
	replyUtf8 := big5ToUtf8(replyBig5)

	editUserID, editNanoTS, editDateTimeStr, editIP, editHost := parseReplyUserIPHost(editDBCS)

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

		DBCS:    origReplyDBCS,
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
func parseReplyUserIPHost(editDBCS []byte) (editUserID bbs.UUserID, editNanoTS types.NanoTS, editDateTimeStr string, editIP string, editHost string) {
	// 1.  get EDIT_PREFIX
	p_editDBCS := editDBCS
	theIdx := bytes.Index(p_editDBCS, MATCH_COMMENT_EDIT_BYTES)
	if theIdx == -1 {
		return "", 0, "", "", ""
	}

	// 2. get (
	p_editDBCS = p_editDBCS[theIdx+len(MATCH_COMMENT_EDIT_BYTES):]

	theIdx = bytes.Index(p_editDBCS, []byte{'('})
	if theIdx == -1 {
		return "", 0, "", "", ""
	}
	editUserID = bbs.UUserID(p_editDBCS[:theIdx-1])

	// 3. get )
	p_editDBCS = p_editDBCS[theIdx+1:]

	theIdx = bytes.Index(p_editDBCS, []byte{')'})
	if theIdx == -1 {
		return "", 0, "", "", ""
	}
	ipHost := types.Big5ToUtf8(p_editDBCS[:theIdx])

	ipHostList := strings.Split(ipHost, " ")
	if len(ipHostList) == 1 {
		editIP = ipHostList[0]
	} else {
		editIP = ipHostList[0]
		editHost = ipHostList[1]
	}

	// 4. get time.
	p_editDBCS = p_editDBCS[theIdx:]

	theIdx = bytes.Index(p_editDBCS, []byte(", "))
	p_editDBCS = p_editDBCS[theIdx+2:]

	editDateTimeStr = string(p_editDBCS[:19])
	editNanoTS = types.NanoTS(0)
	theTime, err := types.DateYearTimeStrToTime(editDateTimeStr)
	if err == nil {
		editNanoTS = types.TimeToNanoTS(theTime)
	}

	return editUserID, editNanoTS, editDateTimeStr, editIP, editHost
}

func CommentUtf8ToDBCS(c *schema.Comment) {
	contentBytes := bytes.Join(Utf8ToDBCS(c.Content), []byte("\n"))
	switch c.TheType {
	case ptttype.COMMENT_TYPE_RECOMMEND:
		c.DBCS = commentUtf8ToDBCSBasic(c, contentBytes, MATCH_COMMENT_RECOMMEND_BYTES)
	case ptttype.COMMENT_TYPE_BOO:
		c.DBCS = commentUtf8ToDBCSBasic(c, contentBytes, MATCH_COMMENT_BOO_BYTES)
	case ptttype.COMMENT_TYPE_COMMENT:
		c.DBCS = commentUtf8ToDBCSBasic(c, contentBytes, MATCH_COMMENT_ARROW_BYTES)
	case ptttype.COMMENT_TYPE_REPLY:
		c.DBCS = contentBytes
	case ptttype.COMMENT_TYPE_EDIT:
		commentBytes := make([]byte, 0, MAX_COMMENT_BYTES)
		commentBytes = append(commentBytes, MATCH_COMMENT_EDIT_BYTES...)
		createTime4 := c.CreateTime.ToTime4()
		theStr := fmt.Sprintf("%s (%s %s), %s", c.Owner, c.IP, c.Host, createTime4.Cdatelite())
		commentBytes = append(commentBytes, []byte(theStr)...)
		c.DBCS = commentBytes
	case ptttype.COMMENT_TYPE_FORWARD:
		commentBytes := make([]byte, 0, MAX_COMMENT_BYTES)
		commentBytes = append(commentBytes, MATCH_COMMENT_FORWARD_PREFIX...)
		commentBytes = append(commentBytes, []byte(c.Owner)...)
		commentBytes = append(commentBytes, MATCH_COMMENT_FORWARD_BYTES...)
		if bytes.Equal(contentBytes, MATCH_COMMENT_FORWARD_HIDDEN_BYTES) {
			commentBytes = append(commentBytes, contentBytes...)
		} else {
			commentBytes = append(commentBytes, MATCH_COMMENT_FORWARD_BOARD_BYTES...)
			commentBytes = append(commentBytes, contentBytes...)
		}
		c.DBCS = commentBytes
	case ptttype.COMMENT_TYPE_DELETED:
		commentBytes := make([]byte, 0, MAX_COMMENT_BYTES)
		commentBytes = append(commentBytes, MATCH_COMMENT_DELETED_PREFIX...)
		commentBytes = append(commentBytes, contentBytes...)
		commentBytes = append(commentBytes, MATCH_COMMENT_DELETED_POSTFIX...)
		c.DBCS = commentBytes
	}
}

func commentUtf8ToDBCSBasic(c *schema.Comment, contentBytes []byte, prefix []byte) (commentBytes []byte) {
	commentBytes = make([]byte, 0, MAX_COMMENT_BYTES)
	commentBytes = append(commentBytes, prefix...)
	commentBytes = append(commentBytes, []byte(c.Owner)...)
	commentBytes = append(commentBytes, MATCH_COMMENT_INFIX...)
	commentBytes = append(commentBytes, contentBytes...)

	return commentBytes
}
