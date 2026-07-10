package api

import (
	"os"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/types"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
)

var (
	testUserInfoResult0 = &GetUserInfoResult{
		UserID:   "SYSOP",
		Username: "SYSOP",
		Realname: "CodingMan",
		Nickname: "神",

		Uflag:        ptttype.UF_CURSOR_ASCII | ptttype.UF_DBCS_DROP_REPEAT | ptttype.UF_DBCS_AWARE | ptttype.UF_ADBANNER | ptttype.UF_BRDSORT,
		Userlevel:    ptttype.PERM_SYSSUBOP | ptttype.PERM_BM | ptttype.PERM_BASIC | ptttype.PERM_CHAT | ptttype.PERM_PAGE | ptttype.PERM_POST | ptttype.PERM_LOGINOK | ptttype.PERM_SYSOP,
		Numlogindays: 2,
		Numposts:     0,
		Firstlogin:   1600681288,
		Lastlogin:    1600756094,
		LastIP:       "59.124.167.226",
		LastHost:     "59.124.167.226",

		Money:    0,
		PttEmail: "",
		Justify:  "",
		Over18:   true,

		PagerUIType: 0,
		Pager:       1,
		Invisible:   false,
		Exmailbox:   0,

		Career:        "全景軟體",
		Role:          0,
		LastSeen:      1600681288,
		TimeSetAngel:  0,
		TimePlayAngel: 0,

		LastSong:  0,
		LoginView: 0,

		Vlcount:  0,
		FiveWin:  0,
		FiveLose: 0,

		Signature: 0,

		BadPost: 0,
		MyAngel: "",

		ChessEloRating: 0,

		TokenUser: "SYSOP",
	}

	testTitle0 = "新的目錄"

	testBoardSummary9 = &apitypes.BoardSummary{
		FBoardID:  "test9",
		Brdname:   "test9",
		Title:     "測試9",
		BrdAttr:   0,
		BoardType: "◎",
		Category:  "測試",
		NUser:     100,
		BMs:       []bbs.UUserID{"okcool", "teemo"},
		Total:     123,

		LastPostTime: 1234567890,
		StatAttr:     ptttype.NBRD_BOARD,
		Idx:          "3",
		Gid:          3,
		Bid:          9,
		Fav:          true,

		URL: "/board/test9/articles",

		TokenUser: "SYSOP",
	}

	testBoardSummary8 = &apitypes.BoardSummary{
		FBoardID:  "test8",
		Brdname:   "test8",
		Title:     "測試8",
		BrdAttr:   0,
		BoardType: "◎",
		Category:  "測試",
		NUser:     101,
		BMs:       []bbs.UUserID{"okcool2", "teemo2"},
		Total:     124,

		LastPostTime: 1300000000,
		StatAttr:     ptttype.NBRD_BOARD,
		Idx:          "5",
		Gid:          3,
		Bid:          8,
		Fav:          true,

		URL: "/board/test8/articles",

		TokenUser: "SYSOP",
	}

	testBoardSummary10 = &apitypes.BoardSummary{
		FBoardID:  "WhoAmI",
		Brdname:   "WhoAmI",
		Title:     "呵呵，猜猜我是誰！",
		BrdAttr:   0,
		BoardType: "◎",
		Category:  "嘰哩",
		NUser:     0,
		BMs:       []bbs.UUserID{},
		Total:     0,

		LastPostTime: 0,
		StatAttr:     ptttype.NBRD_FAV,
		Idx:          "6",
		Gid:          5,
		Bid:          10,

		URL: "/board/WhoAmI/articles",

		TokenUser: "SYSOP",
	}

	testFavoriteBoards0 = []*apitypes.BoardSummary{
		{StatAttr: ptttype.NBRD_LINE, Idx: "0"},
		{Title: testTitle0, StatAttr: ptttype.NBRD_FOLDER, LevelIdx: ":1", Idx: "1", URL: "/user/SYSOP/favorites?level_idx=:1"},
		{StatAttr: ptttype.NBRD_LINE, Idx: "2"},
		testBoardSummary9,
		{Title: testTitle0, StatAttr: ptttype.NBRD_FOLDER, LevelIdx: ":4", Idx: "4", URL: "/user/SYSOP/favorites?level_idx=:4"},
		testBoardSummary8,
	}

	testDeleteBoardSummary9 = &apitypes.BoardSummary{
		FBoardID:  "test9",
		Brdname:   "test9",
		Title:     "測試9",
		BrdAttr:   0,
		BoardType: "◎",
		Category:  "測試",
		NUser:     100,
		BMs:       []bbs.UUserID{"okcool", "teemo"},
		Total:     123,

		LastPostTime: 1234567890,
		StatAttr:     ptttype.NBRD_BOARD,
		Idx:          "2",
		Gid:          3,
		Bid:          9,
		Fav:          true,

		URL: "/board/test9/articles",

		TokenUser: "SYSOP",
	}

	testDeleteBoardSummary8 = &apitypes.BoardSummary{
		FBoardID:  "test8",
		Brdname:   "test8",
		Title:     "測試8",
		BrdAttr:   0,
		BoardType: "◎",
		Category:  "測試",
		NUser:     101,
		BMs:       []bbs.UUserID{"okcool2", "teemo2"},
		Total:     124,

		LastPostTime: 1300000000,
		StatAttr:     ptttype.NBRD_BOARD,
		Idx:          "4",
		Gid:          3,
		Bid:          8,
		Fav:          true,

		URL: "/board/test8/articles",

		TokenUser: "SYSOP",
	}

	testDeleteFavoriteBoards0 = []*apitypes.BoardSummary{
		{Title: testTitle0, StatAttr: ptttype.NBRD_FOLDER, LevelIdx: ":0", Idx: "0", URL: "/user/SYSOP/favorites?level_idx=:0"},
		{StatAttr: ptttype.NBRD_LINE, Idx: "1"},
		testDeleteBoardSummary9,
		{Title: testTitle0, StatAttr: ptttype.NBRD_FOLDER, LevelIdx: ":3", Idx: "3", URL: "/user/SYSOP/favorites?level_idx=:3"},
		testDeleteBoardSummary8,
	}

	testUserSYSOP_b = pttbbsapi.GetUserResult(&bbs.Userec{
		Version:  4194,
		UUserID:  bbs.UUserID("SYSOP"),
		Username: "SYSOP",
		Realname: []byte{ // CodingMan
			0x43, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x6e,
		},
		Nickname: []byte{0xaf, 0xab}, // 神

		Uflag:        ptttype.UF_CURSOR_ASCII | ptttype.UF_DBCS_DROP_REPEAT | ptttype.UF_DBCS_AWARE | ptttype.UF_ADBANNER | ptttype.UF_BRDSORT,
		Userlevel:    ptttype.PERM_SYSSUBOP | ptttype.PERM_BM | ptttype.PERM_BASIC | ptttype.PERM_CHAT | ptttype.PERM_PAGE | ptttype.PERM_POST | ptttype.PERM_LOGINOK | ptttype.PERM_SYSOP,
		Numlogindays: 2,
		Numposts:     0,
		Firstlogin:   1600681288,
		Lastlogin:    1600756094,
		Lasthost:     "59.124.167.226",
		/*
			Address: []byte{ //新竹縣子虛鄉烏有村543號
				0xb7, 0x73, 0xa6, 0xcb, 0xbf, 0xa4, 0xa4, 0x6c, 0xb5, 0xea,
				0xb6, 0x6d, 0xaf, 0x51, 0xa6, 0xb3, 0xa7, 0xf8, 0x35, 0x34,
				0x33, 0xb8, 0xb9,
			},
		*/
		Over18:   true,
		Pager:    ptttype.PAGER_ON,
		Career:   []byte{0xa5, 0xfe, 0xb4, 0xba, 0xb3, 0x6e, 0xc5, 0xe9}, // 全景軟體
		LastSeen: 1600681288,
	})

	testUserChhsiao123_b = pttbbsapi.GetUserResult(&bbs.Userec{
		Version:      ptttype.PASSWD_VERSION,
		UUserID:      bbs.UUserID("chhsiao123"),
		Username:     "chhsiao123",
		Lasthost:     "127.0.0.1",
		Uflag:        ptttype.UF_CURSOR_ASCII | ptttype.UF_DBCS_DROP_REPEAT | ptttype.UF_DBCS_AWARE | ptttype.UF_ADBANNER | ptttype.UF_BRDSORT,
		Userlevel:    ptttype.PERM_BASIC | ptttype.PERM_CHAT | ptttype.PERM_PAGE | ptttype.PERM_POST | ptttype.PERM_LOGINOK,
		Firstlogin:   1600681290,
		Lastlogin:    1600681290,
		Numlogindays: 1,
		Pager:        ptttype.PAGER_ON,
		Over18:       true,
		LastSeen:     1600681290,
	})

	testBoardSummaryWhoAmI_b = &bbs.BoardSummary{
		Gid:      5,
		Bid:      10,
		BBoardID: bbs.BBoardID("10_WhoAmI"),
		StatAttr: ptttype.NBRD_FAV,
		Brdname:  "WhoAmI",
		BoardClass: []byte{
			0xbc, 0x54, 0xad, 0xf9,
		},
		RealTitle: []byte{
			0xa8, 0xfe, 0xa8, 0xfe, 0xa1, 0x41, 0xb2, 0x71, 0xb2, 0x71,
			0xa7, 0xda, 0xac, 0x4f, 0xbd, 0xd6, 0xa1, 0x49,
		},
		BoardType:  []byte{0xa1, 0xb7},
		BM:         []bbs.UUserID{},
		IdxByName:  "WhoAmI",
		IdxByClass: "vFSt-Q@WhoAmI",
	}

	testBoardSummarySYSOP_b = &bbs.BoardSummary{
		Gid:      2,
		Bid:      1,
		BBoardID: bbs.BBoardID("1_SYSOP"),
		StatAttr: ptttype.NBRD_FAV,
		Brdname:  "SYSOP",
		BoardClass: []byte{
			0xbc, 0x54, 0xad, 0xf9,
		},
		RealTitle: []byte{
			0xa1, 0xb7, 0xaf, 0xb8, 0xaa,
			0xf8, 0xa6, 0x6e, 0x21,
		},
		BoardType:  []byte{0xa1, 0xb7},
		BM:         []bbs.UUserID{},
		IdxByName:  "SYSOP",
		IdxByClass: "vFSt-Q@SYSOP",
	}

	testArticleSummary1_b = &bbs.ArticleSummary{
		BBoardID:   bbs.BBoardID("10_WhoAmI"),
		ArticleID:  bbs.ArticleID("1VrooM21"),
		IsDeleted:  false,
		Filename:   "M.1607937174.A.081",
		CreateTime: types.Time4(1607937174),
		MTime:      types.Time4(1607937100),
		Recommend:  3,
		Owner:      bbs.UUserID("teemo"),
		Class:      []byte{0xb0, 0xdd, 0xc3, 0x44},
		FullTitle:  []byte{0x5b, 0xb0, 0xdd, 0xc3, 0x44, 0x5d, 0xa6, 0x41, 0xa8, 0xd3, 0xa9, 0x4f, 0xa1, 0x48, 0xa1, 0xe3}, //[問題]再來呢？～
		Money:      12,
		Filemode:   0,
		Read:       false,
		Idx:        "1607937174@1VrooM21",
		RealTitle:  []byte{0xa6, 0x41, 0xa8, 0xd3, 0xa9, 0x4f, 0xa1, 0x48, 0xa1, 0xe3},
	}

	testArticleSummary2_b = &bbs.ArticleSummary{
		BBoardID:   bbs.BBoardID("10_WhoAmI"),
		ArticleID:  bbs.ArticleID("1VtWRel9"),
		IsDeleted:  false,
		Filename:   "M.1608386280.A.BC9",
		CreateTime: types.Time4(1608386280),
		MTime:      types.Time4(1608386280),
		Recommend:  8,
		Owner:      bbs.UUserID("SYSOP"),
		Class:      []byte{0xb0, 0xdd, 0xc3, 0x44},
		FullTitle:  []byte{0x5b, 0xb0, 0xdd, 0xc3, 0x44, 0x5d, 0xb5, 0x4d, 0xab, 0xe1, 0xa9, 0x4f, 0xa1, 0x48, 0xa1, 0xe3}, //[問題]然後呢？～
		Money:      3,
		Filemode:   0,
		Read:       false,
		Idx:        "1608386280@1VtWRel9",
		RealTitle:  []byte{0xb5, 0x4d, 0xab, 0xe1, 0xa9, 0x4f, 0xa1, 0x48, 0xa1, 0xe3},
	}
)

func initTest() {
	if testContentAll4 != nil {
		return
	}

	initTest3()
	initTestUtf83()
	initTest4()
	initTest11()
	initTestUtf85()
	initTestUtf86()
	initTestUtf88()
}

func loadTest(filename string) (contentAll []byte, content []byte, signature []byte, recommend []byte, firstComments []byte, theRestComments []byte) {
	// content-all
	fullFilename := "testcase/" + filename
	contentAll, err := os.ReadFile(fullFilename)
	if err != nil {
		return nil, nil, nil, nil, nil, nil
	}

	// content
	fullFilename = "testcase/" + filename + ".content"
	content, err = os.ReadFile(fullFilename)
	if err != nil {
		return nil, nil, nil, nil, nil, nil
	}

	if len(content) == 0 {
		content = nil
	}

	// signature
	fullFilename = "testcase/" + filename + ".signature"
	signature, err = os.ReadFile(fullFilename)
	if err != nil {
		return nil, nil, nil, nil, nil, nil
	}

	if len(signature) == 0 {
		signature = nil
	}

	// recommend
	fullFilename = "testcase/" + filename + ".recommend"
	recommend, err = os.ReadFile(fullFilename)
	if err != nil {
		return nil, nil, nil, nil, nil, nil
	}

	if len(recommend) == 0 {
		recommend = nil
	}

	// firstComments
	fullFilename = "testcase/" + filename + ".firstComments"
	firstComments, err = os.ReadFile(fullFilename)
	if err != nil {
		return nil, nil, nil, nil, nil, nil
	}

	if len(firstComments) == 0 {
		firstComments = nil
	}

	// theRestComments
	fullFilename = "testcase/" + filename + ".theRestComments"
	theRestComments, err = os.ReadFile(fullFilename)
	if err != nil {
		return nil, nil, nil, nil, nil, nil
	}

	if len(theRestComments) == 0 {
		theRestComments = nil
	}

	return contentAll, content, signature, recommend, firstComments, theRestComments
}
