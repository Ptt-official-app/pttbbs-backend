package main

import (
	"flag"

	"github.com/Ptt-official-app/pttbbs-backend/boardd"
	"github.com/Ptt-official-app/pttbbs-backend/configutil"
	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/dbcs"
	"github.com/Ptt-official-app/pttbbs-backend/mand"
	"github.com/Ptt-official-app/pttbbs-backend/queue"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"

	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	pttbbstypes "github.com/Ptt-official-app/go-pttbbs/types"
	"github.com/sirupsen/logrus"
	jww "github.com/spf13/jwalterweatherman"
)

// Params
//
//	filename: ini filename
//
// Return
//
//	error: err
func initAllConfig(filename string) (err error) {
	err = configutil.InitViper(filename)
	if err != nil {
		return err
	}

	err = types.InitConfig()
	if err != nil {
		return err
	}

	err = db.InitConfig()
	if err != nil {
		return err
	}

	err = schema.InitConfig()
	if err != nil {
		return err
	}

	err = pttbbsapi.InitConfig()
	if err != nil {
		return err
	}

	err = pttbbstypes.InitConfig()
	if err != nil {
		return err
	}

	err = queue.InitConfig()
	if err != nil {
		return err
	}

	err = boardd.InitConfig()
	if err != nil {
		return err
	}

	err = mand.InitConfig()
	if err != nil {
		return err
	}

	err = dbcs.InitConfig()
	if err != nil {
		return err
	}

	return nil
}

func initMain() error {
	jww.SetLogThreshold(jww.LevelDebug)
	jww.SetStdoutThreshold(jww.LevelDebug)
	logrus.SetLevel(logrus.InfoLevel)

	filename := ""
	flag.StringVar(&filename, "ini", "config.ini", "ini filename")
	flag.Parse()

	err := initAllConfig(filename)
	if err != nil {
		return err
	}

	err = schema.Init()
	if err != nil {
		return err
	}

	err = boardd.Init(false)
	if err != nil {
		return err
	}

	err = mand.Init(false)
	if err != nil {
		return err
	}

	logrus.Infof("initMain: done: MAX_USERS: %v MAX_ACTIVE: %v MAX_BOARD: %v HOTBOARDCACHE: %v", ptttype.MAX_USERS, ptttype.MAX_ACTIVE, ptttype.MAX_BOARD, ptttype.HOTBOARDCACHE)

	return nil
}
