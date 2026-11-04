package fav

import pttbbsfav "github.com/Ptt-official-app/go-pttbbs/ptt/fav"

type FavType struct {
	FavIdx  int
	TheType pttbbsfav.FavT
	Attr    pttbbsfav.Favh
	Fp      any
}

func (ft *FavType) CastFolder() *FavFolder {
	if ft.Fp != nil {
		favFolder, ok := ft.Fp.(*FavFolder)
		if !ok {
			return nil
		}
		return favFolder
	}
	favFolder := &FavFolder{}
	ft.Fp = favFolder
	return favFolder
}

func (ft *FavType) CastBoard() *FavBoard {
	if ft.Fp != nil {
		favBoard, ok := ft.Fp.(*FavBoard)
		if !ok {
			return nil
		}
		return favBoard
	}

	favBoard := &FavBoard{}
	ft.Fp = favBoard
	return favBoard
}

func (ft *FavType) CastLine() *FavLine {
	if ft.Fp != nil {
		favLine, ok := ft.Fp.(*FavLine)
		if !ok {
			return nil
		}
		return favLine
	}

	favLine := &FavLine{}
	ft.Fp = favLine
	return favLine
}

func (ft *FavType) isValid() bool { //nolint
	return ft.Attr&pttbbsfav.FAVH_FAV != 0
}

func (ft *FavType) GetID() (theID int) {
	switch ft.TheType {
	case pttbbsfav.FAVT_BOARD:
		fp := ft.CastBoard()
		return int(fp.Bid)
	case pttbbsfav.FAVT_LINE:
		fp := ft.CastLine()
		return int(fp.Lid)
	case pttbbsfav.FAVT_FOLDER:
		fp := ft.CastFolder()
		return int(fp.Fid)
	}

	return 0
}
