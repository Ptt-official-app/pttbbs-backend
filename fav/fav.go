package fav

import (
	"encoding/binary"
	"io"
	"strconv"

	pttbbsfav "github.com/Ptt-official-app/go-pttbbs/ptt/fav"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	pttbbstypes "github.com/Ptt-official-app/go-pttbbs/types"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

// Fav
//
// It's with it's own serialize method and does not directly copy by struct.
// We can add MTime in FavRaw.
// The content of FavFolder
type Fav struct {
	FavNum int
	Depth  int

	Root   *Fav
	Parent *Fav

	NBoards  int16         /* number of the boards */
	NLines   int8          /* number of the lines */
	NFolders int8          /* number of the folders */
	LineID   pttbbsfav.Lid /* current max line id */
	FolderID pttbbsfav.Fid /* current max folder id */

	Favh []*FavType
}

func NewFav(parent *Fav, root *Fav, depth int) (f *Fav, err error) {
	if depth >= pttbbsfav.FAV_MAXDEPTH {
		return nil, ErrTooMuchDepth
	}
	f = &Fav{
		Parent: parent,
		Root:   root,
		Depth:  depth,
	}
	if root == nil {
		f.Root = f
	}

	return f, nil
}

func (f *Fav) getDataNumber() int {
	return int(f.NBoards) + int(f.NLines) + int(f.NFolders)
}

func (f *Fav) WriteFavrec(file io.Writer) (err error) {
	if f == nil {
		return nil
	}

	if f.Depth == 0 {
		err = binary.Write(file, binary.LittleEndian, FAV_VERSION)
		if err != nil {
			return err
		}
	}

	err = binary.Write(file, binary.LittleEndian, &f.NBoards)
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, &f.NLines)
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, &f.NFolders)
	if err != nil {
		return err
	}
	total := f.getDataNumber()

	for i := range total {
		ft := f.Favh[i]
		err := binary.Write(file, binary.LittleEndian, &ft.TheType)
		if err != nil {
			return err
		}
		err = binary.Write(file, binary.LittleEndian, &ft.Attr)
		if err != nil {
			return err
		}

		switch ft.TheType {
		case pttbbsfav.FAVT_FOLDER:
			favFolder := ft.CastFolder()
			if favFolder == nil {
				return ErrInvalidFavFolder
			}
			err = binary.Write(file, binary.LittleEndian, &favFolder.Fid)
			if err != nil {
				return err
			}
			titleBytes := types.Utf8ToBig5(favFolder.Title)
			title := &ptttype.BoardTitle_t{}
			copy(title[:], titleBytes)
			err = binary.Write(file, binary.LittleEndian, title)
			if err != nil {
				return err
			}
		case pttbbsfav.FAVT_BOARD:
			favBoard := ft.CastBoard()
			if favBoard == nil {
				return ErrInvalidFavBoard
			}
			err = pttbbstypes.BinWrite(file, favBoard, ft.TheType.GetTypeSize())
			if err != nil {
				return err
			}
		case pttbbsfav.FAVT_LINE:
			favLine := ft.CastLine()
			if favLine == nil {
				return ErrInvalidFavLine
			}
			err = pttbbstypes.BinWrite(file, favLine, ft.TheType.GetTypeSize())
			if err != nil {
				return err
			}
		}
	}

	for i := range total {
		ft := f.Favh[i]
		if ft == nil {
			continue
		}
		if ft.TheType == pttbbsfav.FAVT_FOLDER {
			favFolder := ft.CastFolder()
			if favFolder == nil {
				return ErrInvalidFavFolder
			}
			err := favFolder.ThisFolder.WriteFavrec(file)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// ReadFavrec
//
// We need to:
// 1. read version because calling ReadFavrec
// 2. run root.SetFavTypeFavIdx(0)
//
// because ReadFavrec is a recursive-call,
// and the above 2 happens only in the 1st-call.
func ReadFavrec(file io.ReadSeeker, parent *Fav, root *Fav, depth int) (favrec *Fav, err error) {
	// https://github.com/ptt/pttbbs/blob/master/mbbsd/fav.c
	favrec, err = NewFav(parent, root, depth)
	if err != nil {
		return nil, err
	}
	err = binary.Read(file, binary.LittleEndian, &favrec.NBoards)
	if err != nil {
		return nil, ErrInvalidFavRecord
	}

	err = binary.Read(file, binary.LittleEndian, &favrec.NLines)
	if err != nil {
		return nil, ErrInvalidFavRecord
	}

	err = binary.Read(file, binary.LittleEndian, &favrec.NFolders)
	if err != nil {
		return nil, ErrInvalidFavRecord
	}

	nFavh := favrec.getDataNumber()
	favrec.LineID = 0
	favrec.FolderID = 0
	favrec.Favh = make([]*FavType, nFavh)

	for i := range nFavh {
		ft := &FavType{}
		favrec.IncreaseFavNum()

		favrec.Favh[i] = ft

		err = binary.Read(file, binary.LittleEndian, &ft.TheType)
		if err != nil {
			return nil, ErrInvalidFavType
		}
		if !ft.TheType.IsValidFavType() {
			return nil, ErrInvalidFavType
		}

		err = binary.Read(file, binary.LittleEndian, &ft.Attr)
		if err != nil {
			return nil, ErrInvalidFavType
		}

		switch ft.TheType {
		case pttbbsfav.FAVT_FOLDER:
			favFolder := ft.CastFolder()
			if favFolder == nil {
				return nil, ErrInvalidFavFolder
			}
			err = binary.Read(file, binary.LittleEndian, &favFolder.Fid)
			if err != nil {
				return nil, ErrInvalidFavFolder
			}
			title := &ptttype.BoardTitle_t{}
			err = binary.Read(file, binary.LittleEndian, title)
			if err != nil {
				return nil, ErrInvalidFavFolder
			}
			favFolder.Title = types.Big5ToUtf8(pttbbstypes.CstrToBytes(title[:]))
		case pttbbsfav.FAVT_BOARD:
			favBoard := ft.CastBoard()
			if favBoard == nil {
				return nil, ErrInvalidFavBoard
			}
			err = pttbbstypes.BinRead(file, favBoard, ft.TheType.GetTypeSize())
			if err != nil {
				return nil, ErrInvalidFavBoard
			}
		case pttbbsfav.FAVT_LINE:
			favLine := ft.CastLine()
			if favLine == nil {
				return nil, ErrInvalidFavLine
			}
			err = pttbbstypes.BinRead(file, favLine, ft.TheType.GetTypeSize())
			if err != nil {
				return nil, ErrInvalidFavLine
			}
		}
	}

	for i := range nFavh {
		ft := favrec.Favh[i]
		if ft == nil {
			continue
		}

		switch ft.TheType {
		case pttbbsfav.FAVT_FOLDER:
			favFolder := ft.CastFolder()
			if favFolder == nil {
				return nil, ErrInvalidFavFolder
			}
			p, err := ReadFavrec(file, favrec, favrec.Root, favrec.Depth+1)
			if err != nil {
				return nil, ErrInvalidFavFolder
			}
			favFolder.ThisFolder = p
			favrec.FolderID++
			favFolder.Fid = favrec.FolderID
		case pttbbsfav.FAVT_LINE:
			favLine := ft.CastLine()
			if favLine == nil {
				return nil, ErrInvalidFavLine
			}
			favrec.LineID++
			favLine.Lid = favrec.LineID
		}
	}

	return favrec, nil
}

func (f *Fav) IncreaseFavNum() {
	f.FavNum++
	if f.Parent != nil {
		f.Parent.IncreaseFavNum()
	}
}

func (f *Fav) DecreaseFavNum() {
	f.FavNum--
	if f.Parent != nil {
		f.Parent.DecreaseFavNum()
	}
}

func (f *Fav) AddBoard(bid ptttype.Bid) (idx int, favType *FavType, err error) {
	// check whether already added
	idx, favBoard, err := f.GetBoard(bid)
	if err != nil {
		return -1, nil, err
	}

	if favBoard != nil {
		return idx, favBoard, nil
	}

	if f.isMaxSize() {
		return -1, nil, ErrTooManyFavs
	}

	fp := &FavBoard{
		Bid: bid,
	}

	favType, err = f.PreAppend(pttbbsfav.FAVT_BOARD, fp)
	if err != nil {
		return -1, nil, err
	}

	return len(f.Favh) - 1, favType, nil
}

func (f *Fav) GetBoard(bid ptttype.Bid) (idx int, favType *FavType, err error) {
	if !bid.IsValid() {
		return -1, nil, ptttype.ErrInvalidBid
	}

	idx, favType = f.GetFavItem(int(bid), pttbbsfav.FAVT_BOARD)

	return idx, favType, nil
}

func (f *Fav) AddLine() (idx int, favType *FavType, err error) {
	if f.isMaxSize() {
		return -1, nil, ErrTooManyFavs
	}

	if f.NLines >= MAX_LINE {
		return -1, nil, ErrTooManyLines
	}

	fp := &FavLine{}
	favType, err = f.PreAppend(pttbbsfav.FAVT_LINE, fp)
	if err != nil {
		return -1, nil, err
	}

	return len(f.Favh) - 1, favType, nil
}

func (f *Fav) AddFolder(title string) (idx int, favType *FavType, err error) {
	if f.isMaxSize() {
		return -1, nil, ErrTooManyFavs
	}

	if f.NLines >= MAX_FOLDER {
		return -1, nil, ErrTooManyFolders
	}

	newFav, err := NewFav(f, f.Root, f.Depth+1)
	if err != nil {
		return -1, nil, err
	}

	fp := &FavFolder{
		Title:      title,
		ThisFolder: newFav,
	}
	favType, err = f.PreAppend(pttbbsfav.FAVT_FOLDER, fp)
	if err != nil {
		return -1, nil, err
	}

	return len(f.Favh) - 1, favType, nil
}

func (f *Fav) GetFavItem(theID int, theType pttbbsfav.FavT) (idx int, favType *FavType) {
	for idx, each := range f.Favh {
		if each.TheType != theType {
			continue
		}

		eachID := each.GetID()
		if eachID == theID {
			return idx, each
		}
	}

	return -1, nil
}

// PreAppend
//
// https://github.com/ptt/pttbbs/blob/master/mbbsd/fav.c#L804
// Although it is named PreAppend, actually it appends to DataTail
func (f *Fav) PreAppend(theType pttbbsfav.FavT, fp any) (favType *FavType, err error) {
	favt := &FavType{
		TheType: theType,
		Attr:    pttbbsfav.FAVH_FAV,
		Fp:      fp,
	}

	f.Favh = append(f.Favh, favt)
	f.Increase(theType, fp)

	return favt, nil
}

func (f *Fav) Increase(theType pttbbsfav.FavT, fp any) {
	switch theType {
	case pttbbsfav.FAVT_BOARD:
		f.NBoards++
	case pttbbsfav.FAVT_LINE:
		f.NLines++
		f.LineID++
		ft, _ := fp.(*FavLine)
		ft.Lid = f.LineID
	case pttbbsfav.FAVT_FOLDER:
		f.NFolders++
		f.FolderID++
		ft, _ := fp.(*FavFolder)
		ft.Fid = f.FolderID
	}
	f.IncreaseFavNum()
}

func (f *Fav) isMaxSize() bool {
	return f.Root.FavNum >= MAX_FAV
}

func (f *Fav) CleanParentAndRoot() {
	f.Parent = nil
	f.Root = nil
	for _, each := range f.Favh {
		if each.TheType != pttbbsfav.FAVT_FOLDER {
			continue
		}

		ft, _ := each.Fp.(*FavFolder)
		ft.ThisFolder.CleanParentAndRoot()
	}
}

func (f *Fav) SetFavTypeFavIdx(startFavIdx int) (newFavIdx int) {
	newFavIdx = startFavIdx
	for idx, each := range f.Favh {
		each.FavIdx = startFavIdx + idx
	}
	newFavIdx += len(f.Favh)

	for _, each := range f.Favh {
		if each.TheType != pttbbsfav.FAVT_FOLDER {
			continue
		}

		folder := each.CastFolder()
		newFavIdx = folder.ThisFolder.SetFavTypeFavIdx(newFavIdx)
	}

	return newFavIdx
}

// LocateFav
//
// Locate the fav based on levelIdxList.
// Requiring that all the fav / subFav are with Folder type.
//
// There will be 1 levelIdx in levelIdxList if referring to f itself.
// Ex. if referring to f as root, then levelIdxList is []string{""}
func (f *Fav) LocateFav(levelIdxList []string) (newFav *Fav, err error) {
	if len(levelIdxList) < 2 {
		return f, nil
	}

	nextIdxList := levelIdxList[1:]
	nextIdx := nextIdxList[0]
	theIdx := 0
	if nextIdx != "" {
		theIdx, err = strconv.Atoi(nextIdx)
		if err != nil {
			return nil, err
		}
	}

	ftype := f.Favh[theIdx]
	if ftype.TheType != pttbbsfav.FAVT_FOLDER {
		return nil, ErrInvalidLevelIdx
	}

	newFav = ftype.CastFolder().ThisFolder

	return newFav.LocateFav(nextIdxList)
}

func (f *Fav) DeleteIdx(idx int) (err error) {
	if len(f.Favh) <= idx {
		return ErrInvalidFavRecord
	}

	childFav := f.Favh[idx]

	prefixFavs := f.Favh[:idx]
	postFavs := f.Favh[idx+1:]

	favNum := 1
	switch childFav.TheType {
	case pttbbsfav.FAVT_BOARD:
		f.NBoards--
	case pttbbsfav.FAVT_FOLDER:
		childFavFolder := childFav.CastFolder()
		f.NFolders--
		favNum += childFavFolder.ThisFolder.FavNum
		for _, each := range postFavs {
			if each.TheType == pttbbsfav.FAVT_FOLDER {
				eachFolder := each.CastFolder()
				eachFolder.Fid--
			}
		}
	case pttbbsfav.FAVT_LINE:
		f.NLines--
		for _, each := range postFavs {
			if each.TheType == pttbbsfav.FAVT_LINE {
				eachLine := each.CastLine()
				eachLine.Lid--
			}
		}
	}

	f.Favh = append(prefixFavs, postFavs...)

	f.FavNum -= favNum
	if f.Parent != nil {
		f.Parent.DeleteFavNum(favNum)
	}

	return nil
}

func (f *Fav) DeleteFavNum(favNum int) {
	f.FavNum -= favNum
	if f.Parent != nil {
		f.Parent.DeleteFavNum(favNum)
	}
}
