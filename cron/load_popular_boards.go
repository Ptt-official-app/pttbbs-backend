package cron

import (
	"context"
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/api"
	"github.com/sirupsen/logrus"
)

func RetryLoadPopularBoards(ctx context.Context) error {
	time.Sleep(5 * time.Second)

	nBoardID, err := api.RefreshBoardIDWhiteListMap()
	if err != nil {
		logrus.Warnf("RetryLoadPopularBoards: unable to refresh BOARD_ID_WHITE_LIST_MAP: e: %v", err)
	}
	logrus.Infof("RetryLoadPopularBoards: refresh %v board-ids in BOARD_ID_WHITE_LIST_MAP", nBoardID)

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
				logrus.Infof("RetryLoadPopularBoards: to sleep 10 min")
				time.Sleep(10 * time.Minute)
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
