package schema

import (
	"cmp"
	"slices"
	"strconv"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	pttbbsfav "github.com/Ptt-official-app/go-pttbbs/ptt/fav"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/fav"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
)

var UserFavorites_c *db.Collection

const (
	N_USER_FAVORITES_DOUBLE_BUFFER = 2
)

type LevelIdx string

func SetLevelIdx(prefix LevelIdx, idx int) LevelIdx {
	return prefix + ":" + LevelIdx(strconv.Itoa(idx))
}

// favorites is bounded by 1024 items.
// each item can be at most 100 bytes.
// We should be able to have all the favorites inside 1 document.
type UserFavorites struct {
	UserID          bbs.UUserID  `bson:"user_id"`
	DoubleBufferIdx int          `bson:"dbuffer_idx"`
	FavIdx          int          `bson:"fav_idx"`
	LevelIdx        LevelIdx     `bson:"level_idx"`
	Idx             int          `bson:"idx"`
	UpdateNanoTS    types.NanoTS `bson:"update_nano_ts"`
	MTime           types.NanoTS `bson:"mtime_nano_ts"`

	TheType pttbbsfav.FavT `bson:"the_type"`
	Attr    pttbbsfav.Favh `bson:"attr"`
	TheID   int            `bson:"the_id"`

	// for folder
	FolderTitle string      `bson:"folder_title"`
	FolderMeta  *FolderMeta `bson:"folder_meta"`
}

type FolderMeta struct {
	FavNum   int `bson:"fav_num"`
	NBoards  int `bson:"n_boards"`
	NLines   int `bson:"n_lines"`
	NFolders int `bson:"n_folders"`
}

var EMPTY_USER_FAVORITES = &UserFavorites{}

var (
	USER_FAVORITES_USER_ID_b           = getBSONName(EMPTY_USER_FAVORITES, "UserID")
	USER_FAVORITES_DOUBLE_BUFFER_IDX_b = getBSONName(EMPTY_USER_FAVORITES, "DoubleBufferIdx")
	USER_FAVORITES_FAV_IDX_b           = getBSONName(EMPTY_USER_FAVORITES, "FavIdx")
	USER_FAVORITES_LEVEL_IDX_b         = getBSONName(EMPTY_USER_FAVORITES, "LevelIdx")
	USER_FAVORITES_IDX_b               = getBSONName(EMPTY_USER_FAVORITES, "Idx")
	USER_FAVORITES_UPDATE_NANO_TS_b    = getBSONName(EMPTY_USER_FAVORITES, "UpdateNanoTS")
	USER_FAVORITES_MTIME_b             = getBSONName(EMPTY_USER_FAVORITES, "MTime")

	USER_FAVORITES_THE_TYPE_b     = getBSONName(EMPTY_USER_FAVORITES, "TheType")
	USER_FAVORITES_ATTR_b         = getBSONName(EMPTY_USER_FAVORITES, "Attr")
	USER_FAVORITES_THE_ID_b       = getBSONName(EMPTY_USER_FAVORITES, "TheID")
	USER_FAVORITES_FOLDER_TITLE_b = getBSONName(EMPTY_USER_FAVORITES, "FolderTitle")
)

type UserFavoritesQuery struct {
	UserID          bbs.UUserID `bson:"user_id"`
	DoubleBufferIdx int         `bson:"dbuffer_idx"`
	FavIdx          int         `bson:"fav_idx"`
}

var EMPTY_USER_FAVORITES_QUERY = &UserFavoritesQuery{}

func assertUserFavorites() error {
	if err := assertFields(EMPTY_USER_FAVORITES, EMPTY_USER_FAVORITES_QUERY); err != nil {
		return err
	}

	if err := assertFields(EMPTY_USER_FAVORITES, EMPTY_USER_FAVORITE_ID); err != nil {
		return err
	}

	return nil
}

func GetAllUserFavorites(userID bbs.UUserID, doubleBufferIdx int, mTime types.NanoTS) (userFavorites []*UserFavorites, err error) {
	query := bson.M{
		USER_FAVORITES_USER_ID_b:           userID,
		USER_FAVORITES_DOUBLE_BUFFER_IDX_b: doubleBufferIdx,
		USER_FAVORITES_MTIME_b:             mTime,
	}

	sortOpts := bson.D{
		{Key: USER_FAVORITES_FAV_IDX_b, Value: 1},
	}

	err = UserFavorites_c.Find(query, 0, &userFavorites, nil, sortOpts)
	if err != nil {
		return nil, err
	}

	return userFavorites, nil
}

func GetUserFavorites(userID bbs.UUserID, doubleBufferIdx int, levelIdx LevelIdx, startIdx int, ascending bool, limit int, mTime types.NanoTS) (userFavorites []*UserFavorites, err error) {
	var queryIdx bson.M
	var sortOpts bson.D
	if ascending {
		queryIdx = bson.M{
			"$gte": startIdx,
		}
		sortOpts = bson.D{
			{Key: USER_FAVORITES_IDX_b, Value: 1},
		}
	} else {
		queryIdx = bson.M{
			"$lte": startIdx,
		}
		sortOpts = bson.D{
			{Key: USER_FAVORITES_IDX_b, Value: -1},
		}
	}

	query := bson.M{
		USER_FAVORITES_USER_ID_b:           userID,
		USER_FAVORITES_DOUBLE_BUFFER_IDX_b: doubleBufferIdx,
		USER_FAVORITES_LEVEL_IDX_b:         levelIdx,
		USER_FAVORITES_IDX_b:               queryIdx,
		USER_FAVORITES_MTIME_b:             mTime,
	}

	err = UserFavorites_c.Find(query, int64(limit), &userFavorites, nil, sortOpts)
	if err != nil {
		return nil, err
	}

	userFavorites = SortUserFavoritesByFavIdx(userFavorites, ascending)

	return userFavorites, nil
}

func UpdateUserFavorites(userID bbs.UUserID, doubleBufferIdx int, userFavorites []*UserFavorites, mTime types.NanoTS, updateNanoTS types.NanoTS) (err error) {
	if len(userFavorites) == 0 {
		return nil
	}

	// bulk-create-only
	theList := make([]*db.UpdatePair, len(userFavorites))
	for idx, each := range userFavorites {
		query := &UserFavoritesQuery{
			UserID:          userID,
			DoubleBufferIdx: doubleBufferIdx,
			FavIdx:          each.FavIdx,
		}
		theList[idx] = &db.UpdatePair{
			Filter: query,
			Update: each,
		}
	}

	r, err := UserFavorites_c.BulkCreateOnly(theList)
	if err != nil {
		return err
	}
	if r.UpsertedCount == int64(len(userFavorites)) { // all are created
		return nil
	}

	// bulk-update-update-one-only
	upsertedIDs := r.UpsertedIDs
	updateUserFavories := make([]*db.UpdatePair, 0, len(theList))
	for idx, each := range theList {
		_, ok := upsertedIDs[int64(idx)]
		if ok {
			continue
		}

		origFilter := each.Filter.(*UserFavoritesQuery)
		filter := bson.M{
			USER_FAVORITES_USER_ID_b:           origFilter.UserID,
			USER_FAVORITES_DOUBLE_BUFFER_IDX_b: origFilter.DoubleBufferIdx,
			USER_FAVORITES_FAV_IDX_b:           origFilter.FavIdx,
			USER_FAVORITES_MTIME_b: bson.M{
				"$lt": mTime,
			},
		}
		each.Filter = filter
		updateUserFavories = append(updateUserFavories, each)
	}

	_, err = UserFavorites_c.BulkUpdateOneOnly(updateUserFavories)
	return err
}

func SortUserFavoritesByFavIdx(userFavorites []*UserFavorites, ascending bool) (newUserFavorites []*UserFavorites) {
	if ascending {
		slices.SortStableFunc(userFavorites, func(a, b *UserFavorites) int {
			return cmp.Compare(a.FavIdx, b.FavIdx)
		})
	} else {
		slices.SortStableFunc(userFavorites, func(a, b *UserFavorites) int {
			return cmp.Compare(b.FavIdx, a.FavIdx)
		})
	}

	return userFavorites
}

func UserFavoritesToFav(meta *FolderMeta, userFavorites []*UserFavorites, depth int) (f *fav.Fav, err error) {
	if len(userFavorites) != meta.FavNum {
		return nil, ErrInvalidUserFavorites
	}

	SortUserFavoritesByFavIdx(userFavorites, true)

	f, _ = fav.NewFav(nil, nil, depth)

	nLevel0 := meta.NLines + meta.NFolders + meta.NBoards

	first, theRest := userFavorites[:nLevel0], userFavorites[nLevel0:]

	f.FavNum = meta.FavNum
	f.NBoards = int16(meta.NBoards)
	f.NLines = int8(meta.NLines)
	f.NFolders = int8(meta.NFolders)
	f.LineID = pttbbsfav.Lid(meta.NLines)
	f.FolderID = pttbbsfav.Fid(meta.NFolders)
	f.Favh = make([]*fav.FavType, nLevel0)

	for idx, each := range first {
		f.Favh[idx] = userFavoritesToFavType(each)
	}

	var eachFirst []*UserFavorites
	for idx, each := range first {
		if each.TheType != pttbbsfav.FAVT_FOLDER {
			continue
		}

		eachFolderMeta := each.FolderMeta
		eachFavNum := eachFolderMeta.FavNum
		eachFirst, theRest = theRest[:eachFavNum], theRest[eachFavNum:]
		eachFav, err := UserFavoritesToFav(eachFolderMeta, eachFirst, depth+1)
		if err != nil {
			return nil, err
		}

		eachFavT := f.Favh[idx]
		eachFolder := eachFavT.CastFolder()
		eachFolder.ThisFolder = eachFav
	}

	return f, nil
}

// userFavoritestoFavType
//
// deal with only TheType, Attr, Id, Title.
// do not deal with ThisFolder in this function.
func userFavoritesToFavType(userFavorites *UserFavorites) (ft *fav.FavType) {
	ft = &fav.FavType{
		TheType: userFavorites.TheType,
		Attr:    userFavorites.Attr,
	}

	switch userFavorites.TheType {
	case pttbbsfav.FAVT_BOARD:
		fp := &fav.FavBoard{
			Bid: ptttype.Bid(userFavorites.TheID),
		}
		ft.Fp = fp
	case pttbbsfav.FAVT_LINE:
		fp := &fav.FavLine{
			Lid: pttbbsfav.Lid(userFavorites.TheID),
		}
		ft.Fp = fp
	case pttbbsfav.FAVT_FOLDER:
		fp := &fav.FavFolder{
			Fid:   pttbbsfav.Fid(userFavorites.TheID),
			Title: userFavorites.FolderTitle,
		}
		ft.Fp = fp
	}
	return ft
}

func FavToUserFavorites(f *fav.Fav, userID bbs.UUserID, doubleBufferIdx int, updateNanoTS types.NanoTS, mTime types.NanoTS) (meta *UserFavoritesMeta, userFavorites []*UserFavorites) {
	f.SetFavTypeFavIdx(0)

	meta = &UserFavoritesMeta{
		UserID:          userID,
		DoubleBufferIdx: doubleBufferIdx,
		UpdateNanoTS:    updateNanoTS,
		MTime:           mTime,
		FolderMeta: FolderMeta{
			FavNum:   f.FavNum,
			NLines:   int(f.NLines),
			NBoards:  int(f.NBoards),
			NFolders: int(f.NFolders),
		},
	}

	userFavorites = favToUserFavoritesCore(f, userID, doubleBufferIdx, "", updateNanoTS, mTime)

	return meta, userFavorites
}

func favToUserFavoritesCore(f *fav.Fav, userID bbs.UUserID, doubleBufferIdx int, levelIdx LevelIdx, updateNanoTS types.NanoTS, mTime types.NanoTS) (userFavorites []*UserFavorites) {
	userFavorites = make([]*UserFavorites, 0, f.FavNum)
	for idx, each_ft := range f.Favh {
		each := favTypeToUserFavorites(each_ft, userID, doubleBufferIdx, levelIdx, idx, updateNanoTS, mTime)
		userFavorites = append(userFavorites, each)
	}

	for idx, each_ft := range f.Favh {
		if each_ft.TheType != pttbbsfav.FAVT_FOLDER {
			continue
		}

		fp := each_ft.CastFolder()
		eachLevelIdx := SetLevelIdx(levelIdx, idx)
		eachUserFavorites := favToUserFavoritesCore(fp.ThisFolder, userID, doubleBufferIdx, eachLevelIdx, updateNanoTS, mTime)

		userFavorites = append(userFavorites, eachUserFavorites...)
	}
	return userFavorites
}

func favTypeToUserFavorites(ft *fav.FavType, userID bbs.UUserID, doubleBufferIdx int, levelIdx LevelIdx, idx int, updateNanoTS types.NanoTS, mTime types.NanoTS) *UserFavorites {
	theID := 0
	title := ""

	var folderMeta *FolderMeta
	switch ft.TheType {
	case pttbbsfav.FAVT_BOARD:
		fp := ft.CastBoard()
		theID = int(fp.Bid)
	case pttbbsfav.FAVT_LINE:
		fp := ft.CastLine()
		theID = int(fp.Lid)
	case pttbbsfav.FAVT_FOLDER:
		fp := ft.CastFolder()
		theID = int(fp.Fid)
		title = fp.Title

		folderMeta = &FolderMeta{
			FavNum:   fp.ThisFolder.FavNum,
			NLines:   int(fp.ThisFolder.NLines),
			NBoards:  int(fp.ThisFolder.NBoards),
			NFolders: int(fp.ThisFolder.NFolders),
		}
	}
	return &UserFavorites{
		UserID:          userID,
		DoubleBufferIdx: doubleBufferIdx,
		FavIdx:          ft.FavIdx,
		LevelIdx:        levelIdx,
		Idx:             idx,
		UpdateNanoTS:    updateNanoTS,
		MTime:           mTime,

		TheType:     ft.TheType,
		Attr:        ft.Attr,
		TheID:       theID,
		FolderTitle: title,
		FolderMeta:  folderMeta,
	}
}
