package schema

import (
	"reflect"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/mockhttp"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestUpdateArticleSummaries(t *testing.T) {
	setupTest()
	defer teardownTest()

	ret := mockhttp.LoadGeneralArticles(nil)

	updateNanoTS := types.NowNanoTS() - 200

	articleSummaries0 := make([]*ArticleSummary, len(ret.Articles))
	for idx, each_b := range ret.Articles {
		articleSummaries0[idx] = NewArticleSummary(each_b, updateNanoTS)
	}

	query0 := &ArticleQuery{BBoardID: bbs.BBoardID("10_WhoAmI"), ArticleID: bbs.ArticleID("19bWBI4Z")}
	articleSummary0 := &ArticleSummary{
		BBoardID:       bbs.BBoardID("10_WhoAmI"),
		ArticleID:      bbs.ArticleID("19bWBI4Z"),
		BoardArticleID: types.BoardArticleID("10_WhoAmI:19bWBI4Z"),

		IsDeleted:  false,
		CreateTime: types.NanoTS(1234567890000000000),
		MTime:      types.NanoTS(1234567889000000000),

		Recommend:    8,
		Owner:        bbs.UUserID("okcool"),
		FullTitle:    "[問題]然後呢？～",
		Title:        "然後呢？～",
		Money:        3,
		Class:        "問題",
		Filemode:     0,
		UpdateNanoTS: updateNanoTS,
		Idx:          "1234567890@19bWBI4Z",
	}

	articleSummary1 := &ArticleSummary{
		BBoardID:       bbs.BBoardID("10_WhoAmI"),
		ArticleID:      bbs.ArticleID("1VrooM21"),
		BoardArticleID: types.BoardArticleID("10_WhoAmI:1VrooM21"),

		IsDeleted:  false,
		CreateTime: types.NanoTS(1607937174000000000),
		MTime:      types.NanoTS(1607937100000000000),

		Recommend:    3,
		Owner:        bbs.UUserID("teemo"),
		FullTitle:    "[問題]再來呢？～",
		Title:        "再來呢？～",
		Money:        12,
		Class:        "問題",
		Filemode:     0,
		UpdateNanoTS: updateNanoTS,
		Idx:          "1607937174@1VrooM21",
	}

	updateNanoTS = types.NowNanoTS() - 100

	articleSummary2 := &ArticleSummary{
		BBoardID:       bbs.BBoardID("10_WhoAmI"),
		ArticleID:      bbs.ArticleID("1VrooM21"),
		BoardArticleID: types.BoardArticleID("10_WhoAmI:1VrooM21"),

		IsDeleted:  false,
		CreateTime: types.NanoTS(1607937174000000000),
		MTime:      types.NanoTS(1607937100000000000),

		Recommend:    4,
		Owner:        bbs.UUserID("teem2"),
		FullTitle:    "[問題]再來呢2？～",
		Title:        "再來呢2？～",
		Money:        15,
		Class:        "問題",
		Filemode:     0,
		UpdateNanoTS: updateNanoTS,
		Idx:          "1607937174@1VrooM21",
	}

	updateNanoTS1 := types.NowNanoTS()

	articleSummary3 := &ArticleSummary{
		BBoardID:       bbs.BBoardID("10_WhoAmI"),
		ArticleID:      bbs.ArticleID("1VrooM21"),
		BoardArticleID: types.BoardArticleID("10_WhoAmI:1VrooM21"),

		IsDeleted:    false,
		CreateTime:   types.NanoTS(1607937174000000000),
		MTime:        types.NanoTS(1607937100000000000),
		FullTitle:    "[問題]再來呢3？～",
		Recommend:    12,
		Owner:        bbs.UUserID("teemo"),
		Title:        "再來呢3？～",
		Money:        20,
		Class:        "問題",
		Filemode:     0,
		UpdateNanoTS: updateNanoTS1,
		Idx:          "1607937174@1VrooM21",
	}

	articleSummaries1 := []*ArticleSummary{articleSummary2}

	query1 := &ArticleQuery{BBoardID: bbs.BBoardID("10_WhoAmI"), ArticleID: bbs.ArticleID("1VrooM21")}

	query2 := &ArticleQuery{BBoardID: bbs.BBoardID("10_WhoAmI"), ArticleID: bbs.ArticleID("1VrooM21")}

	articleSummaries2 := []*ArticleSummary{articleSummary3}

	type args struct {
		articleSummaries []*ArticleSummary
		updateNanoTS     types.NanoTS
	}
	tests := []struct {
		name     string
		args     args
		query    *ArticleQuery
		expected *ArticleSummary
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{articleSummaries: articleSummaries0, updateNanoTS: updateNanoTS},
			query:    query0,
			expected: articleSummary0,
		},
		{
			args:     args{articleSummaries: articleSummaries0, updateNanoTS: updateNanoTS},
			query:    query1,
			expected: articleSummary1,
		},
		{
			args:     args{articleSummaries: articleSummaries1, updateNanoTS: updateNanoTS},
			query:    query2,
			expected: articleSummary2,
		},
		{
			args:     args{articleSummaries: articleSummaries2, updateNanoTS: updateNanoTS1},
			query:    query1,
			expected: articleSummary3,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if err := UpdateArticleSummaries(tt.args.articleSummaries, tt.args.updateNanoTS); (err != nil) != tt.wantErr {
				t.Errorf("UpdateArticleSummaries() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, _ := GetArticleSummary(tt.query.BBoardID, tt.query.ArticleID)
			testutil.TDeepEqual(t, "got", got, tt.expected)
		})
		wg.Wait()
	}
}

func TestGetArticleSummariesByOwnerID(t *testing.T) {
	setupTest()
	defer teardownTest()

	ret := mockhttp.LoadGeneralArticles(nil)

	updateNanoTS := types.NowNanoTS() - 200

	articleSummaries0 := make([]*ArticleSummary, len(ret.Articles))
	for idx, each_b := range ret.Articles {
		articleSummaries0[idx] = NewArticleSummary(each_b, updateNanoTS)
	}

	UpdateArticleSummaries(articleSummaries0, updateNanoTS)

	expected0 := []*ArticleSummary{articleSummaries0[0]}
	expected1 := []*ArticleSummary{articleSummaries0[1]}

	type args struct {
		ownerID         bbs.UUserID
		startCreateTime types.NanoTS
		descending      bool
		limit           int
	}
	tests := []struct {
		name           string
		args           args
		expectedResult []*ArticleSummary
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			args:           args{ownerID: "teemo", limit: 200},
			expectedResult: expected0,
		},
		{
			args:           args{ownerID: "okcool", limit: 200},
			expectedResult: expected1,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotResult, err := GetArticleSummariesByOwnerID(tt.args.ownerID, tt.args.startCreateTime, tt.args.descending, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArticleSummariesByOwnerID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotResult, tt.expectedResult)
		})
		wg.Wait()
	}
}

func Test_getArticleSummariesByRegexPatternList(t *testing.T) {
	keywords0 := [][]rune{[]rune("abc"), []rune("defgh"), []rune("ijklmnop")}
	expectedFirstPattern0 := "abc"
	expected0 := bson.A{
		bson.M{ARTICLE_BBOARD_ID_b: bbs.BBoardID("test"), ARTICLE_TITLE_REGEX_b: "defgh"},
		bson.M{ARTICLE_BBOARD_ID_b: bbs.BBoardID("test"), ARTICLE_TITLE_REGEX_b: "ijklm"},
	}

	type args struct {
		keywordList [][]rune
	}
	tests := []struct {
		name                 string
		args                 args
		expectedFirstPattern string
		expectedPatternList  bson.A
	}{
		// TODO: Add test cases.
		{
			args:                 args{keywordList: keywords0},
			expectedFirstPattern: expectedFirstPattern0,
			expectedPatternList:  expected0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFirstPattern, gotPatternList := getArticleSummariesByRegexPatternList("test", tt.args.keywordList)

			if gotFirstPattern != tt.expectedFirstPattern {
				t.Errorf("GetArticleSummariesByRegex: firstPattern: %v expected: %v", gotFirstPattern, tt.expectedFirstPattern)
			}
			assert.Equal(t, tt.expectedPatternList, gotPatternList)
		})
	}
}

func Test_getArticleSummariesByRegexSetQuery(t *testing.T) {
	patternList0 := bson.A{
		bson.M{ARTICLE_BBOARD_ID_b: "test", ARTICLE_TITLE_REGEX_b: "defgh"},
		bson.M{ARTICLE_BBOARD_ID_b: "test", ARTICLE_TITLE_REGEX_b: "ijklm"},
	}
	expected0 := bson.M{
		"$and": bson.A{
			bson.M{
				ARTICLE_BBOARD_ID_b:   bbs.BBoardID("test"),
				ARTICLE_TITLE_REGEX_b: "abc",
				ARTICLE_IS_DELETED_b: bson.M{
					"$exists": false,
				},
			},
			bson.M{ARTICLE_BBOARD_ID_b: "test", ARTICLE_TITLE_REGEX_b: "defgh"},
			bson.M{ARTICLE_BBOARD_ID_b: "test", ARTICLE_TITLE_REGEX_b: "ijklm"},
		},
	}

	expected1 := bson.M{
		"$or": bson.A{
			bson.M{
				"$and": bson.A{
					bson.M{
						ARTICLE_BBOARD_ID_b:   bbs.BBoardID("test"),
						ARTICLE_TITLE_REGEX_b: "abc",
						ARTICLE_IS_DELETED_b: bson.M{
							"$exists": false,
						},
						ARTICLE_CREATE_TIME_b: bson.M{
							"$lt": types.NanoTS(1234567890000000000),
						},
					},
					bson.M{ARTICLE_BBOARD_ID_b: "test", ARTICLE_TITLE_REGEX_b: "defgh"},
					bson.M{ARTICLE_BBOARD_ID_b: "test", ARTICLE_TITLE_REGEX_b: "ijklm"},
				},
			},
			bson.M{
				"$and": bson.A{
					bson.M{
						ARTICLE_BBOARD_ID_b:   bbs.BBoardID("test"),
						ARTICLE_TITLE_REGEX_b: "abc",
						ARTICLE_IS_DELETED_b: bson.M{
							"$exists": false,
						},
						ARTICLE_CREATE_TIME_b: types.NanoTS(1234567890000000000),
						ARTICLE_ARTICLE_ID_b: bson.M{
							"$lte": bbs.ArticleID("aid0"),
						},
					},
					bson.M{ARTICLE_BBOARD_ID_b: "test", ARTICLE_TITLE_REGEX_b: "defgh"},
					bson.M{ARTICLE_BBOARD_ID_b: "test", ARTICLE_TITLE_REGEX_b: "ijklm"},
				},
			},
		},
	}
	type args struct {
		boardID      bbs.BBoardID
		firstPattern string
		patternList  bson.A
		createNanoTS types.NanoTS
		articleID    bbs.ArticleID
		descending   bool
	}
	tests := []struct {
		name          string
		args          args
		expectedQuery bson.M
	}{
		// TODO: Add test cases.
		{
			args:          args{boardID: "test", firstPattern: "abc", patternList: patternList0, descending: false},
			expectedQuery: expected0,
		},
		{
			args:          args{boardID: "test", firstPattern: "abc", patternList: patternList0, descending: true, createNanoTS: 1234567890000000000, articleID: "aid0"},
			expectedQuery: expected1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQuery := getArticleSummariesByRegexSetQuery(tt.args.boardID, tt.args.firstPattern, tt.args.patternList, tt.args.createNanoTS, tt.args.articleID, tt.args.descending)

			assert.Equal(t, tt.expectedQuery, gotQuery)
		})
	}
}

func TestGetArticleSummariesByRegex(t *testing.T) {
	setupTest()
	defer teardownTest()

	ret := mockhttp.LoadGeneralArticles(nil)

	updateNanoTS := types.NowNanoTS() - 200

	articleSummaries0 := make([]*ArticleSummary, len(ret.Articles))
	for idx, each_b := range ret.Articles {
		articleSummaries0[idx] = NewArticleSummary(each_b, updateNanoTS)
	}

	articleSummaryWithRegexes0 := make([]*ArticleSummaryWithRegex, len(ret.Articles))
	for idx, each_b := range ret.Articles {
		articleSummaryWithRegexes0[idx] = NewArticleSummaryWithRegex(each_b, updateNanoTS, false)
	}

	UpdateArticleSummaryWithRegexes(articleSummaryWithRegexes0, updateNanoTS)

	title1 := "有沒有這一些事情的八卦呢？～"
	class1 := "問題"
	articleSummaryWithRegex1 := &ArticleSummaryWithRegex{
		BBoardID:       "10_WhoAmI",
		ArticleID:      "testAid0",
		BoardArticleID: "10_WhoAmI:testAid0",

		CreateTime:   1234567890000000000,
		MTime:        1234567890000000000,
		Title:        title1,
		Class:        class1,
		TitleRegex:   articleTitleToTitleRegex(title1),
		UpdateNanoTS: 1234567890000000000,
	}
	articleSummary1 := &ArticleSummary{
		BBoardID:       "10_WhoAmI",
		ArticleID:      "testAid0",
		BoardArticleID: "10_WhoAmI:testAid0",

		CreateTime:   1234567890000000000,
		MTime:        1234567890000000000,
		Title:        title1,
		Class:        class1,
		UpdateNanoTS: 1234567890000000000,
	}

	UpdateArticleSummaryWithRegexes([]*ArticleSummaryWithRegex{articleSummaryWithRegex1}, updateNanoTS)

	article0, _ := GetArticleSummaryWithRegex("10_WhoAmI", "1VrooM21")
	logrus.Infof("article0: (%v/%v)", article0.Title, article0.TitleRegex)

	expected0 := []*ArticleSummary{articleSummaries0[0]}
	expected1 := []*ArticleSummary{articleSummaries0[1]}
	expected2 := []*ArticleSummary{}
	expected3 := []*ArticleSummary{articleSummary1}

	type args struct {
		boardID      bbs.BBoardID
		keywordList  []string
		createNanoTS types.NanoTS
		articleID    bbs.ArticleID
		descending   bool
		limit        int
	}
	tests := []struct {
		name           string
		args           args
		expectedResult []*ArticleSummary
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			args:           args{boardID: "10_WhoAmI", keywordList: []string{"再來"}, descending: true, limit: 200},
			expectedResult: expected0,
		},
		{
			args:           args{boardID: "10_WhoAmI", keywordList: []string{"然後"}, descending: true, limit: 200},
			expectedResult: expected1,
		},
		{
			args:           args{boardID: "10_WhoAmI", keywordList: []string{"然後", "再來"}, descending: true, limit: 200},
			expectedResult: expected2,
		},
		{
			args:           args{boardID: "10_WhoAmI", keywordList: []string{"有沒有這一些", "八卦"}, descending: true, limit: 200},
			expectedResult: expected3,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotResult, err := GetArticleSummariesByRegex(tt.args.boardID, tt.args.keywordList, tt.args.createNanoTS, tt.args.articleID, tt.args.descending, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArticleSummariesByRegex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.expectedResult) {
				t.Errorf("GetArticleSummariesByRegex() = %v, want %v", gotResult, tt.expectedResult)
			}
		})
		wg.Wait()
	}
}

func TestGetBottomArticleSummaries(t *testing.T) {
	setupTest()
	defer teardownTest()

	articleSummary0 := &ArticleSummaryWithRegex{
		BBoardID:       "board0",
		ArticleID:      "1VtW-QXT",
		BoardArticleID: "board0:1VtW-QXT",

		CreateTime:   1608388506000000000,
		MTime:        1608388508000000000,
		Recommend:    12,
		Owner:        "SYSOP",
		FullTitle:    "Re:Re:Re: Fw:[轉錄] [爆卦]這裡是八卦板",
		Title:        "這裡是八卦板",
		Class:        "爆卦",
		TitleRegex:   []string{"這", "裡", "是", "八", "卦", "板", "這裡", "裡是", "是八", "八卦", "卦板", "這裡是", "裡是八", "是八卦", "八卦板", "這裡是八", "裡是八卦", "是八卦板", "這裡是八卦", "裡是八卦板"},
		SubjectType:  ptttype.SUBJECT_REPLY,
		Idx:          "1608388506@1VtW-QXT",
		UpdateNanoTS: 1734567890000000000,
		IsBottom:     true,
	}

	expected0 := &ArticleSummary{
		BBoardID:       "board0",
		ArticleID:      "1VtW-QXT",
		BoardArticleID: "board0:1VtW-QXT",

		CreateTime:   1608388506000000000,
		MTime:        1608388508000000000,
		Recommend:    12,
		Owner:        "SYSOP",
		FullTitle:    "Re:Re:Re: Fw:[轉錄] [爆卦]這裡是八卦板",
		Title:        "這裡是八卦板",
		Class:        "爆卦",
		SubjectType:  ptttype.SUBJECT_REPLY,
		Idx:          "1608388506@1VtW-QXT",
		UpdateNanoTS: 1734567890000000000,
		IsBottom:     true,
	}

	_ = UpdateArticleSummaryWithRegexes([]*ArticleSummaryWithRegex{articleSummary0}, 1734567890000000000)

	type args struct {
		boardID bbs.BBoardID
	}
	tests := []struct {
		name           string
		args           args
		expectedResult []*ArticleSummary
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			args:           args{boardID: "board0"},
			expectedResult: []*ArticleSummary{expected0},
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotResult, err := GetBottomArticleSummaries(tt.args.boardID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBottomArticleSummaries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.expectedResult) {
				t.Errorf("GetBottomArticleSummaries() = %v, want %v", gotResult, tt.expectedResult)
			}
		})
		wg.Wait()
	}
}

func TestGetArticleSummaries(t *testing.T) {
	setupTest()
	defer teardownTest()

	articleSummary0 := &ArticleSummaryWithRegex{
		BBoardID:       "board0",
		ArticleID:      "1VtW-QXT",
		BoardArticleID: "board0:1VtW-QXT",

		CreateTime:   1608388506000000000,
		MTime:        1608388508000000000,
		Recommend:    12,
		Owner:        "SYSOP",
		FullTitle:    "Re:Re:Re: Fw:[轉錄] [爆卦]這裡是八卦板",
		Title:        "這裡是八卦板",
		Class:        "爆卦",
		TitleRegex:   []string{"這", "裡", "是", "八", "卦", "板", "這裡", "裡是", "是八", "八卦", "卦板", "這裡是", "裡是八", "是八卦", "八卦板", "這裡是八", "裡是八卦", "是八卦板", "這裡是八卦", "裡是八卦板"},
		SubjectType:  ptttype.SUBJECT_REPLY,
		Idx:          "1608388506@1VtW-QXT",
		UpdateNanoTS: 1734567890000000000,
		IsBottom:     true,
	}

	expected0 := &ArticleSummary{
		BBoardID:       "board0",
		ArticleID:      "1VtW-QXT",
		BoardArticleID: "board0:1VtW-QXT",

		CreateTime:   1608388506000000000,
		MTime:        1608388508000000000,
		Recommend:    12,
		Owner:        "SYSOP",
		FullTitle:    "Re:Re:Re: Fw:[轉錄] [爆卦]這裡是八卦板",
		Title:        "這裡是八卦板",
		Class:        "爆卦",
		SubjectType:  ptttype.SUBJECT_REPLY,
		Idx:          "1608388506@1VtW-QXT",
		UpdateNanoTS: 1734567890000000000,
		IsBottom:     true,
	}

	_ = UpdateArticleSummaryWithRegexes([]*ArticleSummaryWithRegex{articleSummary0}, 1734567890000000000)

	type args struct {
		boardID    bbs.BBoardID
		startIdx   string
		descending bool
		limit      int
	}
	tests := []struct {
		name           string
		args           args
		expectedResult []*ArticleSummary
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			args:           args{boardID: "board0"},
			expectedResult: []*ArticleSummary{expected0},
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotResult, err := GetArticleSummaries(tt.args.boardID, tt.args.startIdx, tt.args.descending, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArticleSummaries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.expectedResult) {
				t.Errorf("GetArticleSummaries() = %v, want %v", gotResult, tt.expectedResult)
			}
		})
		wg.Wait()
	}
}
