package api

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/Ptt-official-app/pttbbs-backend/types"
)

func TestRefreshBrdnameListMapAllGuestOnlyFromLocal(t *testing.T) {
	dir := t.TempDir()
	whiteListFilename := filepath.Join(dir, "white.yaml")
	blackListFilename := filepath.Join(dir, "black.yaml")
	if err := os.WriteFile(whiteListFilename, []byte("WhoAmI: 1\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(blackListFilename, []byte("Test: true\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	origWhite, origBlack := types.BRDNAME_WHITE_LIST_MAP_FILENAME, types.BRDNAME_BLACK_LIST_MAP_FILENAME
	types.BRDNAME_WHITE_LIST_MAP_FILENAME, types.BRDNAME_BLACK_LIST_MAP_FILENAME = whiteListFilename, blackListFilename
	defer func() {
		types.BRDNAME_WHITE_LIST_MAP_FILENAME, types.BRDNAME_BLACK_LIST_MAP_FILENAME = origWhite, origBlack
		BRDNAME_WHITE_LIST_MAP.Store(nil)
		BRDNAME_BLACK_LIST_MAP.Store(nil)
	}()

	handlers := map[string]APIFunc{
		"white": RefreshBrdnameWhiteListMapAllGuest,
		"black": RefreshBrdnameBlackListMapAllGuest,
	}
	for name, handler := range handlers {
		_, statusCode, err := handler("203.0.113.7", &UserInfo{}, nil, nil)
		if statusCode != 403 || !errors.Is(err, ErrInvalidRemoteAddr) {
			t.Errorf("%v: remote: statusCode: %v e: %v, want 403 ErrInvalidRemoteAddr", name, statusCode, err)
		}
	}
	if len(brdnameWhiteListMap()) != 0 || len(brdnameBlackListMap()) != 0 {
		t.Errorf("lists loaded by a remote request: white: %v black: %v", brdnameWhiteListMap(), brdnameBlackListMap())
	}

	for name, handler := range handlers {
		_, statusCode, err := handler("127.0.0.1", &UserInfo{}, nil, nil)
		if statusCode != 200 || err != nil {
			t.Errorf("%v: local: statusCode: %v e: %v, want 200", name, statusCode, err)
		}
	}
	if brdnameWhiteListMap()["WhoAmI"] != 1 || !brdnameBlackListMap()["Test"] {
		t.Errorf("lists not loaded: white: %v black: %v", brdnameWhiteListMap(), brdnameBlackListMap())
	}
}
