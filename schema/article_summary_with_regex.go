package schema

import (
	"strings"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/boardd"
	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ArticleSummaryWithRegex
type ArticleSummaryWithRegex struct {
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

	TitleRegex []string `bson:"title_regex"`

	Idx string `bson:"pttidx"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`

	NComments int `bson:"n_comments,omitempty"` // n_comments is read-only in article-summary.

	Rank int `bson:"rank,omitempty"` // 評價, read-only

	SubjectType ptttype.SubjectType `bson:"subject_type"`

	IsBottom bool `bson:"is_bottom"`
}

var (
	EMPTY_ARTICLE_SUMMARY_WITH_REGEX = &ArticleSummaryWithRegex{}
	articleSummaryWithRegexFields    = getFields(EMPTY_ARTICLE, EMPTY_ARTICLE_SUMMARY_WITH_REGEX)
)

// NewArticleSummaryWithRegexFromPBArticle
//
// no n_comments in bbs.ArticleSummary from backend.
func NewArticleSummaryWithRegexFromPBArticle(boardID bbs.BBoardID, a_b *boardd.Post, updateNanoTS types.NanoTS, isBottom bool) *ArticleSummaryWithRegex {
	filename := &ptttype.Filename_t{}
	copy(filename[:], []byte(a_b.Filename))
	articleID := bbs.ToArticleID(filename)

	createTime, err := filename.CreateTime()
	if err != nil {
		return nil
	}

	subjectType, realTitleWithClass := parseSubjectEx(a_b.Title)
	theClass, title := parseTitle(realTitleWithClass)
	pttbbsArticleSummary := &bbs.ArticleSummary{
		ArticleID:  articleID,
		CreateTime: createTime,
	}
	idx := bbs.SerializeArticleIdxStr(pttbbsArticleSummary)

	titleRegex := articleTitleToTitleRegex(title)
	return &ArticleSummaryWithRegex{
		BBoardID:       boardID,
		ArticleID:      articleID,
		BoardArticleID: types.ToBoardArticleID(boardID, articleID),

		IsDeleted:  false,
		CreateTime: types.Time4ToNanoTS(createTime),
		MTime:      types.NanoTS(a_b.ModifiedNsec),
		Recommend:  int(a_b.NumRecommends),
		Owner:      bbs.UUserID(a_b.Owner),
		FullTitle:  a_b.Title,
		Title:      title,
		Class:      theClass,
		Filemode:   ptttype.FileMode(a_b.Filemode),

		TitleRegex: titleRegex,

		SubjectType: subjectType,

		UpdateNanoTS: updateNanoTS,
		Idx:          idx,
		IsBottom:     isBottom,
	}
}

func parseSubjectEx(fullTitle string) (subjectType ptttype.SubjectType, realTitleWithClass string) {
	subjectType = ptttype.SUBJECT_NORMAL

	fullTitle = strings.TrimSpace(fullTitle)

	isSet := false
	for len(fullTitle) > 0 {
		prefixLower := strings.ToLower(fullTitle[:len(STR_REPLY)])
		if strings.HasPrefix(prefixLower, STR_REPLY_LOWER) {
			if !isSet {
				isSet = true
				subjectType = ptttype.SUBJECT_REPLY
			}
			fullTitle = strings.TrimSpace(fullTitle[len(STR_REPLY):])
			continue
		}

		prefixLower = strings.ToLower(fullTitle[:len(STR_FORWARD)])
		if strings.HasPrefix(prefixLower, STR_FORWARD_LOWER) {
			if !isSet {
				isSet = true
				subjectType = ptttype.SUBJECT_FORWARD
			}
			fullTitle = strings.TrimSpace(fullTitle[len(STR_FORWARD):])
			continue
		}

		if strings.HasPrefix(fullTitle, STR_LEGACY_FORWARD) {
			if !isSet {
				isSet = true
				subjectType = ptttype.SUBJECT_FORWARD
			}

			fullTitle = strings.TrimSpace(fullTitle[len(STR_LEGACY_FORWARD):])
			continue
		}

		break
	}

	return subjectType, fullTitle
}

func parseTitle(realTitleWithClass string) (theClass string, title string) {
	realTitleWithClass = strings.TrimSpace(realTitleWithClass)
	if !strings.HasPrefix(realTitleWithClass, "[") {
		return "", realTitleWithClass
	}

	count := 0
	for idx, each := range realTitleWithClass {
		if idx == 0 {
			continue
		}

		// count > 4
		if count > 4 {
			return "", strings.TrimSpace(realTitleWithClass)
		}

		// count == 4
		if count == 4 {
			if each != rune(']') {
				return "", realTitleWithClass
			}

			return strings.TrimSpace(realTitleWithClass[1:idx]), strings.TrimSpace(realTitleWithClass[(idx + 1):])
		}

		// count < 4
		if each < 0x80 {
			count++
		} else {
			count += 2
		}
	}

	return "", realTitleWithClass
}

// NewArticleSummaryWithRegex
//
// no n_comments in bbs.ArticleSummary from backend.
func NewArticleSummaryWithRegex(a_b *bbs.ArticleSummary, updateNanoTS types.NanoTS, isBottom bool) *ArticleSummaryWithRegex {
	title := types.Big5ToUtf8(a_b.RealTitle)

	titleRegex := articleTitleToTitleRegex(title)
	return &ArticleSummaryWithRegex{
		BBoardID:       a_b.BBoardID,
		ArticleID:      a_b.ArticleID,
		BoardArticleID: types.ToBoardArticleID(a_b.BBoardID, a_b.ArticleID),

		IsDeleted:  a_b.IsDeleted,
		CreateTime: types.Time4ToNanoTS(a_b.CreateTime),
		MTime:      types.Time4ToNanoTS(a_b.MTime),
		Recommend:  int(a_b.Recommend),
		Owner:      a_b.Owner,
		FullTitle:  types.Big5ToUtf8(a_b.FullTitle),
		Title:      title,
		Money:      int(a_b.Money),
		Class:      types.Big5ToUtf8(a_b.Class),
		Filemode:   a_b.Filemode,

		TitleRegex: titleRegex,

		UpdateNanoTS: updateNanoTS,

		Idx: a_b.Idx,

		SubjectType: a_b.SubjectType,

		IsBottom: isBottom,
	}
}

// articleTitleToTitleRegex
//
// params:
//
//	title:
//	theClass: (with 4-byte length)
func articleTitleToTitleRegex(title string) (titleRegex []string) {
	titleRune := []rune(title)
	nGramTitleRegex := TITLE_REGEX_N_GRAM
	if len(titleRune) < TITLE_REGEX_N_GRAM {
		nGramTitleRegex = len(titleRune)
	}

	lenTitleRegex := nGramTitleRegex * len(titleRune)

	titleRegex = make([]string, 0, lenTitleRegex)
	titleRegex = articleTitleToTitleRegexCore(titleRune, titleRegex, nGramTitleRegex)

	return titleRegex
}

func articleTitleToTitleRegexCore(titleRune []rune, titleRegex []string, nGram int) (newTitleRegex []string) {
	for i := 1; i <= nGram; i++ {
		endPos := len(titleRune) - i
		for idx := 0; idx <= endPos; idx++ {
			titleRegex = append(titleRegex, string(titleRune[idx:(idx+i)]))
		}
	}

	return titleRegex
}

func UpdateArticleSummaryWithRegexes(articleSummaryWithRegexes []*ArticleSummaryWithRegex, updateNanoTS types.NanoTS) (err error) {
	if len(articleSummaryWithRegexes) == 0 {
		return nil
	}

	// create items which do not exists yet.
	theList := make([]*db.UpdatePair, len(articleSummaryWithRegexes))
	for idx, each := range articleSummaryWithRegexes {
		query := &ArticleCreateQuery{
			BBoardID:       each.BBoardID,
			ArticleID:      each.ArticleID,
			BoardArticleID: types.ToBoardArticleID(each.BBoardID, each.ArticleID),
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
	if r.UpsertedCount == int64(len(articleSummaryWithRegexes)) { // all are created
		return nil
	}

	upsertedIDs := r.UpsertedIDs
	updateArticleSummaryWithRegexes := make([]*db.UpdatePair, 0, len(theList))
	for idx, each := range theList {
		_, ok := upsertedIDs[int64(idx)]
		if ok {
			continue
		}

		origFilter := each.Filter.(*ArticleCreateQuery)
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
		updateArticleSummaryWithRegexes = append(updateArticleSummaryWithRegexes, each)
	}

	// update items with comparing update-nano-ts
	_, err = Article_c.BulkUpdateOneOnly(updateArticleSummaryWithRegexes)

	return err
}

func GetArticleSummaryWithRegex(bboardID bbs.BBoardID, articleID bbs.ArticleID) (result *ArticleSummaryWithRegex, err error) {
	query := &ArticleQuery{
		BBoardID:  bboardID,
		ArticleID: articleID,
	}

	result = &ArticleSummaryWithRegex{}
	err = Article_c.FindOne(query, &result, articleSummaryWithRegexFields)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}
