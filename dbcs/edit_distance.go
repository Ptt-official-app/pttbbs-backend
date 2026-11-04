package dbcs

import (
	"slices"
	"strconv"

	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

type EDBlock struct {
	NewComments  []*EDInfo
	OrigComments []*EDInfo
	StartNanoTS  types.NanoTS
	EndNanoTS    types.NanoTS
}

type EDOp uint8

const (
	ED_OP_UNKNOWN EDOp = 0
	ED_OP_SAME    EDOp = 1
	ED_OP_DELETE  EDOp = 2
	ED_OP_ADD     EDOp = 3
)

type INFER_TIMESTAMP_TYPE uint8

const (
	INFER_TIMESTAMP_INVALID INFER_TIMESTAMP_TYPE = 0
	INFER_TIMESTAMP_YMDHM   INFER_TIMESTAMP_TYPE = 1
	INFER_TIMESTAMP_YMD     INFER_TIMESTAMP_TYPE = 2
)

type EDInfo struct {
	Op          EDOp
	NewComment  *schema.Comment // SAME/DELETE: origComments, ADD: newComments
	OrigComment *schema.CommentMD5
	SortTime    types.NanoTS
}

func NewEDInfoFromAddComment(comment *schema.Comment) (edInfo *EDInfo) {
	return &EDInfo{Op: ED_OP_ADD, NewComment: comment, SortTime: comment.SortTime}
}

func NewEDInfoFromDeleteComment(commentMD5 *schema.CommentMD5) (edInfo *EDInfo) {
	return &EDInfo{Op: ED_OP_DELETE, OrigComment: commentMD5, SortTime: commentMD5.SortTime}
}

func NewEDInfoFromSameComment(newComment *schema.Comment, origCommentMD5 *schema.CommentMD5) (edInfo *EDInfo) {
	return &EDInfo{Op: ED_OP_SAME, NewComment: newComment, OrigComment: origCommentMD5, SortTime: origCommentMD5.SortTime}
}

type EDInfoMeta struct {
	// EDInfoMeta

	// StartNanoTS (not included)
	StartNanoTS types.NanoTS

	// EndNanoTS (not included except the last ed-info)
	EndNanoTS types.NanoTS

	// StartIdx (included)
	StartIdx int

	// EndIdx (not incldued)
	EndIdx int
}

// CalcEDBlocks
//
// Must already guarantee that:
// 1. articleCreateTime < all origComments.SortTime
// 2. articleMTime >= all origComments.SortTime
// 3. origComments are sorted by SortTime
// 4. newComments are sorted by the line-idx.
func CalcEDBlocks(newComments []*schema.Comment, origComments []*schema.CommentMD5, articleCreateTime types.NanoTS, articleMTime types.NanoTS) (edBlocks []*EDBlock, err error) {
	if len(newComments) == 0 {
		return nil, nil
	}

	if len(origComments) == 0 {
		newEDInfos := make([]*EDInfo, len(newComments))
		for idx, each := range newComments {
			newEDInfos[idx] = NewEDInfoFromAddComment(each)
		}
		return []*EDBlock{{NewComments: newEDInfos, StartNanoTS: articleCreateTime, EndNanoTS: articleMTime}}, nil
	}

	// 1. do edit-distance

	scoreMat := calcEditDistance(newComments, origComments)

	edInfos := calcEditDistanceBacktracking(scoreMat, newComments, origComments)

	edInfoMetas := calcEDInfoMetas(edInfos, articleCreateTime, articleMTime)

	edBlocks = make([]*EDBlock, len(edInfoMetas))

	for idx, each := range edInfoMetas {
		edBlock := each.ToEDBlock(edInfos)
		edBlocks[idx] = edBlock
	}

	return edBlocks, nil
}

// calcEditDistance
//
// input: abcde (new) ABCDE (orig)
// scoreMat:  _ A B C D E
//
//	_ 0
//	a
//	b
//	c
//	d
//	e
//
// scoreMat[i][j] = min(scoreMat[i-1][j]+1, scoreMat[i][j-1]+1, (if same MD5: scoreMat[i-1][j-1]))
func calcEditDistance(comments []*schema.Comment, allCommentMD5s []*schema.CommentMD5) (scoreMat [][]int) {
	lenNew := len(comments)
	lenOrig := len(allCommentMD5s)

	if lenNew == 0 || lenOrig == 0 {
		return nil
	}

	// init scoreMat
	scoreMat = make([][]int, lenNew+1)
	for idx := 0; idx <= lenNew; idx++ {
		scoreMat[idx] = make([]int, lenOrig+1)
	}

	// 1. init
	for idx := 1; idx <= lenNew; idx++ {
		scoreMat[idx][0] = scoreMat[idx-1][0] + 1
	}

	for idx := 1; idx <= lenOrig; idx++ {
		scoreMat[0][idx] = scoreMat[0][idx-1] + 1
	}

	// 2. dp.
	for idx := 1; idx <= lenNew; idx++ {
		for idx2 := 1; idx2 <= lenOrig; idx2++ {
			minScore := min(scoreMat[idx][idx2-1]+1, scoreMat[idx-1][idx2]+1)

			if comments[idx-1].MD5 == allCommentMD5s[idx2-1].MD5 && scoreMat[idx-1][idx2-1] < minScore {
				minScore = scoreMat[idx-1][idx2-1]
			}

			scoreMat[idx][idx2] = minScore
		}
	}

	return scoreMat
}

// calcEditDistanceBacktracking
//
// # params
//
// * scoreMat: edit-distance score matrix. row: new comments, col: orig-comments
// * comments: new comments sorted by line-idx
// * allCommentMD5s: orig-comments sorted by SortTime
//
// # return
//
// * edInfos: The back-tracking result.
func calcEditDistanceBacktracking(scoreMat [][]int, comments []*schema.Comment, allCommentMD5s []*schema.CommentMD5) (edInfos []*EDInfo) {
	currentNew := len(scoreMat) - 1
	currentOrig := len(scoreMat[0]) - 1

	edInfos = make([]*EDInfo, 0, len(scoreMat)+len(scoreMat[0]))
	for idx := 0; idx <= len(scoreMat)+len(scoreMat[0]); idx++ {
		if currentNew == 0 && currentOrig == 0 {
			break
		}
		currentScore := scoreMat[currentNew][currentOrig]

		if currentNew > 0 && scoreMat[currentNew-1][currentOrig] == currentScore-1 {
			edInfo := NewEDInfoFromAddComment(comments[currentNew-1])
			edInfos = append(edInfos, edInfo)
			currentNew--

			continue
		}

		if currentOrig > 0 && scoreMat[currentNew][currentOrig-1] == currentScore-1 {
			edInfo := NewEDInfoFromDeleteComment(allCommentMD5s[currentOrig-1])
			edInfos = append(edInfos, edInfo)
			currentOrig--

			continue
		}

		if currentOrig > 0 && currentNew > 0 && scoreMat[currentNew-1][currentOrig-1] == currentScore && comments[currentNew-1].MD5 == allCommentMD5s[currentOrig-1].MD5 {

			edInfo := NewEDInfoFromSameComment(comments[currentNew-1], allCommentMD5s[currentOrig-1])
			edInfos = append(edInfos, edInfo)
			currentNew--
			currentOrig--

			continue
		}

	}

	edInfos = reverseEDInfos(edInfos)

	return edInfos
}

func reverseEDInfos(edInfos []*EDInfo) []*EDInfo {
	slices.Reverse(edInfos)
	return edInfos
}

// calcEDInfoMetas
//
// Given the edInfos list, we would like to construct the EDInfoMetas
// by separating the list with EDIT_DISTANCE_OP_SAME
func calcEDInfoMetas(edInfos []*EDInfo, startNanoTS types.NanoTS, endNanoTS types.NanoTS) (edInfoMetas []*EDInfoMeta) {
	// compute number of op-same
	nEDInfoMetas := 1 // we need at least 1 meta
	for _, each := range edInfos {
		if each.Op == ED_OP_SAME {
			nEDInfoMetas++
		}
	}

	edInfoMetas = make([]*EDInfoMeta, 0, nEDInfoMetas)
	currentStartNanoTS := startNanoTS
	currentStartIdx := 0
	for idx, each := range edInfos {
		if each.Op == ED_OP_SAME {
			if currentStartIdx == idx {
				currentStartNanoTS = each.OrigComment.SortTime
				currentStartIdx = idx + 1
				continue
			}
			eachMeta := &EDInfoMeta{StartNanoTS: currentStartNanoTS, EndNanoTS: each.OrigComment.SortTime, StartIdx: currentStartIdx, EndIdx: idx}
			edInfoMetas = append(edInfoMetas, eachMeta)

			currentStartNanoTS = each.OrigComment.SortTime
			currentStartIdx = idx + 1
		}
	}

	if currentStartIdx != len(edInfos) {
		theLastMeta := &EDInfoMeta{StartNanoTS: currentStartNanoTS, EndNanoTS: endNanoTS, StartIdx: currentStartIdx, EndIdx: len(edInfos)}
		edInfoMetas = append(edInfoMetas, theLastMeta)
	}

	return edInfoMetas
}

// ToEDBlock
//
// Given the list of edInfos, where NewComments are OrigComments are already separately sorted,
// construct the corresponding ed-block.
func (meta *EDInfoMeta) ToEDBlock(edInfos []*EDInfo) (edBlock *EDBlock) {
	theEDInfos := edInfos[meta.StartIdx:meta.EndIdx]

	// 1. get nNewComments and nOrigComments
	nNewComments := 0
	nOrigComments := 0
	for _, each := range theEDInfos {
		op := each.Op
		switch op {
		case ED_OP_ADD:
			nNewComments++
		case ED_OP_DELETE:
			nOrigComments++
		}
	}

	// 2. construct edblock
	edBlock = &EDBlock{
		StartNanoTS:  meta.StartNanoTS,
		EndNanoTS:    meta.EndNanoTS,
		NewComments:  make([]*EDInfo, 0, nNewComments),
		OrigComments: make([]*EDInfo, 0, nOrigComments),
	}

	for _, each := range theEDInfos {
		op := each.Op
		switch op {
		case ED_OP_ADD:
			edBlock.NewComments = append(edBlock.NewComments, each)
		case ED_OP_DELETE:
			edBlock.OrigComments = append(edBlock.OrigComments, each)
		}
	}

	if len(edBlock.NewComments) == 0 {
		edBlock.NewComments = nil
	}
	if len(edBlock.OrigComments) == 0 {
		edBlock.OrigComments = nil
	}

	return edBlock
}

func InferTimestamp(edBlocks []*EDBlock, isForwardOnly bool, isLastAlignEndNanoTS bool, articleCreateTime types.NanoTS) (nBlock int) {
	if len(edBlocks) == 0 {
		return
	}

	nBlock = len(edBlocks)
	for idx, each := range edBlocks {
		isAlignEndNanoTS := isLastAlignEndNanoTS && (idx == len(edBlocks)-1)
		lenBeforeComments := len(each.NewComments)
		if lenBeforeComments == 0 {
			continue
		}
		each.InferTimestamp(articleCreateTime, isForwardOnly, isAlignEndNanoTS)
		lenAfterComments := len(each.NewComments)
		if lenAfterComments == 0 {
			return idx
		}
		if lenAfterComments < lenBeforeComments {
			return idx + 1
		}
	}
	return nBlock
}

// InferTimestamp
//
//  1. OrigComments are sorted between ed.StartNanoTS and ed.EndNanoTS
//  2. It's possible that the newComments are with out-of-range time.
//  3. It's possible that multiple comments shares the same date-str,
//     but we still need some way to make the timestamp unique.
//  4. The time from OrigComments should not be moved.
//
// The possibilities that new-comments are in between original-comments:
// XXX 1. delete (try to map the corresponding deleted messages)
//
//	We don't do this to simplify mapping sequence.
//
// 2. reply (previous-appearing-message (currentNanoTS in same or newComments) + REPLY_STEP_NANO_TS)
// 3. new messages. (sort-time should be after the deleted-messages)
// 4. others (the owners accidentally edited something, sort-time should be after the deleted-messages)
func (ed *EDBlock) InferTimestamp(articleCreateTime types.NanoTS, isForwardOnly bool, isLastAlignEndNanoTS bool) {
	// no need to infer timestamp
	if len(ed.NewComments) == 0 {
		return
	}

	// 1. map deleted
	// ed.MapDeletedMessages()

	// 2. forward sort (with reply)
	nextIdx := ed.ForwardInferTS(articleCreateTime)
	if isForwardOnly {
		isLastAlignEndNanoTS = isLastAlignEndNanoTS && (nextIdx == len(ed.NewComments))
		ed.NewComments = ed.NewComments[:nextIdx]
		if isLastAlignEndNanoTS {
			ed.AlignEndNanoTS()
		}
		return
	}

	// 3. backword-sort (with-reply)
	ed.BackwardInferTS(nextIdx, isLastAlignEndNanoTS)
}

func (ed *EDBlock) AlignEndNanoTS() {
	newComments := ed.NewComments
	if len(newComments) == 0 {
		return
	}

	lastComment := newComments[len(newComments)-1].NewComment
	if lastComment.TheType > ptttype.COMMENT_TYPE_BASIC_DATE {
		return
	}

	year := ed.EndNanoTS.ToTime().Year()
	createTime, inferType := inferTSWithYear(lastComment.TheDate, year)

	theDiff := ed.EndNanoTS - createTime
	inferDiff := COMMENT_DIFF_ALIGN_END_NANO_TS
	if inferType == INFER_TIMESTAMP_YMD {
		inferDiff = COMMENT_DIFF2_ALIGN_END_NANO_TS
	}
	if theDiff > -COMMENT_DIFF_ALIGN_END_NANO_TS && theDiff < inferDiff {
		createTime = ed.EndNanoTS
	}

	sortTime := createTime
	if createTime > ed.EndNanoTS {
		sortTime = ed.EndNanoTS
		createTime = lastComment.CreateTime
	}

	lastComment.CreateTime = createTime
	lastComment.SetSortTime(sortTime)
}

// ForwardInferTS
func (ed *EDBlock) ForwardInferTS(articleCreateTime types.NanoTS) (nextIdx int) {
	newComments := ed.NewComments
	preNanoTS := ed.StartNanoTS
	nextIdx = len(newComments)
	var preComment *schema.Comment

	for idx, each := range newComments {

		newComment := each.NewComment

		// 1. already setup, and within expectation.
		if newComment.SortTime != 0 {
			if newComment.SortTime > preNanoTS && newComment.SortTime < ed.EndNanoTS {
				preNanoTS = newComment.SortTime
				continue
			}
		}

		// 2. with reply, but unable to infer the timestamp from edit.
		if newComment.TheType == ptttype.COMMENT_TYPE_REPLY {
			// 2.1 (assuming that currentNanoTS refers to the corresponding content/comment of the reply)
			theNanoTS := preNanoTS + REPLY_STEP_NANO_TS
			if newComment.CreateTime == 0 {
				newComment.CreateTime = theNanoTS
			}

			// 2.2 dealing with exceedingEndNanoTS
			if theNanoTS >= ed.EndNanoTS {
				theNanoTS = forwardExceedingEndNanoTS(idx, len(newComments), preNanoTS, ed.EndNanoTS)
			}

			// 2.3 set reply id.
			if preComment != nil {
				newComment.CommentID = types.ToReplyID(preComment.CommentID)
			}

			// 2.4 set sort-time.
			newComment.SetSortTime(theNanoTS)
			preNanoTS = newComment.SortTime
			preComment = newComment
			continue
		}

		sortNanoTS := newComment.SortTime
		createNanoTS := newComment.CreateTime
		currentYear := preNanoTS.ToTime().Year()

		// 3. if sortNanoTS <= preNanoTS
		if sortNanoTS <= preNanoTS {
			sortNanoTS, createNanoTS = forwardInferTSWithYear(newComment.TheDate, currentYear, preNanoTS)
		}

		// 4. if sortNanoTS > endNanoTS
		if sortNanoTS >= ed.EndNanoTS {
			sortNanoTS = forwardExceedingEndNanoTS(idx, len(newComments), preNanoTS, ed.EndNanoTS)
		}

		// 5. if sortNanoTS is 1-year after article-create-time
		if sortNanoTS >= articleCreateTime+types.YEAR_NANO_TS {
			nextIdx = idx
			break
		}

		// 6. set sort-time.
		newComment.SetSortTime(sortNanoTS)
		if newComment.CreateTime == 0 {
			newComment.CreateTime = createNanoTS
		}

		preNanoTS = newComment.SortTime
		preComment = newComment
	}

	return nextIdx
}

func forwardExceedingEndNanoTS(idx int, total int, currentNanoTS types.NanoTS, endNanoTS types.NanoTS) (nanoTS types.NanoTS) {
	nanoTS = endNanoTS - types.NanoTS(total-idx)*COMMENT_EXCEED_NANO_TS
	if nanoTS < currentNanoTS {
		nanoTS = (currentNanoTS + endNanoTS) / 2
	}
	return nanoTS
}

// BackwardInferTS
func (ed *EDBlock) BackwardInferTS(nextIdx int, isAlignEndNanoTS bool) {
	newComments := ed.NewComments
	startNanoTS := ed.StartNanoTS
	if nextIdx > 0 {
		startNanoTS = newComments[nextIdx-1].NewComment.SortTime
	}
	currentNanoTS := ed.EndNanoTS

	for idx := len(newComments) - 1; idx >= nextIdx; idx-- {
		each := newComments[idx]

		newComment := each.NewComment

		if newComment.SortTime != 0 {
			if newComment.SortTime > currentNanoTS && newComment.SortTime < ed.EndNanoTS {
				currentNanoTS = newComment.SortTime
				continue
			}
		}

		if newComment.TheType == ptttype.COMMENT_TYPE_REPLY { // do not do reply until the last
			continue
		}

		sortNanoTS := newComment.SortTime
		createNanoTS := newComment.CreateTime
		currentYear := currentNanoTS.ToTime().Year()
		if sortNanoTS == 0 || sortNanoTS >= currentNanoTS {
			sortNanoTS, createNanoTS = backwardInferTSWithYear(newComment.TheDate, currentYear, currentNanoTS)
		}

		if sortNanoTS <= startNanoTS {
			sortNanoTS = backwardExceedingStartNanoTS(idx, len(newComments), currentNanoTS, startNanoTS)
		}

		newComment.SetSortTime(sortNanoTS)
		if newComment.CreateTime == 0 {
			newComment.CreateTime = createNanoTS
		}

		currentNanoTS = newComment.SortTime
	}

	// deal with reply
	currentNanoTS = startNanoTS
	var preComment *schema.Comment
	for idx := nextIdx; idx < len(newComments); idx++ {
		newComment := newComments[idx].NewComment
		if newComment.TheType != ptttype.COMMENT_TYPE_REPLY {
			currentNanoTS = newComment.SortTime
			preComment = newComment
			continue
		}
		theNanoTS := currentNanoTS + REPLY_STEP_NANO_TS
		if newComment.CreateTime == 0 {
			newComment.CreateTime = theNanoTS
		}

		endNanoTS := ed.EndNanoTS
		if idx != len(newComments)-1 {
			endNanoTS = newComments[idx+1].NewComment.SortTime
		}

		if theNanoTS >= endNanoTS {
			theNanoTS = forwardExceedingEndNanoTS(idx, idx+1, currentNanoTS, endNanoTS)
		}

		if preComment != nil {
			newComment.CommentID = types.ToReplyID(preComment.CommentID)
		} else {
			newComment.CommentID = ""
		}
		newComment.SetSortTime(theNanoTS)

		currentNanoTS = newComment.SortTime
		preComment = newComment
	}

	if isAlignEndNanoTS {
		ed.AlignEndNanoTS()
	}
}

func backwardExceedingStartNanoTS(idx int, total int, currentNanoTS types.NanoTS, startNanoTS types.NanoTS) (nanoTS types.NanoTS) {
	nanoTS = startNanoTS + types.NanoTS(idx+1)*COMMENT_EXCEED_NANO_TS
	if nanoTS >= currentNanoTS {
		nanoTS = (currentNanoTS + startNanoTS) / 2
	}
	return nanoTS
}

// forwardInferTS
func forwardInferTSWithYear(theDate string, year int, startNanoTS types.NanoTS) (sortNanoTS types.NanoTS, createNanoTS types.NanoTS) {
	// 1. infer TS
	sortNanoTS, inferType := inferTSWithYear(theDate, year)
	createNanoTS = sortNanoTS

	// 2. if sortNanoTS < startNanoTS
	if sortNanoTS <= startNanoTS {
		// 2.1. setup infer-diff
		inferDiff := COMMENT_STEP_DIFF_NANO_TS
		if inferType == INFER_TIMESTAMP_YMD {
			inferDiff = COMMENT_STEP_DIFF2_NANO_TS
		}

		// 2.2. if the differences is < inferDiff (very close to the startNanoTS / endNanoTS):
		if startNanoTS-sortNanoTS < inferDiff {
			sortNanoTS = startNanoTS + COMMENT_STEP_NANO_TS
		} else {
			// 2.3. else: add 1 more year.
			sortNanoTS, _ = inferTSWithYear(theDate, year+1)
			createNanoTS = sortNanoTS
		}
	}
	if sortNanoTS <= startNanoTS {
		sortNanoTS = startNanoTS + COMMENT_STEP_NANO_TS
	}
	if createNanoTS == 0 {
		createNanoTS = sortNanoTS
	}

	return sortNanoTS, createNanoTS
}

// backwardInferTS
func backwardInferTSWithYear(theDate string, year int, endNanoTS types.NanoTS) (sortNanoTS types.NanoTS, createNanoTS types.NanoTS) {
	sortNanoTS, inferType := inferTSWithYear(theDate, year)
	createNanoTS = sortNanoTS
	sortNanoTS += COMMENT_BACKWARD_OFFSET_NANO_TS

	if sortNanoTS >= endNanoTS {
		inferDiff := COMMENT_STEP_DIFF_NANO_TS
		if inferType == INFER_TIMESTAMP_YMD {
			inferDiff = COMMENT_STEP_DIFF2_NANO_TS
		}
		if sortNanoTS-endNanoTS < inferDiff {
			sortNanoTS = endNanoTS - COMMENT_STEP_NANO_TS
		} else {
			sortNanoTS, _ = inferTSWithYear(theDate, year-1)
			createNanoTS = sortNanoTS
			sortNanoTS += COMMENT_BACKWARD_OFFSET_NANO_TS
		}
	}

	// unable to infer sortNanoTS
	if sortNanoTS == COMMENT_BACKWARD_OFFSET_NANO_TS || sortNanoTS >= endNanoTS {
		sortNanoTS = endNanoTS - COMMENT_STEP_NANO_TS
	}
	if createNanoTS == 0 {
		createNanoTS = sortNanoTS
	}

	return sortNanoTS, createNanoTS
}

// inferTSWithYear
//
// theDate: possibly MM/DD or MM/DD hh:mm
// alg: 1. add as YYYY/MM/DD or YYYY/MM/DD hh:mm
func inferTSWithYear(theDate string, year int) (nanoTS types.NanoTS, theType INFER_TIMESTAMP_TYPE) {
	theDateWithYear := strconv.Itoa(year) + "/" + theDate

	// 1. try YYYY/MM/DD hh:mm
	theTime, err := types.DateMinStrToTime(theDateWithYear)
	if err == nil {
		return types.TimeToNanoTS(theTime), INFER_TIMESTAMP_YMDHM
	}

	// 2. try YYYY/MM/DD
	theTime, err = types.DateStrToTime(theDateWithYear)
	if err == nil {
		return types.TimeToNanoTS(theTime), INFER_TIMESTAMP_YMD
	}

	return 0, INFER_TIMESTAMP_INVALID
}
