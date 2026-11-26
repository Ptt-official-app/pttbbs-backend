package cron

import (
	"context"
	"fmt"
	"time"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/pttbbs-backend/api"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/sirupsen/logrus"
)

func RetryLoadGeneralBoards(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			logrus.Infof("RetryLoadGeneralBoards: to LoadGeneralBoards")
			_ = LoadGeneralBoards()
			logrus.Infof("RetryLoadGeneralBoards: to sleep 1 min")
			if !sleepCtx(ctx, 1*time.Minute) {
				return nil
			}
		}
	}
}

func LoadGeneralBoards() (err error) {
	nextIdx := ""
	count := 0
	for {
		boardDetails, newNextIdx, err := loadGeneralBoards(nextIdx)
		if err != nil {
			logrus.Errorf("cron.LoadGeneralBoards: unable to loadGeneralBoards: nextIdx: %v e: %v", nextIdx, err)
			return err
		}

		count += len(boardDetails)

		if newNextIdx == "" {
			logrus.Infof("cron.LoadGeneralBoards: load %v boards", count)
			return nil
		}

		nextIdx = newNextIdx
	}
}

func loadGeneralBoards(startIdx string) (boardDetails []*schema.BoardDetail, nextIdx string, err error) {
	// backend load-general-baords
	theParams_b := &pttbbsapi.LoadGeneralBoardDetailsParams{
		StartIdx: startIdx,
		NBoards:  N_BOARDS,
		Asc:      true,
		IsSystem: true,
	}
	var result_b *pttbbsapi.LoadGeneralBoardDetailsResult

	statusCode, err := utils.BackendGet(nil, pttbbsapi.LOAD_GENERAL_BOARD_DETAILS_R, theParams_b, nil, &result_b)
	if err != nil {
		return nil, "", err
	}
	if statusCode != 200 {
		return nil, "", fmt.Errorf("invalid statusCode: statusCode: %v", statusCode)
	}

	// update to db
	updateNanoTS := types.NowNanoTS()
	boardDetails, err = api.DeserializeBoardDetailsAndUpdateDB(result_b.Boards, updateNanoTS)
	if err != nil {
		return nil, "", err
	}

	return boardDetails, result_b.NextIdx, nil
}
