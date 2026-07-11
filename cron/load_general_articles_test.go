package cron

import (
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

func Test_loadGeneralArticlesCorePtt(t *testing.T) {
	setupTest()
	defer teardownTest()

	updateNanoTS := types.NowNanoTS()
	want := []*schema.ArticleSummaryWithRegex{
		{
			BBoardID:       "10_WhoAmI",
			ArticleID:      "1VrooM21",
			BoardArticleID: "10_WhoAmI:1VrooM21",

			IsDeleted:    false,
			CreateTime:   1607937174000000000,
			MTime:        1607937100000000000,
			Recommend:    3,
			Owner:        "teemo",
			Title:        "再來呢？～",
			Class:        "問題",
			Money:        12,
			Filemode:     0,
			Idx:          "1607937174@1VrooM21",
			FullTitle:    "[問題]再來呢？～",
			TitleRegex:   []string{"再", "來", "呢", "？", "～", "再來", "來呢", "呢？", "？～", "再來呢", "來呢？", "呢？～", "再來呢？", "來呢？～", "再來呢？～"},
			UpdateNanoTS: updateNanoTS,
		},
		{
			BBoardID:       "10_WhoAmI",
			ArticleID:      "19bWBI4Z",
			BoardArticleID: "10_WhoAmI:19bWBI4Z",

			IsDeleted:    false,
			CreateTime:   1234567890000000000,
			MTime:        1234567889000000000,
			Recommend:    8,
			Owner:        "okcool",
			Title:        "然後呢？～",
			Class:        "問題",
			Money:        3,
			Filemode:     0,
			Idx:          "1234567890@19bWBI4Z",
			FullTitle:    "[問題]然後呢？～",
			TitleRegex:   []string{"然", "後", "呢", "？", "～", "然後", "後呢", "呢？", "？～", "然後呢", "後呢？", "呢？～", "然後呢？", "後呢？～", "然後呢？～"},
			UpdateNanoTS: updateNanoTS,
		},
	}
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		boardID      bbs.BBoardID
		startIdx     string
		updateNanoTS types.NanoTS
		want         []*schema.ArticleSummaryWithRegex
		want2        string
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			boardID:      "10_WhoAmI",
			updateNanoTS: updateNanoTS,
			want:         want,
			want2:        "1234560000@19bUG021",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2, gotErr := loadGeneralArticlesCorePtt(tt.boardID, tt.startIdx, tt.updateNanoTS)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("loadGeneralArticlesCorePtt() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Errorf("loadGeneralArticlesCorePtt() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			testutil.TDeepEqual(t, "want", got, tt.want)
			if got2 != tt.want2 {
				t.Errorf("loadGeneralArticlesCorePtt() = %v, want2 %v", got2, tt.want2)
			}
		})
	}
}

func Test_loadBottomArticlesPtt(t *testing.T) {
	setupTest()
	defer teardownTest()

	updateNanoTS := types.NowNanoTS()
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		boardID      bbs.BBoardID
		updateNanoTS types.NanoTS
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			boardID:      "10_WhoAmI",
			updateNanoTS: updateNanoTS,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := loadBottomArticlesPtt(tt.boardID, tt.updateNanoTS)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("loadBottomArticlesPtt() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("loadBottomArticlesPtt() succeeded unexpectedly")
			}
		})
	}
}
