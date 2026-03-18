package schema

import (
	"reflect"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/boardd"
	"github.com/Ptt-official-app/pttbbs-backend/mockhttp"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

func Test_articleTitleToTitleRegexCore(t *testing.T) {
	setupTest()
	defer teardownTest()

	title0 := []rune("abcd")
	nGram0 := 3
	expected0 := []string{"a", "b", "c", "d", "ab", "bc", "cd", "abc", "bcd"}

	title1 := []rune("abcd")
	nGram1 := 4
	expected1 := []string{"a", "b", "c", "d", "ab", "bc", "cd", "abc", "bcd", "abcd"}

	title2 := []rune("abcd")
	nGram2 := 5
	expected2 := []string{"a", "b", "c", "d", "ab", "bc", "cd", "abc", "bcd", "abcd"}

	title3 := []rune("再來呢？～")
	nGram3 := 5
	expected3 := []string{"再", "來", "呢", "？", "～", "再來", "來呢", "呢？", "？～", "再來呢", "來呢？", "呢？～", "再來呢？", "來呢？～", "再來呢？～"}

	type args struct {
		title      []rune
		titleRegex []string
		nGram      int
	}
	tests := []struct {
		name                  string
		args                  args
		expectedNewTitleRegex []string
	}{
		// TODO: Add test cases.
		{
			name:                  "abcd with 3",
			args:                  args{title: title0, nGram: nGram0},
			expectedNewTitleRegex: expected0,
		},
		{
			name:                  "abcd with 4",
			args:                  args{title: title1, nGram: nGram1},
			expectedNewTitleRegex: expected1,
		},
		{
			name:                  "abcd with 5",
			args:                  args{title: title2, nGram: nGram2},
			expectedNewTitleRegex: expected2,
		},
		{
			name:                  "再來呢？～ with 5",
			args:                  args{title: title3, nGram: nGram3},
			expectedNewTitleRegex: expected3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewTitleRegex := articleTitleToTitleRegexCore(tt.args.title, tt.args.titleRegex, tt.args.nGram); !reflect.DeepEqual(gotNewTitleRegex, tt.expectedNewTitleRegex) {
				t.Errorf("articleTitleToTitleRegexCore() = %v, want %v", gotNewTitleRegex, tt.expectedNewTitleRegex)
			}
		})
	}
}

func Test_articleTitleToTitleRegex(t *testing.T) {
	setupTest()
	defer teardownTest()

	title0 := "abcd"
	expected0 := []string{"a", "b", "c", "d", "ab", "bc", "cd", "abc", "bcd", "abcd"}

	title1 := "abcde"
	expected1 := []string{"a", "b", "c", "d", "e", "ab", "bc", "cd", "de", "abc", "bcd", "cde", "abcd", "bcde", "abcde"}

	title2 := "abcdef"
	expected2 := []string{"a", "b", "c", "d", "e", "f", "ab", "bc", "cd", "de", "ef", "abc", "bcd", "cde", "def", "abcd", "bcde", "cdef", "abcde", "bcdef"}

	title3 := "abcdef"
	expected3 := []string{"a", "b", "c", "d", "e", "f", "ab", "bc", "cd", "de", "ef", "abc", "bcd", "cde", "def", "abcd", "bcde", "cdef", "abcde", "bcdef"}

	type args struct {
		title string
	}
	tests := []struct {
		name               string
		args               args
		expectedTitleRegex []string
	}{
		// TODO: Add test cases.
		{
			name:               "title only with abcd",
			args:               args{title: title0},
			expectedTitleRegex: expected0,
		},
		{
			name:               "title only with abcde",
			args:               args{title: title1},
			expectedTitleRegex: expected1,
		},
		{
			name:               "title only with abcdef",
			args:               args{title: title2},
			expectedTitleRegex: expected2,
		},
		{
			name:               "title with class",
			args:               args{title: title3},
			expectedTitleRegex: expected3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTitleRegex := articleTitleToTitleRegex(tt.args.title); !reflect.DeepEqual(gotTitleRegex, tt.expectedTitleRegex) {
				t.Errorf("articleTitleToTitleRegex() = %v, want %v", gotTitleRegex, tt.expectedTitleRegex)
			}
		})
	}
}

func TestUpdateArticleSummaryWithRegexes(t *testing.T) {
	setupTest()
	defer teardownTest()

	ret := mockhttp.LoadGeneralArticles(nil)

	updateNanoTS0 := types.NowNanoTS() - 200

	articleSummaryWithRegexes0 := make([]*ArticleSummaryWithRegex, len(ret.Articles))
	for idx, each_b := range ret.Articles {
		articleSummaryWithRegexes0[idx] = NewArticleSummaryWithRegex(each_b, updateNanoTS0, false)
	}

	updateNanoTS1 := types.NowNanoTS()
	articleSummaryWithRegexes1 := make([]*ArticleSummaryWithRegex, len(ret.Articles))
	for idx, each_b := range ret.Articles {
		articleSummaryWithRegexes1[idx] = NewArticleSummaryWithRegex(each_b, updateNanoTS1, false)
	}

	type args struct {
		articleSummaryWithRegexes []*ArticleSummaryWithRegex
		updateNanoTS              types.NanoTS
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{articleSummaryWithRegexes: articleSummaryWithRegexes0, updateNanoTS: updateNanoTS0},
		},
		{
			args: args{articleSummaryWithRegexes: articleSummaryWithRegexes1, updateNanoTS: updateNanoTS1},
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if err := UpdateArticleSummaryWithRegexes(tt.args.articleSummaryWithRegexes, tt.args.updateNanoTS); (err != nil) != tt.wantErr {
				t.Errorf("UpdateArticleSummaryWithRegexes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		wg.Wait()
	}
}

func Test_parseTitle(t *testing.T) {
	type args struct {
		fullTitle string
	}
	tests := []struct {
		name             string
		args             args
		expectedTheClass string
		expectedTitle    string
	}{
		// TODO: Add test cases.
		{
			args:             args{fullTitle: "[CPBL] 這裡是棒球板"},
			expectedTheClass: "CPBL",
			expectedTitle:    "這裡是棒球板",
		},
		{
			args:             args{fullTitle: "[八卦] 這裡是八卦板"},
			expectedTheClass: "八卦",
			expectedTitle:    "這裡是八卦板",
		},
		{
			args:             args{fullTitle: "[a一d]abcd"},
			expectedTheClass: "a一d",
			expectedTitle:    "abcd",
		},
		{
			args:             args{fullTitle: "[ad一]abcd"},
			expectedTheClass: "ad一",
			expectedTitle:    "abcd",
		},
		{
			args:             args{fullTitle: "[ad  ]abcd"},
			expectedTheClass: "ad",
			expectedTitle:    "abcd",
		},
		{
			args:             args{fullTitle: "[NBA ]這裡是 NBA 板"},
			expectedTheClass: "NBA",
			expectedTitle:    "這裡是 NBA 板",
		},
		{
			args:             args{fullTitle: "[0123]這裡是 NBA 板"},
			expectedTheClass: "0123",
			expectedTitle:    "這裡是 NBA 板",
		},
		{
			args:             args{fullTitle: "[0123 ]這裡是 NBA 板"},
			expectedTheClass: "",
			expectedTitle:    "[0123 ]這裡是 NBA 板",
		},
		{
			args:             args{fullTitle: "[NBA  ]這裡是 NBA 板"},
			expectedTheClass: "",
			expectedTitle:    "[NBA  ]這裡是 NBA 板",
		},
		{
			args:             args{fullTitle: "[NBA 這裡是 NBA 板"},
			expectedTheClass: "",
			expectedTitle:    "[NBA 這裡是 NBA 板",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTheClass, gotTitle := parseTitle(tt.args.fullTitle)
			if gotTheClass != tt.expectedTheClass {
				t.Errorf("parseTitle() gotTheClass = %v, want %v", gotTheClass, tt.expectedTheClass)
			}
			if gotTitle != tt.expectedTitle {
				t.Errorf("parseTitle() gotTitle = %v, want %v", gotTitle, tt.expectedTitle)
			}
		})
	}
}

func Test_parseSubjectEx(t *testing.T) {
	type args struct {
		fullTitle string
	}
	tests := []struct {
		name                       string
		args                       args
		expectedSubjectType        ptttype.SubjectType
		expectedRealTitleWithClass string
	}{
		// TODO: Add test cases.
		{
			args:                       args{fullTitle: "Re: Fw: [轉錄] [NBA ] 這裡是 NBA 板"},
			expectedSubjectType:        ptttype.SUBJECT_REPLY,
			expectedRealTitleWithClass: "[NBA ] 這裡是 NBA 板",
		},
		{
			args:                       args{fullTitle: "Fw: [轉錄] Re: [NBA ] 這裡是 NBA 板"},
			expectedSubjectType:        ptttype.SUBJECT_FORWARD,
			expectedRealTitleWithClass: "[NBA ] 這裡是 NBA 板",
		},

		{
			args:                       args{fullTitle: "[轉錄] Re:Re:Re: [NBA ] 這裡是 NBA 板"},
			expectedSubjectType:        ptttype.SUBJECT_FORWARD,
			expectedRealTitleWithClass: "[NBA ] 這裡是 NBA 板",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSubjectType, gotRealTitleWithClass := parseSubjectEx(tt.args.fullTitle)
			if !reflect.DeepEqual(gotSubjectType, tt.expectedSubjectType) {
				t.Errorf("parseSubjectEx() gotSubjectType = %v, want %v", gotSubjectType, tt.expectedSubjectType)
			}
			if gotRealTitleWithClass != tt.expectedRealTitleWithClass {
				t.Errorf("parseSubjectEx() gotRealTitleWithClass = %v, want %v", gotRealTitleWithClass, tt.expectedRealTitleWithClass)
			}
		})
	}
}

func TestNewArticleSummaryWithRegexFromPBArticle(t *testing.T) {
	pbArticle0 := &boardd.Post{
		Filename:      "M.1608388506.A.85D",
		Title:         "Re:Re:Re: Fw:[轉錄] [爆卦]這裡是八卦板",
		ModifiedNsec:  1608388508000000000,
		NumRecommends: 12,
		Owner:         "SYSOP",
	}
	expected0 := &ArticleSummaryWithRegex{
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
	}

	type args struct {
		boardID      bbs.BBoardID
		a_b          *boardd.Post
		updateNanoTS types.NanoTS
		isBottom     bool
	}
	tests := []struct {
		name     string
		args     args
		expected *ArticleSummaryWithRegex
	}{
		// TODO: Add test cases.
		{
			args:     args{boardID: "board0", a_b: pbArticle0, updateNanoTS: 1734567890000000000},
			expected: expected0,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			got := NewArticleSummaryWithRegexFromPBArticle(tt.args.boardID, tt.args.a_b, tt.args.updateNanoTS, tt.args.isBottom)

			testutil.TDeepEqual(t, "got", got, tt.expected)
		})
		wg.Wait()
	}
}
