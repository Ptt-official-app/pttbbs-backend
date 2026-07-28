package dbcs

import (
	"reflect"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/sirupsen/logrus"
)

func Test_dbcsToBig5(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		dbcs []byte
	}
	tests := []struct {
		name         string
		args         args
		expectedBig5 [][]*types.Rune
	}{
		// TODO: Add test cases.
		{
			name:         "0_" + testFilename0 + "_content",
			args:         args{dbcs: testContent0},
			expectedBig5: testContent0Big5,
		},
		{
			name:         "1_" + testFilename1 + "_content",
			args:         args{dbcs: testContent1},
			expectedBig5: testContent1Big5,
		},
		{
			name:         "2_" + testFilename2 + "_content",
			args:         args{dbcs: testContent2},
			expectedBig5: testContent2Big5,
		},
		{
			name:         "3_" + testFilename3 + "_content",
			args:         args{dbcs: testContent3},
			expectedBig5: testContent3Big5,
		},
		{
			name:         "4_" + testFilename4 + "_content",
			args:         args{dbcs: testContent4},
			expectedBig5: testContent4Big5,
		},
		{
			name:         "5_" + testFilename5 + "_content",
			args:         args{dbcs: testContent5},
			expectedBig5: testContent5Big5,
		},
		{
			name:         "6_" + testFilename6 + "_content",
			args:         args{dbcs: testContent6},
			expectedBig5: testContent6Big5,
		},
		{
			name:         "7_" + testFilename7 + "_content",
			args:         args{dbcs: testContent7},
			expectedBig5: testContent7Big5,
		},
		{
			name:         "8_" + testFilename8 + "_content",
			args:         args{dbcs: testContent8},
			expectedBig5: testContent8Big5,
		},
		{
			name:         "9_" + testFilename9 + "_content",
			args:         args{dbcs: testContent9},
			expectedBig5: testContent9Big5,
		},
		{
			name:         "10_" + testFilename10 + "_content",
			args:         args{dbcs: testContent10},
			expectedBig5: testContent10Big5,
		},
		{
			name:         "11_" + testFilename11 + "_content",
			args:         args{dbcs: testContent11},
			expectedBig5: testContent11Big5,
		},
		{
			name:         "12_" + testFilename12 + "_content",
			args:         args{dbcs: testContent12},
			expectedBig5: testContent12Big5,
		},
		{
			name:         "13_" + testFilename13 + "_content",
			args:         args{dbcs: testContent13},
			expectedBig5: testContent13Big5,
		},
		{
			name:         "14_" + testFilename14 + "_content",
			args:         args{dbcs: testContent14},
			expectedBig5: testContent14Big5,
		},
		{
			name:         "15_" + testFilename15 + "_content",
			args:         args{dbcs: testContent15},
			expectedBig5: testContent15Big5,
		},
		{
			name:         "16_" + testFilename16 + "_content",
			args:         args{dbcs: testContent16},
			expectedBig5: testContent16Big5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBig5 := dbcsToBig5(tt.args.dbcs)

			if len(gotBig5) != len(tt.expectedBig5) {
				t.Errorf("len(Big5): %v expected: %v", len(gotBig5), len(tt.expectedBig5))
			}

			for idx, each := range gotBig5 {
				if idx >= len(tt.expectedBig5) {
					t.Errorf("Big5: (%v/%v) %v", idx, len(gotBig5), each)
					continue
				}

				if len(each) != len(tt.expectedBig5[idx]) {
					t.Errorf("Big5: (%v/%v): len: %v expected: %v", idx, len(gotBig5), len(each), len(tt.expectedBig5[idx]))
				}
				for idx2, each2 := range each {
					if idx2 >= len(tt.expectedBig5[idx]) {
						t.Errorf("Big5: (%v/%v/%v/%v): %v", idx, len(gotBig5), idx2, len(each), each2)
						continue
					}
					if !reflect.DeepEqual(each2, tt.expectedBig5[idx][idx2]) {
						t.Errorf("Big5: (%v/%v/%v/%v) %v expected: %v", idx, len(gotBig5), idx2, len(each), each2, tt.expectedBig5[idx][idx2])
					}
				}
			}
		})
	}
}

func Test_big5ToUtf8(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		a [][]*types.Rune
	}
	tests := []struct {
		name     string
		args     args
		expected [][]*types.Rune
	}{
		// TODO: Add test cases.
		{
			name:     "0_" + testFilename0 + "_content",
			args:     args{a: testContent0Big5},
			expected: testContent0Utf8,
		},
		{
			name:     "1_" + testFilename1 + "_content",
			args:     args{a: testContent1Big5},
			expected: testContent1Utf8,
		},
		{
			name:     "2_" + testFilename2 + "_content",
			args:     args{a: testContent2Big5},
			expected: testContent2Utf8,
		},
		{
			name:     "3_" + testFilename3 + "_content",
			args:     args{a: testContent3Big5},
			expected: testContent3Utf8,
		},
		{
			name:     "4_" + testFilename4 + "_content",
			args:     args{a: testContent4Big5},
			expected: testContent4Utf8,
		},
		{
			name:     "5_" + testFilename5 + "_content",
			args:     args{a: testContent5Big5},
			expected: testContent5Utf8,
		},
		{
			name:     "6_" + testFilename6 + "_content",
			args:     args{a: testContent6Big5},
			expected: testContent6Utf8,
		},
		{
			name:     "7_" + testFilename7 + "_content",
			args:     args{a: testContent7Big5},
			expected: testContent7Utf8,
		},
		{
			name:     "8_" + testFilename8 + "_content",
			args:     args{a: testContent8Big5},
			expected: testContent8Utf8,
		},
		{
			name:     "9_" + testFilename9 + "_content",
			args:     args{a: testContent9Big5},
			expected: testContent9Utf8,
		},
		{
			name:     "10_" + testFilename10 + "_content",
			args:     args{a: testContent10Big5},
			expected: testContent10Utf8,
		},
		{
			name:     "11_" + testFilename11 + "_content",
			args:     args{a: testContent11Big5},
			expected: testContent11Utf8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := big5ToUtf8(tt.args.a)

			if len(got) != len(tt.expected) {
				t.Errorf("len(utf8): %v expected: %v", len(got), len(tt.expected))
			}

			for idx, each := range got {
				if idx >= len(tt.expected) {
					t.Errorf("utf8: (%v/%v) %v", idx, len(got), each)
					continue
				}

				if len(each) != len(tt.expected[idx]) {
					t.Errorf("utf8: (%v/%v): len: %v expected: %v", idx, len(got), len(each), len(tt.expected[idx]))
				}
				for idx2, each2 := range each {
					if idx2 >= len(tt.expected[idx]) {
						t.Errorf("utf8: (%v/%v/%v/%v): %v", idx, len(got), idx2, len(each), each2)
						continue
					}
					if !reflect.DeepEqual(each2, tt.expected[idx][idx2]) {
						t.Errorf("utf8: (%v/%v/%v/%v) %v expected: %v", idx, len(got), idx2, len(each), each2, tt.expected[idx][idx2])
					}
				}
			}
		})
	}
}

func Test_utf8ToBig5ByRune(t *testing.T) {
	setupTest()
	defer teardownTest()

	rune0 := &types.Rune{
		Utf8:   "測試0",
		Color0: types.DefaultColor,
		Color1: types.DefaultColor,
	}
	color0 := types.DefaultColor
	expected0 := []byte("\xb4\xfa\xb8\xd50")

	rune1 := &types.Rune{
		Utf8: "測試0",
		Color0: types.Color{
			Foreground: types.COLOR_FOREGROUND_YELLOW,
			Background: types.COLOR_BACKGROUND_BLUE,
		},
		Color1: types.Color{
			Foreground: types.COLOR_FOREGROUND_YELLOW,
			Background: types.COLOR_BACKGROUND_BLUE,
		},
	}
	color1 := types.DefaultColor
	expected1 := []byte("\x1b[33;44m\xb4\xfa\xb8\xd50")
	expectedColor1 := rune1.Color1

	rune2 := &types.Rune{
		Utf8: "測試0",
		Color0: types.Color{
			Foreground: types.COLOR_FOREGROUND_YELLOW,
			Background: types.COLOR_BACKGROUND_BLACK,
		},
		Color1: types.Color{
			Foreground: types.COLOR_FOREGROUND_YELLOW,
			Background: types.COLOR_BACKGROUND_BLACK,
		},
	}
	color2 := types.DefaultColor
	expected2 := []byte("\x1b[33m\xb4\xfa\xb8\xd50")
	expectedColor2 := rune2.Color1

	rune3 := &types.Rune{
		Utf8: "測試0",
		Color0: types.Color{
			Foreground: types.COLOR_FOREGROUND_YELLOW,
			Background: types.COLOR_BACKGROUND_BLACK,
		},
		Color1: types.Color{
			Foreground: types.COLOR_FOREGROUND_YELLOW,
			Background: types.COLOR_BACKGROUND_BLUE,
		},
	}
	color3 := types.DefaultColor
	expected3 := []byte("\x1b[33m\xb4\xfa\xb8\xd5\x1b[44m0")
	expectedColor3 := rune3.Color1

	rune4 := &types.Rune{
		Utf8: "",
		Color0: types.Color{
			Foreground: types.COLOR_FOREGROUND_WHITE,
			Background: types.COLOR_BACKGROUND_BLACK,
			IsReset:    true,
		},
		Color1: types.Color{
			Foreground: types.COLOR_FOREGROUND_WHITE,
			Background: types.COLOR_BACKGROUND_BLACK,
			IsReset:    true,
		},
	}
	color4 := types.DefaultColor
	expected4 := []byte("\x1b[m")
	expectedColor4 := rune4.Color1

	rune5 := &types.Rune{
		Utf8: "",
		Color0: types.Color{
			Foreground: types.COLOR_FOREGROUND_YELLOW,
			Background: types.COLOR_BACKGROUND_BLUE,
		},
		Color1: types.Color{
			Foreground: types.COLOR_FOREGROUND_YELLOW,
			Background: types.COLOR_BACKGROUND_GREEN,
		},
	}
	color5 := types.DefaultColor
	expected5 := []byte("\x1b[33;44m\x1b[42m")
	expectedColor5 := rune5.Color1

	type args struct {
		theRune *types.Rune
		color   types.Color
		isAddCR bool
	}
	tests := []struct {
		name             string
		args             args
		expectedTheBig5  []byte
		expectedNewColor types.Color
	}{
		// TODO: Add test cases.
		{
			args:             args{theRune: rune0, color: color0},
			expectedNewColor: types.DefaultColor,
			expectedTheBig5:  expected0,
		},
		{
			args:             args{theRune: rune1, color: color1},
			expectedNewColor: expectedColor1,
			expectedTheBig5:  expected1,
		},
		{
			args:             args{theRune: rune2, color: color2},
			expectedNewColor: expectedColor2,
			expectedTheBig5:  expected2,
		},
		{
			args:             args{theRune: rune3, color: color3},
			expectedNewColor: expectedColor3,
			expectedTheBig5:  expected3,
		},
		{
			args:             args{theRune: rune4, color: color4},
			expectedNewColor: expectedColor4,
			expectedTheBig5:  expected4,
		},
		{
			args:             args{theRune: rune5, color: color5},
			expectedNewColor: expectedColor5,
			expectedTheBig5:  expected5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTheBig5, gotNewColor := utf8ToDBCSByRune(tt.args.theRune, tt.args.color, tt.args.isAddCR)
			logrus.Infof("gotTheBig5: %x", gotTheBig5)
			testutil.TDeepEqual(t, "got", gotTheBig5, tt.expectedTheBig5)

			if !reflect.DeepEqual(gotNewColor, tt.expectedNewColor) {
				t.Errorf("utf8ToBig5ByRune() gotNewColor = %v, want %v", gotNewColor, tt.expectedNewColor)
			}
		})
	}
}

func Test_utf8ToBig5ByLine(t *testing.T) {
	setupTest()
	defer teardownTest()

	rune0 := &types.Rune{
		Utf8:   "測試0",
		Color0: types.DefaultColor,
		Color1: types.DefaultColor,
	}

	rune1 := &types.Rune{
		Utf8: "測試1",
		Color0: types.Color{
			Foreground: types.COLOR_FOREGROUND_YELLOW,
			Background: types.COLOR_BACKGROUND_BLUE,
		},
		Color1: types.Color{
			Foreground: types.COLOR_FOREGROUND_YELLOW,
			Background: types.COLOR_BACKGROUND_BLUE,
		},
	}

	line0 := []*types.Rune{rune0, rune1}
	color0 := types.DefaultColor
	expected0 := []byte("\xb4\xfa\xb8\xd50\x1b[33;44m\xb4\xfa\xb8\xd51\r")
	expectedColor0 := rune1.Color1

	type args struct {
		line  []*types.Rune
		color types.Color
	}
	tests := []struct {
		name             string
		args             args
		expectedLineBig5 []byte
		expectedNewColor types.Color
	}{
		// TODO: Add test cases.
		{
			args:             args{line: line0, color: color0},
			expectedLineBig5: expected0,
			expectedNewColor: expectedColor0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLineBig5, gotNewColor := utf8ToDBCSByLine(tt.args.line, tt.args.color)
			if !reflect.DeepEqual(gotLineBig5, tt.expectedLineBig5) {
				t.Errorf("utf8ToBig5ByLine() gotLineBig5 = %v, want %v", gotLineBig5, tt.expectedLineBig5)
			}
			if !reflect.DeepEqual(gotNewColor, tt.expectedNewColor) {
				t.Errorf("utf8ToBig5ByLine() gotNewColor = %v, want %v", gotNewColor, tt.expectedNewColor)
			}
		})
	}
}

func TestUtf8ToDBCS(t *testing.T) {
	setupTest()
	defer teardownTest()

	rune0 := &types.Rune{
		Utf8:   "測試0",
		Color0: types.DefaultColor,
		Color1: types.DefaultColor,
	}

	rune1 := &types.Rune{
		Utf8: "測試1",
		Color0: types.Color{
			Foreground: types.COLOR_FOREGROUND_YELLOW,
			Background: types.COLOR_BACKGROUND_BLUE,
		},
		Color1: types.Color{
			Foreground: types.COLOR_FOREGROUND_YELLOW,
			Background: types.COLOR_BACKGROUND_BLUE,
		},
	}

	rune2 := &types.Rune{
		Utf8:   "測試2",
		Color0: types.DefaultColor,
		Color1: types.DefaultColor,
	}

	rune3 := &types.Rune{
		Utf8: "測試3",
		Color0: types.Color{
			Foreground: types.COLOR_FOREGROUND_YELLOW,
			Background: types.COLOR_BACKGROUND_BLUE,
		},
		Color1: types.Color{
			Foreground: types.COLOR_FOREGROUND_YELLOW,
			Background: types.COLOR_BACKGROUND_BLUE,
		},
	}

	utf80 := [][]*types.Rune{
		{rune0, rune1},
		{rune2, rune3},
	}
	expected0 := [][]byte{
		[]byte("\xb4\xfa\xb8\xd50\x1b[33;44m\xb4\xfa\xb8\xd51\r"),
		[]byte("\x1b[37;40m\xb4\xfa\xb8\xd52\x1b[33;44m\xb4\xfa\xb8\xd53\r"),
	}

	type args struct {
		utf8 [][]*types.Rune
	}
	tests := []struct {
		name         string
		args         args
		expectedBig5 [][]byte
	}{
		// TODO: Add test cases.
		{
			args:         args{utf8: utf80},
			expectedBig5: expected0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBig5 := Utf8ToDBCS(tt.args.utf8); !reflect.DeepEqual(gotBig5, tt.expectedBig5) {
				t.Errorf("Utf8ToBig5() = %v, want %v", gotBig5, tt.expectedBig5)
			}
		})
	}
}

func Test_dbcsParseColor(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		dbcs  []byte
		want  types.Color
		want2 int
	}{
		// TODO: Add test cases.
		{
			dbcs:  []byte{'\x1b', '[', '0', 'm'},
			want:  types.ResetColor,
			want2: 4,
		},
		{
			dbcs:  []byte{'\x1b', '[', 'm'},
			want:  types.ResetColor,
			want2: 3,
		},
		{
			dbcs:  []byte{'\x1b', '[', ';', '3', '2', 'm'},
			want:  types.Color{Foreground: types.COLOR_FOREGROUND_GREEN, Background: types.COLOR_BACKGROUND_BLACK, IsReset: true},
			want2: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2 := dbcsParseColor(tt.dbcs)
			// TODO: update the condition below to compare got with tt.want.
			testutil.TDeepEqual(t, tt.name, &got, &tt.want)
			if got2 != tt.want2 {
				t.Errorf("got2 = %v, want2 %v", got2, tt.want2)
			}
		})
	}
}
