package schema

import (
	"errors"
	"strings"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ArticleSummary
type ArticleSummary struct {
	// ArticleSummary
	BBoardID       bbs.BBoardID         `bson:"bid"`
	ArticleID      bbs.ArticleID        `bson:"aid"`
	BoardArticleID types.BoardArticleID `bson:"baid"`

	IsDeleted  bool         `bson:"deleted,omitempty"`
	CreateTime types.NanoTS `bson:"create_time_nano_ts"`
	MTime      types.NanoTS `bson:"mtime_nano_ts"`

	Recommend int              `bson:"recommend"`
	Owner     bbs.UUserID      `bson:"owner"`
	FullTitle string           `bson:"full_title"`
	Title     string           `bson:"title"`
	Money     int              `bson:"money"`
	Class     string           `bson:"class"`
	Filemode  ptttype.FileMode `bson:"mode"`

	Idx string `bson:"pttidx"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`

	NComments int `bson:"n_comments,omitempty"` // n_comments is read-only in article-summary.

	Rank int `bson:"rank,omitempty"` // 評價, read-only

	SubjectType ptttype.SubjectType `bson:"subject_type"`

	IsBottom bool `bson:"is_bottom"`
}

var (
	EMPTY_ARTICLE_SUMMARY = &ArticleSummary{}
	articleSummaryFields  = getFields(EMPTY_ARTICLE, EMPTY_ARTICLE_SUMMARY)
)

func GetArticleSummary(bboardID bbs.BBoardID, articleID bbs.ArticleID) (result *ArticleSummary, err error) {
	query := &ArticleQuery{
		BBoardID:  bboardID,
		ArticleID: articleID,
	}

	result = &ArticleSummary{}
	err = Article_c.FindOne(query, &result, articleSummaryFields)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetArticleSummariesByBoardArticleIDs(boardArticleIDs []types.BoardArticleID) (result []*ArticleSummary, err error) {
	query := bson.M{
		ARTICLE_BOARD_ARTICLE_ID_b: bson.M{
			"$in": boardArticleIDs,
		},
	}

	// find
	err = Article_c.Find(query, 0, &result, articleSummaryFields, nil)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetArticleSummariesByOwnerID(ownerID bbs.UUserID, startCreateTime types.NanoTS, descending bool, limit int) (result []*ArticleSummary, err error) {
	// setup query
	var query bson.M
	if startCreateTime == 0 {
		query = bson.M{
			ARTICLE_OWNER_b: ownerID,
			ARTICLE_IS_DELETED_b: bson.M{
				"$exists": false,
			},
		}
	} else {
		theDir := "$gte"
		if descending {
			theDir = "$lte"
		}

		query = bson.M{
			ARTICLE_OWNER_b: ownerID,
			ARTICLE_CREATE_TIME_b: bson.M{
				theDir: startCreateTime,
			},
			ARTICLE_IS_DELETED_b: bson.M{
				"$exists": false,
			},
		}
	}
	// sort opts
	var sortOpts bson.D
	if descending {
		sortOpts = bson.D{
			{Key: ARTICLE_CREATE_TIME_b, Value: -1},
		}
	} else {
		sortOpts = bson.D{
			{Key: ARTICLE_CREATE_TIME_b, Value: 1},
		}
	}

	// find
	err = Article_c.Find(query, int64(limit), &result, articleSummaryFields, sortOpts)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetBottomArticleSummaries(boardID bbs.BBoardID) (result []*ArticleSummary, err error) {
	query := bson.M{
		ARTICLE_BBOARD_ID_b: boardID,
		ARTICLE_IS_BOTTOM_b: true,
	}
	sortOpts := bson.D{
		{Key: ARTICLE_CREATE_TIME_b, Value: 1},
		{Key: ARTICLE_ARTICLE_ID_b, Value: 1},
	}
	err = Article_c.Find(query, 0, &result, articleSummaryFields, sortOpts)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetArticleSummaries(boardID bbs.BBoardID, startIdx string, descending bool, limit int) (result []*ArticleSummary, err error) {
	var query bson.M
	if startIdx == "" {
		query = bson.M{
			ARTICLE_BBOARD_ID_b: boardID,
		}
	} else {
		theDir := "$gte"
		if descending {
			theDir = "$lte"
		}
		query = bson.M{
			ARTICLE_BBOARD_ID_b: boardID,
			ARTICLE_IDX_b: bson.M{
				theDir: startIdx,
			},
		}
	}
	query[ARTICLE_IS_DELETED_b] = bson.M{
		"$exists": false,
	}

	// sort opts
	var sortOpts bson.D
	if descending {
		sortOpts = bson.D{
			{Key: ARTICLE_IDX_b, Value: -1},
		}
	} else {
		sortOpts = bson.D{
			{Key: ARTICLE_IDX_b, Value: 1},
		}
	}

	// find
	err = Article_c.Find(query, int64(limit), &result, articleSummaryFields, sortOpts)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetArticleSummariesByRegex(boardID bbs.BBoardID, keywordList []string, createNanoTS types.NanoTS, articleID bbs.ArticleID, descending bool, limit int) (result []*ArticleSummary, err error) {
	result = make([]*ArticleSummary, 0, limit)

	// keyword-list
	runeList := make([][]rune, len(keywordList))
	for idx, each := range keywordList {
		runeList[idx] = []rune(each)
	}
	isToValidate := getArticleSummariesByRegexIsToValidate(runeList)

	firstPattern, patternList := getArticleSummariesByRegexPatternList(boardID, runeList)

	query := getArticleSummariesByRegexSetQuery(boardID, firstPattern, patternList, createNanoTS, articleID, descending)

	// sort opts
	var sortOpts bson.D
	if descending {
		sortOpts = bson.D{
			{Key: ARTICLE_CREATE_TIME_b, Value: -1},
			{Key: ARTICLE_ARTICLE_ID_b, Value: -1},
		}
	} else {
		sortOpts = bson.D{
			{Key: ARTICLE_CREATE_TIME_b, Value: 1},
			{Key: ARTICLE_ARTICLE_ID_b, Value: 1},
		}
	}

	nextCreateTS := createNanoTS
	nextArticleID := articleID

	isEndLoop := false
	remaining := limit
	for !isEndLoop && remaining > 0 {
		// find
		var eachResult []*ArticleSummary
		err = Article_c.Find(query, int64(limit+1), &eachResult, articleSummaryFields, sortOpts)
		if err != nil {
			return nil, err
		}
		if len(eachResult) < limit+1 {
			isEndLoop = true
		} else {
			nextArticleSummary := eachResult[limit]
			eachResult = eachResult[:limit]

			nextCreateTS = nextArticleSummary.CreateTime
			nextArticleID = nextArticleSummary.ArticleID
		}

		validResult := eachResult
		if isToValidate {
			validResult = getArticleSummariesByRegexIsValidResult(eachResult, keywordList, runeList)
		}

		// append
		if len(validResult) > remaining {
			validResult = validResult[:remaining]
		}

		result = append(result, validResult...)
		remaining -= len(validResult)

		query = getArticleSummariesByRegexSetQuery(boardID, firstPattern, patternList, nextCreateTS, nextArticleID, descending)
	}

	return result, nil
}

func getArticleSummariesByRegexSetQuery(boardID bbs.BBoardID, firstPattern string, patternList bson.A, createNanoTS types.NanoTS, articleID bbs.ArticleID, descending bool) (query bson.M) {
	if createNanoTS == 0 {
		firstQuery := bson.M{
			ARTICLE_BBOARD_ID_b:   boardID,
			ARTICLE_TITLE_REGEX_b: firstPattern,
			ARTICLE_IS_DELETED_b: bson.M{
				"$exists": false,
			},
		}

		theList := make(bson.A, 0, len(patternList)+1)
		theList = append(theList, firstQuery)
		theList = append(theList, patternList...)

		return bson.M{
			"$and": theList,
		}
	}

	theDir := "$gt"
	if descending {
		theDir = "$lt"
	}

	firstQuery := bson.M{
		ARTICLE_BBOARD_ID_b:   boardID,
		ARTICLE_TITLE_REGEX_b: firstPattern,
		ARTICLE_CREATE_TIME_b: bson.M{
			theDir: createNanoTS,
		},
		ARTICLE_IS_DELETED_b: bson.M{
			"$exists": false,
		},
	}

	theList := make(bson.A, 0, len(patternList)+1)
	theList = append(theList, firstQuery)
	theList = append(theList, patternList...)

	theDir2 := "$gte"
	if descending {
		theDir2 = "$lte"
	}

	secondQuery := bson.M{
		ARTICLE_BBOARD_ID_b:   boardID,
		ARTICLE_TITLE_REGEX_b: firstPattern,
		ARTICLE_CREATE_TIME_b: createNanoTS,
		ARTICLE_ARTICLE_ID_b: bson.M{
			theDir2: articleID,
		},
		ARTICLE_IS_DELETED_b: bson.M{
			"$exists": false,
		},
	}

	theList2 := make(bson.A, 0, len(patternList)+1)
	theList2 = append(theList2, secondQuery)
	theList2 = append(theList2, patternList...)

	return bson.M{
		"$or": bson.A{
			bson.M{
				"$and": theList,
			},
			bson.M{
				"$and": theList2,
			},
		},
	}
}

func getArticleSummariesByRegexIsToValidate(keywordList [][]rune) (isToValidate bool) {
	for _, each := range keywordList {
		if len(each) > TITLE_REGEX_N_GRAM {
			return true
		}
	}

	return false
}

func getArticleSummariesByRegexPatternList(boardID bbs.BBoardID, keywordList [][]rune) (firstPattern string, patternList bson.A) {
	if len(keywordList) == 0 {
		return "", nil
	}

	lenFirst := min(len(keywordList[0]), TITLE_REGEX_N_GRAM)
	firstPattern = string(keywordList[0][:lenFirst])

	patternList = make(bson.A, 0, len(keywordList)-1)
	for _, each := range keywordList[1:] {
		lenEach := min(len(each), TITLE_REGEX_N_GRAM)

		patternList = append(patternList, bson.M{ARTICLE_BBOARD_ID_b: boardID, ARTICLE_TITLE_REGEX_b: string(each[:lenEach])})
	}

	return firstPattern, patternList
}

// getArticleSummariesByRegexIsValidResult
//
// Assume:
func getArticleSummariesByRegexIsValidResult(articleSummaries []*ArticleSummary, keywordList []string, runeList [][]rune) (validArticleSummaries []*ArticleSummary) {
	validArticleSummaries = make([]*ArticleSummary, 0, len(articleSummaries))
	for _, each := range articleSummaries {
		if getArticleSummariesByRegexIsValidTitle(each.Title, keywordList, runeList) {
			validArticleSummaries = append(validArticleSummaries, each)
		}
	}

	return validArticleSummaries
}

func getArticleSummariesByRegexIsValidTitle(title string, keywordList []string, runeList [][]rune) bool {
	for idx, each := range runeList {
		if len(each) <= TITLE_REGEX_N_GRAM {
			continue
		}

		if !strings.Contains(title, keywordList[idx]) {
			return false
		}
	}

	return true
}

// NewARticleSummary
//
// no n_comments in bbs.ArticleSummary from backend.
func NewArticleSummary(a_b *bbs.ArticleSummary, updateNanoTS types.NanoTS) *ArticleSummary {
	return &ArticleSummary{
		BBoardID:       a_b.BBoardID,
		ArticleID:      a_b.ArticleID,
		BoardArticleID: types.ToBoardArticleID(a_b.BBoardID, a_b.ArticleID),

		IsDeleted:  a_b.IsDeleted,
		CreateTime: types.Time4ToNanoTS(a_b.CreateTime),
		MTime:      types.Time4ToNanoTS(a_b.MTime),
		Recommend:  int(a_b.Recommend),
		Owner:      a_b.Owner,
		FullTitle:  types.Big5ToUtf8(a_b.FullTitle),
		Title:      types.Big5ToUtf8(a_b.RealTitle),
		Money:      int(a_b.Money),
		Class:      types.Big5ToUtf8(a_b.Class),
		Filemode:   a_b.Filemode,

		UpdateNanoTS: updateNanoTS,

		Idx: a_b.Idx,

		SubjectType: a_b.SubjectType,
	}
}

func UpdateArticleSummaries(articleSummaries []*ArticleSummary, updateNanoTS types.NanoTS) (err error) {
	if len(articleSummaries) == 0 {
		return nil
	}

	// create items which do not exists yet.
	theList := make([]*db.UpdatePair, len(articleSummaries))
	for idx, each := range articleSummaries {
		query := &ArticleQuery{
			BBoardID:  each.BBoardID,
			ArticleID: each.ArticleID,
		}

		theList[idx] = &db.UpdatePair{
			Filter: query,
			Update: each,
		}
	}

	r, err := Article_c.BulkCreateOnly(theList)
	if err != nil {
		return err
	}
	if r.UpsertedCount == int64(len(articleSummaries)) { // all are created
		return nil
	}

	upsertedIDs := r.UpsertedIDs
	updateArticleSummaries := make([]*db.UpdatePair, 0, len(theList))
	for idx, each := range theList {
		_, ok := upsertedIDs[int64(idx)]
		if ok {
			continue
		}

		origFilter := each.Filter.(*ArticleQuery)
		filter := bson.M{
			"$or": bson.A{
				bson.M{
					ARTICLE_BBOARD_ID_b:  origFilter.BBoardID,
					ARTICLE_ARTICLE_ID_b: origFilter.ArticleID,
					ARTICLE_UPDATE_NANO_TS_b: bson.M{
						"$exists": false,
					},

					ARTICLE_IS_DELETED_b: bson.M{"$exists": false},
				},
				bson.M{
					ARTICLE_BBOARD_ID_b:  origFilter.BBoardID,
					ARTICLE_ARTICLE_ID_b: origFilter.ArticleID,
					ARTICLE_UPDATE_NANO_TS_b: bson.M{
						"$lt": updateNanoTS,
					},

					ARTICLE_IS_DELETED_b: bson.M{"$exists": false},
				},
			},
		}
		each.Filter = filter
		updateArticleSummaries = append(updateArticleSummaries, each)
	}

	// update items with comparing update-nano-ts
	_, err = Article_c.BulkUpdateOneOnly(updateArticleSummaries)

	return err
}
