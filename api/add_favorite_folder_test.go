package api

import (
	"reflect"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/gin-gonic/gin"
)

func TestAddFavoriteFolder(t *testing.T) {
	setupTest()
	defer teardownTest()

	path0 := &AddFavoriteFolderPath{
		UserID: "SYSOP",
	}

	params0 := &AddFavoriteFolderParams{}

	ret0 := AddFavoriteResult(&apitypes.BoardSummary{
		Title:    "新的目錄",
		StatAttr: ptttype.NBRD_FOLDER,
		LevelIdx: ":6",
		Idx:      "6",
		URL:      "/user/SYSOP/favorites?level_idx=:6",
		BMs:      []bbs.UUserID{},

		TokenUser: "SYSOP",
	})

	type args struct {
		remoteAddr string
		userID     bbs.UUserID
		params     any
		path       any
		c          *gin.Context
	}
	tests := []struct {
		name           string
		args           args
		wantResult     any
		wantStatusCode int
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			args:           args{userID: "SYSOP", params: params0, path: path0},
			wantResult:     ret0,
			wantStatusCode: 200,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			user := &UserInfo{UserID: tt.args.userID, IsOver18: true}
			gotResult, gotStatusCode, err := AddFavoriteFolder(tt.args.remoteAddr, user, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddFavoriteFolder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("AddFavoriteFolder() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotStatusCode != tt.wantStatusCode {
				t.Errorf("AddFavoriteFolder() gotStatusCode = %v, want %v", gotStatusCode, tt.wantStatusCode)
			}
		})
		wg.Wait()
	}
}
