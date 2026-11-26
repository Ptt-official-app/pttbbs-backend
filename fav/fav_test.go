package fav

import (
	"bytes"
	"encoding/binary"
	"io"
	"os"
	"reflect"
	"sync"
	"testing"

	pttbbsfav "github.com/Ptt-official-app/go-pttbbs/ptt/fav"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestFavRaw_AddBoard(t *testing.T) {
	setupTest()
	defer teardownTest()

	f0, _ := NewFav(nil, nil, 0)
	logrus.Infof("TestAddBoard: f0: %v", f0)
	ft0 := &FavType{0, pttbbsfav.FAVT_BOARD, 1, &FavBoard{1, 0, 0}}

	f1 := &Fav{
		FavNum:  1,
		NBoards: 1,
		Favh: []*FavType{
			{0, pttbbsfav.FAVT_BOARD, 1, &FavBoard{Bid: 1}},
		},
	}
	f1.Root = f1

	f1_1 := &Fav{
		FavNum:  1,
		NBoards: 1,
		Favh: []*FavType{
			{0, pttbbsfav.FAVT_BOARD, 1, &FavBoard{Bid: 1}},
		},
	}
	f1_1.Root = f1_1
	ft1 := &FavType{1, pttbbsfav.FAVT_BOARD, 1, &FavBoard{2, 0, 0}}

	f1_2 := &Fav{
		FavNum:  1,
		NBoards: 1,
		Favh: []*FavType{
			{0, pttbbsfav.FAVT_BOARD, 1, &FavBoard{Bid: 1}},
		},
	}
	f1_2.Root = f1_2

	f2 := &Fav{
		FavNum:  2,
		NBoards: 2,
		Favh: []*FavType{
			{0, pttbbsfav.FAVT_BOARD, 1, &FavBoard{Bid: 1}},
			{1, pttbbsfav.FAVT_BOARD, 1, &FavBoard{Bid: 2}},
		},
	}
	f2.Root = f1_2

	f3, _ := NewFav(nil, nil, 0)
	_, _, _ = f3.AddBoard(1)
	_, favType3, _ := f3.AddFolder("")
	favFolder3 := favType3.Fp.(*FavFolder)

	ft3 := &FavType{2, pttbbsfav.FAVT_BOARD, 1, &FavBoard{1, 0, 0}}

	sub3 := &Fav{
		NBoards: 1,
		FavNum:  1,
		Depth:   1,
		Favh: []*FavType{
			{2, pttbbsfav.FAVT_BOARD, 1, &FavBoard{Bid: 1}},
		},
	}

	expected3 := &Fav{
		FavNum:   3,
		NBoards:  1,
		NFolders: 1,
		FolderID: 1,
		Favh: []*FavType{
			{0, pttbbsfav.FAVT_BOARD, 1, &FavBoard{Bid: 1}},
			{1, pttbbsfav.FAVT_FOLDER, 1, &FavFolder{Fid: 1, ThisFolder: sub3}},
		},
	}

	filename3 := "./testcase/home1/t/testUser3/.fav"

	theBytes3, _ := os.ReadFile(filename3)

	buf3 := bytes.NewReader(theBytes3)

	filename4 := "./testcase/home1/t/testUser4/.fav"

	theBytes4, _ := os.ReadFile(filename4)

	buf4 := bytes.NewReader(theBytes4)

	// version
	version := int16(0)
	_ = binary.Read(buf3, binary.LittleEndian, &version)

	_ = binary.Read(buf4, binary.LittleEndian, &version)

	fav3, _ := ReadFavrec(buf3, nil, nil, 0)

	fav3Ft9 := &FavType{9, pttbbsfav.FAVT_BOARD, 1, &FavBoard{21, 0, 0}}

	fav3Ft10 := &FavType{10, pttbbsfav.FAVT_BOARD, 1, &FavBoard{23, 0, 0}}

	fav3.AddBoard(21)

	type args struct {
		bid ptttype.Bid
	}
	tests := []struct {
		name            string
		f               *Fav
		args            args
		expectedFavType *FavType
		cmp             *Fav
		expected        *Fav
		expectedIdx     int
		wantErr         bool
	}{
		// TODO: Add test cases.
		{
			name:            "add bid = 1",
			f:               f0,
			args:            args{bid: 1},
			expectedFavType: ft0,
			cmp:             f0,
			expected:        f1,
		},
		{
			name:            "add bid = 1 (again)",
			f:               f1_1,
			args:            args{bid: 1},
			expectedFavType: ft0,
			cmp:             f1_1,
			expected:        f1_1,
		},
		{
			name:            "add bid = 2",
			f:               f1_2,
			args:            args{bid: 2},
			expectedFavType: ft1,
			cmp:             f1_2,
			expected:        f2,
			expectedIdx:     1,
		},
		{
			name:            "add bid = 1 (sub-folder)",
			f:               favFolder3.ThisFolder,
			args:            args{bid: 1},
			expectedFavType: ft3,
			cmp:             f3,
			expected:        expected3,
		},

		{
			name:            "add bid = 23 (after setting up sub-folder)",
			f:               fav3,
			args:            args{bid: 23},
			expectedFavType: fav3Ft10,
			cmp:             fav3,
			expected:        testFav5,
			expectedIdx:     10,
		},
		{
			name:            "add bid = 23 (after setting up sub-folder) (again)",
			f:               fav3,
			args:            args{bid: 23},
			expectedFavType: fav3Ft10,
			cmp:             fav3,
			expected:        testFav5,
			expectedIdx:     10,
		},
		{
			name:            "add bid = 21 (after setting up sub-folder) (again)",
			f:               fav3,
			args:            args{bid: 21},
			expectedFavType: fav3Ft9,
			cmp:             fav3,
			expected:        testFav5,
			expectedIdx:     9,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			f := tt.f
			gotIdx, gotFavType, err := f.AddBoard(tt.args.bid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fav.AddBoard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			tt.cmp.SetFavTypeFavIdx(0)

			if !reflect.DeepEqual(gotFavType, tt.expectedFavType) {
				t.Errorf("Fav.AddBoard() = %v, want %v", gotFavType, tt.expectedFavType)
			}

			if gotIdx != tt.expectedIdx {
				t.Errorf("idx: got: %v expected: %v", gotIdx, tt.expectedIdx)
			}

			tt.cmp.CleanParentAndRoot()
			tt.expected.CleanParentAndRoot()

			testutil.TDeepEqual(t, "expected", tt.cmp, tt.expected)
		})
		wg.Wait()
	}
}

func TestFavRaw_AddLine(t *testing.T) {
	setupTest()
	defer teardownTest()

	f0, _ := NewFav(nil, nil, 0)
	ft0 := &FavType{0, pttbbsfav.FAVT_LINE, 1, &FavLine{1}}

	expected0 := &Fav{
		FavNum: 1,
		NLines: 1,
		LineID: 1,
		Favh: []*FavType{
			{0, pttbbsfav.FAVT_LINE, 1, &FavLine{1}},
		},
	}

	f1 := &Fav{
		FavNum: 1,
		NLines: 1,
		LineID: 1,
		Favh: []*FavType{
			{0, pttbbsfav.FAVT_LINE, 1, &FavLine{1}},
		},
	}
	f1.Root = f1

	ft1 := &FavType{1, pttbbsfav.FAVT_LINE, 1, &FavLine{2}}

	expected1 := &Fav{
		Parent: f1,
		FavNum: 2,
		NLines: 2,
		LineID: 2,
		Favh: []*FavType{
			{0, pttbbsfav.FAVT_LINE, 1, &FavLine{1}},
			{1, pttbbsfav.FAVT_LINE, 1, &FavLine{2}},
		},
	}

	tests := []struct {
		name            string
		f               *Fav
		expectedFavType *FavType
		expected        *Fav
		wantErr         bool
	}{
		// TODO: Add test cases.
		{
			f:               f0,
			expectedFavType: ft0,
			expected:        expected0,
		},
		{
			name:            "add 2nd line",
			f:               f1,
			expectedFavType: ft1,
			expected:        expected1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.f
			_, gotFavType, err := f.AddLine()
			if (err != nil) != tt.wantErr {
				t.Errorf("FavRaw.AddLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			f.SetFavTypeFavIdx(0)

			f.CleanParentAndRoot()
			tt.expected.CleanParentAndRoot()

			assert.Equal(t, gotFavType, tt.expectedFavType)

			testutil.TDeepEqual(t, "expected", f, tt.expected)
		})
	}
}

func TestFavRaw_AddFolder(t *testing.T) {
	setupTest()
	defer teardownTest()

	f0, _ := NewFav(nil, nil, 0)

	folder0, _ := NewFav(nil, nil, 1)
	ft0 := &FavType{0, pttbbsfav.FAVT_FOLDER, 1, &FavFolder{Fid: 1, ThisFolder: folder0}}
	expected0 := &Fav{
		FavNum:   1,
		NFolders: 1,
		FolderID: 1,
		Favh: []*FavType{
			{0, pttbbsfav.FAVT_FOLDER, 1, &FavFolder{Fid: 1, ThisFolder: folder0}},
		},
	}

	f1, _ := NewFav(nil, nil, 0)
	_, _, _ = f1.AddFolder("")

	folder1, _ := NewFav(nil, nil, 1)
	ft1 := &FavType{1, pttbbsfav.FAVT_FOLDER, 1, &FavFolder{Fid: 2, ThisFolder: folder1}}
	expected1 := &Fav{
		FavNum:   2,
		NFolders: 2,
		FolderID: 2,
		Favh: []*FavType{
			{0, pttbbsfav.FAVT_FOLDER, 1, &FavFolder{Fid: 1, ThisFolder: folder1}},
			{1, pttbbsfav.FAVT_FOLDER, 1, &FavFolder{Fid: 2, ThisFolder: folder1}},
		},
	}

	tests := []struct {
		name            string
		f               *Fav
		expectedFavType *FavType
		expected        *Fav
		wantErr         bool
	}{
		// TODO: Add test cases.
		{
			f:               f0,
			expectedFavType: ft0,
			expected:        expected0,
		},
		{
			f:               f1,
			expectedFavType: ft1,
			expected:        expected1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.f
			_, gotFavType, err := f.AddFolder("")
			if (err != nil) != tt.wantErr {
				t.Errorf("FavRaw.AddFolder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			f.SetFavTypeFavIdx(0)
			f.CleanParentAndRoot()
			tt.expected.CleanParentAndRoot()

			gotFolder := gotFavType.Fp.(*FavFolder)
			gotFolder.ThisFolder.CleanParentAndRoot()

			expectedFolder := tt.expectedFavType.Fp.(*FavFolder)
			expectedFolder.ThisFolder.CleanParentAndRoot()

			assert.Equal(t, gotFavType, tt.expectedFavType)

			testutil.TDeepEqual(t, "expected", f, tt.expected)
		})
	}
}

func TestReadFavrec(t *testing.T) {
	setupTest()
	defer teardownTest()

	filename0 := "./testcase/home1/t/testUser/.fav"

	theBytes0, _ := os.ReadFile(filename0)

	buf0 := bytes.NewReader(theBytes0)

	filename1 := "./testcase/home1/t/testUser2/.fav"

	theBytes1, _ := os.ReadFile(filename1)

	buf1 := bytes.NewReader(theBytes1)

	filename3 := "./testcase/home1/t/testUser3/.fav"

	theBytes3, _ := os.ReadFile(filename3)

	buf3 := bytes.NewReader(theBytes3)

	filename4 := "./testcase/home1/t/testUser4/.fav"

	theBytes4, _ := os.ReadFile(filename4)

	buf4 := bytes.NewReader(theBytes4)

	// version
	version := int16(0)
	_ = binary.Read(buf0, binary.LittleEndian, &version)

	_ = binary.Read(buf1, binary.LittleEndian, &version)

	_ = binary.Read(buf3, binary.LittleEndian, &version)

	_ = binary.Read(buf4, binary.LittleEndian, &version)

	type args struct {
		file io.ReadSeeker
	}
	tests := []struct {
		name           string
		args           args
		expectedFavrec *Fav
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			args:           args{file: buf0},
			expectedFavrec: testFav0,
		},
		{
			args:           args{file: buf1},
			expectedFavrec: testFav1,
		},
		{
			args:           args{file: buf3},
			expectedFavrec: testFav3,
		},
		{
			args:           args{file: buf4},
			expectedFavrec: testFav4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFavrec, err := ReadFavrec(tt.args.file, nil, nil, 0)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFavrec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotFavrec.SetFavTypeFavIdx(0)
			gotFavrec.CleanParentAndRoot()

			testutil.TDeepEqual(t, "got", gotFavrec, tt.expectedFavrec)
		})
	}
}

func TestFav_WriteFavrec(t *testing.T) {
	setupTest()
	defer teardownTest()

	filename := "./testcase/home1/t/testUser/.fav"

	theBytes0, _ := os.ReadFile(filename)

	filename1 := "./testcase/home1/t/testUser2/.fav"

	theBytes1, _ := os.ReadFile(filename1)

	tests := []struct {
		name         string
		favrec       *Fav
		expectedFile []byte
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			favrec:       testFav0,
			expectedFile: theBytes0,
		},
		{
			favrec:       testFav1,
			expectedFile: theBytes1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.favrec
			file := &bytes.Buffer{}
			if err := f.WriteFavrec(file); (err != nil) != tt.wantErr {
				t.Errorf("Fav.WriteFavrec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			gotFile := file.Bytes()

			testutil.TDeepEqual(t, "got", gotFile, tt.expectedFile)
		})
	}
}

func TestFav_LocateFav(t *testing.T) {
	setupTest()
	defer teardownTest()

	filename0 := "./testcase/home1/t/testUser/.fav"

	theBytes0, _ := os.ReadFile(filename0)

	buf0 := bytes.NewReader(theBytes0)

	// version
	version := int16(0)
	_ = binary.Read(buf0, binary.LittleEndian, &version)

	// fav
	fav0, _ := ReadFavrec(buf0, nil, nil, 0)

	subFav0 := fav0.Favh[2].CastFolder().ThisFolder

	logrus.Infof("TestFav_LocateFav: fav0: %v", fav0)

	type args struct {
		levelIdxList []string
	}
	tests := []struct {
		name       string
		f          *Fav
		args       args
		wantNewFav *Fav
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			f:          fav0,
			args:       args{levelIdxList: []string{""}},
			wantNewFav: fav0,
		},
		{
			f:       fav0,
			args:    args{levelIdxList: []string{"", "1"}},
			wantErr: true,
		},
		{
			f:          fav0,
			args:       args{levelIdxList: []string{"", "2"}},
			wantNewFav: subFav0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.f
			gotNewFav, err := f.LocateFav(tt.args.levelIdxList)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fav.LocateFav() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			logrus.Infof("TestFav_LocateFav: to TDeepEqual")
			if !reflect.DeepEqual(gotNewFav, tt.wantNewFav) {
				t.Errorf("Fav.LocateFav: got: %v want: %v", gotNewFav, tt.wantNewFav)
				return
			}
		})
	}
}

func TestFav_DeleteIdx(t *testing.T) {
	setupTest()
	defer teardownTest()

	filename0 := "./testcase/home1/t/testUser/.fav"

	theBytes0, _ := os.ReadFile(filename0)

	// fav0
	buf0 := bytes.NewReader(theBytes0)
	version := int16(0)
	_ = binary.Read(buf0, binary.LittleEndian, &version)
	fav0, _ := ReadFavrec(buf0, nil, nil, 0)

	logrus.Infof("fav0.Favh: (%v/%v)", fav0.Favh, len(fav0.Favh))
	for idx, each := range fav0.Favh {
		logrus.Infof("(%v/%v) %v", idx, len(fav0.Favh), each)
	}

	buf0 = bytes.NewReader(theBytes0)
	_ = binary.Read(buf0, binary.LittleEndian, &version)
	ret0, err := ReadFavrec(buf0, nil, nil, 0)

	logrus.Infof("ret0.Favh (after read): e: %v ret0: (%v/%v)", err, ret0, len(ret0.Favh))

	ret0.Favh = ret0.Favh[1:]
	ret0.FavNum--
	ret0.NBoards--
	logrus.Infof("ret0.Favh (after 1:): (%v/%v)", ret0.Favh, len(ret0.Favh))

	// fav1
	buf1 := bytes.NewReader(theBytes0)
	_ = binary.Read(buf1, binary.LittleEndian, &version)
	fav1, _ := ReadFavrec(buf1, nil, nil, 0)

	fav1Folder := fav1.Favh[2].CastFolder()
	fav1Child := fav1Folder.ThisFolder

	// ret1
	buf1 = bytes.NewReader(theBytes0)
	_ = binary.Read(buf1, binary.LittleEndian, &version)
	ret1, _ := ReadFavrec(buf1, nil, nil, 0)
	ret1.FavNum -= fav1Child.FavNum

	ret1Folder := ret1.Favh[2].CastFolder()
	ret1Folder.ThisFolder.FavNum--
	ret1Folder.ThisFolder.NBoards--
	ret1Folder.ThisFolder.Favh = []*FavType{}

	// fav2
	buf2 := bytes.NewReader(theBytes0)
	_ = binary.Read(buf2, binary.LittleEndian, &version)
	fav2, _ := ReadFavrec(buf2, nil, nil, 0)

	// ret1
	buf2 = bytes.NewReader(theBytes0)
	_ = binary.Read(buf2, binary.LittleEndian, &version)
	ret2, _ := ReadFavrec(buf2, nil, nil, 0)

	logrus.Infof("ret2.Favh (after read): %v", ret2.Favh)

	ret2Folder := ret2.Favh[2].CastFolder()
	ret2.FavNum -= 1 + ret2Folder.ThisFolder.FavNum
	ret2.NFolders--
	ret2.Favh = append(ret2.Favh[:2], ret2.Favh[3:]...)

	logrus.Infof("ret2 (after remove folder): %v", ret2.Favh)

	type args struct {
		idx int
	}
	tests := []struct {
		name    string
		f       *Fav
		rootF   *Fav
		args    args
		wantF   *Fav
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:  "remove a board",
			f:     fav0,
			rootF: fav0,
			args:  args{idx: 0},
			wantF: ret0,
		},
		{
			name:  "remove board in a folder",
			f:     fav1Child,
			rootF: fav1,
			args:  args{idx: 0},
			wantF: ret1,
		},
		{
			name:  "remove folder",
			f:     fav2,
			rootF: fav2,
			args:  args{idx: 2},
			wantF: ret2,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			f := tt.f
			logrus.Infof("Fav.DeleteIdx: to DeleteIdx: f: (%v/%v)", f.Favh, len(f.Favh))
			if err := f.DeleteIdx(tt.args.idx); (err != nil) != tt.wantErr {
				t.Errorf("Fav.DeleteIdx() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.Equal(t, tt.wantF, tt.rootF, "rootF")
		})
		wg.Wait()
	}
}

func TestFavRaw_AddFolder_maxFolder(t *testing.T) {
	setupTest()
	defer teardownTest()

	f, _ := NewFav(nil, nil, 0)
	// NFolders is the limit to check, not NLines.
	f.NFolders = MAX_FOLDER
	_, _, err := f.AddFolder("")
	assert.ErrorIs(t, err, ErrTooManyFolders)

	f, _ = NewFav(nil, nil, 0)
	f.NLines = MAX_FOLDER
	_, _, err = f.AddFolder("")
	assert.NoError(t, err)
}

func TestReadFavrec_negativeCount(t *testing.T) {
	setupTest()
	defer teardownTest()

	tests := []struct {
		name  string
		bytes []byte
	}{
		{name: "negative boards", bytes: []byte{0xff, 0xff, 0, 0}},
		{name: "negative lines", bytes: []byte{0, 0, 0xff, 0}},
		{name: "negative folders", bytes: []byte{0, 0, 0, 0xff}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				_, err := ReadFavrec(bytes.NewReader(tt.bytes), nil, nil, 0)
				assert.ErrorIs(t, err, ErrInvalidFavRecord)
			})
		})
	}
}

func TestFav_LocateFav_outOfRange(t *testing.T) {
	setupTest()
	defer teardownTest()

	f, _ := NewFav(nil, nil, 0)
	_, _, _ = f.AddFolder("")

	for _, idx := range []string{"99", "-1", "1"} {
		t.Run(idx, func(t *testing.T) {
			assert.NotPanics(t, func() {
				_, err := f.LocateFav([]string{"", idx})
				assert.ErrorIs(t, err, ErrInvalidLevelIdx)
			})
		})
	}

	got, err := f.LocateFav([]string{"", "0"})
	assert.NoError(t, err)
	assert.NotNil(t, got)
}

func TestFav_DeleteIdx_negative(t *testing.T) {
	setupTest()
	defer teardownTest()

	f, _ := NewFav(nil, nil, 0)
	_, _, _ = f.AddFolder("")

	assert.NotPanics(t, func() {
		assert.ErrorIs(t, f.DeleteIdx(-1), ErrInvalidFavRecord)
	})
}
