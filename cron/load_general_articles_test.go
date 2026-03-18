package cron

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

func Test_loadGeneralArticles(t *testing.T) {
	setupTest()
	defer teardownTest()

	expected0 := []*schema.ArticleSummaryWithRegex{
		{ // 0
			BBoardID:       "10_WhoAmI",
			ArticleID:      "1UT16-xG",
			BoardArticleID: "10_WhoAmI:1UT16-xG",

			IsDeleted:  false,
			CreateTime: 1584665022000000000,
			MTime:      1644506386000000000,
			Recommend:  17,
			Owner:      "hellohiro",
			Title:      "為何打麻將叫賭博但買股票叫投資？",
			Class:      "問卦",
			Money:      0,
			Filemode:   0,
			Idx:        "1584665022@1UT16-xG",
			FullTitle:  "[問卦] 為何打麻將叫賭博但買股票叫投資？",
			TitleRegex: []string{"為", "何", "打", "麻", "將", "叫", "賭", "博", "但", "買", "股", "票", "叫", "投", "資", "？", "為何", "何打", "打麻", "麻將", "將叫", "叫賭", "賭博", "博但", "但買", "買股", "股票", "票叫", "叫投", "投資", "資？", "為何打", "何打麻", "打麻將", "麻將叫", "將叫賭", "叫賭博", "賭博但", "博但買", "但買股", "買股票", "股票叫", "票叫投", "叫投資", "投資？", "為何打麻", "何打麻將", "打麻將叫", "麻將叫賭", "將叫賭博", "叫賭博但", "賭博但買", "博但買股", "但買股票", "買股票叫", "股票叫投", "票叫投資", "叫投資？", "為何打麻將", "何打麻將叫", "打麻將叫賭", "麻將叫賭博", "將叫賭博但", "叫賭博但買", "賭博但買股", "博但買股票", "但買股票叫", "買股票叫投", "股票叫投資", "票叫投資？"},
		},
		{ // 1
			BBoardID:       "10_WhoAmI",
			ArticleID:      "1VtWRel9",
			BoardArticleID: "10_WhoAmI:1VtWRel9",

			IsDeleted:  false,
			CreateTime: 1608386280000000000,
			MTime:      1608386280000000000,
			Recommend:  9,
			Owner:      "SYSOP",
			Title:      "測試一下特殊字～",
			Class:      "心得",
			Money:      0,
			Filemode:   0,
			Idx:        "1608386280@1VtWRel9",
			FullTitle:  "[心得] 測試一下特殊字～",
			TitleRegex: []string{"測", "試", "一", "下", "特", "殊", "字", "～", "測試", "試一", "一下", "下特", "特殊", "殊字", "字～", "測試一", "試一下", "一下特", "下特殊", "特殊字", "殊字～", "測試一下", "試一下特", "一下特殊", "下特殊字", "特殊字～", "測試一下特", "試一下特殊", "一下特殊字", "下特殊字～"},
		},
		{ // 2
			BBoardID:       "10_WhoAmI",
			ArticleID:      "1VtW-QXT",
			BoardArticleID: "10_WhoAmI:1VtW-QXT",

			IsDeleted:  false,
			CreateTime: 1608388506000000000,
			MTime:      1608386280000000000,
			Recommend:  9,
			Owner:      "SYSOP",
			Title:      "所以特殊字真的是有綠色的∼",
			Class:      "閒聊",
			Money:      0,
			Filemode:   0,
			Idx:        "1608388506@1VtW-QXT",
			FullTitle:  "[閒聊] 所以特殊字真的是有綠色的∼",
			TitleRegex: []string{"所", "以", "特", "殊", "字", "真", "的", "是", "有", "綠", "色", "的", "∼", "所以", "以特", "特殊", "殊字", "字真", "真的", "的是", "是有", "有綠", "綠色", "色的", "的∼", "所以特", "以特殊", "特殊字", "殊字真", "字真的", "真的是", "的是有", "是有綠", "有綠色", "綠色的", "色的∼", "所以特殊", "以特殊字", "特殊字真", "殊字真的", "字真的是", "真的是有", "的是有綠", "是有綠色", "有綠色的", "綠色的∼", "所以特殊字", "以特殊字真", "特殊字真的", "殊字真的是", "字真的是有", "真的是有綠", "的是有綠色", "是有綠色的", "有綠色的∼"},
		},
		{ // 3
			BBoardID:       "10_WhoAmI",
			ArticleID:      "1Vo_N0CD",
			BoardArticleID: "10_WhoAmI:1Vo_N0CD",

			IsDeleted:  false,
			CreateTime: 1607202240000000000,
			MTime:      1607202240000000000,
			Recommend:  23,
			Owner:      "cheinshin",
			Title:      "TVBS六都民調 侯奪冠、盧升第四、柯墊底",
			Class:      "新聞",
			Money:      0,
			Filemode:   0,
			Idx:        "1607202240@1Vo_N0CD",
			FullTitle:  "[新聞] TVBS六都民調 侯奪冠、盧升第四、柯墊底",
			TitleRegex: []string{"T", "V", "B", "S", "六", "都", "民", "調", " ", "侯", "奪", "冠", "、", "盧", "升", "第", "四", "、", "柯", "墊", "底", "TV", "VB", "BS", "S六", "六都", "都民", "民調", "調 ", " 侯", "侯奪", "奪冠", "冠、", "、盧", "盧升", "升第", "第四", "四、", "、柯", "柯墊", "墊底", "TVB", "VBS", "BS六", "S六都", "六都民", "都民調", "民調 ", "調 侯", " 侯奪", "侯奪冠", "奪冠、", "冠、盧", "、盧升", "盧升第", "升第四", "第四、", "四、柯", "、柯墊", "柯墊底", "TVBS", "VBS六", "BS六都", "S六都民", "六都民調", "都民調 ", "民調 侯", "調 侯奪", " 侯奪冠", "侯奪冠、", "奪冠、盧", "冠、盧升", "、盧升第", "盧升第四", "升第四、", "第四、柯", "四、柯墊", "、柯墊底", "TVBS六", "VBS六都", "BS六都民", "S六都民調", "六都民調 ", "都民調 侯", "民調 侯奪", "調 侯奪冠", " 侯奪冠、", "侯奪冠、盧", "奪冠、盧升", "冠、盧升第", "、盧升第四", "盧升第四、", "升第四、柯", "第四、柯墊", "四、柯墊底"},
		},
		{ // 4
			BBoardID:       "10_WhoAmI",
			ArticleID:      "1VrooM21",
			BoardArticleID: "10_WhoAmI:1VrooM21",

			IsDeleted:  false,
			CreateTime: 1607937174000000000,
			MTime:      1607937100000000000,
			Recommend:  3,
			Owner:      "teemo",
			Title:      "新書的情報",
			Class:      "閒聊",
			Money:      0,
			Filemode:   0,
			Idx:        "1607937174@1VrooM21",
			FullTitle:  "[閒聊] 新書的情報",
			TitleRegex: []string{"新", "書", "的", "情", "報", "新書", "書的", "的情", "情報", "新書的", "書的情", "的情報", "新書的情", "書的情報", "新書的情報"},
		},
		{ // 5
			BBoardID:       "10_WhoAmI",
			ArticleID:      "19bWBI4Z",
			BoardArticleID: "10_WhoAmI:19bWBI4Z",

			IsDeleted:  false,
			CreateTime: 1234567890000000000,
			MTime:      1234567889000000000,
			Recommend:  8,
			Owner:      "okcool",
			Title:      "然後呢？～",
			Class:      "問題",
			Money:      0,
			Filemode:   0,
			Idx:        "1234567890@19bWBI4Z",
			FullTitle:  "[問題]然後呢？～",
			TitleRegex: []string{"然", "後", "呢", "？", "～", "然後", "後呢", "呢？", "？～", "然後呢", "後呢？", "呢？～", "然後呢？", "後呢？～", "然後呢？～"},
		},
		{ // 6
			BBoardID:       "10_WhoAmI",
			ArticleID:      "19bUG021",
			BoardArticleID: "10_WhoAmI:19bUG021",

			IsDeleted:  false,
			CreateTime: 1234560000000000000,
			MTime:      1234560000000000000,
			Recommend:  13,
			Owner:      "SYSOP",
			Title:      "這是 SYSOP",
			Class:      "問題",
			Money:      0,
			Filemode:   0,
			Idx:        "1234560000@19bUG021",
			FullTitle:  "[問題]這是 SYSOP",
			TitleRegex: []string{"這", "是", " ", "S", "Y", "S", "O", "P", "這是", "是 ", " S", "SY", "YS", "SO", "OP", "這是 ", "是 S", " SY", "SYS", "YSO", "SOP", "這是 S", "是 SY", " SYS", "SYSO", "YSOP", "這是 SY", "是 SYS", " SYSO", "SYSOP"},
		},
	}

	type args struct {
		boardID  bbs.BBoardID
		startIdx int32
	}
	tests := []struct {
		name                     string
		args                     args
		expectedArticleSummaries []*schema.ArticleSummaryWithRegex
		expectedNextIdx          int32
		wantErr                  bool
	}{
		// TODO: Add test cases.
		{
			args:                     args{boardID: "10_WhoAmI"},
			expectedArticleSummaries: expected0,
			expectedNextIdx:          -1,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotArticleSummaries, gotNextIdx, err := loadGeneralArticlesCoreBoardd(tt.args.boardID, tt.args.startIdx)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadGeneralArticles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, each := range gotArticleSummaries {
				each.UpdateNanoTS = 0
			}
			testutil.TDeepEqual(t, "got", gotArticleSummaries, tt.expectedArticleSummaries)
			if gotNextIdx != tt.expectedNextIdx {
				t.Errorf("loadGeneralArticles() gotNextIdx = %v, want %v", gotNextIdx, tt.expectedNextIdx)
			}
		})
		wg.Wait()
	}
}

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
