package cron

import (
	"context"
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/api"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/sirupsen/logrus"
)

func RetryLoadPopularBoards(ctx context.Context) error {
	time.Sleep(5 * time.Second)

	nBoardID, err := api.RefreshBrdnameWhiteListMap()
	if err != nil {
		logrus.Warnf("RetryLoadPopularBoards: unable to refresh BRDNAME_WHITE_LIST_MAP: e: %v", err)
	}
	logrus.Infof("RetryLoadPopularBoards: refresh %v board-ids in BRDNAME_WHITE_LIST_MAP", nBoardID)

	nBoardID, err = api.RefreshBrdnameBlackListMap()
	if err != nil {
		logrus.Warnf("RetryLoadPopularBoards: unable to refresh BRDNAME_BLACK_LIST_MAP: e: %v", err)
	}
	logrus.Infof("RetryLoadPopularBoards: refresh %v board-ids in BRDNAME_BLACK_LIST_MAP", nBoardID)

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			logrus.Infof("RetryLoadPopularBoards: to LoadPopularBoards")
			_ = LoadPopularBoards()
			select {
			case <-ctx.Done():
				return nil
			default:
				logrus.Infof("RetryLoadPopularBoards: to sleep %v seconds", types.SLEEP_RETRY_LOAD_POPULAR_BOARDS_TS_DURATION.Seconds())
				time.Sleep(types.SLEEP_RETRY_LOAD_POPULAR_BOARDS_TS_DURATION)
			}
		}
	}
}

func LoadPopularBoards() (err error) {
	boardSummaries, err := api.TryLoadPttWebPopularBoards(nil)
	if err != nil {
		logrus.Errorf("cron.LoadPopularBoards: unable to LoadPopularBoards: e: %v", err)
		return err
	}
	logrus.Infof("cron.LoadPopularBoards: boardSummaries: %v", len(boardSummaries))

	return nil
}
