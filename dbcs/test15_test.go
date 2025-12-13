package dbcs

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"

	"github.com/sirupsen/logrus"
)

var (
	testFilename15            = "temp10"
	testContentAll15          []byte
	testContent15             []byte
	testSignature15           []byte
	testComment15             []byte
	testFirstCommentsDBCS15   []byte
	testTheRestCommentsDBCS15 []byte
	testContent15Big5         [][]*types.Rune
	testContent15Utf8         [][]*types.Rune

	testFirstComments15     []*schema.Comment
	testFullFirstComments15 []*schema.Comment
)

func initTest15() {
	testContentAll15, testContent15, testSignature15, testComment15, testFirstCommentsDBCS15, testTheRestCommentsDBCS15 = loadTest(testFilename15)
	logrus.Infof("initTest15: testContentAll15: %v testContent15: %v testSignature15: %v testComment15: %v", len(testContentAll15), len(testContent15), len(testSignature15), len(testComment15))

	testContent15Big5 = [][]*types.Rune{
		{ // 0
			{
				Big5:   []byte("\xa7@\xaa\xcc: AAstar (\xa6\xd1\xad\xf4\xa4@\xaaT) \xac\xdd\xaaO: NBA"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: AAstar (\xa6\xd1\xad\xf4\xa4@\xaaT) \xac\xdd\xaaO: NBA\r"),
			},
		},
		{ // 1
			{
				Big5:   []byte("\xbc\xd0\xc3D: [\xc2\xe0\xbf\xfd][\xbc\xc6\xbe\xda] \xc0\xb0\xbe\xe3\xb2z\xa4F\xa4@\xa4UMJ\xabO\xab\xf9\xaa\xba\xa4E\xa4Q\xb6\xb5NBA\xb2\xc4\xa4@"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xc2\xe0\xbf\xfd][\xbc\xc6\xbe\xda] \xc0\xb0\xbe\xe3\xb2z\xa4F\xa4@\xa4UMJ\xabO\xab\xf9\xaa\xba\xa4E\xa4Q\xb6\xb5NBA\xb2\xc4\xa4@\r"),
			},
		},
		{ // 2
			{
				Big5:   []byte("\xae\xc9\xb6\xa1: Thu Jun 23 12:14:48 2016"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Thu Jun 23 12:14:48 2016\r"),
			},
		},
		{ // 3
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 4
			{
				Big5:   []byte("\xa1\xb0 [\xa5\xbb\xa4\xe5\xc2\xe0\xbf\xfd\xa6\xdb AIR_JORDAN \xac\xdd\xaaO #19UFg6e0 ]"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xb0 [\xa5\xbb\xa4\xe5\xc2\xe0\xbf\xfd\xa6\xdb AIR_JORDAN \xac\xdd\xaaO #19UFg6e0 ]\r"),
			},
		},
		{ // 5
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 6
			{
				Big5:   []byte("\xa7@\xaa\xcc: hate2004 (\xaf\xac\xa7A\xa6\xf2\xbd\xcf\xa4\xe9\xa7\xd6\xbc\xd6) \xac\xdd\xaaO: AIR_JORDAN"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: hate2004 (\xaf\xac\xa7A\xa6\xf2\xbd\xcf\xa4\xe9\xa7\xd6\xbc\xd6) \xac\xdd\xaaO: AIR_JORDAN\r"),
			},
		},
		{ // 7
			{
				Big5:   []byte("\xbc\xd0\xc3D: [\xc2\xe0\xbf\xfd][\xbc\xc6\xbe\xda] \xc0\xb0\xbe\xe3\xb2z\xa4F\xa4@\xa4UMJ\xabO\xab\xf9\xaa\xba\xa4E\xa4Q\xb6\xb5NBA\xb2\xc4\xa4@"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xc2\xe0\xbf\xfd][\xbc\xc6\xbe\xda] \xc0\xb0\xbe\xe3\xb2z\xa4F\xa4@\xa4UMJ\xabO\xab\xf9\xaa\xba\xa4E\xa4Q\xb6\xb5NBA\xb2\xc4\xa4@\r"),
			},
		},
		{ // 8
			{
				Big5:   []byte("\xae\xc9\xb6\xa1: Fri Jan 23 07:00:21 2009"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Fri Jan 23 07:00:21 2009\r"),
			},
		},
		{ // 9
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 10
			{
				Big5:   []byte("\xa1\xb0 [\xa5\xbb\xa4\xe5\xc2\xe0\xbf\xfd\xa6\xdb NBA \xac\xdd\xaaO]"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xb0 [\xa5\xbb\xa4\xe5\xc2\xe0\xbf\xfd\xa6\xdb NBA \xac\xdd\xaaO]\r"),
			},
		},
		{ // 11
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 12
			{
				Big5:   []byte("\xa7@\xaa\xcc: sk2g (\xb0\xb2\xafZ\xa5N\xa5\xac\xa5\xac) \xac\xdd\xaaO: NBA"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: sk2g (\xb0\xb2\xafZ\xa5N\xa5\xac\xa5\xac) \xac\xdd\xaaO: NBA\r"),
			},
		},
		{ // 13
			{
				Big5:   []byte("\xbc\xd0\xc3D: [\xbc\xc6\xbe\xda] MJ\xabO\xab\xf9\xaa\xba\xa4E\xa4Q\xb6\xb5NBA\xbe\xfa\xa5v\xb2\xc4\xa4@\xb0O\xbf\xfd"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xbc\xc6\xbe\xda] MJ\xabO\xab\xf9\xaa\xba\xa4E\xa4Q\xb6\xb5NBA\xbe\xfa\xa5v\xb2\xc4\xa4@\xb0O\xbf\xfd\r"),
			},
		},
		{ // 14
			{
				Big5:   []byte("\xae\xc9\xb6\xa1: Thu Jan 22 23:34:14 2009"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Thu Jan 22 23:34:14 2009\r"),
			},
		},
		{ // 15
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 16
			{
				Big5:   []byte("\xa4p\xa7\xcc\xa4j\xb7\xa7\xb4\xc0\xb3o\xa6\xec\xa4j\xa4j\xbe\xe3\xb2z\xa4@\xa4U"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa4p\xa7\xcc\xa4j\xb7\xa7\xb4\xc0\xb3o\xa6\xec\xa4j\xa4j\xbe\xe3\xb2z\xa4@\xa4U\r"),
			},
		},
		{ // 17
			{
				Big5:   []byte("\xa1\xb0 \xa4\xde\xadz\xa1mLaban (\xb0\xc7\xa5\xbb)\xa1n\xa4\xa7\xbb\xca\xa8\xa5\xa1G"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xb0 \xa4\xde\xadz\xa1mLaban (\xb0\xc7\xa5\xbb)\xa1n\xa4\xa7\xbb\xca\xa8\xa5\xa1G\r"),
			},
		},
		{ // 18
			{
				Big5:   []byte("\xb1`\xb3W\xc1\xc9\xb0O\xbf\xfd\xa1G"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33;1m\xb1`\xb3W\xc1\xc9\xb0O\xbf\xfd\xa1G"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 19
			{
				Big5:   []byte("1.\xc2\xbe\xb7~\xa5\xcd\xb2P\xb3\xf5\xa7\xa130.123\xa4\xc0\xa1A\xbe\xfa\xa5v\xb2\xc4\xa4@\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("1.\xc2\xbe\xb7~\xa5\xcd\xb2P\xb3\xf5\xa7\xa130.123\xa4\xc0\xa1A\xbe\xfa\xa5v\xb2\xc4\xa4@\xa1C\r"),
			},
		},
		{ // 20
			{
				Big5:   []byte(" "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte(" "),
			},
			{
				Big5:   []byte("\xa1]\xb5\xf9:Chamberlain\xa5H30.06\xa9e\xa6C\xb2\xc4\xa4G\xa1^"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m\xa1]\xb5\xf9:Chamberlain\xa5H30.06\xa9e\xa6C\xb2\xc4\xa4G\xa1^"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 22
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 23
			{
				Big5:   []byte("2.\xad\xd3\xa4H10\xa6\xb8\xb1o\xa4\xc0\xa4\xfd\xa1A\xbe\xfa\xa5v\xb2\xc4\xa4@\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("2.\xad\xd3\xa4H10\xa6\xb8\xb1o\xa4\xc0\xa4\xfd\xa1A\xbe\xfa\xa5v\xb2\xc4\xa4@\xa1C\r"),
			},
		},
		{ // 24
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 25
			{
				Big5:   []byte("3.\xa5\xad\xa4FChamberlain\xb3s\xc4\xf2\xa4C\xa6\xb8\xb1o\xa4\xc0\xa4\xfd\xac\xf6\xbf\xfd\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("3.\xa5\xad\xa4FChamberlain\xb3s\xc4\xf2\xa4C\xa6\xb8\xb1o\xa4\xc0\xa4\xfd\xac\xf6\xbf\xfd\xa1C\r"),
			},
		},
		{ // 26
			{
				Big5:   []byte(" "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte(" "),
			},
			{
				Big5:   []byte("\xa1]\xb5\xf9:1987 - 1993\xa1^"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m\xa1]\xb5\xf9:1987 - 1993\xa1^"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 27
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 28
			{
				Big5:   []byte("4.\xb3s\xc4\xf2866\xb3\xf5\xa4\xf1\xc1\xc9\xb1o\xa4\xc0\xb6W\xb9L\xa8\xe2\xa6\xec\xbc\xc6\xa1A\xbe\xfa\xa5v\xb2\xc4\xa4@\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("4.\xb3s\xc4\xf2866\xb3\xf5\xa4\xf1\xc1\xc9\xb1o\xa4\xc0\xb6W\xb9L\xa8\xe2\xa6\xec\xbc\xc6\xa1A\xbe\xfa\xa5v\xb2\xc4\xa4@\xa1C\r"),
			},
		},
		{ // 29
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 30
			{
				Big5:   []byte("5.\xabO\xab\xf9\xb3\xe6\xb8`\xbb@\xb2y\xa9R\xa4\xa4NBA\xb3\xcc\xb0\xaa\xb0O\xbf\xfd14\xa6\xb8\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("5.\xabO\xab\xf9\xb3\xe6\xb8`\xbb@\xb2y\xa9R\xa4\xa4NBA\xb3\xcc\xb0\xaa\xb0O\xbf\xfd14\xa6\xb8\xa1C\r"),
			},
		},
		{ // 31
			{
				Big5:   []byte(" "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte(" "),
			},
			{
				Big5:   []byte("\xa1]\xb5\xf9\xa1G\xa6\xb3\xa8\xe2\xa6\xb8,\xa4@\xa6\xb8\xacO1989\xa6~11\xa4\xeb15\xa4\xe9\xa1A\xb9\xefUtah Jazz\xb3\xe6\xb8`\xbb@\xa4\xa414\xb2y\xa1A"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m\xa1]\xb5\xf9\xa1G\xa6\xb3\xa8\xe2\xa6\xb8,\xa4@\xa6\xb8\xacO1989\xa6~11\xa4\xeb15\xa4\xe9\xa1A\xb9\xefUtah Jazz\xb3\xe6\xb8`\xbb@\xa4\xa414\xb2y\xa1A\r"),
			},
		},
		{ // 32
			{
				Big5:   []byte("   \xa5t\xa4@\xa6\xb8\xacO1993\xa6~\xb9\xefMiami Heat\xa1C\xa1^"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("   \xa5t\xa4@\xa6\xb8\xacO1993\xa6~\xb9\xefMiami Heat\xa1C\xa1^"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 32
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 33
			{
				Big5:   []byte("6.\xabO\xab\xf9\xb5\xdb\xb3\xe6\xb8`\xbb@\xb2y\xa6\xb8\xbc\xc6\xb3\xcc\xb0\xaa16\xa6\xb8\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("6.\xabO\xab\xf9\xb5\xdb\xb3\xe6\xb8`\xbb@\xb2y\xa6\xb8\xbc\xc6\xb3\xcc\xb0\xaa16\xa6\xb8\xa1C\r"),
			},
		},
		{ // 34
			{
				Big5:   []byte(" \xa1]\xb5\xf9:1992\xa6~12\xa4\xeb30\xa4\xe9\xb9\xefMiami Heat\xaa\xba\xb2\xc44\xb8`\xa1AJordan\xa5\xfe\xb3\xf524\xbb@\xb2y\xa1C\xa1^"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m \xa1]\xb5\xf9:1992\xa6~12\xa4\xeb30\xa4\xe9\xb9\xefMiami Heat\xaa\xba\xb2\xc44\xb8`\xa1AJordan\xa5\xfe\xb3\xf524\xbb@\xb2y\xa1C\xa1^"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 35
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 36
			{
				Big5:   []byte("7.NBA\xbe\xfa\xa5v\xb3\xcc\xa6h\xaa\xba10\xad\xd3\xc1\xc9\xa9u\xaa\xba\xa7\xeb\xc4x\xa6\xb8\xbc\xc6\xa9M\xa7\xeb\xa4\xa4\xa6\xb8\xbc\xc6\xa7\xa1\xa6CNBA\xbe\xfa\xa5v\xb2\xc4\xa4@\xa6\xec\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("7.NBA\xbe\xfa\xa5v\xb3\xcc\xa6h\xaa\xba10\xad\xd3\xc1\xc9\xa9u\xaa\xba\xa7\xeb\xc4x\xa6\xb8\xbc\xc6\xa9M\xa7\xeb\xa4\xa4\xa6\xb8\xbc\xc6\xa7\xa1\xa6CNBA\xbe\xfa\xa5v\xb2\xc4\xa4@\xa6\xec\xa1C\r"),
			},
		},
		{ // 39
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 38
			{
				Big5:   []byte("8.1987\xa6~\xb9\xefAtlanta Hawks\xa4@\xbe\xd4\xb3\xd0\xa4U\xb3s\xc4\xf2\xb1o23\xa4\xc0\xaa\xbaNBA\xb0O\xbf\xfd\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("8.1987\xa6~\xb9\xefAtlanta Hawks\xa4@\xbe\xd4\xb3\xd0\xa4U\xb3s\xc4\xf2\xb1o23\xa4\xc0\xaa\xbaNBA\xb0O\xbf\xfd\xa1C\r"),
			},
		},
		{ // 39
			{
				Big5:   []byte(" "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte(" "),
			},
			{
				Big5:   []byte("\xa1]\xb5\xf91:\xb4N\xacO\xa7\xeb\xc4x\xbb@\xb2y\xa5\xfe\xb3\xa1\xa9R\xa4\xa4\xa1A\xa8\xc3\xa5B\xa6b\xb3o\xb4\xc1\xb6\xa1\xa5\xbb\xb6\xa4\xb5L\xa5\xf4\xa6\xf3\xa8\xe4\xa5L\xb2y\xad\xfb\xb1o\xa4\xc0\xa1C\xa1^"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m\xa1]\xb5\xf91:\xb4N\xacO\xa7\xeb\xc4x\xbb@\xb2y\xa5\xfe\xb3\xa1\xa9R\xa4\xa4\xa1A\xa8\xc3\xa5B\xa6b\xb3o\xb4\xc1\xb6\xa1\xa5\xbb\xb6\xa4\xb5L\xa5\xf4\xa6\xf3\xa8\xe4\xa5L\xb2y\xad\xfb\xb1o\xa4\xc0\xa1C\xa1^\r"),
			},
		},
		{ // 40
			{
				Big5:   []byte(" "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte(" "),
			},
			{
				Big5:   []byte(" (\xb5\xf92:\xb3\xe6\xb3\xf5\xb3s\xc4\xf2\xb1o\xa4\xc0\xb1\xc6\xa6\xe6\xba]\xabe\xa5|\xa7\xa1\xac\xb0Jordan\xa1G23\xa4\xc0"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m (\xb5\xf92:\xb3\xe6\xb3\xf5\xb3s\xc4\xf2\xb1o\xa4\xc0\xb1\xc6\xa6\xe6\xba]\xabe\xa5|\xa7\xa1\xac\xb0Jordan\xa1G23\xa4\xc0"),
			},
			{
				Big5:   []byte("\xa1]"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1]"),
			},
			{
				Big5:   []byte("1987"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[31m1987"),
			},
			{
				Big5:   []byte("\xa1^"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1^"),
			},
			{
				Big5:   []byte("\xa1A22\xa4\xc0"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m\xa1A22\xa4\xc0"),
			},
			{
				Big5:   []byte("\xa1]"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\x1b[m\xa1]"),
			},
			{
				Big5:   []byte("2002"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[31m2002"),
			},
			{
				Big5:   []byte("\xa1^"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1^\r"),
			},
		},
		{ // 41
			{
				Big5:   []byte(" "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte(" "),
			},
			{
				Big5:   []byte("      \xa1A19\xa4\xc0"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m      \xa1A19\xa4\xc0"),
			},
			{
				Big5:   []byte("\xa1]"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1]"),
			},
			{
				Big5:   []byte("1996"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[31m1996"),
			},
			{
				Big5:   []byte("\xa1^"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1^"),
			},
			{
				Big5:   []byte("\xa1A18\xa4\xc0"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m\xa1A18\xa4\xc0"),
			},
			{
				Big5:   []byte("\xa1]"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\x1b[m\xa1]"),
			},
			{
				Big5:   []byte("1987"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[31m1987"),
			},
			{
				Big5:   []byte("\xa1^"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1^"),
			},
			{
				Big5:   []byte("\xa1C)"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m\xa1C)"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 44
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 45
			{
				Big5:   []byte("19.10\xa6\xb8\xa4J\xbf\xefNBA\xa6~\xab\xd7\xb2\xc4\xa4@\xb6\xa4\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("19.10\xa6\xb8\xa4J\xbf\xefNBA\xa6~\xab\xd7\xb2\xc4\xa4@\xb6\xa4\xa1C\r"),
			},
		},
		{ // 46
			{
				Big5:   []byte("   (\xb5\xf9:1986-87\xa6\xdc1992-93\xc1\xc9\xa9u, 1995-96\xa6\xdc1997-98\xc1\xc9\xa9u\xa1C)"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m   (\xb5\xf9:1986-87\xa6\xdc1992-93\xc1\xc9\xa9u, 1995-96\xa6\xdc1997-98\xc1\xc9\xa9u\xa1C)"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 47
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 48
			{
				Big5:   []byte("20.87-88\xc1\xc9\xa9u\xaa\xfd\xa7\xf0\xa6\xb8\xbc\xc6\xb9F\xa8\xec131\xa6\xb8,\xacO\xab\xe1\xbd\xc3\xaa\xba\xaa\xfd\xa7\xf0\xb0O\xbf\xfd\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("20.87-88\xc1\xc9\xa9u\xaa\xfd\xa7\xf0\xa6\xb8\xbc\xc6\xb9F\xa8\xec131\xa6\xb8,\xacO\xab\xe1\xbd\xc3\xaa\xba\xaa\xfd\xa7\xf0\xb0O\xbf\xfd\xa1C\r"),
			},
		},
		{ // 49
			{
				Big5:   []byte("  \xa1]\xb5\xf9:\xb3\xe6\xb3\xf5\xb3\xcc\xb0\xaa6\xa6\xb8\xa1A\xc2\xbe\xb7~\xa5\xcd\xb2P\xa6\xb38\xb3\xf5\xa4\xf1\xc1\xc9\xb3\xe6\xb3\xf5\xaa\xfd\xa7\xf05\xa6\xb8\xa5H\xa4W\xa1C\xa1^"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m  \xa1]\xb5\xf9:\xb3\xe6\xb3\xf5\xb3\xcc\xb0\xaa6\xa6\xb8\xa1A\xc2\xbe\xb7~\xa5\xcd\xb2P\xa6\xb38\xb3\xf5\xa4\xf1\xc1\xc9\xb3\xe6\xb3\xf5\xaa\xfd\xa7\xf05\xa6\xb8\xa5H\xa4W\xa1C\xa1^"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 50
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 51
			{
				Big5:   []byte("39.\xb3\xe6\xa4@\xc1\xc9\xa9u\xa9u\xab\xe1\xc1\xc9\xa5\xad\xa7\xa1\xb3\xcc\xb0\xaa43.7\xa4\xc0\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("39.\xb3\xe6\xa4@\xc1\xc9\xa9u\xa9u\xab\xe1\xc1\xc9\xa5\xad\xa7\xa1\xb3\xcc\xb0\xaa43.7\xa4\xc0\xa1C\r"),
			},
		},
		{ // 52
			{
				Big5:   []byte("   (\xb5\xf9:1986\xa6~\xa1AJordan\xa6b\xad\xb1\xb9\xef\xb7\xed\xa6~67\xb3\xd3\xa1]\xa5D\xb3\xf540\xb3\xd31\xb1\xd1)\xaa\xba\xc1`\xaba\xadxBoston Celtics"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m   (\xb5\xf9:1986\xa6~\xa1AJordan\xa6b\xad\xb1\xb9\xef\xb7\xed\xa6~67\xb3\xd3\xa1]\xa5D\xb3\xf540\xb3\xd31\xb1\xd1)\xaa\xba\xc1`\xaba\xadxBoston Celtics\r"),
			},
		},
		{ // 53
			{
				Big5:   []byte("       \xaa\xba\xa4T\xb3\xf5\xa9u\xab\xe1\xc1\xc9\xa1A\xa4\xc0\xa7O\xae\xb3\xa4U49\xa4\xc0\xa1A63\xa4\xc0\xa1A19\xa4\xc0\xa1A\xb3\xf5\xa7\xa143.7\xa4\xc0\xacO\xb3\xe6\xa4@\xc1\xc9\xa9u\xa9u\xab\xe1\xc1\xc9\xb3\xf5\xa7\xa1"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("       \xaa\xba\xa4T\xb3\xf5\xa9u\xab\xe1\xc1\xc9\xa1A\xa4\xc0\xa7O\xae\xb3\xa4U49\xa4\xc0\xa1A63\xa4\xc0\xa1A19\xa4\xc0\xa1A\xb3\xf5\xa7\xa143.7\xa4\xc0\xacO\xb3\xe6\xa4@\xc1\xc9\xa9u\xa9u\xab\xe1\xc1\xc9\xb3\xf5\xa7\xa1\r"),
			},
		},
		{ // 54
			{
				Big5:   []byte("       \xb3\xcc\xb0\xaa\xb0O\xbf\xfd\xa1A\xbe\xfa\xa5v\xa4W\xa5u\xa6\xb3\xa8\xe2\xad\xd3\xa4H\xa6b\xa9u\xab\xe1\xc1\xc9\xb3\xf5\xa7\xa140\xa4\xc0\xa5H\xa4W\xa1A\xa5t\xa4@\xa6\xec\xacOJerry West"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("       \xb3\xcc\xb0\xaa\xb0O\xbf\xfd\xa1A\xbe\xfa\xa5v\xa4W\xa5u\xa6\xb3\xa8\xe2\xad\xd3\xa4H\xa6b\xa9u\xab\xe1\xc1\xc9\xb3\xf5\xa7\xa140\xa4\xc0\xa5H\xa4W\xa1A\xa5t\xa4@\xa6\xec\xacOJerry West\r"),
			},
		},
		{ // 55
			{
				Big5:   []byte("      \xa1]\xb3\xf5\xa7\xa140.6\xa4\xc0\xa1^)"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("      \xa1]\xb3\xf5\xa7\xa140.6\xa4\xc0\xa1^)"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 56
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 57
			{
				Big5:   []byte("40.\xa9u\xab\xe1\xc1\xc9\xb3\xe6\xb3\xf5\xb1o\xa4\xc063\xa4\xc0\xa1A\xbe\xfa\xa5v\xb2\xc4\xa4@\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("40.\xa9u\xab\xe1\xc1\xc9\xb3\xe6\xb3\xf5\xb1o\xa4\xc063\xa4\xc0\xa1A\xbe\xfa\xa5v\xb2\xc4\xa4@\xa1C\r"),
			},
		},
		{ // 58
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 59
			{
				Big5:   []byte("41.\xa9u\xab\xe1\xc1\xc9\xb0\xdf\xa4@\xa4@\xad\xd3\xae\xb3\xb9Lback to back 50\xa4\xc0\xaa\xba\xb2y\xad\xfb\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("41.\xa9u\xab\xe1\xc1\xc9\xb0\xdf\xa4@\xa4@\xad\xd3\xae\xb3\xb9Lback to back 50\xa4\xc0\xaa\xba\xb2y\xad\xfb\xa1C\r"),
			},
		},
		{ // 60
			{
				Big5:   []byte("   (\xb5\xf9:1988\xa6~\xa9u\xab\xe1\xc1\xc9\xa1AJordan\xa6b\xb9\xefCleveland Cavaliers\xaa\xba\xa4\xf1\xc1\xc9\xa4\xa4\xa1A\xb2\xc4\xa4@\xb3\xf550\xa4\xc0\xa1A"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m   (\xb5\xf9:1988\xa6~\xa9u\xab\xe1\xc1\xc9\xa1AJordan\xa6b\xb9\xefCleveland Cavaliers\xaa\xba\xa4\xf1\xc1\xc9\xa4\xa4\xa1A\xb2\xc4\xa4@\xb3\xf550\xa4\xc0\xa1A\r"),
			},
		},
		{ // 61
			{
				Big5:   []byte("       \xb2\xc4\xa4G\xb3\xf555\xa4\xc0\xa1A\xb3s\xc4\xf2\xa8\xe2\xb3\xf5\xae\xb3\xa4U50\xa4\xc0\xa1A\xbe\xfa\xa5v\xa4W\xa5u\xa6\xb3Jordan\xa1C)"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("       \xb2\xc4\xa4G\xb3\xf555\xa4\xc0\xa1A\xb3s\xc4\xf2\xa8\xe2\xb3\xf5\xae\xb3\xa4U50\xa4\xc0\xa1A\xbe\xfa\xa5v\xa4W\xa5u\xa6\xb3Jordan\xa1C)"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 62
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 63
			{
				Big5:   []byte("42.\xa9u\xab\xe1\xc1\xc9\xb0\xdf\xa4@\xa4@\xad\xd3\xae\xb3\xb9L\xb3s\xc4\xf23\xb3\xf545\xa4\xc0\xaa\xba\xb2y\xad\xfb\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("42.\xa9u\xab\xe1\xc1\xc9\xb0\xdf\xa4@\xa4@\xad\xd3\xae\xb3\xb9L\xb3s\xc4\xf23\xb3\xf545\xa4\xc0\xaa\xba\xb2y\xad\xfb\xa1C\r"),
			},
		},
		{ // 64
			{
				Big5:   []byte("   (\xb5\xf91:1990\xa6~\xa6b\xb9\xefPhiladelphia 76ers\xaa\xba\xa5b\xa8M\xc1\xc9\xa4\xa4\xa1AJordan5\xb3\xf5\xa4\xf1\xc1\xc9\xb3\xf5\xa7\xa1\xae\xb3\xa4U43.0\xa4\xc0"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m   (\xb5\xf91:1990\xa6~\xa6b\xb9\xefPhiladelphia 76ers\xaa\xba\xa5b\xa8M\xc1\xc9\xa4\xa4\xa1AJordan5\xb3\xf5\xa4\xf1\xc1\xc9\xb3\xf5\xa7\xa1\xae\xb3\xa4U43.0\xa4\xc0\r"),
			},
		},
		{ // 65
			{
				Big5:   []byte("        \xa1A\xa8\xe4\xa4\xa4\xa6b\xb2\xc4\xa4G\xa8\xec\xb2\xc44\xb3\xf5\xa4\xf1\xc1\xc9\xa4\xa4\xa1A\xa4\xc0\xa7O\xb1o\xa8\xec45\xa4\xc0\xa1A49\xa4\xc0\xa1A45\xa4\xc0\xb3s\xc4\xf2\xa4T\xb3\xf5\xa4\xf1\xc1\xc9\xb6W\xb9L45+\xa1A"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("        \xa1A\xa8\xe4\xa4\xa4\xa6b\xb2\xc4\xa4G\xa8\xec\xb2\xc44\xb3\xf5\xa4\xf1\xc1\xc9\xa4\xa4\xa1A\xa4\xc0\xa7O\xb1o\xa8\xec45\xa4\xc0\xa1A49\xa4\xc0\xa1A45\xa4\xc0\xb3s\xc4\xf2\xa4T\xb3\xf5\xa4\xf1\xc1\xc9\xb6W\xb9L45+\xa1A\r"),
			},
		},
		{ // 66
			{
				Big5:   []byte("        \xbe\xfa\xa5v\xa4W\xa5u\xa6\xb3Jordan\xa1C)"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("        \xbe\xfa\xa5v\xa4W\xa5u\xa6\xb3Jordan\xa1C)\r"),
			},
		},
		{ // 67
			{
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\r"),
			},
		},
		{ // 68
			{
				Big5:   []byte("  \xa1]\xb5\xf92:Jordan\xc1\xd9\xa6\xb3\xa6\xb8\xb3s\xc4\xf2\xb1o\xa8\xec44\xa4\xc0\xa1A50\xa4\xc0\xa9M44\xa4\xc0\xa1AIverson\xb4\xbf\xb8g\xb3s\xc4\xf2\xb1o\xa8\xec46\xa4\xc0\xa1A44\xa4\xc0\xa9M"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("  \xa1]\xb5\xf92:Jordan\xc1\xd9\xa6\xb3\xa6\xb8\xb3s\xc4\xf2\xb1o\xa8\xec44\xa4\xc0\xa1A50\xa4\xc0\xa9M44\xa4\xc0\xa1AIverson\xb4\xbf\xb8g\xb3s\xc4\xf2\xb1o\xa8\xec46\xa4\xc0\xa1A44\xa4\xc0\xa9M\r"),
			},
		},
		{ // 69
			{
				Big5:   []byte("        48\xa4\xc0\xa1A\xb3\xa3\xa5u\xaet\xa4@\xc2I\xa1C\xa1^"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("        48\xa4\xc0\xa1A\xb3\xa3\xa5u\xaet\xa4@\xc2I\xa1C\xa1^"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 70
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 69
			{
				Big5:   []byte("90.\xb0\xdf\xa4@\xaf\xe0\xc5\xfd\xc1p\xb7\xf9\xaf}\xa8\xd2\xa7\xe2\xa4w\xb8g\xb0h\xa7\xd0\xaa\xba\xb2y\xa6\xe7"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("90.\xb0\xdf\xa4@\xaf\xe0\xc5\xfd\xc1p\xb7\xf9\xaf}\xa8\xd2\xa7\xe2\xa4w\xb8g\xb0h\xa7\xd0\xaa\xba\xb2y\xa6\xe7"),
			},
			{
				Big5:   []byte("23\xb8\xb9"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[31;1m23\xb8\xb9"),
			},
			{
				Big5:   []byte("\xad\xab\xb7s\xac\xef\xa6^\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xad\xab\xb7s\xac\xef\xa6^\xa1C\r"),
			},
		},
		{ // 70
			{
				Big5:   []byte("   (\xb5\xf9:\xa4U\xad\xb1\xa4w\xa6\xb3\xaa\xa9\xa4\xcd\xb8\xc9\xa5R\xa4\xa3\xa5u\xa5L\xa4@\xa4H\xa4F\xa1AMagic Johnson 1996\xa6~\xb4_\xa5X\xae\xc9\xa4]\xacO\xa1C)"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m   (\xb5\xf9:\xa4U\xad\xb1\xa4w\xa6\xb3\xaa\xa9\xa4\xcd\xb8\xc9\xa5R\xa4\xa3\xa5u\xa5L\xa4@\xa4H\xa4F\xa1AMagic Johnson 1996\xa6~\xb4_\xa5X\xae\xc9\xa4]\xacO\xa1C)"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 71
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 72
			{
				Big5:   []byte("91.MJ\xacO\xb0\xdf\xa4@\xa6b\xa5L\xa4\xa3\xb4\xbf\xae\xc4\xa4O\xaa\xba\xb2y\xb6\xa4\xb3Q\xb1\xbe\xa4W\xb2y\xa6\xe7\xa6b\xb2y\xc0]("),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("91.MJ\xacO\xb0\xdf\xa4@\xa6b\xa5L\xa4\xa3\xb4\xbf\xae\xc4\xa4O\xaa\xba\xb2y\xb6\xa4\xb3Q\xb1\xbe\xa4W\xb2y\xa6\xe7\xa6b\xb2y\xc0]("),
			},
			{
				Big5:   []byte("Miami Heat"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[31;1mMiami Heat"),
			},
			{
				Big5:   []byte("\xaa\xba\xb2y\xc0])\xaa\xba\xb2y\xad\xfb~"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xaa\xba\xb2y\xc0])\xaa\xba\xb2y\xad\xfb~\r"),
			},
		},
		{ // 73
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 74
			{
				Big5:   []byte("92.Jordan\xc1\xd9\xacO\xb0\xdf\xa4@\xa5N\xaa\xed\xac\xfc\xb0\xea\xa4j\xbe\xc7\xc4x\xb2y\xb6\xa4\xa9M\xac\xfc\xb0\xea\xb9\xda\xa4\xa7\xb6\xa4\xa4\xc0\xa7O\xc0\xf2\xb1o\xb6\xf8\xb9B\xb7|\xaa\xf7\xb5P\xaa\xba\xb2y\xad\xfb~\xa6\xb9\xb0O\xbf\xfd\xaa`\xa9w"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("92.Jordan\xc1\xd9\xacO\xb0\xdf\xa4@\xa5N\xaa\xed\xac\xfc\xb0\xea\xa4j\xbe\xc7\xc4x\xb2y\xb6\xa4\xa9M\xac\xfc\xb0\xea\xb9\xda\xa4\xa7\xb6\xa4\xa4\xc0\xa7O\xc0\xf2\xb1o\xb6\xf8\xb9B\xb7|\xaa\xf7\xb5P\xaa\xba\xb2y\xad\xfb~\xa6\xb9\xb0O\xbf\xfd\xaa`\xa9w\r"),
			},
		},
		{ // 77
			{
				Big5:   []byte("   \xacO\xb5\xb4\xab\xe1\xaa\xba\xa4F,\xa6]\xac\xb0\xb2{\xa6b\xac\xfc\xb0\xea\xa4\xa3\xa5i\xaf\xe0\xa6A\xac\xa3\xa4j\xbe\xc7\xb6\xa4\xa5\xee\xb0\xd1\xa5[\xb6\xf8\xb9B\xb7|\xa4F~"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("   \xacO\xb5\xb4\xab\xe1\xaa\xba\xa4F,\xa6]\xac\xb0\xb2{\xa6b\xac\xfc\xb0\xea\xa4\xa3\xa5i\xaf\xe0\xa6A\xac\xa3\xa4j\xbe\xc7\xb6\xa4\xa5\xee\xb0\xd1\xa5[\xb6\xf8\xb9B\xb7|\xa4F~\r"),
			},
		},
		{ // 78
			{
				Big5:   []byte("   (\xb5\xf9:\xa4w\xa6\xb3\xaa\xa9\xa4\xcd\xb8\xc9\xa5R\xa4\xa3\xa5uJordan\xa4@\xa4H\xa1APatrick Ewing\xb8\xf2Chris Mullin\xb3\xa3\xb8\xf2Jordan\xa4@\xb0_\xa1C)"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m   (\xb5\xf9:\xa4w\xa6\xb3\xaa\xa9\xa4\xcd\xb8\xc9\xa5R\xa4\xa3\xa5uJordan\xa4@\xa4H\xa1APatrick Ewing\xb8\xf2Chris Mullin\xb3\xa3\xb8\xf2Jordan\xa4@\xb0_\xa1C)"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 79
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 80
			{
				Big5:   []byte("93.Jordan\xacO\xbe\xfa\xa5v\xa4W\xb2\xc4\xa4@\xad\xd3\xaf\xe0\xa6b\xb4_\xa5X\xab\xe1\xc1\xd9\xb1a\xbb\xe2\xb2y\xb6\xa4\xb3s\xb1o3\xa6\xb8NBA\xc1`\xaba\xadx\xaa\xba\xa4H\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("93.Jordan\xacO\xbe\xfa\xa5v\xa4W\xb2\xc4\xa4@\xad\xd3\xaf\xe0\xa6b\xb4_\xa5X\xab\xe1\xc1\xd9\xb1a\xbb\xe2\xb2y\xb6\xa4\xb3s\xb1o3\xa6\xb8NBA\xc1`\xaba\xadx\xaa\xba\xa4H\xa1C\r"),
			},
		},
		{ // 81
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 82
			{
				Big5:   []byte("94.02\xa6~\xb4_\xa5X\xae\xc9\xb1N\xa8\xba\xa6~\xaa\xba\xa9\xd2\xa6\xb3\xa4u\xb8\xea\xa1]\xa4j\xac\xf9100\xb8U\xac\xfc\xa4\xb8\xa1^\xae\xbd\xb5\xb9911\xa8\xfc\xc3\xf8\xaa\xcc\xa1A\xacONBA\xa5v\xa4W\xb0\xdf\xa4@\xa4@\xad\xd3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("94.02\xa6~\xb4_\xa5X\xae\xc9\xb1N\xa8\xba\xa6~\xaa\xba\xa9\xd2\xa6\xb3\xa4u\xb8\xea\xa1]\xa4j\xac\xf9100\xb8U\xac\xfc\xa4\xb8\xa1^\xae\xbd\xb5\xb9911\xa8\xfc\xc3\xf8\xaa\xcc\xa1A\xacONBA\xa5v\xa4W\xb0\xdf\xa4@\xa4@\xad\xd3\r"),
			},
		},
		{ // 83
			{
				Big5:   []byte("   \xae\xbd\xa6\xdb\xa4v\xa5\xfe\xa6~\xa4u\xb8\xea\xaa\xba\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("   \xae\xbd\xa6\xdb\xa4v\xa5\xfe\xa6~\xa4u\xb8\xea\xaa\xba\xa1C\r"),
			},
		},
		{ // 84
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 85
			{
				Big5:   []byte(": ------------------------------------------------------------------------------"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte(": ------------------------------------------------------------------------------\r"),
			},
		},
		{ // 86
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 87
			{
				Big5:   []byte(" "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte(" "),
			},
			{
				Big5:   []byte("Jordan : \xa7\xda\xb4N\xacO\xaf\xab!"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[31;1mJordan : \xa7\xda\xb4N\xacO\xaf\xab!"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 88
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 89
			{
				Big5:   []byte("--"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("--\r"),
			},
		},
		{ // 90
			{
				Big5:   []byte("\xa1\xb0 \xb5o\xabH\xaf\xb8: \xa7\xe5\xbd\xf0\xbd\xf0\xb9\xea\xb7~\xa7{(ptt.cc)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xb0 \xb5o\xabH\xaf\xb8: \xa7\xe5\xbd\xf0\xbd\xf0\xb9\xea\xb7~\xa7{(ptt.cc)\r"),
			},
		},
		{ // 91
			{
				Big5:   []byte("\xa1\xbb From: 61.228.169.158"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xbb From: 61.228.169.158\r"),
			},
		},
		{ // 92
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("chinhan1216"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mchinhan1216"),
			},
			{
				Big5:   []byte(":\xa7A\xcex....                                           "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa7A\xcex....                                           "),
			},
			{
				Big5:   []byte(" 01/22 23:35"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/22 23:35\r"),
			},
		},
		{ // 93
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("jasonlin68"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mjasonlin68"),
			},
			{
				Big5:   []byte(":\xa6n\xa4\xe5\xb1\xc0 \xa8\xaf\xadW\xa4F                                       "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa6n\xa4\xe5\xb1\xc0 \xa8\xaf\xadW\xa4F                                       "),
			},
			{
				Big5:   []byte(" 01/22 23:36"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/22 23:36\r"),
			},
		},
		{ // 92
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("Gigabye"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mGigabye"),
			},
			{
				Big5:   []byte(":\xa9u\xab\xe1\xc1\xc9\xa4~\xb3\xcc\xacO\xa5i\xa9\xc8\xaa\xba\xac\xf6\xbf\xfd \xaf\xe0\xb6q\xa5\xce\xa6b\xb3\xcc\xc3\xf6\xc1\xe4\xaa\xba\xae\xc9\xa8\xe8 \xb4N\xacO\xaf\xab     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa9u\xab\xe1\xc1\xc9\xa4~\xb3\xcc\xacO\xa5i\xa9\xc8\xaa\xba\xac\xf6\xbf\xfd \xaf\xe0\xb6q\xa5\xce\xa6b\xb3\xcc\xc3\xf6\xc1\xe4\xaa\xba\xae\xc9\xa8\xe8 \xb4N\xacO\xaf\xab     "),
			},
			{
				Big5:   []byte(" 01/22 23:37"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/22 23:37\r"),
			},
		},
		{ // 93
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("RonArtest93"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mRonArtest93"),
			},
			{
				Big5:   []byte(":PUSH                                               "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:PUSH                                               "),
			},
			{
				Big5:   []byte(" 01/22 23:38"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/22 23:38\r"),
			},
		},
		{ // 97
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("DenyPedrosa"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mDenyPedrosa"),
			},
			{
				Big5:   []byte(":\xaaG\xb5M\xacO\xaf\xab= = \xa4\xd3\xb1j\xa4F~                                "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xaaG\xb5M\xacO\xaf\xab= = \xa4\xd3\xb1j\xa4F~                                "),
			},
			{
				Big5:   []byte(" 01/22 23:39"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/22 23:39\r"),
			},
		},
		{ // 98
			{
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Big5:   []byte("iamdada"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33miamdada"),
			},
			{
				Big5:   []byte(":\xa4\xd3\xa5i\xa9\xc8\xa4F...\xa7A\xafu\xa6\xb3\xad@\xa9\xca                                  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa4\xd3\xa5i\xa9\xc8\xa4F...\xa7A\xafu\xa6\xb3\xad@\xa9\xca                                  "),
			},
			{
				Big5:   []byte(" 01/22 23:44"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/22 23:44\r"),
			},
		},
		{ // 96
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("ThBasketball"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mThBasketball"),
			},
			{
				Big5:   []byte(":\xb0\xdf\xa4@\xa6b\xa6P\xa4@\xb0\xa6\xb2y\xb6\xa4\xa6P\xa4@\xb3\xf5\xa4\xf1\xc1\xc9\xa4\xa4\xa1A2\xa4H\xa4@\xb0_\xa4j\xa4T\xa4\xb8\xa1C     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb0\xdf\xa4@\xa6b\xa6P\xa4@\xb0\xa6\xb2y\xb6\xa4\xa6P\xa4@\xb3\xf5\xa4\xf1\xc1\xc9\xa4\xa4\xa1A2\xa4H\xa4@\xb0_\xa4j\xa4T\xa4\xb8\xa1C     "),
			},
			{
				Big5:   []byte(" 01/22 23:44"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/22 23:44\r"),
			},
		},
		{ // 100
			{
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Big5:   []byte("ThBasketball"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mThBasketball"),
			},
			{
				Big5:   []byte(":\xb3o\xad\xd3Carter\xb8\xf2Kidd\xb0\xb5\xa8\xec\xa4F                            "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb3o\xad\xd3Carter\xb8\xf2Kidd\xb0\xb5\xa8\xec\xa4F                            "),
			},
			{
				Big5:   []byte(" 01/22 23:44"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/22 23:44\r"),
			},
		},
		{ // 101
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("mark0628"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mmark0628"),
			},
			{
				Big5:   []byte(":\xb4\xc2\xb8t\xb1\xc0 \xa7A\xa6n\xbb{\xafu                                       "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb4\xc2\xb8t\xb1\xc0 \xa7A\xa6n\xbb{\xafu                                       "),
			},
			{
				Big5:   []byte(" 01/22 23:47"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/22 23:47\r"),
			},
		},
		{ // 102
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("grf7186"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mgrf7186"),
			},
			{
				Big5:   []byte(":\xbb{\xafu \xb1\xc0\xa4@\xad\xd3~~                                          "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xbb{\xafu \xb1\xc0\xa4@\xad\xd3~~                                          "),
			},
			{
				Big5:   []byte(" 01/22 23:49"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/22 23:49\r"),
			},
		},
		{ // 100
			{
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Big5:   []byte("sk2g"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33msk2g"),
			},
			{
				Big5:   []byte(":\xa7\xda\xac\xddDarious Miles\xb4N\xa4\xa3\xa5u3\xa6\xb8\xb0\xd5XD  \xa5i\xacO\xa5L\xb8\xf2Jordan\xb2\xa6\xb3\xba\xa4\xa3\xa6P:)  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa7\xda\xac\xddDarious Miles\xb4N\xa4\xa3\xa5u3\xa6\xb8\xb0\xd5XD  \xa5i\xacO\xa5L\xb8\xf2Jordan\xb2\xa6\xb3\xba\xa4\xa3\xa6P:)  "),
			},
			{
				Big5:   []byte(" 01/23 00:15"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:15\r"),
			},
		},
		{ // 101
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("fanix"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mfanix"),
			},
			{
				Big5:   []byte(":\xa8\xd3\xa4H\xc2\xe0\xa5h\xaa\xfc\xc0\xef\xaaO!!                                         "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa8\xd3\xa4H\xc2\xe0\xa5h\xaa\xfc\xc0\xef\xaaO!!                                         "),
			},
			{
				Big5:   []byte(" 01/23 00:17"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:17\r"),
			},
		},
		{ // 106
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("CenaC"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mCenaC"),
			},
			{
				Big5:   []byte(":\xb1\xc0\"\xaf\xab\" \xaf\xab\xb8\xf2\xa4H\xacO\xa4\xa3\xa6P\xaa\xba \xa5L\xacO\xa6b\xa5t\xa4@\xad\xd3\xbch\xa6\xb8\xa5\xb4\xc4x\xb2y             "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb1\xc0\"\xaf\xab\" \xaf\xab\xb8\xf2\xa4H\xacO\xa4\xa3\xa6P\xaa\xba \xa5L\xacO\xa6b\xa5t\xa4@\xad\xd3\xbch\xa6\xb8\xa5\xb4\xc4x\xb2y             "),
			},
			{
				Big5:   []byte(" 01/23 00:22"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:22\r"),
			},
		},
		{ // 107
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("jjj1004"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mjjj1004"),
			},
			{
				Big5:   []byte(":\xab\xe7\xbb\xf2\xa4S\xb1qJONDAN\xc5\xdc\xa6\xa8KOBE\xbb\xb9\xa6\xe2\xa4\xa3\xa4\xd6                         "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xab\xe7\xbb\xf2\xa4S\xb1qJONDAN\xc5\xdc\xa6\xa8KOBE\xbb\xb9\xa6\xe2\xa4\xa3\xa4\xd6                         "),
			},
			{
				Big5:   []byte(" 01/23 00:22"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:22\r"),
			},
		},
		{ // 108
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("EXDOG"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mEXDOG"),
			},
			{
				Big5:   []byte(":\xb0\xdf\xa4@\xa5\xb4\xb9LNBA \xa5[\xa4WMLB\xb3o\xbc\xcb\xba\xe2\xa4\xa3\xba\xe2\xac\xf6\xbf\xfd?                       "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb0\xdf\xa4@\xa5\xb4\xb9LNBA \xa5[\xa4WMLB\xb3o\xbc\xcb\xba\xe2\xa4\xa3\xba\xe2\xac\xf6\xbf\xfd?                       "),
			},
			{
				Big5:   []byte(" 01/23 00:25"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:25\r"),
			},
		},
		{ // 105
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("AlmaMater"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mAlmaMater"),
			},
			{
				Big5:   []byte(":\xb3\xec\xa4\xa6\xacO\xaf\xab                                             "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb3\xec\xa4\xa6\xacO\xaf\xab                                             "),
			},
			{
				Big5:   []byte(" 01/23 00:25"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:25\r"),
			},
		},
		{ // 106
			{
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Big5:   []byte("fox0922"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mfox0922"),
			},
			{
				Big5:   []byte(":MLB\xa6\xb3\xad\xd3\xa7\xeb\xa4\xe2\xa6b\xa4p\xa4\xfb\xa5\xb4\xb9L Jordan\xa8S\xb6i\xb9LMLB\xa7a...             "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:MLB\xa6\xb3\xad\xd3\xa7\xeb\xa4\xe2\xa6b\xa4p\xa4\xfb\xa5\xb4\xb9L Jordan\xa8S\xb6i\xb9LMLB\xa7a...             "),
			},
			{
				Big5:   []byte(" 01/23 00:26"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:26\r"),
			},
		},
		{ // 112
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("luxylu"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mluxylu"),
			},
			{
				Big5:   []byte(":2002... \xb3\xa338\xb7\xb3\xa4F\xc1\xd9\xb3s\xc4\xf2\xb1o\xa4\xc022\xa4\xc0    \xaaG\xb5M\xbcF\xae`              "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:2002... \xb3\xa338\xb7\xb3\xa4F\xc1\xd9\xb3s\xc4\xf2\xb1o\xa4\xc022\xa4\xc0    \xaaG\xb5M\xbcF\xae`              "),
			},
			{
				Big5:   []byte(" 01/23 00:27"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:27\r"),
			},
		},
		{ // 108
			{
				Big5:   []byte("\xa1\xb0 \xbds\xbf\xe8: sk2g            \xa8\xd3\xa6\xdb: 61.228.169.158       (01/23 00:40)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xb0 \xbds\xbf\xe8: sk2g            \xa8\xd3\xa6\xdb: 61.228.169.158       (01/23 00:40)\r"),
			},
		},
		{ // 109
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("accprote"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33maccprote"),
			},
			{
				Big5:   []byte(":\xa5L\xa5u\xa6b\xabi\xa4h\xc1\xd9\xacO\xa5\xd5\xc4\xfb\xaa\xba2A\xab\xdd\xb9L\xa6\xd3\xa4w,\xae\xda\xa5\xbb\xa8S\xa4WMajor          "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa5L\xa5u\xa6b\xabi\xa4h\xc1\xd9\xacO\xa5\xd5\xc4\xfb\xaa\xba2A\xab\xdd\xb9L\xa6\xd3\xa4w,\xae\xda\xa5\xbb\xa8S\xa4WMajor          "),
			},
			{
				Big5:   []byte(" 01/23 00:31"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:31\r"),
			},
		},
		{ // 115
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("boatbear"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mboatbear"),
			},
			{
				Big5:   []byte(":\xa7\xda\xb7Q\xb0\xdd91~\xac\xb0\xa4\xb0\xbb\xf2\xbc\xf6\xa4\xf5\xadn\xb1\xbe\xa5L\xb2y\xa6\xe7?                        "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa7\xda\xb7Q\xb0\xdd91~\xac\xb0\xa4\xb0\xbb\xf2\xbc\xf6\xa4\xf5\xadn\xb1\xbe\xa5L\xb2y\xa6\xe7?                        "),
			},
			{
				Big5:   []byte(" 01/23 00:34"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:34\r"),
			},
		},
		{ // 111
			{
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Big5:   []byte("fox0922"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mfox0922"),
			},
			{
				Big5:   []byte(":\xa7\xda\xb0O\xb1o\xaao\xc0Y\xa6n\xb9\xb3\xbb\xa1\xa6]\xac\xb0\xa5L\xa4\xd3\xb0\xb6\xa4j\xa4\xa7\xc3\xfe\xaa\xba                     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa7\xda\xb0O\xb1o\xaao\xc0Y\xa6n\xb9\xb3\xbb\xa1\xa6]\xac\xb0\xa5L\xa4\xd3\xb0\xb6\xa4j\xa4\xa7\xc3\xfe\xaa\xba                     "),
			},
			{
				Big5:   []byte(" 01/23 00:37"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:37\r"),
			},
		},
		{ // 117
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("jiunyilee"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mjiunyilee"),
			},
			{
				Big5:   []byte(":http://0rz.tw/t1Qpb                                  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:http://0rz.tw/t1Qpb                                  "),
			},
			{
				Big5:   []byte(" 01/23 00:47"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:47\r"),
			},
		},
		{ // 118
			{
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Big5:   []byte("jiunyilee"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mjiunyilee"),
			},
			{
				Big5:   []byte(":http://0rz.tw/Ksn7u \xa5|\xb9\xef\xa4j\xa4T\xa4\xb8\xaa\xba\xbc\xc6\xbe\xda                 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:http://0rz.tw/Ksn7u \xa5|\xb9\xef\xa4j\xa4T\xa4\xb8\xaa\xba\xbc\xc6\xbe\xda                 "),
			},
			{
				Big5:   []byte(" 01/23 00:49"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:49\r"),
			},
		},
		{ // 119
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("joloucow"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mjoloucow"),
			},
			{
				Big5:   []byte(":58\xacO1993 :) \xb5\xa7\xbb~\xc5o~                                   "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:58\xacO1993 :) \xb5\xa7\xbb~\xc5o~                                   "),
			},
			{
				Big5:   []byte(" 01/23 00:57"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:57\r"),
			},
		},
		{ // 120
			{
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Big5:   []byte("sk2g"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33msk2g"),
			},
			{
				Big5:   []byte(":OOPS~~\xb9\xef\xadC!!\xc1\xc2\xc1\xc2\xb0\xd5XD                                      "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:OOPS~~\xb9\xef\xadC!!\xc1\xc2\xc1\xc2\xb0\xd5XD                                      "),
			},
			{
				Big5:   []byte(" 01/23 00:59"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:59\r"),
			},
		},
		{ // 121
			{
				Big5:   []byte("\xa1\xb0 \xbds\xbf\xe8: sk2g            \xa8\xd3\xa6\xdb: 61.228.169.158       (01/23 01:00)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xb0 \xbds\xbf\xe8: sk2g            \xa8\xd3\xa6\xdb: 61.228.169.158       (01/23 01:00)\r"),
			},
		},
		{ // 122
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("timohu"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mtimohu"),
			},
			{
				Big5:   []byte(":\xb2\xc458\xc1`\xa8M\xc1\xc9\xc0\xb3\xb8\xd3\xacO1993\xb9\xef\xa4\xd3\xb6\xa7\xaa\xba\xa6\xd1\xa4\xda\xa1\xe3\xa5i\xaf\xe0\xad\xecPO\xb5\xa7\xbb~\xa5\xb4\xa6\xa81983  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb2\xc458\xc1`\xa8M\xc1\xc9\xc0\xb3\xb8\xd3\xacO1993\xb9\xef\xa4\xd3\xb6\xa7\xaa\xba\xa6\xd1\xa4\xda\xa1\xe3\xa5i\xaf\xe0\xad\xecPO\xb5\xa7\xbb~\xa5\xb4\xa6\xa81983  "),
			},
			{
				Big5:   []byte(" 01/23 01:00"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:00\r"),
			},
		},
		{ // 123
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("kaiDX"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mkaiDX"),
			},
			{
				Big5:   []byte(":\xb1\xc0\xa6n\xa4\xe5 \xb7P\xc1\xc2\xad\xecPO                                          "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb1\xc0\xa6n\xa4\xe5 \xb7P\xc1\xc2\xad\xecPO                                          "),
			},
			{
				Big5:   []byte(" 01/23 01:09"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:09\r"),
			},
		},
		{ // 124
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("kalfan1"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mkalfan1"),
			},
			{
				Big5:   []byte(":\xa4\xd3\xaf\xab\xa4F NBA\xb4N\xac\xdd\xa5L\xaa\xed\xbat                                   "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa4\xd3\xaf\xab\xa4F NBA\xb4N\xac\xdd\xa5L\xaa\xed\xbat                                   "),
			},
			{
				Big5:   []byte(" 01/23 01:12"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:12\r"),
			},
		},
		{ // 120
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("clenny"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mclenny"),
			},
			{
				Big5:   []byte(":\xa6~\xab\xd7\xb2\xc4\xa4@\xb6\xa4\xb0O\xbf\xfd\xa6\xb3\xa5i\xaf\xe0\xb3QTD\xa5\xb4\xaf}\xa1ATD\xa5\xd8\xabe\xb2\xd6\xadp9\xa6\xb8(2\xa6\xb8\xb2\xc4\xa4G\xb6\xa4)  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa6~\xab\xd7\xb2\xc4\xa4@\xb6\xa4\xb0O\xbf\xfd\xa6\xb3\xa5i\xaf\xe0\xb3QTD\xa5\xb4\xaf}\xa1ATD\xa5\xd8\xabe\xb2\xd6\xadp9\xa6\xb8(2\xa6\xb8\xb2\xc4\xa4G\xb6\xa4)  "),
			},
			{
				Big5:   []byte(" 01/23 01:12"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:12\r"),
			},
		},
		{ // 121
			{
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Big5:   []byte("clenny"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mclenny"),
			},
			{
				Big5:   []byte(":\xa7\xf3\xa5\xbf\xa4@\xa4U\xa1A\xa5\xcd\xb2P\xb2\xc4\xa4@\xb6\xa4\xa1G\xa6\xd1\xb0\xa811\xa6\xb8\xa1B\xb8\xeb\xc5Q10\xa6\xb8\xa1A\xa9\xd2\xa5H\xb3o\xc0\xb3\xb8\xd3    "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa7\xf3\xa5\xbf\xa4@\xa4U\xa1A\xa5\xcd\xb2P\xb2\xc4\xa4@\xb6\xa4\xa1G\xa6\xd1\xb0\xa811\xa6\xb8\xa1B\xb8\xeb\xc5Q10\xa6\xb8\xa1A\xa9\xd2\xa5H\xb3o\xc0\xb3\xb8\xd3    "),
			},
			{
				Big5:   []byte(" 01/23 01:25"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:25\r"),
			},
		},
		{ // 122
			{
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Big5:   []byte("clenny"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mclenny"),
			},
			{
				Big5:   []byte(":\xa4\xa3\xaf\xe0\xba\xe2\xacO\xa5L\xbfW\xa6\xb3\xaa\xba\xb0O\xbf\xfdXD (\xc5]\xb3N\xa9M\xa4j\xb3\xbe\xa4]\xb3\xa3\xacO\xa5\xcd\xb2P9\xa6\xb8)        "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa4\xa3\xaf\xe0\xba\xe2\xacO\xa5L\xbfW\xa6\xb3\xaa\xba\xb0O\xbf\xfdXD (\xc5]\xb3N\xa9M\xa4j\xb3\xbe\xa4]\xb3\xa3\xacO\xa5\xcd\xb2P9\xa6\xb8)        "),
			},
			{
				Big5:   []byte(" 01/23 01:27"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:27\r"),
			},
		},
		{ // 123
			{
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Big5:   []byte("urinmymind"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33murinmymind"),
			},
			{
				Big5:   []byte(":\xa6\xb3\xb6m\xa5\xc1100\xb1\xf8\xb3\xa3\xac\xdd\xa7\xb9\xb6\xdc\xa1H  ~\"~                          "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa6\xb3\xb6m\xa5\xc1100\xb1\xf8\xb3\xa3\xac\xdd\xa7\xb9\xb6\xdc\xa1H  ~\"~                          "),
			},
			{
				Big5:   []byte(" 01/23 01:32"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:32\r"),
			},
		},
		{ // 124
			{
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Big5:   []byte("sk2g"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33msk2g"),
			},
			{
				Big5:   []byte(":\xa6\xb3\xb0\xda...\xa7\xda ...=_=....\xa6\xd3\xa5B\xacO94\xad\xf2                            "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa6\xb3\xb0\xda...\xa7\xda ...=_=....\xa6\xd3\xa5B\xacO94\xad\xf2                            "),
			},
			{
				Big5:   []byte(" 01/23 01:35"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:35\r"),
			},
		},
		{ // 125
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("MousePads"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mMousePads"),
			},
			{
				Big5:   []byte(":\xc3\xf6\xa9\xf3\xb2\xc422\xc2I\xa1A\xb3\xcc\xb0\xaa\xc0W\xb2v\xaa\xba\xb7N\xab\xe4\xacO\xa1HBig O\xaa\xba\xac\xf6\xbf\xfd\xacO\xa6h\xa4\xd6\xb0\xda\xa1H  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xc3\xf6\xa9\xf3\xb2\xc422\xc2I\xa1A\xb3\xcc\xb0\xaa\xc0W\xb2v\xaa\xba\xb7N\xab\xe4\xacO\xa1HBig O\xaa\xba\xac\xf6\xbf\xfd\xacO\xa6h\xa4\xd6\xb0\xda\xa1H  "),
			},
			{
				Big5:   []byte(" 01/23 01:36"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:36\r"),
			},
		},
		{ // 126
			{
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Big5:   []byte("sk2g"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33msk2g"),
			},
			{
				Big5:   []byte(":\xa7\xda\xb2q\xb4\xfa...\xa5i\xaf\xe0\xacO\xbb\xa1\xa5L\xa4j\xa4T\xa4\xb8\xaa\xba\xb3\xf5\xa6\xb8\xa4\xa4\xa1A\xb3s\xc4\xf2\xaa\xba\xa6\xb8\xbc\xc6\xa4\xf1\xa8\xd2\xa4\xa7\xb0\xaa\xa7a!  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa7\xda\xb2q\xb4\xfa...\xa5i\xaf\xe0\xacO\xbb\xa1\xa5L\xa4j\xa4T\xa4\xb8\xaa\xba\xb3\xf5\xa6\xb8\xa4\xa4\xa1A\xb3s\xc4\xf2\xaa\xba\xa6\xb8\xbc\xc6\xa4\xf1\xa8\xd2\xa4\xa7\xb0\xaa\xa7a!  "),
			},
			{
				Big5:   []byte(" 01/23 01:39"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:39\r"),
			},
		},
		{ // 127
			{
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Big5:   []byte("sk2g"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33msk2g"),
			},
			{
				Big5:   []byte(":\xa4j\xab\xd378\xa6\xb8\xb8\xcc\xad\xb1\xa6\xb39\xa6\xb8\xb3s\xc4\xf2\xa1A\xaf\xab30\xa6\xb8(\xa7tplayoff)\xa6\xb37\xa6\xb8...          "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa4j\xab\xd378\xa6\xb8\xb8\xcc\xad\xb1\xa6\xb39\xa6\xb8\xb3s\xc4\xf2\xa1A\xaf\xab30\xa6\xb8(\xa7tplayoff)\xa6\xb37\xa6\xb8...          "),
			},
			{
				Big5:   []byte(" 01/23 01:44"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:44\r"),
			},
		},
		{ // 128
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("MousePads"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mMousePads"),
			},
			{
				Big5:   []byte(":\xb3o\xbc\xcb\xaa\xba\xbb\xa1\xaak\xa6n\xa9\xc7\xa1A\xa8\xba\xa5\xcd\xb2P\xa5u\xa6\xb31\xa6\xb8\xaa\xba\xa4\xa3\xacO\xc0\xb3\xb8\xd3\xb3\xcc\xb0\xaa @@\"      "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb3o\xbc\xcb\xaa\xba\xbb\xa1\xaak\xa6n\xa9\xc7\xa1A\xa8\xba\xa5\xcd\xb2P\xa5u\xa6\xb31\xa6\xb8\xaa\xba\xa4\xa3\xacO\xc0\xb3\xb8\xd3\xb3\xcc\xb0\xaa @@\"      "),
			},
			{
				Big5:   []byte(" 01/23 01:57"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:57\r"),
			},
		},
		{ // 129
			{
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Big5:   []byte("sk2g"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33msk2g"),
			},
			{
				Big5:   []byte(":\xa5i\xacO\xb3s\xc4\xf21\xb3\xf5\xa9\xc7\xa9\xc7\xaa\xba\xa7aXD   \xb3s\xc4\xf22\xb3\xf5\xaa\xba\xa8\xba\xa7\xda\xc1\xd9\xc4\xb1\xb1o\xa4\xf1\xb8\xfb\xa9_\xa9\xc7@@     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa5i\xacO\xb3s\xc4\xf21\xb3\xf5\xa9\xc7\xa9\xc7\xaa\xba\xa7aXD   \xb3s\xc4\xf22\xb3\xf5\xaa\xba\xa8\xba\xa7\xda\xc1\xd9\xc4\xb1\xb1o\xa4\xf1\xb8\xfb\xa9_\xa9\xc7@@     "),
			},
			{
				Big5:   []byte(" 01/23 01:58"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:58\r"),
			},
		},
		{ // 130
			{
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Big5:   []byte("sk2g"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33msk2g"),
			},
			{
				Big5:   []byte(":\xba\xe2\xa4F\xa1A\xc5\xfd\xa7\xda\xad\xcc\xbe\xd6\xa9\xea\xaaO\xa4W\xaa\xba\xafu\xa3\xbb\xb1j\xaa\xcc\xa7a                          "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xba\xe2\xa4F\xa1A\xc5\xfd\xa7\xda\xad\xcc\xbe\xd6\xa9\xea\xaaO\xa4W\xaa\xba\xafu\xa3\xbb\xb1j\xaa\xcc\xa7a                          "),
			},
			{
				Big5:   []byte(" 01/23 01:59"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:59\r"),
			},
		},
		{ // 131
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("molemilk"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mmolemilk"),
			},
			{
				Big5:   []byte(":\xaf\xab                                                    "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xaf\xab                                                    "),
			},
			{
				Big5:   []byte(" 01/23 02:01"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 02:01\r"),
			},
		},
		{ // 132
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("blinder"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mblinder"),
			},
			{
				Big5:   []byte(":\xb6W\xbb{\xafu\xa6n\xa4\xe5~                                            "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb6W\xbb{\xafu\xa6n\xa4\xe5~                                            "),
			},
			{
				Big5:   []byte(" 01/23 02:13"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 02:13\r"),
			},
		},
		{ // 133
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("samprus"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33msamprus"),
			},
			{
				Big5:   []byte(":\xa7\xda\xc4\xb1\xb1o\xab\xe1\xbd\xc3\xaa\xba\xaa\xfd\xa7\xf0\xb0O\xbf\xfd\xa4\xb5\xa6~\xb7|\xb3Qwade\xa5\xb4\xaf}XD                 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa7\xda\xc4\xb1\xb1o\xab\xe1\xbd\xc3\xaa\xba\xaa\xfd\xa7\xf0\xb0O\xbf\xfd\xa4\xb5\xa6~\xb7|\xb3Qwade\xa5\xb4\xaf}XD                 "),
			},
			{
				Big5:   []byte(" 01/23 02:29"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 02:29\r"),
			},
		},
		{ // 134
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("avrilrock"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mavrilrock"),
			},
			{
				Big5:   []byte(":\xa6n\xa4\xe5                                                 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa6n\xa4\xe5                                                 "),
			},
			{
				Big5:   []byte(" 01/23 02:46"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 02:46\r"),
			},
		},
		{ // 135
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("Kaverson"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mKaverson"),
			},
			{
				Big5:   []byte(":\xb1\xc0\xa6n\xa4\xe5..XD...\xa8\xaf\xadW\xa7A\xa4F..:)                             "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb1\xc0\xa6n\xa4\xe5..XD...\xa8\xaf\xadW\xa7A\xa4F..:)                             "),
			},
			{
				Big5:   []byte(" 01/23 03:36"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 03:36\r"),
			},
		},
		{ // 136
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("Yao1218"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mYao1218"),
			},
			{
				Big5:   []byte(":\xc0\xb0\xb1\xc0 \xa6n\xa4\xe5                                              "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xc0\xb0\xb1\xc0 \xa6n\xa4\xe5                                              "),
			},
			{
				Big5:   []byte(" 01/23 03:42"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 03:42\r"),
			},
		},
		{ // 137
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("BrentRoy"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mBrentRoy"),
			},
			{
				Big5:   []byte(":\xb1\xc0\xa6n\xa4\xe5! \xa6\xfd\xacO\xa5i\xaf\xe0\xa8S\xb4X\xa4\xd1\xb4N\xb7|\xb3Q\xb0O\xaa\xcc\xa7\xdb\xa8\xab\xa4F= =             "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb1\xc0\xa6n\xa4\xe5! \xa6\xfd\xacO\xa5i\xaf\xe0\xa8S\xb4X\xa4\xd1\xb4N\xb7|\xb3Q\xb0O\xaa\xcc\xa7\xdb\xa8\xab\xa4F= =             "),
			},
			{
				Big5:   []byte(" 01/23 04:52"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 04:52\r"),
			},
		},
		{ // 138
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("ksk0516"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mksk0516"),
			},
			{
				Big5:   []byte(":\xa8C\xa6\xb8\xac\xdd\xa8\xec\xa6\xb3\xa4H\xbb\xa1Kobe\xa1BLBJ\xb6W\xb6VJordan\xb4N\xab\xdc\xb7Q\xaf\xba XD           "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa8C\xa6\xb8\xac\xdd\xa8\xec\xa6\xb3\xa4H\xbb\xa1Kobe\xa1BLBJ\xb6W\xb6VJordan\xb4N\xab\xdc\xb7Q\xaf\xba XD           "),
			},
			{
				Big5:   []byte(" 01/23 05:25"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 05:25\r"),
			},
		},
		{ // 139
			{
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Big5:   []byte("rumourmonger"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mrumourmonger"),
			},
			{
				Big5:   []byte(":Kobe\xb8\xf2LBJ\xa6p\xaaG\xa8S\xa6\xb3\xa6]\xb6\xcb\xb0h\xa5\xf0 \xb1N\xa8\xd3\xa5\xcd\xb2P\xa6\xa8\xc1Z\xa4\xf1\xaa\xd3MJ\xa5\xbf\xb1`  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:Kobe\xb8\xf2LBJ\xa6p\xaaG\xa8S\xa6\xb3\xa6]\xb6\xcb\xb0h\xa5\xf0 \xb1N\xa8\xd3\xa5\xcd\xb2P\xa6\xa8\xc1Z\xa4\xf1\xaa\xd3MJ\xa5\xbf\xb1`  "),
			},
			{
				Big5:   []byte(" 01/23 05:44"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 05:44\r"),
			},
		},
		{ // 140
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("chengyuyang"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mchengyuyang"),
			},
			{
				Big5:   []byte(":\xb1\xc0\xa6n\xa4\xe5n\xb3\xcc\xab\xe1\xa4@\xa5y      \xa6b\xa7O\xa4H\xaea\xb1\xbe\xb2y\xa6\xe7........\xaf\xab      "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb1\xc0\xa6n\xa4\xe5n\xb3\xcc\xab\xe1\xa4@\xa5y      \xa6b\xa7O\xa4H\xaea\xb1\xbe\xb2y\xa6\xe7........\xaf\xab      "),
			},
			{
				Big5:   []byte(" 01/23 05:45"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 05:45\r"),
			},
		},
		{ // 141
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("kobe872125"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mkobe872125"),
			},
			{
				Big5:   []byte(":\xa7\xda\xb0O\xb1o\xb4\xbf\xa6\xb3\xa4@\xb3\xf5\xb3\xec\xaf\xab\xb8\xf2\xa5\xd6\xaaB2\xad\xd3\xb3\xe6\xb3\xf540+ \xa6\xb3\xb6m\xa5\xc1\xb0O\xb1o\xb6\xdc\xa1H   "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa7\xda\xb0O\xb1o\xb4\xbf\xa6\xb3\xa4@\xb3\xf5\xb3\xec\xaf\xab\xb8\xf2\xa5\xd6\xaaB2\xad\xd3\xb3\xe6\xb3\xf540+ \xa6\xb3\xb6m\xa5\xc1\xb0O\xb1o\xb6\xdc\xa1H   "),
			},
			{
				Big5:   []byte(" 01/23 06:09"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 06:09\r"),
			},
		},
		{ // 142
			{
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Big5:   []byte("kobe872125"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mkobe872125"),
			},
			{
				Big5:   []byte(":\xb4N\xacO\xa4@\xb3\xf5\xa4\xf1\xc1\xc9\xa8\xe2\xad\xd3\xa4H\xb3\xa3\xb1o40+\xa4\xc0.\xa4\xa3\xaa\xbe\xa7\xda\xa6\xb3\xa8S\xa6\xb3\xb0O\xbf\xf9        "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb4N\xacO\xa4@\xb3\xf5\xa4\xf1\xc1\xc9\xa8\xe2\xad\xd3\xa4H\xb3\xa3\xb1o40+\xa4\xc0.\xa4\xa3\xaa\xbe\xa7\xda\xa6\xb3\xa8S\xa6\xb3\xb0O\xbf\xf9        "),
			},
			{
				Big5:   []byte(" 01/23 06:11"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 06:11\r"),
			},
		},
		{ // 143
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("henrie"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mhenrie"),
			},
			{
				Big5:   []byte(":\xa8\xbe...\xa8\xbe\xa6u\xa4\xfd?!                                           "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa8\xbe...\xa8\xbe\xa6u\xa4\xfd?!                                           "),
			},
			{
				Big5:   []byte(" 01/23 06:39"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 06:39\r"),
			},
		},
		{ // 144
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("chengyuyang"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mchengyuyang"),
			},
			{
				Big5:   []byte(":\xa8\xba\xb3\xf5\xadW\xa5DPacers \xb3\xec44 \xa5\xd640                           "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa8\xba\xb3\xf5\xadW\xa5DPacers \xb3\xec44 \xa5\xd640                           "),
			},
			{
				Big5:   []byte(" 01/23 06:42"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 06:42\r"),
			},
		},
		{ // 145
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("chengyuyang"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mchengyuyang"),
			},
			{
				Big5:   []byte(":\xac\xa1\xc1\xc9\xa9M\xaa\xf7\xb6\xf4 \xa5v\xa4W\xb3\xcc\xb0\xaa\xa4\xc0\xa8\xba\xb3\xf5 \xa8\xe2\xb6\xa4\xb3\xa3\xa6\xb3\xa8\xe2\xad\xd3\xa4H\xb6W\xb9L40     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xac\xa1\xc1\xc9\xa9M\xaa\xf7\xb6\xf4 \xa5v\xa4W\xb3\xcc\xb0\xaa\xa4\xc0\xa8\xba\xb3\xf5 \xa8\xe2\xb6\xa4\xb3\xa3\xa6\xb3\xa8\xe2\xad\xd3\xa4H\xb6W\xb9L40     "),
			},
			{
				Big5:   []byte(" 01/23 06:53"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 06:53\r"),
			},
		},
		{ // 146
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 147
			{
				Big5:   []byte("--"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("--\r"),
			},
		},
		{ // 148
			{
				Big5:   []byte("\xa1\xb0 \xb5o\xabH\xaf\xb8: \xa7\xe5\xbd\xf0\xbd\xf0\xb9\xea\xb7~\xa7{(ptt.cc) "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xb0 \xb5o\xabH\xaf\xb8: \xa7\xe5\xbd\xf0\xbd\xf0\xb9\xea\xb7~\xa7{(ptt.cc) \r"),
			},
		},
		{ // 149
			{
				Big5:   []byte("\xa1\xbb From: 61.228.174.209"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xbb From: 61.228.174.209\r"),
			},
		},
		{ // 150
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("airLfly"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mairLfly"),
			},
			{
				Big5:   []byte(":\xafu\xa4\xd4\xb9\xd8 \xa4\xd3\xb2r\xa4F                                          "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xafu\xa4\xd4\xb9\xd8 \xa4\xd3\xb2r\xa4F                                          "),
			},
			{
				Big5:   []byte(" 01/23 21:32"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 21:32\r"),
			},
		},
		{ // 151
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("chron"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mchron"),
			},
			{
				Big5:   []byte(":\xa5\xc3\xbb\xb7\xaa\xba\xaf\xab.....                                            "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa5\xc3\xbb\xb7\xaa\xba\xaf\xab.....                                            "),
			},
			{
				Big5:   []byte(" 01/29 12:40"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/29 12:40\r"),
			},
		},
		{ // 152
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("s9588008"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33ms9588008"),
			},
			{
				Big5:   []byte(":\xa9u\xab\xe1\xc1\xc9\xac\xdd\xa8\xec\xa4@\xa5b\xb4N\xa8\xd3\xb1\xc0\xa4F~ \xb9\xea\xa6b\xa4\xd3\xb1j\xa4j\xa4F!!   Q_Q          "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa9u\xab\xe1\xc1\xc9\xac\xdd\xa8\xec\xa4@\xa5b\xb4N\xa8\xd3\xb1\xc0\xa4F~ \xb9\xea\xa6b\xa4\xd3\xb1j\xa4j\xa4F!!   Q_Q          "),
			},
			{
				Big5:   []byte(" 02/05 06:23"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 02/05 06:23\r"),
			},
		},
		{ // 153
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("s9588008"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33ms9588008"),
			},
			{
				Big5:   []byte(":\xbd\xec\xa8\xfd\xac\xf6\xbf\xfd:  89.\xb0\xdf\xa4@\xb0\xb5\xa8\xec\xa4T\xb6i\xa4T\xa5X\xaa\xba\xa4H\xa1C  <-   XD         "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xbd\xec\xa8\xfd\xac\xf6\xbf\xfd:  89.\xb0\xdf\xa4@\xb0\xb5\xa8\xec\xa4T\xb6i\xa4T\xa5X\xaa\xba\xa4H\xa1C  <-   XD         "),
			},
			{
				Big5:   []byte(" 02/05 06:28"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 02/05 06:28\r"),
			},
		},
		{ // 154
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("bestshow500"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mbestshow500"),
			},
			{
				Big5:   []byte(":\xa5@\xac\xf6\xb9B\xb0\xca\xad\xfb\xb7\xed\xa4\xa7\xb5L\xb7\\\xa1@\xa2\xb1\xa2\xaf\xa5@\xac\xf6\xacO\xb3\xec\xa4\xa6\xaa\xba\xc4x\xb2y\xab\xd2\xb0\xea\xae\xc9\xa5N   "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa5@\xac\xf6\xb9B\xb0\xca\xad\xfb\xb7\xed\xa4\xa7\xb5L\xb7\\\xa1@\xa2\xb1\xa2\xaf\xa5@\xac\xf6\xacO\xb3\xec\xa4\xa6\xaa\xba\xc4x\xb2y\xab\xd2\xb0\xea\xae\xc9\xa5N   "),
			},
			{
				Big5:   []byte(" 02/12 00:35"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 02/12 00:35\r"),
			},
		},
		{ // 155
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("KGhuang"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mKGhuang"),
			},
			{
				Big5:   []byte(":\xac\xdd\xa8\xec\xa4U\xa4\xda\xb3\xa3\xa7\xd6\xb1\xbc\xa4F...                                    "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xac\xdd\xa8\xec\xa4U\xa4\xda\xb3\xa3\xa7\xd6\xb1\xbc\xa4F...                                    "),
			},
			{
				Big5:   []byte(" 03/08 21:12"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 03/08 21:12\r"),
			},
		},
		{ // 156
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("sklyn"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33msklyn"),
			},
			{
				Big5:   []byte(":\xa1K \xa7\xda\xa4\xa3\xaa\xbe\xb9D\xadn\xbb\xa1\xa4\xb0\xbb\xf2\xa1K                                    "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa1K \xa7\xda\xa4\xa3\xaa\xbe\xb9D\xadn\xbb\xa1\xa4\xb0\xbb\xf2\xa1K                                    "),
			},
			{
				Big5:   []byte(" 03/17 00:58"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 03/17 00:58\r"),
			},
		},
		{ // 157
			{
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Big5:   []byte("sklyn"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33msklyn"),
			},
			{
				Big5:   []byte(":Jordan\xa6b88\xa6~\xb9\xef76\xa4H29\xa7\xeb24\xa4\xa4\xa6n\xb9\xb3\xacO\xa9\xd2\xa6\xb350+\xa4\xf1\xc1\xc9FG%\xb3\xcc\xb0\xaa\xaa\xba\xa4@\xb3\xf5 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:Jordan\xa6b88\xa6~\xb9\xef76\xa4H29\xa7\xeb24\xa4\xa4\xa6n\xb9\xb3\xacO\xa9\xd2\xa6\xb350+\xa4\xf1\xc1\xc9FG%\xb3\xcc\xb0\xaa\xaa\xba\xa4@\xb3\xf5 "),
			},
			{
				Big5:   []byte(" 04/08 12:01"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 04/08 12:01\r"),
			},
		},
		{ // 158
			{
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Big5:   []byte("sklyn"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33msklyn"),
			},
			{
				Big5:   []byte(":\xb2\xc4\xa4T\xa4Q\xa4\xad\xb6\xb5\xa1K                                             "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb2\xc4\xa4T\xa4Q\xa4\xad\xb6\xb5\xa1K                                             "),
			},
			{
				Big5:   []byte(" 04/08 12:05"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 04/08 12:05\r"),
			},
		},
	}

	testContent15Utf8 = [][]*types.Rune{
		{ // 0
			{
				Utf8:   "作者: AAstar (老哥一枚) 看板: NBA",
				Big5:   []byte("\xa7@\xaa\xcc: AAstar (\xa6\xd1\xad\xf4\xa4@\xaaT) \xac\xdd\xaaO: NBA"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: AAstar (\xa6\xd1\xad\xf4\xa4@\xaaT) \xac\xdd\xaaO: NBA\r"),
			},
		},
		{ // 1
			{
				Utf8:   "標題: [轉錄][數據] 幫整理了一下MJ保持的九十項NBA第一",
				Big5:   []byte("\xbc\xd0\xc3D: [\xc2\xe0\xbf\xfd][\xbc\xc6\xbe\xda] \xc0\xb0\xbe\xe3\xb2z\xa4F\xa4@\xa4UMJ\xabO\xab\xf9\xaa\xba\xa4E\xa4Q\xb6\xb5NBA\xb2\xc4\xa4@"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xc2\xe0\xbf\xfd][\xbc\xc6\xbe\xda] \xc0\xb0\xbe\xe3\xb2z\xa4F\xa4@\xa4UMJ\xabO\xab\xf9\xaa\xba\xa4E\xa4Q\xb6\xb5NBA\xb2\xc4\xa4@\r"),
			},
		},
		{ // 2
			{
				Utf8:   "時間: Thu Jun 23 12:14:48 2016",
				Big5:   []byte("\xae\xc9\xb6\xa1: Thu Jun 23 12:14:48 2016"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Thu Jun 23 12:14:48 2016\r"),
			},
		},
		{ // 3
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 4
			{
				Utf8:   "※ [本文轉錄自 AIR_JORDAN 看板 #19UFg6e0 ]",
				Big5:   []byte("\xa1\xb0 [\xa5\xbb\xa4\xe5\xc2\xe0\xbf\xfd\xa6\xdb AIR_JORDAN \xac\xdd\xaaO #19UFg6e0 ]"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xb0 [\xa5\xbb\xa4\xe5\xc2\xe0\xbf\xfd\xa6\xdb AIR_JORDAN \xac\xdd\xaaO #19UFg6e0 ]\r"),
			},
		},
		{ // 5
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 6
			{
				Utf8:   "作者: hate2004 (祝你佛誕日快樂) 看板: AIR_JORDAN",
				Big5:   []byte("\xa7@\xaa\xcc: hate2004 (\xaf\xac\xa7A\xa6\xf2\xbd\xcf\xa4\xe9\xa7\xd6\xbc\xd6) \xac\xdd\xaaO: AIR_JORDAN"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: hate2004 (\xaf\xac\xa7A\xa6\xf2\xbd\xcf\xa4\xe9\xa7\xd6\xbc\xd6) \xac\xdd\xaaO: AIR_JORDAN\r"),
			},
		},
		{ // 7
			{
				Utf8:   "標題: [轉錄][數據] 幫整理了一下MJ保持的九十項NBA第一",
				Big5:   []byte("\xbc\xd0\xc3D: [\xc2\xe0\xbf\xfd][\xbc\xc6\xbe\xda] \xc0\xb0\xbe\xe3\xb2z\xa4F\xa4@\xa4UMJ\xabO\xab\xf9\xaa\xba\xa4E\xa4Q\xb6\xb5NBA\xb2\xc4\xa4@"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xc2\xe0\xbf\xfd][\xbc\xc6\xbe\xda] \xc0\xb0\xbe\xe3\xb2z\xa4F\xa4@\xa4UMJ\xabO\xab\xf9\xaa\xba\xa4E\xa4Q\xb6\xb5NBA\xb2\xc4\xa4@\r"),
			},
		},
		{ // 8
			{
				Utf8:   "時間: Fri Jan 23 07:00:21 2009",
				Big5:   []byte("\xae\xc9\xb6\xa1: Fri Jan 23 07:00:21 2009"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Fri Jan 23 07:00:21 2009\r"),
			},
		},
		{ // 9
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 10
			{
				Utf8:   "※ [本文轉錄自 NBA 看板]",
				Big5:   []byte("\xa1\xb0 [\xa5\xbb\xa4\xe5\xc2\xe0\xbf\xfd\xa6\xdb NBA \xac\xdd\xaaO]"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xb0 [\xa5\xbb\xa4\xe5\xc2\xe0\xbf\xfd\xa6\xdb NBA \xac\xdd\xaaO]\r"),
			},
		},
		{ // 11
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 12
			{
				Utf8:   "作者: sk2g (假班代布布) 看板: NBA",
				Big5:   []byte("\xa7@\xaa\xcc: sk2g (\xb0\xb2\xafZ\xa5N\xa5\xac\xa5\xac) \xac\xdd\xaaO: NBA"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa7@\xaa\xcc: sk2g (\xb0\xb2\xafZ\xa5N\xa5\xac\xa5\xac) \xac\xdd\xaaO: NBA\r"),
			},
		},
		{ // 13
			{
				Utf8:   "標題: [數據] MJ保持的九十項NBA歷史第一記錄",
				Big5:   []byte("\xbc\xd0\xc3D: [\xbc\xc6\xbe\xda] MJ\xabO\xab\xf9\xaa\xba\xa4E\xa4Q\xb6\xb5NBA\xbe\xfa\xa5v\xb2\xc4\xa4@\xb0O\xbf\xfd"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xbc\xd0\xc3D: [\xbc\xc6\xbe\xda] MJ\xabO\xab\xf9\xaa\xba\xa4E\xa4Q\xb6\xb5NBA\xbe\xfa\xa5v\xb2\xc4\xa4@\xb0O\xbf\xfd\r"),
			},
		},
		{ // 14
			{
				Utf8:   "時間: Thu Jan 22 23:34:14 2009",
				Big5:   []byte("\xae\xc9\xb6\xa1: Thu Jan 22 23:34:14 2009"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xae\xc9\xb6\xa1: Thu Jan 22 23:34:14 2009\r"),
			},
		},
		{ // 15
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 16
			{
				Utf8:   "小弟大概替這位大大整理一下",
				Big5:   []byte("\xa4p\xa7\xcc\xa4j\xb7\xa7\xb4\xc0\xb3o\xa6\xec\xa4j\xa4j\xbe\xe3\xb2z\xa4@\xa4U"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa4p\xa7\xcc\xa4j\xb7\xa7\xb4\xc0\xb3o\xa6\xec\xa4j\xa4j\xbe\xe3\xb2z\xa4@\xa4U\r"),
			},
		},
		{ // 17
			{
				Utf8:   "※ 引述《Laban (勒本)》之銘言：",
				Big5:   []byte("\xa1\xb0 \xa4\xde\xadz\xa1mLaban (\xb0\xc7\xa5\xbb)\xa1n\xa4\xa7\xbb\xca\xa8\xa5\xa1G"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xb0 \xa4\xde\xadz\xa1mLaban (\xb0\xc7\xa5\xbb)\xa1n\xa4\xa7\xbb\xca\xa8\xa5\xa1G\r"),
			},
		},
		{ // 18
			{
				Utf8:   "常規賽記錄：",
				Big5:   []byte("\xb1`\xb3W\xc1\xc9\xb0O\xbf\xfd\xa1G"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33;1m\xb1`\xb3W\xc1\xc9\xb0O\xbf\xfd\xa1G"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 19
			{
				Utf8:   "1.職業生涯場均30.123分，歷史第一。",
				Big5:   []byte("1.\xc2\xbe\xb7~\xa5\xcd\xb2P\xb3\xf5\xa7\xa130.123\xa4\xc0\xa1A\xbe\xfa\xa5v\xb2\xc4\xa4@\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("1.\xc2\xbe\xb7~\xa5\xcd\xb2P\xb3\xf5\xa7\xa130.123\xa4\xc0\xa1A\xbe\xfa\xa5v\xb2\xc4\xa4@\xa1C\r"),
			},
		},
		{ // 20
			{
				Utf8:   " ",
				Big5:   []byte(" "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte(" "),
			},
			{
				Utf8:   "（註:Chamberlain以30.06委列第二）",
				Big5:   []byte("\xa1]\xb5\xf9:Chamberlain\xa5H30.06\xa9e\xa6C\xb2\xc4\xa4G\xa1^"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m\xa1]\xb5\xf9:Chamberlain\xa5H30.06\xa9e\xa6C\xb2\xc4\xa4G\xa1^"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 22
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 23
			{
				Utf8:   "2.個人10次得分王，歷史第一。",
				Big5:   []byte("2.\xad\xd3\xa4H10\xa6\xb8\xb1o\xa4\xc0\xa4\xfd\xa1A\xbe\xfa\xa5v\xb2\xc4\xa4@\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("2.\xad\xd3\xa4H10\xa6\xb8\xb1o\xa4\xc0\xa4\xfd\xa1A\xbe\xfa\xa5v\xb2\xc4\xa4@\xa1C\r"),
			},
		},
		{ // 24
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 25
			{
				Utf8:   "3.平了Chamberlain連續七次得分王紀錄。",
				Big5:   []byte("3.\xa5\xad\xa4FChamberlain\xb3s\xc4\xf2\xa4C\xa6\xb8\xb1o\xa4\xc0\xa4\xfd\xac\xf6\xbf\xfd\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("3.\xa5\xad\xa4FChamberlain\xb3s\xc4\xf2\xa4C\xa6\xb8\xb1o\xa4\xc0\xa4\xfd\xac\xf6\xbf\xfd\xa1C\r"),
			},
		},
		{ // 26
			{
				Utf8:   " ",
				Big5:   []byte(" "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte(" "),
			},
			{
				Utf8:   "（註:1987 - 1993）",
				Big5:   []byte("\xa1]\xb5\xf9:1987 - 1993\xa1^"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m\xa1]\xb5\xf9:1987 - 1993\xa1^"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 27
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 28
			{
				Utf8:   "4.連續866場比賽得分超過兩位數，歷史第一。",
				Big5:   []byte("4.\xb3s\xc4\xf2866\xb3\xf5\xa4\xf1\xc1\xc9\xb1o\xa4\xc0\xb6W\xb9L\xa8\xe2\xa6\xec\xbc\xc6\xa1A\xbe\xfa\xa5v\xb2\xc4\xa4@\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("4.\xb3s\xc4\xf2866\xb3\xf5\xa4\xf1\xc1\xc9\xb1o\xa4\xc0\xb6W\xb9L\xa8\xe2\xa6\xec\xbc\xc6\xa1A\xbe\xfa\xa5v\xb2\xc4\xa4@\xa1C\r"),
			},
		},
		{ // 29
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 30
			{
				Utf8:   "5.保持單節罰球命中NBA最高記錄14次。",
				Big5:   []byte("5.\xabO\xab\xf9\xb3\xe6\xb8`\xbb@\xb2y\xa9R\xa4\xa4NBA\xb3\xcc\xb0\xaa\xb0O\xbf\xfd14\xa6\xb8\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("5.\xabO\xab\xf9\xb3\xe6\xb8`\xbb@\xb2y\xa9R\xa4\xa4NBA\xb3\xcc\xb0\xaa\xb0O\xbf\xfd14\xa6\xb8\xa1C\r"),
			},
		},
		{ // 31
			{
				Utf8:   " ",
				Big5:   []byte(" "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte(" "),
			},
			{
				Utf8:   "（註：有兩次,一次是1989年11月15日，對Utah Jazz單節罰中14球，",
				Big5:   []byte("\xa1]\xb5\xf9\xa1G\xa6\xb3\xa8\xe2\xa6\xb8,\xa4@\xa6\xb8\xacO1989\xa6~11\xa4\xeb15\xa4\xe9\xa1A\xb9\xefUtah Jazz\xb3\xe6\xb8`\xbb@\xa4\xa414\xb2y\xa1A"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m\xa1]\xb5\xf9\xa1G\xa6\xb3\xa8\xe2\xa6\xb8,\xa4@\xa6\xb8\xacO1989\xa6~11\xa4\xeb15\xa4\xe9\xa1A\xb9\xefUtah Jazz\xb3\xe6\xb8`\xbb@\xa4\xa414\xb2y\xa1A\r"),
			},
		},
		{ // 32
			{
				Utf8:   "   另一次是1993年對Miami Heat。）",
				Big5:   []byte("   \xa5t\xa4@\xa6\xb8\xacO1993\xa6~\xb9\xefMiami Heat\xa1C\xa1^"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("   \xa5t\xa4@\xa6\xb8\xacO1993\xa6~\xb9\xefMiami Heat\xa1C\xa1^"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 32
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 33
			{
				Utf8:   "6.保持著單節罰球次數最高16次。",
				Big5:   []byte("6.\xabO\xab\xf9\xb5\xdb\xb3\xe6\xb8`\xbb@\xb2y\xa6\xb8\xbc\xc6\xb3\xcc\xb0\xaa16\xa6\xb8\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("6.\xabO\xab\xf9\xb5\xdb\xb3\xe6\xb8`\xbb@\xb2y\xa6\xb8\xbc\xc6\xb3\xcc\xb0\xaa16\xa6\xb8\xa1C\r"),
			},
		},
		{ // 34
			{
				Utf8:   " （註:1992年12月30日對Miami Heat的第4節，Jordan全場24罰球。）",
				Big5:   []byte(" \xa1]\xb5\xf9:1992\xa6~12\xa4\xeb30\xa4\xe9\xb9\xefMiami Heat\xaa\xba\xb2\xc44\xb8`\xa1AJordan\xa5\xfe\xb3\xf524\xbb@\xb2y\xa1C\xa1^"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m \xa1]\xb5\xf9:1992\xa6~12\xa4\xeb30\xa4\xe9\xb9\xefMiami Heat\xaa\xba\xb2\xc44\xb8`\xa1AJordan\xa5\xfe\xb3\xf524\xbb@\xb2y\xa1C\xa1^"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 35
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 36
			{
				Utf8:   "7.NBA歷史最多的10個賽季的投籃次數和投中次數均列NBA歷史第一位。",
				Big5:   []byte("7.NBA\xbe\xfa\xa5v\xb3\xcc\xa6h\xaa\xba10\xad\xd3\xc1\xc9\xa9u\xaa\xba\xa7\xeb\xc4x\xa6\xb8\xbc\xc6\xa9M\xa7\xeb\xa4\xa4\xa6\xb8\xbc\xc6\xa7\xa1\xa6CNBA\xbe\xfa\xa5v\xb2\xc4\xa4@\xa6\xec\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("7.NBA\xbe\xfa\xa5v\xb3\xcc\xa6h\xaa\xba10\xad\xd3\xc1\xc9\xa9u\xaa\xba\xa7\xeb\xc4x\xa6\xb8\xbc\xc6\xa9M\xa7\xeb\xa4\xa4\xa6\xb8\xbc\xc6\xa7\xa1\xa6CNBA\xbe\xfa\xa5v\xb2\xc4\xa4@\xa6\xec\xa1C\r"),
			},
		},
		{ // 39
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 38
			{
				Utf8:   "8.1987年對Atlanta Hawks一戰創下連續得23分的NBA記錄。",
				Big5:   []byte("8.1987\xa6~\xb9\xefAtlanta Hawks\xa4@\xbe\xd4\xb3\xd0\xa4U\xb3s\xc4\xf2\xb1o23\xa4\xc0\xaa\xbaNBA\xb0O\xbf\xfd\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("8.1987\xa6~\xb9\xefAtlanta Hawks\xa4@\xbe\xd4\xb3\xd0\xa4U\xb3s\xc4\xf2\xb1o23\xa4\xc0\xaa\xbaNBA\xb0O\xbf\xfd\xa1C\r"),
			},
		},
		{ // 39
			{
				Utf8:   " ",
				Big5:   []byte(" "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte(" "),
			},
			{
				Utf8:   "（註1:就是投籃罰球全部命中，並且在這期間本隊無任何其他球員得分。）",
				Big5:   []byte("\xa1]\xb5\xf91:\xb4N\xacO\xa7\xeb\xc4x\xbb@\xb2y\xa5\xfe\xb3\xa1\xa9R\xa4\xa4\xa1A\xa8\xc3\xa5B\xa6b\xb3o\xb4\xc1\xb6\xa1\xa5\xbb\xb6\xa4\xb5L\xa5\xf4\xa6\xf3\xa8\xe4\xa5L\xb2y\xad\xfb\xb1o\xa4\xc0\xa1C\xa1^"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m\xa1]\xb5\xf91:\xb4N\xacO\xa7\xeb\xc4x\xbb@\xb2y\xa5\xfe\xb3\xa1\xa9R\xa4\xa4\xa1A\xa8\xc3\xa5B\xa6b\xb3o\xb4\xc1\xb6\xa1\xa5\xbb\xb6\xa4\xb5L\xa5\xf4\xa6\xf3\xa8\xe4\xa5L\xb2y\xad\xfb\xb1o\xa4\xc0\xa1C\xa1^\r"),
			},
		},
		{ // 40
			{
				Utf8:   " ",
				Big5:   []byte(" "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte(" "),
			},
			{
				Utf8:   " (註2:單場連續得分排行榜前四均為Jordan：23分",
				Big5:   []byte(" (\xb5\xf92:\xb3\xe6\xb3\xf5\xb3s\xc4\xf2\xb1o\xa4\xc0\xb1\xc6\xa6\xe6\xba]\xabe\xa5|\xa7\xa1\xac\xb0Jordan\xa1G23\xa4\xc0"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m (\xb5\xf92:\xb3\xe6\xb3\xf5\xb3s\xc4\xf2\xb1o\xa4\xc0\xb1\xc6\xa6\xe6\xba]\xabe\xa5|\xa7\xa1\xac\xb0Jordan\xa1G23\xa4\xc0"),
			},
			{
				Utf8:   "（",
				Big5:   []byte("\xa1]"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1]"),
			},
			{
				Utf8:   "1987",
				Big5:   []byte("1987"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[31m1987"),
			},
			{
				Utf8:   "）",
				Big5:   []byte("\xa1^"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1^"),
			},
			{
				Utf8:   "，22分",
				Big5:   []byte("\xa1A22\xa4\xc0"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m\xa1A22\xa4\xc0"),
			},
			{
				Utf8:   "（",
				Big5:   []byte("\xa1]"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\x1b[m\xa1]"),
			},
			{
				Utf8:   "2002",
				Big5:   []byte("2002"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[31m2002"),
			},
			{
				Utf8:   "）",
				Big5:   []byte("\xa1^"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1^\r"),
			},
		},
		{ // 41
			{
				Utf8:   " ",
				Big5:   []byte(" "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte(" "),
			},
			{
				Utf8:   "      ，19分",
				Big5:   []byte("      \xa1A19\xa4\xc0"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m      \xa1A19\xa4\xc0"),
			},
			{
				Utf8:   "（",
				Big5:   []byte("\xa1]"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1]"),
			},
			{
				Utf8:   "1996",
				Big5:   []byte("1996"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[31m1996"),
			},
			{
				Utf8:   "）",
				Big5:   []byte("\xa1^"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1^"),
			},
			{
				Utf8:   "，18分",
				Big5:   []byte("\xa1A18\xa4\xc0"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m\xa1A18\xa4\xc0"),
			},
			{
				Utf8:   "（",
				Big5:   []byte("\xa1]"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\x1b[m\xa1]"),
			},
			{
				Utf8:   "1987",
				Big5:   []byte("1987"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[31m1987"),
			},
			{
				Utf8:   "）",
				Big5:   []byte("\xa1^"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xa1^"),
			},
			{
				Utf8:   "。)",
				Big5:   []byte("\xa1C)"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m\xa1C)"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 44
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 45
			{
				Utf8:   "19.10次入選NBA年度第一隊。",
				Big5:   []byte("19.10\xa6\xb8\xa4J\xbf\xefNBA\xa6~\xab\xd7\xb2\xc4\xa4@\xb6\xa4\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("19.10\xa6\xb8\xa4J\xbf\xefNBA\xa6~\xab\xd7\xb2\xc4\xa4@\xb6\xa4\xa1C\r"),
			},
		},
		{ // 46
			{
				Utf8:   "   (註:1986-87至1992-93賽季, 1995-96至1997-98賽季。)",
				Big5:   []byte("   (\xb5\xf9:1986-87\xa6\xdc1992-93\xc1\xc9\xa9u, 1995-96\xa6\xdc1997-98\xc1\xc9\xa9u\xa1C)"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m   (\xb5\xf9:1986-87\xa6\xdc1992-93\xc1\xc9\xa9u, 1995-96\xa6\xdc1997-98\xc1\xc9\xa9u\xa1C)"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 47
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 48
			{
				Utf8:   "20.87-88賽季阻攻次數達到131次,是後衛的阻攻記錄。",
				Big5:   []byte("20.87-88\xc1\xc9\xa9u\xaa\xfd\xa7\xf0\xa6\xb8\xbc\xc6\xb9F\xa8\xec131\xa6\xb8,\xacO\xab\xe1\xbd\xc3\xaa\xba\xaa\xfd\xa7\xf0\xb0O\xbf\xfd\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("20.87-88\xc1\xc9\xa9u\xaa\xfd\xa7\xf0\xa6\xb8\xbc\xc6\xb9F\xa8\xec131\xa6\xb8,\xacO\xab\xe1\xbd\xc3\xaa\xba\xaa\xfd\xa7\xf0\xb0O\xbf\xfd\xa1C\r"),
			},
		},
		{ // 49
			{
				Utf8:   "  （註:單場最高6次，職業生涯有8場比賽單場阻攻5次以上。）",
				Big5:   []byte("  \xa1]\xb5\xf9:\xb3\xe6\xb3\xf5\xb3\xcc\xb0\xaa6\xa6\xb8\xa1A\xc2\xbe\xb7~\xa5\xcd\xb2P\xa6\xb38\xb3\xf5\xa4\xf1\xc1\xc9\xb3\xe6\xb3\xf5\xaa\xfd\xa7\xf05\xa6\xb8\xa5H\xa4W\xa1C\xa1^"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m  \xa1]\xb5\xf9:\xb3\xe6\xb3\xf5\xb3\xcc\xb0\xaa6\xa6\xb8\xa1A\xc2\xbe\xb7~\xa5\xcd\xb2P\xa6\xb38\xb3\xf5\xa4\xf1\xc1\xc9\xb3\xe6\xb3\xf5\xaa\xfd\xa7\xf05\xa6\xb8\xa5H\xa4W\xa1C\xa1^"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 50
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 51
			{
				Utf8:   "39.單一賽季季後賽平均最高43.7分。",
				Big5:   []byte("39.\xb3\xe6\xa4@\xc1\xc9\xa9u\xa9u\xab\xe1\xc1\xc9\xa5\xad\xa7\xa1\xb3\xcc\xb0\xaa43.7\xa4\xc0\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("39.\xb3\xe6\xa4@\xc1\xc9\xa9u\xa9u\xab\xe1\xc1\xc9\xa5\xad\xa7\xa1\xb3\xcc\xb0\xaa43.7\xa4\xc0\xa1C\r"),
			},
		},
		{ // 52
			{
				Utf8:   "   (註:1986年，Jordan在面對當年67勝（主場40勝1敗)的總冠軍Boston Celtics",
				Big5:   []byte("   (\xb5\xf9:1986\xa6~\xa1AJordan\xa6b\xad\xb1\xb9\xef\xb7\xed\xa6~67\xb3\xd3\xa1]\xa5D\xb3\xf540\xb3\xd31\xb1\xd1)\xaa\xba\xc1`\xaba\xadxBoston Celtics"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m   (\xb5\xf9:1986\xa6~\xa1AJordan\xa6b\xad\xb1\xb9\xef\xb7\xed\xa6~67\xb3\xd3\xa1]\xa5D\xb3\xf540\xb3\xd31\xb1\xd1)\xaa\xba\xc1`\xaba\xadxBoston Celtics\r"),
			},
		},
		{ // 53
			{
				Utf8:   "       的三場季後賽，分別拿下49分，63分，19分，場均43.7分是單一賽季季後賽場均",
				Big5:   []byte("       \xaa\xba\xa4T\xb3\xf5\xa9u\xab\xe1\xc1\xc9\xa1A\xa4\xc0\xa7O\xae\xb3\xa4U49\xa4\xc0\xa1A63\xa4\xc0\xa1A19\xa4\xc0\xa1A\xb3\xf5\xa7\xa143.7\xa4\xc0\xacO\xb3\xe6\xa4@\xc1\xc9\xa9u\xa9u\xab\xe1\xc1\xc9\xb3\xf5\xa7\xa1"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("       \xaa\xba\xa4T\xb3\xf5\xa9u\xab\xe1\xc1\xc9\xa1A\xa4\xc0\xa7O\xae\xb3\xa4U49\xa4\xc0\xa1A63\xa4\xc0\xa1A19\xa4\xc0\xa1A\xb3\xf5\xa7\xa143.7\xa4\xc0\xacO\xb3\xe6\xa4@\xc1\xc9\xa9u\xa9u\xab\xe1\xc1\xc9\xb3\xf5\xa7\xa1\r"),
			},
		},
		{ // 54
			{
				Utf8:   "       最高記錄，歷史上只有兩個人在季後賽場均40分以上，另一位是Jerry West",
				Big5:   []byte("       \xb3\xcc\xb0\xaa\xb0O\xbf\xfd\xa1A\xbe\xfa\xa5v\xa4W\xa5u\xa6\xb3\xa8\xe2\xad\xd3\xa4H\xa6b\xa9u\xab\xe1\xc1\xc9\xb3\xf5\xa7\xa140\xa4\xc0\xa5H\xa4W\xa1A\xa5t\xa4@\xa6\xec\xacOJerry West"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("       \xb3\xcc\xb0\xaa\xb0O\xbf\xfd\xa1A\xbe\xfa\xa5v\xa4W\xa5u\xa6\xb3\xa8\xe2\xad\xd3\xa4H\xa6b\xa9u\xab\xe1\xc1\xc9\xb3\xf5\xa7\xa140\xa4\xc0\xa5H\xa4W\xa1A\xa5t\xa4@\xa6\xec\xacOJerry West\r"),
			},
		},
		{ // 55
			{
				Utf8:   "      （場均40.6分）)",
				Big5:   []byte("      \xa1]\xb3\xf5\xa7\xa140.6\xa4\xc0\xa1^)"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("      \xa1]\xb3\xf5\xa7\xa140.6\xa4\xc0\xa1^)"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 56
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 57
			{
				Utf8:   "40.季後賽單場得分63分，歷史第一。",
				Big5:   []byte("40.\xa9u\xab\xe1\xc1\xc9\xb3\xe6\xb3\xf5\xb1o\xa4\xc063\xa4\xc0\xa1A\xbe\xfa\xa5v\xb2\xc4\xa4@\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("40.\xa9u\xab\xe1\xc1\xc9\xb3\xe6\xb3\xf5\xb1o\xa4\xc063\xa4\xc0\xa1A\xbe\xfa\xa5v\xb2\xc4\xa4@\xa1C\r"),
			},
		},
		{ // 58
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 59
			{
				Utf8:   "41.季後賽唯一一個拿過back to back 50分的球員。",
				Big5:   []byte("41.\xa9u\xab\xe1\xc1\xc9\xb0\xdf\xa4@\xa4@\xad\xd3\xae\xb3\xb9Lback to back 50\xa4\xc0\xaa\xba\xb2y\xad\xfb\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("41.\xa9u\xab\xe1\xc1\xc9\xb0\xdf\xa4@\xa4@\xad\xd3\xae\xb3\xb9Lback to back 50\xa4\xc0\xaa\xba\xb2y\xad\xfb\xa1C\r"),
			},
		},
		{ // 60
			{
				Utf8:   "   (註:1988年季後賽，Jordan在對Cleveland Cavaliers的比賽中，第一場50分，",
				Big5:   []byte("   (\xb5\xf9:1988\xa6~\xa9u\xab\xe1\xc1\xc9\xa1AJordan\xa6b\xb9\xefCleveland Cavaliers\xaa\xba\xa4\xf1\xc1\xc9\xa4\xa4\xa1A\xb2\xc4\xa4@\xb3\xf550\xa4\xc0\xa1A"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m   (\xb5\xf9:1988\xa6~\xa9u\xab\xe1\xc1\xc9\xa1AJordan\xa6b\xb9\xefCleveland Cavaliers\xaa\xba\xa4\xf1\xc1\xc9\xa4\xa4\xa1A\xb2\xc4\xa4@\xb3\xf550\xa4\xc0\xa1A\r"),
			},
		},
		{ // 61
			{
				Utf8:   "       第二場55分，連續兩場拿下50分，歷史上只有Jordan。)",
				Big5:   []byte("       \xb2\xc4\xa4G\xb3\xf555\xa4\xc0\xa1A\xb3s\xc4\xf2\xa8\xe2\xb3\xf5\xae\xb3\xa4U50\xa4\xc0\xa1A\xbe\xfa\xa5v\xa4W\xa5u\xa6\xb3Jordan\xa1C)"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("       \xb2\xc4\xa4G\xb3\xf555\xa4\xc0\xa1A\xb3s\xc4\xf2\xa8\xe2\xb3\xf5\xae\xb3\xa4U50\xa4\xc0\xa1A\xbe\xfa\xa5v\xa4W\xa5u\xa6\xb3Jordan\xa1C)"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 62
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 63
			{
				Utf8:   "42.季後賽唯一一個拿過連續3場45分的球員。",
				Big5:   []byte("42.\xa9u\xab\xe1\xc1\xc9\xb0\xdf\xa4@\xa4@\xad\xd3\xae\xb3\xb9L\xb3s\xc4\xf23\xb3\xf545\xa4\xc0\xaa\xba\xb2y\xad\xfb\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("42.\xa9u\xab\xe1\xc1\xc9\xb0\xdf\xa4@\xa4@\xad\xd3\xae\xb3\xb9L\xb3s\xc4\xf23\xb3\xf545\xa4\xc0\xaa\xba\xb2y\xad\xfb\xa1C\r"),
			},
		},
		{ // 64
			{
				Utf8:   "   (註1:1990年在對Philadelphia 76ers的半決賽中，Jordan5場比賽場均拿下43.0分",
				Big5:   []byte("   (\xb5\xf91:1990\xa6~\xa6b\xb9\xefPhiladelphia 76ers\xaa\xba\xa5b\xa8M\xc1\xc9\xa4\xa4\xa1AJordan5\xb3\xf5\xa4\xf1\xc1\xc9\xb3\xf5\xa7\xa1\xae\xb3\xa4U43.0\xa4\xc0"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m   (\xb5\xf91:1990\xa6~\xa6b\xb9\xefPhiladelphia 76ers\xaa\xba\xa5b\xa8M\xc1\xc9\xa4\xa4\xa1AJordan5\xb3\xf5\xa4\xf1\xc1\xc9\xb3\xf5\xa7\xa1\xae\xb3\xa4U43.0\xa4\xc0\r"),
			},
		},
		{ // 65
			{
				Utf8:   "        ，其中在第二到第4場比賽中，分別得到45分，49分，45分連續三場比賽超過45+，",
				Big5:   []byte("        \xa1A\xa8\xe4\xa4\xa4\xa6b\xb2\xc4\xa4G\xa8\xec\xb2\xc44\xb3\xf5\xa4\xf1\xc1\xc9\xa4\xa4\xa1A\xa4\xc0\xa7O\xb1o\xa8\xec45\xa4\xc0\xa1A49\xa4\xc0\xa1A45\xa4\xc0\xb3s\xc4\xf2\xa4T\xb3\xf5\xa4\xf1\xc1\xc9\xb6W\xb9L45+\xa1A"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("        \xa1A\xa8\xe4\xa4\xa4\xa6b\xb2\xc4\xa4G\xa8\xec\xb2\xc44\xb3\xf5\xa4\xf1\xc1\xc9\xa4\xa4\xa1A\xa4\xc0\xa7O\xb1o\xa8\xec45\xa4\xc0\xa1A49\xa4\xc0\xa1A45\xa4\xc0\xb3s\xc4\xf2\xa4T\xb3\xf5\xa4\xf1\xc1\xc9\xb6W\xb9L45+\xa1A\r"),
			},
		},
		{ // 66
			{
				Utf8:   "        歷史上只有Jordan。)",
				Big5:   []byte("        \xbe\xfa\xa5v\xa4W\xa5u\xa6\xb3Jordan\xa1C)"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("        \xbe\xfa\xa5v\xa4W\xa5u\xa6\xb3Jordan\xa1C)\r"),
			},
		},
		{ // 67
			{
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\r"),
			},
		},
		{ // 68
			{
				Utf8:   "  （註2:Jordan還有次連續得到44分，50分和44分，Iverson曾經連續得到46分，44分和",
				Big5:   []byte("  \xa1]\xb5\xf92:Jordan\xc1\xd9\xa6\xb3\xa6\xb8\xb3s\xc4\xf2\xb1o\xa8\xec44\xa4\xc0\xa1A50\xa4\xc0\xa9M44\xa4\xc0\xa1AIverson\xb4\xbf\xb8g\xb3s\xc4\xf2\xb1o\xa8\xec46\xa4\xc0\xa1A44\xa4\xc0\xa9M"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("  \xa1]\xb5\xf92:Jordan\xc1\xd9\xa6\xb3\xa6\xb8\xb3s\xc4\xf2\xb1o\xa8\xec44\xa4\xc0\xa1A50\xa4\xc0\xa9M44\xa4\xc0\xa1AIverson\xb4\xbf\xb8g\xb3s\xc4\xf2\xb1o\xa8\xec46\xa4\xc0\xa1A44\xa4\xc0\xa9M\r"),
			},
		},
		{ // 69
			{
				Utf8:   "        48分，都只差一點。）",
				Big5:   []byte("        48\xa4\xc0\xa1A\xb3\xa3\xa5u\xaet\xa4@\xc2I\xa1C\xa1^"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("        48\xa4\xc0\xa1A\xb3\xa3\xa5u\xaet\xa4@\xc2I\xa1C\xa1^"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 70
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 69
			{
				Utf8:   "90.唯一能讓聯盟破例把已經退役的球衣",
				Big5:   []byte("90.\xb0\xdf\xa4@\xaf\xe0\xc5\xfd\xc1p\xb7\xf9\xaf}\xa8\xd2\xa7\xe2\xa4w\xb8g\xb0h\xa7\xd0\xaa\xba\xb2y\xa6\xe7"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("90.\xb0\xdf\xa4@\xaf\xe0\xc5\xfd\xc1p\xb7\xf9\xaf}\xa8\xd2\xa7\xe2\xa4w\xb8g\xb0h\xa7\xd0\xaa\xba\xb2y\xa6\xe7"),
			},
			{
				Utf8:   "23號",
				Big5:   []byte("23\xb8\xb9"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[31;1m23\xb8\xb9"),
			},
			{
				Utf8:   "重新穿回。",
				Big5:   []byte("\xad\xab\xb7s\xac\xef\xa6^\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xad\xab\xb7s\xac\xef\xa6^\xa1C\r"),
			},
		},
		{ // 70
			{
				Utf8:   "   (註:下面已有版友補充不只他一人了，Magic Johnson 1996年復出時也是。)",
				Big5:   []byte("   (\xb5\xf9:\xa4U\xad\xb1\xa4w\xa6\xb3\xaa\xa9\xa4\xcd\xb8\xc9\xa5R\xa4\xa3\xa5u\xa5L\xa4@\xa4H\xa4F\xa1AMagic Johnson 1996\xa6~\xb4_\xa5X\xae\xc9\xa4]\xacO\xa1C)"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m   (\xb5\xf9:\xa4U\xad\xb1\xa4w\xa6\xb3\xaa\xa9\xa4\xcd\xb8\xc9\xa5R\xa4\xa3\xa5u\xa5L\xa4@\xa4H\xa4F\xa1AMagic Johnson 1996\xa6~\xb4_\xa5X\xae\xc9\xa4]\xacO\xa1C)"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 71
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 72
			{
				Utf8:   "91.MJ是唯一在他不曾效力的球隊被掛上球衣在球館(",
				Big5:   []byte("91.MJ\xacO\xb0\xdf\xa4@\xa6b\xa5L\xa4\xa3\xb4\xbf\xae\xc4\xa4O\xaa\xba\xb2y\xb6\xa4\xb3Q\xb1\xbe\xa4W\xb2y\xa6\xe7\xa6b\xb2y\xc0]("),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("91.MJ\xacO\xb0\xdf\xa4@\xa6b\xa5L\xa4\xa3\xb4\xbf\xae\xc4\xa4O\xaa\xba\xb2y\xb6\xa4\xb3Q\xb1\xbe\xa4W\xb2y\xa6\xe7\xa6b\xb2y\xc0]("),
			},
			{
				Utf8:   "Miami Heat",
				Big5:   []byte("Miami Heat"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[31;1mMiami Heat"),
			},
			{
				Utf8:   "的球館)的球員~",
				Big5:   []byte("\xaa\xba\xb2y\xc0])\xaa\xba\xb2y\xad\xfb~"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\xaa\xba\xb2y\xc0])\xaa\xba\xb2y\xad\xfb~\r"),
			},
		},
		{ // 73
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 74
			{
				Utf8:   "92.Jordan還是唯一代表美國大學籃球隊和美國夢之隊分別獲得奧運會金牌的球員~此記錄注定",
				Big5:   []byte("92.Jordan\xc1\xd9\xacO\xb0\xdf\xa4@\xa5N\xaa\xed\xac\xfc\xb0\xea\xa4j\xbe\xc7\xc4x\xb2y\xb6\xa4\xa9M\xac\xfc\xb0\xea\xb9\xda\xa4\xa7\xb6\xa4\xa4\xc0\xa7O\xc0\xf2\xb1o\xb6\xf8\xb9B\xb7|\xaa\xf7\xb5P\xaa\xba\xb2y\xad\xfb~\xa6\xb9\xb0O\xbf\xfd\xaa`\xa9w"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("92.Jordan\xc1\xd9\xacO\xb0\xdf\xa4@\xa5N\xaa\xed\xac\xfc\xb0\xea\xa4j\xbe\xc7\xc4x\xb2y\xb6\xa4\xa9M\xac\xfc\xb0\xea\xb9\xda\xa4\xa7\xb6\xa4\xa4\xc0\xa7O\xc0\xf2\xb1o\xb6\xf8\xb9B\xb7|\xaa\xf7\xb5P\xaa\xba\xb2y\xad\xfb~\xa6\xb9\xb0O\xbf\xfd\xaa`\xa9w\r"),
			},
		},
		{ // 77
			{
				Utf8:   "   是絕後的了,因為現在美國不可能再派大學隊伍參加奧運會了~",
				Big5:   []byte("   \xacO\xb5\xb4\xab\xe1\xaa\xba\xa4F,\xa6]\xac\xb0\xb2{\xa6b\xac\xfc\xb0\xea\xa4\xa3\xa5i\xaf\xe0\xa6A\xac\xa3\xa4j\xbe\xc7\xb6\xa4\xa5\xee\xb0\xd1\xa5[\xb6\xf8\xb9B\xb7|\xa4F~"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("   \xacO\xb5\xb4\xab\xe1\xaa\xba\xa4F,\xa6]\xac\xb0\xb2{\xa6b\xac\xfc\xb0\xea\xa4\xa3\xa5i\xaf\xe0\xa6A\xac\xa3\xa4j\xbe\xc7\xb6\xa4\xa5\xee\xb0\xd1\xa5[\xb6\xf8\xb9B\xb7|\xa4F~\r"),
			},
		},
		{ // 78
			{
				Utf8:   "   (註:已有版友補充不只Jordan一人，Patrick Ewing跟Chris Mullin都跟Jordan一起。)",
				Big5:   []byte("   (\xb5\xf9:\xa4w\xa6\xb3\xaa\xa9\xa4\xcd\xb8\xc9\xa5R\xa4\xa3\xa5uJordan\xa4@\xa4H\xa1APatrick Ewing\xb8\xf2Chris Mullin\xb3\xa3\xb8\xf2Jordan\xa4@\xb0_\xa1C)"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[32m   (\xb5\xf9:\xa4w\xa6\xb3\xaa\xa9\xa4\xcd\xb8\xc9\xa5R\xa4\xa3\xa5uJordan\xa4@\xa4H\xa1APatrick Ewing\xb8\xf2Chris Mullin\xb3\xa3\xb8\xf2Jordan\xa4@\xb0_\xa1C)"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 79
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 80
			{
				Utf8:   "93.Jordan是歷史上第一個能在復出後還帶領球隊連得3次NBA總冠軍的人。",
				Big5:   []byte("93.Jordan\xacO\xbe\xfa\xa5v\xa4W\xb2\xc4\xa4@\xad\xd3\xaf\xe0\xa6b\xb4_\xa5X\xab\xe1\xc1\xd9\xb1a\xbb\xe2\xb2y\xb6\xa4\xb3s\xb1o3\xa6\xb8NBA\xc1`\xaba\xadx\xaa\xba\xa4H\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("93.Jordan\xacO\xbe\xfa\xa5v\xa4W\xb2\xc4\xa4@\xad\xd3\xaf\xe0\xa6b\xb4_\xa5X\xab\xe1\xc1\xd9\xb1a\xbb\xe2\xb2y\xb6\xa4\xb3s\xb1o3\xa6\xb8NBA\xc1`\xaba\xadx\xaa\xba\xa4H\xa1C\r"),
			},
		},
		{ // 81
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 82
			{
				Utf8:   "94.02年復出時將那年的所有工資（大約100萬美元）捐給911受難者，是NBA史上唯一一個",
				Big5:   []byte("94.02\xa6~\xb4_\xa5X\xae\xc9\xb1N\xa8\xba\xa6~\xaa\xba\xa9\xd2\xa6\xb3\xa4u\xb8\xea\xa1]\xa4j\xac\xf9100\xb8U\xac\xfc\xa4\xb8\xa1^\xae\xbd\xb5\xb9911\xa8\xfc\xc3\xf8\xaa\xcc\xa1A\xacONBA\xa5v\xa4W\xb0\xdf\xa4@\xa4@\xad\xd3"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("94.02\xa6~\xb4_\xa5X\xae\xc9\xb1N\xa8\xba\xa6~\xaa\xba\xa9\xd2\xa6\xb3\xa4u\xb8\xea\xa1]\xa4j\xac\xf9100\xb8U\xac\xfc\xa4\xb8\xa1^\xae\xbd\xb5\xb9911\xa8\xfc\xc3\xf8\xaa\xcc\xa1A\xacONBA\xa5v\xa4W\xb0\xdf\xa4@\xa4@\xad\xd3\r"),
			},
		},
		{ // 83
			{
				Utf8:   "   捐自己全年工資的。",
				Big5:   []byte("   \xae\xbd\xa6\xdb\xa4v\xa5\xfe\xa6~\xa4u\xb8\xea\xaa\xba\xa1C"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("   \xae\xbd\xa6\xdb\xa4v\xa5\xfe\xa6~\xa4u\xb8\xea\xaa\xba\xa1C\r"),
			},
		},
		{ // 84
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 85
			{
				Utf8:   ": ------------------------------------------------------------------------------",
				Big5:   []byte(": ------------------------------------------------------------------------------"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte(": ------------------------------------------------------------------------------\r"),
			},
		},
		{ // 86
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 87
			{
				Utf8:   " ",
				Big5:   []byte(" "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte(" "),
			},
			{
				Utf8:   "Jordan : 我就是神!",
				Big5:   []byte("Jordan : \xa7\xda\xb4N\xacO\xaf\xab!"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[31;1mJordan : \xa7\xda\xb4N\xacO\xaf\xab!"),
			},
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m\r"),
			},
		},
		{ // 88
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 89
			{
				Utf8:   "--",
				Big5:   []byte("--"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("--\r"),
			},
		},
		{ // 90
			{
				Utf8:   "※ 發信站: 批踢踢實業坊(ptt.cc)",
				Big5:   []byte("\xa1\xb0 \xb5o\xabH\xaf\xb8: \xa7\xe5\xbd\xf0\xbd\xf0\xb9\xea\xb7~\xa7{(ptt.cc)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xb0 \xb5o\xabH\xaf\xb8: \xa7\xe5\xbd\xf0\xbd\xf0\xb9\xea\xb7~\xa7{(ptt.cc)\r"),
			},
		},
		{ // 91
			{
				Utf8:   "◆ From: 61.228.169.158",
				Big5:   []byte("\xa1\xbb From: 61.228.169.158"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xbb From: 61.228.169.158\r"),
			},
		},
		{ // 92
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "chinhan1216",
				Big5:   []byte("chinhan1216"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mchinhan1216"),
			},
			{
				Utf8:   ":你屌....                                           ",
				Big5:   []byte(":\xa7A\xcex....                                           "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa7A\xcex....                                           "),
			},
			{
				Utf8:   " 01/22 23:35",
				Big5:   []byte(" 01/22 23:35"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/22 23:35\r"),
			},
		},
		{ // 93
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "jasonlin68",
				Big5:   []byte("jasonlin68"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mjasonlin68"),
			},
			{
				Utf8:   ":好文推 辛苦了                                       ",
				Big5:   []byte(":\xa6n\xa4\xe5\xb1\xc0 \xa8\xaf\xadW\xa4F                                       "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa6n\xa4\xe5\xb1\xc0 \xa8\xaf\xadW\xa4F                                       "),
			},
			{
				Utf8:   " 01/22 23:36",
				Big5:   []byte(" 01/22 23:36"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/22 23:36\r"),
			},
		},
		{ // 92
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "Gigabye",
				Big5:   []byte("Gigabye"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mGigabye"),
			},
			{
				Utf8:   ":季後賽才最是可怕的紀錄 能量用在最關鍵的時刻 就是神     ",
				Big5:   []byte(":\xa9u\xab\xe1\xc1\xc9\xa4~\xb3\xcc\xacO\xa5i\xa9\xc8\xaa\xba\xac\xf6\xbf\xfd \xaf\xe0\xb6q\xa5\xce\xa6b\xb3\xcc\xc3\xf6\xc1\xe4\xaa\xba\xae\xc9\xa8\xe8 \xb4N\xacO\xaf\xab     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa9u\xab\xe1\xc1\xc9\xa4~\xb3\xcc\xacO\xa5i\xa9\xc8\xaa\xba\xac\xf6\xbf\xfd \xaf\xe0\xb6q\xa5\xce\xa6b\xb3\xcc\xc3\xf6\xc1\xe4\xaa\xba\xae\xc9\xa8\xe8 \xb4N\xacO\xaf\xab     "),
			},
			{
				Utf8:   " 01/22 23:37",
				Big5:   []byte(" 01/22 23:37"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/22 23:37\r"),
			},
		},
		{ // 93
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "RonArtest93",
				Big5:   []byte("RonArtest93"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mRonArtest93"),
			},
			{
				Utf8:   ":PUSH                                               ",
				Big5:   []byte(":PUSH                                               "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:PUSH                                               "),
			},
			{
				Utf8:   " 01/22 23:38",
				Big5:   []byte(" 01/22 23:38"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/22 23:38\r"),
			},
		},
		{ // 97
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "DenyPedrosa",
				Big5:   []byte("DenyPedrosa"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mDenyPedrosa"),
			},
			{
				Utf8:   ":果然是神= = 太強了~                                ",
				Big5:   []byte(":\xaaG\xb5M\xacO\xaf\xab= = \xa4\xd3\xb1j\xa4F~                                "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xaaG\xb5M\xacO\xaf\xab= = \xa4\xd3\xb1j\xa4F~                                "),
			},
			{
				Utf8:   " 01/22 23:39",
				Big5:   []byte(" 01/22 23:39"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/22 23:39\r"),
			},
		},
		{ // 98
			{
				Utf8:   "→ ",
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Utf8:   "iamdada",
				Big5:   []byte("iamdada"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33miamdada"),
			},
			{
				Utf8:   ":太可怕了...你真有耐性                                  ",
				Big5:   []byte(":\xa4\xd3\xa5i\xa9\xc8\xa4F...\xa7A\xafu\xa6\xb3\xad@\xa9\xca                                  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa4\xd3\xa5i\xa9\xc8\xa4F...\xa7A\xafu\xa6\xb3\xad@\xa9\xca                                  "),
			},
			{
				Utf8:   " 01/22 23:44",
				Big5:   []byte(" 01/22 23:44"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/22 23:44\r"),
			},
		},
		{ // 96
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "ThBasketball",
				Big5:   []byte("ThBasketball"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mThBasketball"),
			},
			{
				Utf8:   ":唯一在同一隻球隊同一場比賽中，2人一起大三元。     ",
				Big5:   []byte(":\xb0\xdf\xa4@\xa6b\xa6P\xa4@\xb0\xa6\xb2y\xb6\xa4\xa6P\xa4@\xb3\xf5\xa4\xf1\xc1\xc9\xa4\xa4\xa1A2\xa4H\xa4@\xb0_\xa4j\xa4T\xa4\xb8\xa1C     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb0\xdf\xa4@\xa6b\xa6P\xa4@\xb0\xa6\xb2y\xb6\xa4\xa6P\xa4@\xb3\xf5\xa4\xf1\xc1\xc9\xa4\xa4\xa1A2\xa4H\xa4@\xb0_\xa4j\xa4T\xa4\xb8\xa1C     "),
			},
			{
				Utf8:   " 01/22 23:44",
				Big5:   []byte(" 01/22 23:44"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/22 23:44\r"),
			},
		},
		{ // 100
			{
				Utf8:   "→ ",
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Utf8:   "ThBasketball",
				Big5:   []byte("ThBasketball"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mThBasketball"),
			},
			{
				Utf8:   ":這個Carter跟Kidd做到了                            ",
				Big5:   []byte(":\xb3o\xad\xd3Carter\xb8\xf2Kidd\xb0\xb5\xa8\xec\xa4F                            "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb3o\xad\xd3Carter\xb8\xf2Kidd\xb0\xb5\xa8\xec\xa4F                            "),
			},
			{
				Utf8:   " 01/22 23:44",
				Big5:   []byte(" 01/22 23:44"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/22 23:44\r"),
			},
		},
		{ // 101
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "mark0628",
				Big5:   []byte("mark0628"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mmark0628"),
			},
			{
				Utf8:   ":朝聖推 你好認真                                       ",
				Big5:   []byte(":\xb4\xc2\xb8t\xb1\xc0 \xa7A\xa6n\xbb{\xafu                                       "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb4\xc2\xb8t\xb1\xc0 \xa7A\xa6n\xbb{\xafu                                       "),
			},
			{
				Utf8:   " 01/22 23:47",
				Big5:   []byte(" 01/22 23:47"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/22 23:47\r"),
			},
		},
		{ // 102
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "grf7186",
				Big5:   []byte("grf7186"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mgrf7186"),
			},
			{
				Utf8:   ":認真 推一個~~                                          ",
				Big5:   []byte(":\xbb{\xafu \xb1\xc0\xa4@\xad\xd3~~                                          "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xbb{\xafu \xb1\xc0\xa4@\xad\xd3~~                                          "),
			},
			{
				Utf8:   " 01/22 23:49",
				Big5:   []byte(" 01/22 23:49"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/22 23:49\r"),
			},
		},
		{ // 100
			{
				Utf8:   "→ ",
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Utf8:   "sk2g",
				Big5:   []byte("sk2g"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33msk2g"),
			},
			{
				Utf8:   ":我看Darious Miles就不只3次啦XD  可是他跟Jordan畢竟不同:)  ",
				Big5:   []byte(":\xa7\xda\xac\xddDarious Miles\xb4N\xa4\xa3\xa5u3\xa6\xb8\xb0\xd5XD  \xa5i\xacO\xa5L\xb8\xf2Jordan\xb2\xa6\xb3\xba\xa4\xa3\xa6P:)  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa7\xda\xac\xddDarious Miles\xb4N\xa4\xa3\xa5u3\xa6\xb8\xb0\xd5XD  \xa5i\xacO\xa5L\xb8\xf2Jordan\xb2\xa6\xb3\xba\xa4\xa3\xa6P:)  "),
			},
			{
				Utf8:   " 01/23 00:15",
				Big5:   []byte(" 01/23 00:15"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:15\r"),
			},
		},
		{ // 101
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "fanix",
				Big5:   []byte("fanix"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mfanix"),
			},
			{
				Utf8:   ":來人轉去阿爵板!!                                         ",
				Big5:   []byte(":\xa8\xd3\xa4H\xc2\xe0\xa5h\xaa\xfc\xc0\xef\xaaO!!                                         "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa8\xd3\xa4H\xc2\xe0\xa5h\xaa\xfc\xc0\xef\xaaO!!                                         "),
			},
			{
				Utf8:   " 01/23 00:17",
				Big5:   []byte(" 01/23 00:17"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:17\r"),
			},
		},
		{ // 106
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "CenaC",
				Big5:   []byte("CenaC"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mCenaC"),
			},
			{
				Utf8:   ":推\"神\" 神跟人是不同的 他是在另一個層次打籃球             ",
				Big5:   []byte(":\xb1\xc0\"\xaf\xab\" \xaf\xab\xb8\xf2\xa4H\xacO\xa4\xa3\xa6P\xaa\xba \xa5L\xacO\xa6b\xa5t\xa4@\xad\xd3\xbch\xa6\xb8\xa5\xb4\xc4x\xb2y             "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb1\xc0\"\xaf\xab\" \xaf\xab\xb8\xf2\xa4H\xacO\xa4\xa3\xa6P\xaa\xba \xa5L\xacO\xa6b\xa5t\xa4@\xad\xd3\xbch\xa6\xb8\xa5\xb4\xc4x\xb2y             "),
			},
			{
				Utf8:   " 01/23 00:22",
				Big5:   []byte(" 01/23 00:22"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:22\r"),
			},
		},
		{ // 107
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "jjj1004",
				Big5:   []byte("jjj1004"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mjjj1004"),
			},
			{
				Utf8:   ":怎麼又從JONDAN變成KOBE遜色不少                         ",
				Big5:   []byte(":\xab\xe7\xbb\xf2\xa4S\xb1qJONDAN\xc5\xdc\xa6\xa8KOBE\xbb\xb9\xa6\xe2\xa4\xa3\xa4\xd6                         "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xab\xe7\xbb\xf2\xa4S\xb1qJONDAN\xc5\xdc\xa6\xa8KOBE\xbb\xb9\xa6\xe2\xa4\xa3\xa4\xd6                         "),
			},
			{
				Utf8:   " 01/23 00:22",
				Big5:   []byte(" 01/23 00:22"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:22\r"),
			},
		},
		{ // 108
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "EXDOG",
				Big5:   []byte("EXDOG"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mEXDOG"),
			},
			{
				Utf8:   ":唯一打過NBA 加上MLB這樣算不算紀錄?                       ",
				Big5:   []byte(":\xb0\xdf\xa4@\xa5\xb4\xb9LNBA \xa5[\xa4WMLB\xb3o\xbc\xcb\xba\xe2\xa4\xa3\xba\xe2\xac\xf6\xbf\xfd?                       "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb0\xdf\xa4@\xa5\xb4\xb9LNBA \xa5[\xa4WMLB\xb3o\xbc\xcb\xba\xe2\xa4\xa3\xba\xe2\xac\xf6\xbf\xfd?                       "),
			},
			{
				Utf8:   " 01/23 00:25",
				Big5:   []byte(" 01/23 00:25"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:25\r"),
			},
		},
		{ // 105
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "AlmaMater",
				Big5:   []byte("AlmaMater"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mAlmaMater"),
			},
			{
				Utf8:   ":喬丹是神                                             ",
				Big5:   []byte(":\xb3\xec\xa4\xa6\xacO\xaf\xab                                             "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb3\xec\xa4\xa6\xacO\xaf\xab                                             "),
			},
			{
				Utf8:   " 01/23 00:25",
				Big5:   []byte(" 01/23 00:25"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:25\r"),
			},
		},
		{ // 106
			{
				Utf8:   "→ ",
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Utf8:   "fox0922",
				Big5:   []byte("fox0922"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mfox0922"),
			},
			{
				Utf8:   ":MLB有個投手在小牛打過 Jordan沒進過MLB吧...             ",
				Big5:   []byte(":MLB\xa6\xb3\xad\xd3\xa7\xeb\xa4\xe2\xa6b\xa4p\xa4\xfb\xa5\xb4\xb9L Jordan\xa8S\xb6i\xb9LMLB\xa7a...             "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:MLB\xa6\xb3\xad\xd3\xa7\xeb\xa4\xe2\xa6b\xa4p\xa4\xfb\xa5\xb4\xb9L Jordan\xa8S\xb6i\xb9LMLB\xa7a...             "),
			},
			{
				Utf8:   " 01/23 00:26",
				Big5:   []byte(" 01/23 00:26"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:26\r"),
			},
		},
		{ // 112
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "luxylu",
				Big5:   []byte("luxylu"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mluxylu"),
			},
			{
				Utf8:   ":2002... 都38歲了還連續得分22分    果然厲害              ",
				Big5:   []byte(":2002... \xb3\xa338\xb7\xb3\xa4F\xc1\xd9\xb3s\xc4\xf2\xb1o\xa4\xc022\xa4\xc0    \xaaG\xb5M\xbcF\xae`              "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:2002... \xb3\xa338\xb7\xb3\xa4F\xc1\xd9\xb3s\xc4\xf2\xb1o\xa4\xc022\xa4\xc0    \xaaG\xb5M\xbcF\xae`              "),
			},
			{
				Utf8:   " 01/23 00:27",
				Big5:   []byte(" 01/23 00:27"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:27\r"),
			},
		},
		{ // 108
			{
				Utf8:   "※ 編輯: sk2g            來自: 61.228.169.158       (01/23 00:40)",
				Big5:   []byte("\xa1\xb0 \xbds\xbf\xe8: sk2g            \xa8\xd3\xa6\xdb: 61.228.169.158       (01/23 00:40)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xb0 \xbds\xbf\xe8: sk2g            \xa8\xd3\xa6\xdb: 61.228.169.158       (01/23 00:40)\r"),
			},
		},
		{ // 109
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "accprote",
				Big5:   []byte("accprote"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33maccprote"),
			},
			{
				Utf8:   ":他只在勇士還是白襪的2A待過而已,根本沒上Major          ",
				Big5:   []byte(":\xa5L\xa5u\xa6b\xabi\xa4h\xc1\xd9\xacO\xa5\xd5\xc4\xfb\xaa\xba2A\xab\xdd\xb9L\xa6\xd3\xa4w,\xae\xda\xa5\xbb\xa8S\xa4WMajor          "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa5L\xa5u\xa6b\xabi\xa4h\xc1\xd9\xacO\xa5\xd5\xc4\xfb\xaa\xba2A\xab\xdd\xb9L\xa6\xd3\xa4w,\xae\xda\xa5\xbb\xa8S\xa4WMajor          "),
			},
			{
				Utf8:   " 01/23 00:31",
				Big5:   []byte(" 01/23 00:31"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:31\r"),
			},
		},
		{ // 115
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "boatbear",
				Big5:   []byte("boatbear"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mboatbear"),
			},
			{
				Utf8:   ":我想問91~為什麼熱火要掛他球衣?                        ",
				Big5:   []byte(":\xa7\xda\xb7Q\xb0\xdd91~\xac\xb0\xa4\xb0\xbb\xf2\xbc\xf6\xa4\xf5\xadn\xb1\xbe\xa5L\xb2y\xa6\xe7?                        "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa7\xda\xb7Q\xb0\xdd91~\xac\xb0\xa4\xb0\xbb\xf2\xbc\xf6\xa4\xf5\xadn\xb1\xbe\xa5L\xb2y\xa6\xe7?                        "),
			},
			{
				Utf8:   " 01/23 00:34",
				Big5:   []byte(" 01/23 00:34"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:34\r"),
			},
		},
		{ // 111
			{
				Utf8:   "→ ",
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Utf8:   "fox0922",
				Big5:   []byte("fox0922"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mfox0922"),
			},
			{
				Utf8:   ":我記得油頭好像說因為他太偉大之類的                     ",
				Big5:   []byte(":\xa7\xda\xb0O\xb1o\xaao\xc0Y\xa6n\xb9\xb3\xbb\xa1\xa6]\xac\xb0\xa5L\xa4\xd3\xb0\xb6\xa4j\xa4\xa7\xc3\xfe\xaa\xba                     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa7\xda\xb0O\xb1o\xaao\xc0Y\xa6n\xb9\xb3\xbb\xa1\xa6]\xac\xb0\xa5L\xa4\xd3\xb0\xb6\xa4j\xa4\xa7\xc3\xfe\xaa\xba                     "),
			},
			{
				Utf8:   " 01/23 00:37",
				Big5:   []byte(" 01/23 00:37"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:37\r"),
			},
		},
		{ // 117
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "jiunyilee",
				Big5:   []byte("jiunyilee"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mjiunyilee"),
			},
			{
				Utf8:   ":http://0rz.tw/t1Qpb                                  ",
				Big5:   []byte(":http://0rz.tw/t1Qpb                                  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:http://0rz.tw/t1Qpb                                  "),
			},
			{
				Utf8:   " 01/23 00:47",
				Big5:   []byte(" 01/23 00:47"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:47\r"),
			},
		},
		{ // 118
			{
				Utf8:   "→ ",
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Utf8:   "jiunyilee",
				Big5:   []byte("jiunyilee"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mjiunyilee"),
			},
			{
				Utf8:   ":http://0rz.tw/Ksn7u 四對大三元的數據                 ",
				Big5:   []byte(":http://0rz.tw/Ksn7u \xa5|\xb9\xef\xa4j\xa4T\xa4\xb8\xaa\xba\xbc\xc6\xbe\xda                 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:http://0rz.tw/Ksn7u \xa5|\xb9\xef\xa4j\xa4T\xa4\xb8\xaa\xba\xbc\xc6\xbe\xda                 "),
			},
			{
				Utf8:   " 01/23 00:49",
				Big5:   []byte(" 01/23 00:49"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:49\r"),
			},
		},
		{ // 119
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "joloucow",
				Big5:   []byte("joloucow"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mjoloucow"),
			},
			{
				Utf8:   ":58是1993 :) 筆誤囉~                                   ",
				Big5:   []byte(":58\xacO1993 :) \xb5\xa7\xbb~\xc5o~                                   "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:58\xacO1993 :) \xb5\xa7\xbb~\xc5o~                                   "),
			},
			{
				Utf8:   " 01/23 00:57",
				Big5:   []byte(" 01/23 00:57"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:57\r"),
			},
		},
		{ // 120
			{
				Utf8:   "→ ",
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Utf8:   "sk2g",
				Big5:   []byte("sk2g"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33msk2g"),
			},
			{
				Utf8:   ":OOPS~~對耶!!謝謝啦XD                                      ",
				Big5:   []byte(":OOPS~~\xb9\xef\xadC!!\xc1\xc2\xc1\xc2\xb0\xd5XD                                      "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:OOPS~~\xb9\xef\xadC!!\xc1\xc2\xc1\xc2\xb0\xd5XD                                      "),
			},
			{
				Utf8:   " 01/23 00:59",
				Big5:   []byte(" 01/23 00:59"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 00:59\r"),
			},
		},
		{ // 121
			{
				Utf8:   "※ 編輯: sk2g            來自: 61.228.169.158       (01/23 01:00)",
				Big5:   []byte("\xa1\xb0 \xbds\xbf\xe8: sk2g            \xa8\xd3\xa6\xdb: 61.228.169.158       (01/23 01:00)"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xb0 \xbds\xbf\xe8: sk2g            \xa8\xd3\xa6\xdb: 61.228.169.158       (01/23 01:00)\r"),
			},
		},
		{ // 122
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "timohu",
				Big5:   []byte("timohu"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mtimohu"),
			},
			{
				Utf8:   ":第58總決賽應該是1993對太陽的老巴～可能原PO筆誤打成1983  ",
				Big5:   []byte(":\xb2\xc458\xc1`\xa8M\xc1\xc9\xc0\xb3\xb8\xd3\xacO1993\xb9\xef\xa4\xd3\xb6\xa7\xaa\xba\xa6\xd1\xa4\xda\xa1\xe3\xa5i\xaf\xe0\xad\xecPO\xb5\xa7\xbb~\xa5\xb4\xa6\xa81983  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb2\xc458\xc1`\xa8M\xc1\xc9\xc0\xb3\xb8\xd3\xacO1993\xb9\xef\xa4\xd3\xb6\xa7\xaa\xba\xa6\xd1\xa4\xda\xa1\xe3\xa5i\xaf\xe0\xad\xecPO\xb5\xa7\xbb~\xa5\xb4\xa6\xa81983  "),
			},
			{
				Utf8:   " 01/23 01:00",
				Big5:   []byte(" 01/23 01:00"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:00\r"),
			},
		},
		{ // 123
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "kaiDX",
				Big5:   []byte("kaiDX"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mkaiDX"),
			},
			{
				Utf8:   ":推好文 感謝原PO                                          ",
				Big5:   []byte(":\xb1\xc0\xa6n\xa4\xe5 \xb7P\xc1\xc2\xad\xecPO                                          "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb1\xc0\xa6n\xa4\xe5 \xb7P\xc1\xc2\xad\xecPO                                          "),
			},
			{
				Utf8:   " 01/23 01:09",
				Big5:   []byte(" 01/23 01:09"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:09\r"),
			},
		},
		{ // 124
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "kalfan1",
				Big5:   []byte("kalfan1"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mkalfan1"),
			},
			{
				Utf8:   ":太神了 NBA就看他表演                                   ",
				Big5:   []byte(":\xa4\xd3\xaf\xab\xa4F NBA\xb4N\xac\xdd\xa5L\xaa\xed\xbat                                   "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa4\xd3\xaf\xab\xa4F NBA\xb4N\xac\xdd\xa5L\xaa\xed\xbat                                   "),
			},
			{
				Utf8:   " 01/23 01:12",
				Big5:   []byte(" 01/23 01:12"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:12\r"),
			},
		},
		{ // 120
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "clenny",
				Big5:   []byte("clenny"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mclenny"),
			},
			{
				Utf8:   ":年度第一隊記錄有可能被TD打破，TD目前累計9次(2次第二隊)  ",
				Big5:   []byte(":\xa6~\xab\xd7\xb2\xc4\xa4@\xb6\xa4\xb0O\xbf\xfd\xa6\xb3\xa5i\xaf\xe0\xb3QTD\xa5\xb4\xaf}\xa1ATD\xa5\xd8\xabe\xb2\xd6\xadp9\xa6\xb8(2\xa6\xb8\xb2\xc4\xa4G\xb6\xa4)  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa6~\xab\xd7\xb2\xc4\xa4@\xb6\xa4\xb0O\xbf\xfd\xa6\xb3\xa5i\xaf\xe0\xb3QTD\xa5\xb4\xaf}\xa1ATD\xa5\xd8\xabe\xb2\xd6\xadp9\xa6\xb8(2\xa6\xb8\xb2\xc4\xa4G\xb6\xa4)  "),
			},
			{
				Utf8:   " 01/23 01:12",
				Big5:   []byte(" 01/23 01:12"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:12\r"),
			},
		},
		{ // 121
			{
				Utf8:   "→ ",
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Utf8:   "clenny",
				Big5:   []byte("clenny"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mclenny"),
			},
			{
				Utf8:   ":更正一下，生涯第一隊：老馬11次、賈霸10次，所以這應該    ",
				Big5:   []byte(":\xa7\xf3\xa5\xbf\xa4@\xa4U\xa1A\xa5\xcd\xb2P\xb2\xc4\xa4@\xb6\xa4\xa1G\xa6\xd1\xb0\xa811\xa6\xb8\xa1B\xb8\xeb\xc5Q10\xa6\xb8\xa1A\xa9\xd2\xa5H\xb3o\xc0\xb3\xb8\xd3    "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa7\xf3\xa5\xbf\xa4@\xa4U\xa1A\xa5\xcd\xb2P\xb2\xc4\xa4@\xb6\xa4\xa1G\xa6\xd1\xb0\xa811\xa6\xb8\xa1B\xb8\xeb\xc5Q10\xa6\xb8\xa1A\xa9\xd2\xa5H\xb3o\xc0\xb3\xb8\xd3    "),
			},
			{
				Utf8:   " 01/23 01:25",
				Big5:   []byte(" 01/23 01:25"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:25\r"),
			},
		},
		{ // 122
			{
				Utf8:   "→ ",
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Utf8:   "clenny",
				Big5:   []byte("clenny"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mclenny"),
			},
			{
				Utf8:   ":不能算是他獨有的記錄XD (魔術和大鳥也都是生涯9次)        ",
				Big5:   []byte(":\xa4\xa3\xaf\xe0\xba\xe2\xacO\xa5L\xbfW\xa6\xb3\xaa\xba\xb0O\xbf\xfdXD (\xc5]\xb3N\xa9M\xa4j\xb3\xbe\xa4]\xb3\xa3\xacO\xa5\xcd\xb2P9\xa6\xb8)        "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa4\xa3\xaf\xe0\xba\xe2\xacO\xa5L\xbfW\xa6\xb3\xaa\xba\xb0O\xbf\xfdXD (\xc5]\xb3N\xa9M\xa4j\xb3\xbe\xa4]\xb3\xa3\xacO\xa5\xcd\xb2P9\xa6\xb8)        "),
			},
			{
				Utf8:   " 01/23 01:27",
				Big5:   []byte(" 01/23 01:27"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:27\r"),
			},
		},
		{ // 123
			{
				Utf8:   "→ ",
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Utf8:   "urinmymind",
				Big5:   []byte("urinmymind"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33murinmymind"),
			},
			{
				Utf8:   ":有鄉民100條都看完嗎？  ~\"~                          ",
				Big5:   []byte(":\xa6\xb3\xb6m\xa5\xc1100\xb1\xf8\xb3\xa3\xac\xdd\xa7\xb9\xb6\xdc\xa1H  ~\"~                          "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa6\xb3\xb6m\xa5\xc1100\xb1\xf8\xb3\xa3\xac\xdd\xa7\xb9\xb6\xdc\xa1H  ~\"~                          "),
			},
			{
				Utf8:   " 01/23 01:32",
				Big5:   []byte(" 01/23 01:32"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:32\r"),
			},
		},
		{ // 124
			{
				Utf8:   "→ ",
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Utf8:   "sk2g",
				Big5:   []byte("sk2g"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33msk2g"),
			},
			{
				Utf8:   ":有啊...我 ...=_=....而且是94唷                            ",
				Big5:   []byte(":\xa6\xb3\xb0\xda...\xa7\xda ...=_=....\xa6\xd3\xa5B\xacO94\xad\xf2                            "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa6\xb3\xb0\xda...\xa7\xda ...=_=....\xa6\xd3\xa5B\xacO94\xad\xf2                            "),
			},
			{
				Utf8:   " 01/23 01:35",
				Big5:   []byte(" 01/23 01:35"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:35\r"),
			},
		},
		{ // 125
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "MousePads",
				Big5:   []byte("MousePads"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mMousePads"),
			},
			{
				Utf8:   ":關於第22點，最高頻率的意思是？Big O的紀錄是多少啊？  ",
				Big5:   []byte(":\xc3\xf6\xa9\xf3\xb2\xc422\xc2I\xa1A\xb3\xcc\xb0\xaa\xc0W\xb2v\xaa\xba\xb7N\xab\xe4\xacO\xa1HBig O\xaa\xba\xac\xf6\xbf\xfd\xacO\xa6h\xa4\xd6\xb0\xda\xa1H  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xc3\xf6\xa9\xf3\xb2\xc422\xc2I\xa1A\xb3\xcc\xb0\xaa\xc0W\xb2v\xaa\xba\xb7N\xab\xe4\xacO\xa1HBig O\xaa\xba\xac\xf6\xbf\xfd\xacO\xa6h\xa4\xd6\xb0\xda\xa1H  "),
			},
			{
				Utf8:   " 01/23 01:36",
				Big5:   []byte(" 01/23 01:36"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:36\r"),
			},
		},
		{ // 126
			{
				Utf8:   "→ ",
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Utf8:   "sk2g",
				Big5:   []byte("sk2g"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33msk2g"),
			},
			{
				Utf8:   ":我猜測...可能是說他大三元的場次中，連續的次數比例之高吧!  ",
				Big5:   []byte(":\xa7\xda\xb2q\xb4\xfa...\xa5i\xaf\xe0\xacO\xbb\xa1\xa5L\xa4j\xa4T\xa4\xb8\xaa\xba\xb3\xf5\xa6\xb8\xa4\xa4\xa1A\xb3s\xc4\xf2\xaa\xba\xa6\xb8\xbc\xc6\xa4\xf1\xa8\xd2\xa4\xa7\xb0\xaa\xa7a!  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa7\xda\xb2q\xb4\xfa...\xa5i\xaf\xe0\xacO\xbb\xa1\xa5L\xa4j\xa4T\xa4\xb8\xaa\xba\xb3\xf5\xa6\xb8\xa4\xa4\xa1A\xb3s\xc4\xf2\xaa\xba\xa6\xb8\xbc\xc6\xa4\xf1\xa8\xd2\xa4\xa7\xb0\xaa\xa7a!  "),
			},
			{
				Utf8:   " 01/23 01:39",
				Big5:   []byte(" 01/23 01:39"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:39\r"),
			},
		},
		{ // 127
			{
				Utf8:   "→ ",
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Utf8:   "sk2g",
				Big5:   []byte("sk2g"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33msk2g"),
			},
			{
				Utf8:   ":大帥78次裡面有9次連續，神30次(含playoff)有7次...          ",
				Big5:   []byte(":\xa4j\xab\xd378\xa6\xb8\xb8\xcc\xad\xb1\xa6\xb39\xa6\xb8\xb3s\xc4\xf2\xa1A\xaf\xab30\xa6\xb8(\xa7tplayoff)\xa6\xb37\xa6\xb8...          "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa4j\xab\xd378\xa6\xb8\xb8\xcc\xad\xb1\xa6\xb39\xa6\xb8\xb3s\xc4\xf2\xa1A\xaf\xab30\xa6\xb8(\xa7tplayoff)\xa6\xb37\xa6\xb8...          "),
			},
			{
				Utf8:   " 01/23 01:44",
				Big5:   []byte(" 01/23 01:44"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:44\r"),
			},
		},
		{ // 128
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "MousePads",
				Big5:   []byte("MousePads"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mMousePads"),
			},
			{
				Utf8:   ":這樣的說法好怪，那生涯只有1次的不是應該最高 @@\"      ",
				Big5:   []byte(":\xb3o\xbc\xcb\xaa\xba\xbb\xa1\xaak\xa6n\xa9\xc7\xa1A\xa8\xba\xa5\xcd\xb2P\xa5u\xa6\xb31\xa6\xb8\xaa\xba\xa4\xa3\xacO\xc0\xb3\xb8\xd3\xb3\xcc\xb0\xaa @@\"      "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb3o\xbc\xcb\xaa\xba\xbb\xa1\xaak\xa6n\xa9\xc7\xa1A\xa8\xba\xa5\xcd\xb2P\xa5u\xa6\xb31\xa6\xb8\xaa\xba\xa4\xa3\xacO\xc0\xb3\xb8\xd3\xb3\xcc\xb0\xaa @@\"      "),
			},
			{
				Utf8:   " 01/23 01:57",
				Big5:   []byte(" 01/23 01:57"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:57\r"),
			},
		},
		{ // 129
			{
				Utf8:   "→ ",
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Utf8:   "sk2g",
				Big5:   []byte("sk2g"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33msk2g"),
			},
			{
				Utf8:   ":可是連續1場怪怪的吧XD   連續2場的那我還覺得比較奇怪@@     ",
				Big5:   []byte(":\xa5i\xacO\xb3s\xc4\xf21\xb3\xf5\xa9\xc7\xa9\xc7\xaa\xba\xa7aXD   \xb3s\xc4\xf22\xb3\xf5\xaa\xba\xa8\xba\xa7\xda\xc1\xd9\xc4\xb1\xb1o\xa4\xf1\xb8\xfb\xa9_\xa9\xc7@@     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa5i\xacO\xb3s\xc4\xf21\xb3\xf5\xa9\xc7\xa9\xc7\xaa\xba\xa7aXD   \xb3s\xc4\xf22\xb3\xf5\xaa\xba\xa8\xba\xa7\xda\xc1\xd9\xc4\xb1\xb1o\xa4\xf1\xb8\xfb\xa9_\xa9\xc7@@     "),
			},
			{
				Utf8:   " 01/23 01:58",
				Big5:   []byte(" 01/23 01:58"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:58\r"),
			},
		},
		{ // 130
			{
				Utf8:   "→ ",
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Utf8:   "sk2g",
				Big5:   []byte("sk2g"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33msk2g"),
			},
			{
				Utf8:   ":算了，讓我們擁抱板上的真˙強者吧                          ",
				Big5:   []byte(":\xba\xe2\xa4F\xa1A\xc5\xfd\xa7\xda\xad\xcc\xbe\xd6\xa9\xea\xaaO\xa4W\xaa\xba\xafu\xa3\xbb\xb1j\xaa\xcc\xa7a                          "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xba\xe2\xa4F\xa1A\xc5\xfd\xa7\xda\xad\xcc\xbe\xd6\xa9\xea\xaaO\xa4W\xaa\xba\xafu\xa3\xbb\xb1j\xaa\xcc\xa7a                          "),
			},
			{
				Utf8:   " 01/23 01:59",
				Big5:   []byte(" 01/23 01:59"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 01:59\r"),
			},
		},
		{ // 131
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "molemilk",
				Big5:   []byte("molemilk"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mmolemilk"),
			},
			{
				Utf8:   ":神                                                    ",
				Big5:   []byte(":\xaf\xab                                                    "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xaf\xab                                                    "),
			},
			{
				Utf8:   " 01/23 02:01",
				Big5:   []byte(" 01/23 02:01"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 02:01\r"),
			},
		},
		{ // 132
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "blinder",
				Big5:   []byte("blinder"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mblinder"),
			},
			{
				Utf8:   ":超認真好文~                                            ",
				Big5:   []byte(":\xb6W\xbb{\xafu\xa6n\xa4\xe5~                                            "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb6W\xbb{\xafu\xa6n\xa4\xe5~                                            "),
			},
			{
				Utf8:   " 01/23 02:13",
				Big5:   []byte(" 01/23 02:13"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 02:13\r"),
			},
		},
		{ // 133
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "samprus",
				Big5:   []byte("samprus"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33msamprus"),
			},
			{
				Utf8:   ":我覺得後衛的阻攻記錄今年會被wade打破XD                 ",
				Big5:   []byte(":\xa7\xda\xc4\xb1\xb1o\xab\xe1\xbd\xc3\xaa\xba\xaa\xfd\xa7\xf0\xb0O\xbf\xfd\xa4\xb5\xa6~\xb7|\xb3Qwade\xa5\xb4\xaf}XD                 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: false},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa7\xda\xc4\xb1\xb1o\xab\xe1\xbd\xc3\xaa\xba\xaa\xfd\xa7\xf0\xb0O\xbf\xfd\xa4\xb5\xa6~\xb7|\xb3Qwade\xa5\xb4\xaf}XD                 "),
			},
			{
				Utf8:   " 01/23 02:29",
				Big5:   []byte(" 01/23 02:29"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 02:29\r"),
			},
		},
		{ // 134
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "avrilrock",
				Big5:   []byte("avrilrock"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mavrilrock"),
			},
			{
				Utf8:   ":好文                                                 ",
				Big5:   []byte(":\xa6n\xa4\xe5                                                 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa6n\xa4\xe5                                                 "),
			},
			{
				Utf8:   " 01/23 02:46",
				Big5:   []byte(" 01/23 02:46"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 02:46\r"),
			},
		},
		{ // 135
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "Kaverson",
				Big5:   []byte("Kaverson"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mKaverson"),
			},
			{
				Utf8:   ":推好文..XD...辛苦你了..:)                             ",
				Big5:   []byte(":\xb1\xc0\xa6n\xa4\xe5..XD...\xa8\xaf\xadW\xa7A\xa4F..:)                             "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb1\xc0\xa6n\xa4\xe5..XD...\xa8\xaf\xadW\xa7A\xa4F..:)                             "),
			},
			{
				Utf8:   " 01/23 03:36",
				Big5:   []byte(" 01/23 03:36"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 03:36\r"),
			},
		},
		{ // 136
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "Yao1218",
				Big5:   []byte("Yao1218"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mYao1218"),
			},
			{
				Utf8:   ":幫推 好文                                              ",
				Big5:   []byte(":\xc0\xb0\xb1\xc0 \xa6n\xa4\xe5                                              "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xc0\xb0\xb1\xc0 \xa6n\xa4\xe5                                              "),
			},
			{
				Utf8:   " 01/23 03:42",
				Big5:   []byte(" 01/23 03:42"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 03:42\r"),
			},
		},
		{ // 137
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "BrentRoy",
				Big5:   []byte("BrentRoy"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mBrentRoy"),
			},
			{
				Utf8:   ":推好文! 但是可能沒幾天就會被記者抄走了= =             ",
				Big5:   []byte(":\xb1\xc0\xa6n\xa4\xe5! \xa6\xfd\xacO\xa5i\xaf\xe0\xa8S\xb4X\xa4\xd1\xb4N\xb7|\xb3Q\xb0O\xaa\xcc\xa7\xdb\xa8\xab\xa4F= =             "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb1\xc0\xa6n\xa4\xe5! \xa6\xfd\xacO\xa5i\xaf\xe0\xa8S\xb4X\xa4\xd1\xb4N\xb7|\xb3Q\xb0O\xaa\xcc\xa7\xdb\xa8\xab\xa4F= =             "),
			},
			{
				Utf8:   " 01/23 04:52",
				Big5:   []byte(" 01/23 04:52"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 04:52\r"),
			},
		},
		{ // 138
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "ksk0516",
				Big5:   []byte("ksk0516"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mksk0516"),
			},
			{
				Utf8:   ":每次看到有人說Kobe、LBJ超越Jordan就很想笑 XD           ",
				Big5:   []byte(":\xa8C\xa6\xb8\xac\xdd\xa8\xec\xa6\xb3\xa4H\xbb\xa1Kobe\xa1BLBJ\xb6W\xb6VJordan\xb4N\xab\xdc\xb7Q\xaf\xba XD           "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa8C\xa6\xb8\xac\xdd\xa8\xec\xa6\xb3\xa4H\xbb\xa1Kobe\xa1BLBJ\xb6W\xb6VJordan\xb4N\xab\xdc\xb7Q\xaf\xba XD           "),
			},
			{
				Utf8:   " 01/23 05:25",
				Big5:   []byte(" 01/23 05:25"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 05:25\r"),
			},
		},
		{ // 139
			{
				Utf8:   "→ ",
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Utf8:   "rumourmonger",
				Big5:   []byte("rumourmonger"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mrumourmonger"),
			},
			{
				Utf8:   ":Kobe跟LBJ如果沒有因傷退休 將來生涯成績比肩MJ正常  ",
				Big5:   []byte(":Kobe\xb8\xf2LBJ\xa6p\xaaG\xa8S\xa6\xb3\xa6]\xb6\xcb\xb0h\xa5\xf0 \xb1N\xa8\xd3\xa5\xcd\xb2P\xa6\xa8\xc1Z\xa4\xf1\xaa\xd3MJ\xa5\xbf\xb1`  "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:Kobe\xb8\xf2LBJ\xa6p\xaaG\xa8S\xa6\xb3\xa6]\xb6\xcb\xb0h\xa5\xf0 \xb1N\xa8\xd3\xa5\xcd\xb2P\xa6\xa8\xc1Z\xa4\xf1\xaa\xd3MJ\xa5\xbf\xb1`  "),
			},
			{
				Utf8:   " 01/23 05:44",
				Big5:   []byte(" 01/23 05:44"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 05:44\r"),
			},
		},
		{ // 140
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "chengyuyang",
				Big5:   []byte("chengyuyang"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mchengyuyang"),
			},
			{
				Utf8:   ":推好文n最後一句      在別人家掛球衣........神      ",
				Big5:   []byte(":\xb1\xc0\xa6n\xa4\xe5n\xb3\xcc\xab\xe1\xa4@\xa5y      \xa6b\xa7O\xa4H\xaea\xb1\xbe\xb2y\xa6\xe7........\xaf\xab      "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb1\xc0\xa6n\xa4\xe5n\xb3\xcc\xab\xe1\xa4@\xa5y      \xa6b\xa7O\xa4H\xaea\xb1\xbe\xb2y\xa6\xe7........\xaf\xab      "),
			},
			{
				Utf8:   " 01/23 05:45",
				Big5:   []byte(" 01/23 05:45"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 05:45\r"),
			},
		},
		{ // 141
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "kobe872125",
				Big5:   []byte("kobe872125"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mkobe872125"),
			},
			{
				Utf8:   ":我記得曾有一場喬神跟皮朋2個單場40+ 有鄉民記得嗎？   ",
				Big5:   []byte(":\xa7\xda\xb0O\xb1o\xb4\xbf\xa6\xb3\xa4@\xb3\xf5\xb3\xec\xaf\xab\xb8\xf2\xa5\xd6\xaaB2\xad\xd3\xb3\xe6\xb3\xf540+ \xa6\xb3\xb6m\xa5\xc1\xb0O\xb1o\xb6\xdc\xa1H   "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa7\xda\xb0O\xb1o\xb4\xbf\xa6\xb3\xa4@\xb3\xf5\xb3\xec\xaf\xab\xb8\xf2\xa5\xd6\xaaB2\xad\xd3\xb3\xe6\xb3\xf540+ \xa6\xb3\xb6m\xa5\xc1\xb0O\xb1o\xb6\xdc\xa1H   "),
			},
			{
				Utf8:   " 01/23 06:09",
				Big5:   []byte(" 01/23 06:09"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 06:09\r"),
			},
		},
		{ // 142
			{
				Utf8:   "→ ",
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Utf8:   "kobe872125",
				Big5:   []byte("kobe872125"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mkobe872125"),
			},
			{
				Utf8:   ":就是一場比賽兩個人都得40+分.不知我有沒有記錯        ",
				Big5:   []byte(":\xb4N\xacO\xa4@\xb3\xf5\xa4\xf1\xc1\xc9\xa8\xe2\xad\xd3\xa4H\xb3\xa3\xb1o40+\xa4\xc0.\xa4\xa3\xaa\xbe\xa7\xda\xa6\xb3\xa8S\xa6\xb3\xb0O\xbf\xf9        "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb4N\xacO\xa4@\xb3\xf5\xa4\xf1\xc1\xc9\xa8\xe2\xad\xd3\xa4H\xb3\xa3\xb1o40+\xa4\xc0.\xa4\xa3\xaa\xbe\xa7\xda\xa6\xb3\xa8S\xa6\xb3\xb0O\xbf\xf9        "),
			},
			{
				Utf8:   " 01/23 06:11",
				Big5:   []byte(" 01/23 06:11"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 06:11\r"),
			},
		},
		{ // 143
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "henrie",
				Big5:   []byte("henrie"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mhenrie"),
			},
			{
				Utf8:   ":防...防守王?!                                           ",
				Big5:   []byte(":\xa8\xbe...\xa8\xbe\xa6u\xa4\xfd?!                                           "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa8\xbe...\xa8\xbe\xa6u\xa4\xfd?!                                           "),
			},
			{
				Utf8:   " 01/23 06:39",
				Big5:   []byte(" 01/23 06:39"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 06:39\r"),
			},
		},
		{ // 144
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "chengyuyang",
				Big5:   []byte("chengyuyang"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mchengyuyang"),
			},
			{
				Utf8:   ":那場苦主Pacers 喬44 皮40                           ",
				Big5:   []byte(":\xa8\xba\xb3\xf5\xadW\xa5DPacers \xb3\xec44 \xa5\xd640                           "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa8\xba\xb3\xf5\xadW\xa5DPacers \xb3\xec44 \xa5\xd640                           "),
			},
			{
				Utf8:   " 01/23 06:42",
				Big5:   []byte(" 01/23 06:42"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 06:42\r"),
			},
		},
		{ // 145
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "chengyuyang",
				Big5:   []byte("chengyuyang"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mchengyuyang"),
			},
			{
				Utf8:   ":活賽和金塊 史上最高分那場 兩隊都有兩個人超過40     ",
				Big5:   []byte(":\xac\xa1\xc1\xc9\xa9M\xaa\xf7\xb6\xf4 \xa5v\xa4W\xb3\xcc\xb0\xaa\xa4\xc0\xa8\xba\xb3\xf5 \xa8\xe2\xb6\xa4\xb3\xa3\xa6\xb3\xa8\xe2\xad\xd3\xa4H\xb6W\xb9L40     "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xac\xa1\xc1\xc9\xa9M\xaa\xf7\xb6\xf4 \xa5v\xa4W\xb3\xcc\xb0\xaa\xa4\xc0\xa8\xba\xb3\xf5 \xa8\xe2\xb6\xa4\xb3\xa3\xa6\xb3\xa8\xe2\xad\xd3\xa4H\xb6W\xb9L40     "),
			},
			{
				Utf8:   " 01/23 06:53",
				Big5:   []byte(" 01/23 06:53"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 06:53\r"),
			},
		},
		{ // 146
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{ // 147
			{
				Utf8:   "--",
				Big5:   []byte("--"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("--\r"),
			},
		},
		{ // 148
			{
				Utf8:   "※ 發信站: 批踢踢實業坊(ptt.cc) ",
				Big5:   []byte("\xa1\xb0 \xb5o\xabH\xaf\xb8: \xa7\xe5\xbd\xf0\xbd\xf0\xb9\xea\xb7~\xa7{(ptt.cc) "),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xb0 \xb5o\xabH\xaf\xb8: \xa7\xe5\xbd\xf0\xbd\xf0\xb9\xea\xb7~\xa7{(ptt.cc) \r"),
			},
		},
		{ // 149
			{
				Utf8:   "◆ From: 61.228.174.209",
				Big5:   []byte("\xa1\xbb From: 61.228.174.209"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xa1\xbb From: 61.228.174.209\r"),
			},
		},
		{ // 150
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "airLfly",
				Big5:   []byte("airLfly"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mairLfly"),
			},
			{
				Utf8:   ":真夭壽 太猛了                                          ",
				Big5:   []byte(":\xafu\xa4\xd4\xb9\xd8 \xa4\xd3\xb2r\xa4F                                          "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xafu\xa4\xd4\xb9\xd8 \xa4\xd3\xb2r\xa4F                                          "),
			},
			{
				Utf8:   " 01/23 21:32",
				Big5:   []byte(" 01/23 21:32"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/23 21:32\r"),
			},
		},
		{ // 151
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "chron",
				Big5:   []byte("chron"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mchron"),
			},
			{
				Utf8:   ":永遠的神.....                                            ",
				Big5:   []byte(":\xa5\xc3\xbb\xb7\xaa\xba\xaf\xab.....                                            "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa5\xc3\xbb\xb7\xaa\xba\xaf\xab.....                                            "),
			},
			{
				Utf8:   " 01/29 12:40",
				Big5:   []byte(" 01/29 12:40"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 01/29 12:40\r"),
			},
		},
		{ // 152
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "s9588008",
				Big5:   []byte("s9588008"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33ms9588008"),
			},
			{
				Utf8:   ":季後賽看到一半就來推了~ 實在太強大了!!   Q_Q          ",
				Big5:   []byte(":\xa9u\xab\xe1\xc1\xc9\xac\xdd\xa8\xec\xa4@\xa5b\xb4N\xa8\xd3\xb1\xc0\xa4F~ \xb9\xea\xa6b\xa4\xd3\xb1j\xa4j\xa4F!!   Q_Q          "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa9u\xab\xe1\xc1\xc9\xac\xdd\xa8\xec\xa4@\xa5b\xb4N\xa8\xd3\xb1\xc0\xa4F~ \xb9\xea\xa6b\xa4\xd3\xb1j\xa4j\xa4F!!   Q_Q          "),
			},
			{
				Utf8:   " 02/05 06:23",
				Big5:   []byte(" 02/05 06:23"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 02/05 06:23\r"),
			},
		},
		{ // 153
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "s9588008",
				Big5:   []byte("s9588008"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33ms9588008"),
			},
			{
				Utf8:   ":趣味紀錄:  89.唯一做到三進三出的人。  <-   XD         ",
				Big5:   []byte(":\xbd\xec\xa8\xfd\xac\xf6\xbf\xfd:  89.\xb0\xdf\xa4@\xb0\xb5\xa8\xec\xa4T\xb6i\xa4T\xa5X\xaa\xba\xa4H\xa1C  <-   XD         "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xbd\xec\xa8\xfd\xac\xf6\xbf\xfd:  89.\xb0\xdf\xa4@\xb0\xb5\xa8\xec\xa4T\xb6i\xa4T\xa5X\xaa\xba\xa4H\xa1C  <-   XD         "),
			},
			{
				Utf8:   " 02/05 06:28",
				Big5:   []byte(" 02/05 06:28"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 02/05 06:28\r"),
			},
		},
		{ // 154
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "bestshow500",
				Big5:   []byte("bestshow500"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mbestshow500"),
			},
			{
				Utf8:   ":世紀運動員當之無愧　２０世紀是喬丹的籃球帝國時代   ",
				Big5:   []byte(":\xa5@\xac\xf6\xb9B\xb0\xca\xad\xfb\xb7\xed\xa4\xa7\xb5L\xb7\\\xa1@\xa2\xb1\xa2\xaf\xa5@\xac\xf6\xacO\xb3\xec\xa4\xa6\xaa\xba\xc4x\xb2y\xab\xd2\xb0\xea\xae\xc9\xa5N   "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa5@\xac\xf6\xb9B\xb0\xca\xad\xfb\xb7\xed\xa4\xa7\xb5L\xb7\\\xa1@\xa2\xb1\xa2\xaf\xa5@\xac\xf6\xacO\xb3\xec\xa4\xa6\xaa\xba\xc4x\xb2y\xab\xd2\xb0\xea\xae\xc9\xa5N   "),
			},
			{
				Utf8:   " 02/12 00:35",
				Big5:   []byte(" 02/12 00:35"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 02/12 00:35\r"),
			},
		},
		{ // 155
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "KGhuang",
				Big5:   []byte("KGhuang"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33mKGhuang"),
			},
			{
				Utf8:   ":看到下巴都快掉了...                                    ",
				Big5:   []byte(":\xac\xdd\xa8\xec\xa4U\xa4\xda\xb3\xa3\xa7\xd6\xb1\xbc\xa4F...                                    "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xac\xdd\xa8\xec\xa4U\xa4\xda\xb3\xa3\xa7\xd6\xb1\xbc\xa4F...                                    "),
			},
			{
				Utf8:   " 03/08 21:12",
				Big5:   []byte(" 03/08 21:12"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 03/08 21:12\r"),
			},
		},
		{ // 156
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "sklyn",
				Big5:   []byte("sklyn"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33msklyn"),
			},
			{
				Utf8:   ":… 我不知道要說什麼…                                    ",
				Big5:   []byte(":\xa1K \xa7\xda\xa4\xa3\xaa\xbe\xb9D\xadn\xbb\xa1\xa4\xb0\xbb\xf2\xa1K                                    "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xa1K \xa7\xda\xa4\xa3\xaa\xbe\xb9D\xadn\xbb\xa1\xa4\xb0\xbb\xf2\xa1K                                    "),
			},
			{
				Utf8:   " 03/17 00:58",
				Big5:   []byte(" 03/17 00:58"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 03/17 00:58\r"),
			},
		},
		{ // 157
			{
				Utf8:   "推 ",
				Big5:   []byte("\xb1\xc0 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_WHITE, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;37m\xb1\xc0 "),
			},
			{
				Utf8:   "sklyn",
				Big5:   []byte("sklyn"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33msklyn"),
			},
			{
				Utf8:   ":Jordan在88年對76人29投24中好像是所有50+比賽FG%最高的一場 ",
				Big5:   []byte(":Jordan\xa6b88\xa6~\xb9\xef76\xa4H29\xa7\xeb24\xa4\xa4\xa6n\xb9\xb3\xacO\xa9\xd2\xa6\xb350+\xa4\xf1\xc1\xc9FG%\xb3\xcc\xb0\xaa\xaa\xba\xa4@\xb3\xf5 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:Jordan\xa6b88\xa6~\xb9\xef76\xa4H29\xa7\xeb24\xa4\xa4\xa6n\xb9\xb3\xacO\xa9\xd2\xa6\xb350+\xa4\xf1\xc1\xc9FG%\xb3\xcc\xb0\xaa\xaa\xba\xa4@\xb3\xf5 "),
			},
			{
				Utf8:   " 04/08 12:01",
				Big5:   []byte(" 04/08 12:01"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 04/08 12:01\r"),
			},
		},
		{ // 158
			{
				Utf8:   "→ ",
				Big5:   []byte("\xa1\xf7 "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_RED, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[1;31m\xa1\xf7 "),
			},
			{
				Utf8:   "sklyn",
				Big5:   []byte("sklyn"),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK, Highlight: true},
				DBCS:   []byte("\x1b[33msklyn"),
			},
			{
				Utf8:   ":第三十五項…                                             ",
				Big5:   []byte(":\xb2\xc4\xa4T\xa4Q\xa4\xad\xb6\xb5\xa1K                                             "),
				Color0: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				Color1: types.Color{Foreground: types.COLOR_FOREGROUND_YELLOW, Background: types.COLOR_BACKGROUND_BLACK},
				DBCS:   []byte("\x1b[m\x1b[33m:\xb2\xc4\xa4T\xa4Q\xa4\xad\xb6\xb5\xa1K                                             "),
			},
			{
				Utf8:   " 04/08 12:05",
				Big5:   []byte(" 04/08 12:05"),
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\x1b[m 04/08 12:05\r"),
			},
		},
	}

	testFirstComments15 = []*schema.Comment{
		{ // 0
			TheType: ptttype.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("hunt5566"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "http://i.imgur.com/4PQq6rd.jpg",
						Big5:   []byte("http://i.imgur.com/4PQq6rd.jpg                   "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("http://i.imgur.com/4PQq6rd.jpg                   "),
					},
				},
			},
			MD5:     "xNzbuQWYR2ZPDBIsHyfyNA",
			TheDate: "06/23 12:16",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mhunt5566    \x1b[m\x1b[33m: http://i.imgur.com/4PQq6rd.jpg                   \x1b[m 06/23 12:16\r"),
		},
		{ // 1
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("sunnyyoung"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "第23項是怎樣，驕傲什麼啦~~~",
						Big5:   []byte("\xb2\xc423\xb6\xb5\xacO\xab\xe7\xbc\xcb\xa1A\xc5\xba\xb6\xc6\xa4\xb0\xbb\xf2\xb0\xd5~~~                      "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb2\xc423\xb6\xb5\xacO\xab\xe7\xbc\xcb\xa1A\xc5\xba\xb6\xc6\xa4\xb0\xbb\xf2\xb0\xd5~~~                      "),
					},
				},
			},
			MD5:     "jXWE4sLfiC2rGLr1FRdUQQ",
			TheDate: "06/23 12:18",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33msunnyyoung  \x1b[m\x1b[33m: \xb2\xc423\xb6\xb5\xacO\xab\xe7\xbc\xcb\xa1A\xc5\xba\xb6\xc6\xa4\xb0\xbb\xf2\xb0\xd5~~~                      \x1b[m 06/23 12:18\r"),
		},
		{ // 2
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("YO8BO10"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "神！！！跪！！！",
						Big5:   []byte("\xaf\xab\xa1I\xa1I\xa1I\xb8\xf7\xa1I\xa1I\xa1I                                 "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xaf\xab\xa1I\xa1I\xa1I\xb8\xf7\xa1I\xa1I\xa1I                                 "),
					},
				},
			},
			MD5:     "YSEKMIj--mQbv17h2JAXTA",
			TheDate: "06/23 12:18",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mYO8BO10     \x1b[m\x1b[33m: \xaf\xab\xa1I\xa1I\xa1I\xb8\xf7\xa1I\xa1I\xa1I                                 \x1b[m 06/23 12:18\r"),
		},
		{ // 3
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("Aggro"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "得分王同時有防守王是啥概念……",
						Big5:   []byte("\xb1o\xa4\xc0\xa4\xfd\xa6P\xae\xc9\xa6\xb3\xa8\xbe\xa6u\xa4\xfd\xacO\xd4\xa3\xb7\xa7\xa9\xc0\xa1K\xa1K                   "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb1o\xa4\xc0\xa4\xfd\xa6P\xae\xc9\xa6\xb3\xa8\xbe\xa6u\xa4\xfd\xacO\xd4\xa3\xb7\xa7\xa9\xc0\xa1K\xa1K                   "),
					},
				},
			},
			MD5:     "XbsC1DBGI67-hVLniVcQVA",
			TheDate: "06/23 12:20",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mAggro       \x1b[m\x1b[33m: \xb1o\xa4\xc0\xa4\xfd\xa6P\xae\xc9\xa6\xb3\xa8\xbe\xa6u\xa4\xfd\xacO\xd4\xa3\xb7\xa7\xa9\xc0\xa1K\xa1K                   \x1b[m 06/23 12:20\r"),
		},
		{ // 4
			TheType: ptttype.COMMENT_TYPE_BOO,
			Owner:   bbs.UUserID("lycium"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "bug",
						Big5:   []byte("bug                                              "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("bug                                              "),
					},
				},
			},
			MD5:     "JiR9wAMpshqmSUZigf4uug",
			TheDate: "06/23 12:21",
			DBCS:    []byte("\x1b[1;31m\xbcN \x1b[33mlycium      \x1b[m\x1b[33m: bug                                              \x1b[m 06/23 12:21\r"),
		},
		{ // 5
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("XaviYang"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "矮哦...威猛",
						Big5:   []byte("\xb8G\xae@...\xab\xc2\xb2r                                      "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb8G\xae@...\xab\xc2\xb2r                                      "),
					},
				},
			},
			MD5:     "-Zei7Rg7TeLUyWcUIDKowQ",
			TheDate: "06/23 12:21",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mXaviYang    \x1b[m\x1b[33m: \xb8G\xae@...\xab\xc2\xb2r                                      \x1b[m 06/23 12:21\r"),
		},
		{ // 6
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("katanakiller"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "單節13罰 鬍子也甘拜下風",
						Big5:   []byte("\xb3\xe6\xb8`13\xbb@ \xc4G\xa4l\xa4]\xa5\xcc\xab\xf4\xa4U\xad\xb7                          "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb3\xe6\xb8`13\xbb@ \xc4G\xa4l\xa4]\xa5\xcc\xab\xf4\xa4U\xad\xb7                          "),
					},
				},
			},
			MD5:     "xYcELntlC-hefUK7kfp-dA",
			TheDate: "06/23 12:21",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mkatanakiller\x1b[m\x1b[33m: \xb3\xe6\xb8`13\xbb@ \xc4G\xa4l\xa4]\xa5\xcc\xab\xf4\xa4U\xad\xb7                          \x1b[m 06/23 12:21\r"),
		},
		{ // 7
			TheType: ptttype.COMMENT_TYPE_COMMENT,
			Owner:   bbs.UUserID("lycium"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "外星人來地球虐菜不值得一提",
						Big5:   []byte("\xa5~\xacP\xa4H\xa8\xd3\xa6a\xb2y\xadh\xb5\xe6\xa4\xa3\xad\xc8\xb1o\xa4@\xb4\xa3                       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa5~\xacP\xa4H\xa8\xd3\xa6a\xb2y\xadh\xb5\xe6\xa4\xa3\xad\xc8\xb1o\xa4@\xb4\xa3                       "),
					},
				},
			},
			MD5:     "G03hZoEVJ0R8PZlEcKKxGg",
			TheDate: "06/23 12:21",
			DBCS:    []byte("\x1b[1;31m\xa1\xf7 \x1b[33mlycium      \x1b[m\x1b[33m: \xa5~\xacP\xa4H\xa8\xd3\xa6a\xb2y\xadh\xb5\xe6\xa4\xa3\xad\xc8\xb1o\xa4@\xb4\xa3                       \x1b[m 06/23 12:21\r"),
		},
		{ // 8
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("yankees733"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "說MJ要走下神壇的人呢?",
						Big5:   []byte("\xbb\xa1MJ\xadn\xa8\xab\xa4U\xaf\xab\xbe\xc2\xaa\xba\xa4H\xa9O?                            "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xbb\xa1MJ\xadn\xa8\xab\xa4U\xaf\xab\xbe\xc2\xaa\xba\xa4H\xa9O?                            "),
					},
				},
			},
			MD5:     "wHOypI5hcuHS4wVztWmYyg",
			TheDate: "06/23 12:23",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33myankees733  \x1b[m\x1b[33m: \xbb\xa1MJ\xadn\xa8\xab\xa4U\xaf\xab\xbe\xc2\xaa\xba\xa4H\xa9O?                            \x1b[m 06/23 12:23\r"),
		},
		{ // 9
			TheType: ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:   bbs.UUserID("juniorpenny"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "鞋子出到30代了啦",
						Big5:   []byte("\xbec\xa4l\xa5X\xa8\xec30\xa5N\xa4F\xb0\xd5                                 "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xbec\xa4l\xa5X\xa8\xec30\xa5N\xa4F\xb0\xd5                                 "),
					},
				},
			},
			MD5:     "DuNQJilqP07PXmhfjw3Dfw",
			TheDate: "06/23 12:24",
			DBCS:    []byte("\x1b[1;37m\xb1\xc0 \x1b[33mjuniorpenny \x1b[m\x1b[33m: \xbec\xa4l\xa5X\xa8\xec30\xa5N\xa4F\xb0\xd5                                 \x1b[m 06/23 12:24\r"),
		},
	}

	testFullFirstComments15 = []*schema.Comment{
		{ // 0
			BBoardID:  "test",
			ArticleID: "test15",
			TheType:   ptttype.COMMENT_TYPE_COMMENT,
			Owner:     bbs.UUserID("hunt5566"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "http://i.imgur.com/4PQq6rd.jpg",
						Big5:   []byte("http://i.imgur.com/4PQq6rd.jpg                   "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("http://i.imgur.com/4PQq6rd.jpg                   "),
					},
				},
			},
			MD5:        "xNzbuQWYR2ZPDBIsHyfyNA",
			TheDate:    "06/23 12:16",
			DBCS:       []byte("\x1b[1;31m\xa1\xf7 \x1b[33mhunt5566    \x1b[m\x1b[33m: http://i.imgur.com/4PQq6rd.jpg                   \x1b[m 06/23 12:16\r"),
			CommentID:  "FFqbR3y5AAA:xNzbuQWYR2ZPDBIsHyfyNA",
			CreateTime: 1466655360000000000,
			SortTime:   1466655360000000000,
		},
		{ // 1
			BBoardID:  "test",
			ArticleID: "test15",
			TheType:   ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:     bbs.UUserID("sunnyyoung"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "第23項是怎樣，驕傲什麼啦~~~",
						Big5:   []byte("\xb2\xc423\xb6\xb5\xacO\xab\xe7\xbc\xcb\xa1A\xc5\xba\xb6\xc6\xa4\xb0\xbb\xf2\xb0\xd5~~~                      "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb2\xc423\xb6\xb5\xacO\xab\xe7\xbc\xcb\xa1A\xc5\xba\xb6\xc6\xa4\xb0\xbb\xf2\xb0\xd5~~~                      "),
					},
				},
			},
			MD5:        "jXWE4sLfiC2rGLr1FRdUQQ",
			TheDate:    "06/23 12:18",
			DBCS:       []byte("\x1b[1;37m\xb1\xc0 \x1b[33msunnyyoung  \x1b[m\x1b[33m: \xb2\xc423\xb6\xb5\xacO\xab\xe7\xbc\xcb\xa1A\xc5\xba\xb6\xc6\xa4\xb0\xbb\xf2\xb0\xd5~~~                      \x1b[m 06/23 12:18\r"),
			CommentID:  "FFqbY21HsAA:jXWE4sLfiC2rGLr1FRdUQQ",
			CreateTime: 1466655480000000000,
			SortTime:   1466655480000000000,
		},
		{ // 2
			BBoardID:  "test",
			ArticleID: "test15",
			TheType:   ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:     bbs.UUserID("YO8BO10"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "神！！！跪！！！",
						Big5:   []byte("\xaf\xab\xa1I\xa1I\xa1I\xb8\xf7\xa1I\xa1I\xa1I                                 "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xaf\xab\xa1I\xa1I\xa1I\xb8\xf7\xa1I\xa1I\xa1I                                 "),
					},
				},
			},
			MD5:        "YSEKMIj--mQbv17h2JAXTA",
			TheDate:    "06/23 12:18",
			DBCS:       []byte("\x1b[1;37m\xb1\xc0 \x1b[33mYO8BO10     \x1b[m\x1b[33m: \xaf\xab\xa1I\xa1I\xa1I\xb8\xf7\xa1I\xa1I\xa1I                                 \x1b[m 06/23 12:18\r"),
			CommentID:  "FFqbY21W8kA:YSEKMIj--mQbv17h2JAXTA",
			CreateTime: 1466655480000000000,
			SortTime:   1466655480001000000,
		},
		{ // 3
			BBoardID:  "test",
			ArticleID: "test15",
			TheType:   ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:     bbs.UUserID("Aggro"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "得分王同時有防守王是啥概念……",
						Big5:   []byte("\xb1o\xa4\xc0\xa4\xfd\xa6P\xae\xc9\xa6\xb3\xa8\xbe\xa6u\xa4\xfd\xacO\xd4\xa3\xb7\xa7\xa9\xc0\xa1K\xa1K                   "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb1o\xa4\xc0\xa4\xfd\xa6P\xae\xc9\xa6\xb3\xa8\xbe\xa6u\xa4\xfd\xacO\xd4\xa3\xb7\xa7\xa9\xc0\xa1K\xa1K                   "),
					},
				},
			},
			MD5:        "XbsC1DBGI67-hVLniVcQVA",
			TheDate:    "06/23 12:20",
			DBCS:       []byte("\x1b[1;37m\xb1\xc0 \x1b[33mAggro       \x1b[m\x1b[33m: \xb1o\xa4\xc0\xa4\xfd\xa6P\xae\xc9\xa6\xb3\xa8\xbe\xa6u\xa4\xfd\xacO\xd4\xa3\xb7\xa7\xa9\xc0\xa1K\xa1K                   \x1b[m 06/23 12:20\r"),
			CommentID:  "FFqbf13WYAA:XbsC1DBGI67-hVLniVcQVA",
			CreateTime: 1466655600000000000,
			SortTime:   1466655600000000000,
		},
		{ // 4
			BBoardID:  "test",
			ArticleID: "test15",
			TheType:   ptttype.COMMENT_TYPE_BOO,
			Owner:     bbs.UUserID("lycium"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "bug",
						Big5:   []byte("bug                                              "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("bug                                              "),
					},
				},
			},
			MD5:        "JiR9wAMpshqmSUZigf4uug",
			TheDate:    "06/23 12:21",
			DBCS:       []byte("\x1b[1;31m\xbcN \x1b[33mlycium      \x1b[m\x1b[33m: bug                                              \x1b[m 06/23 12:21\r"),
			CommentID:  "FFqbjVYduAA:JiR9wAMpshqmSUZigf4uug",
			CreateTime: 1466655660000000000,
			SortTime:   1466655660000000000,
		},
		{ // 5
			BBoardID:  "test",
			ArticleID: "test15",
			TheType:   ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:     bbs.UUserID("XaviYang"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "矮哦...威猛",
						Big5:   []byte("\xb8G\xae@...\xab\xc2\xb2r                                      "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb8G\xae@...\xab\xc2\xb2r                                      "),
					},
				},
			},
			MD5:        "-Zei7Rg7TeLUyWcUIDKowQ",
			TheDate:    "06/23 12:21",
			DBCS:       []byte("\x1b[1;37m\xb1\xc0 \x1b[33mXaviYang    \x1b[m\x1b[33m: \xb8G\xae@...\xab\xc2\xb2r                                      \x1b[m 06/23 12:21\r"),
			CommentID:  "FFqbjVYs-kA:-Zei7Rg7TeLUyWcUIDKowQ",
			CreateTime: 1466655660000000000,
			SortTime:   1466655660001000000,
		},
		{ // 6
			BBoardID:  "test",
			ArticleID: "test15",
			TheType:   ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:     bbs.UUserID("katanakiller"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "單節13罰 鬍子也甘拜下風",
						Big5:   []byte("\xb3\xe6\xb8`13\xbb@ \xc4G\xa4l\xa4]\xa5\xcc\xab\xf4\xa4U\xad\xb7                          "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xb3\xe6\xb8`13\xbb@ \xc4G\xa4l\xa4]\xa5\xcc\xab\xf4\xa4U\xad\xb7                          "),
					},
				},
			},
			MD5:        "xYcELntlC-hefUK7kfp-dA",
			TheDate:    "06/23 12:21",
			DBCS:       []byte("\x1b[1;37m\xb1\xc0 \x1b[33mkatanakiller\x1b[m\x1b[33m: \xb3\xe6\xb8`13\xbb@ \xc4G\xa4l\xa4]\xa5\xcc\xab\xf4\xa4U\xad\xb7                          \x1b[m 06/23 12:21\r"),
			CommentID:  "FFqbjVY8PIA:xYcELntlC-hefUK7kfp-dA",
			CreateTime: 1466655660000000000,
			SortTime:   1466655660002000000,
		},
		{ // 7
			BBoardID:  "test",
			ArticleID: "test15",
			TheType:   ptttype.COMMENT_TYPE_COMMENT,
			Owner:     bbs.UUserID("lycium"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "外星人來地球虐菜不值得一提",
						Big5:   []byte("\xa5~\xacP\xa4H\xa8\xd3\xa6a\xb2y\xadh\xb5\xe6\xa4\xa3\xad\xc8\xb1o\xa4@\xb4\xa3                       "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xa5~\xacP\xa4H\xa8\xd3\xa6a\xb2y\xadh\xb5\xe6\xa4\xa3\xad\xc8\xb1o\xa4@\xb4\xa3                       "),
					},
				},
			},
			MD5:        "G03hZoEVJ0R8PZlEcKKxGg",
			TheDate:    "06/23 12:21",
			DBCS:       []byte("\x1b[1;31m\xa1\xf7 \x1b[33mlycium      \x1b[m\x1b[33m: \xa5~\xacP\xa4H\xa8\xd3\xa6a\xb2y\xadh\xb5\xe6\xa4\xa3\xad\xc8\xb1o\xa4@\xb4\xa3                       \x1b[m 06/23 12:21\r"),
			CommentID:  "FFqbjVZLfsA:G03hZoEVJ0R8PZlEcKKxGg",
			CreateTime: 1466655660000000000,
			SortTime:   1466655660003000000,
		},
		{ // 8
			BBoardID:  "test",
			ArticleID: "test15",
			TheType:   ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:     bbs.UUserID("yankees733"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "說MJ要走下神壇的人呢?",
						Big5:   []byte("\xbb\xa1MJ\xadn\xa8\xab\xa4U\xaf\xab\xbe\xc2\xaa\xba\xa4H\xa9O?                            "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xbb\xa1MJ\xadn\xa8\xab\xa4U\xaf\xab\xbe\xc2\xaa\xba\xa4H\xa9O?                            "),
					},
				},
			},
			MD5:        "wHOypI5hcuHS4wVztWmYyg",
			TheDate:    "06/23 12:23",
			DBCS:       []byte("\x1b[1;37m\xb1\xc0 \x1b[33myankees733  \x1b[m\x1b[33m: \xbb\xa1MJ\xadn\xa8\xab\xa4U\xaf\xab\xbe\xc2\xaa\xba\xa4H\xa9O?                            \x1b[m 06/23 12:23\r"),
			CommentID:  "FFqbqUasaAA:wHOypI5hcuHS4wVztWmYyg",
			CreateTime: 1466655780000000000,
			SortTime:   1466655780000000000,
		},
		{ // 9
			BBoardID:  "test",
			ArticleID: "test15",
			TheType:   ptttype.COMMENT_TYPE_RECOMMEND,
			Owner:     bbs.UUserID("juniorpenny"),
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "鞋子出到30代了啦",
						Big5:   []byte("\xbec\xa4l\xa5X\xa8\xec30\xa5N\xa4F\xb0\xd5                                 "),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
						DBCS:   []byte("\xbec\xa4l\xa5X\xa8\xec30\xa5N\xa4F\xb0\xd5                                 "),
					},
				},
			},
			MD5:        "DuNQJilqP07PXmhfjw3Dfw",
			TheDate:    "06/23 12:24",
			DBCS:       []byte("\x1b[1;37m\xb1\xc0 \x1b[33mjuniorpenny \x1b[m\x1b[33m: \xbec\xa4l\xa5X\xa8\xec30\xa5N\xa4F\xb0\xd5                                 \x1b[m 06/23 12:24\r"),
			CommentID:  "FFqbtz7zwAA:DuNQJilqP07PXmhfjw3Dfw",
			CreateTime: 1466655840000000000,
			SortTime:   1466655840000000000,
		},
	}
}
