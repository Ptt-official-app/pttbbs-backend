package cron

import (
	"context"
	"time"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/api"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/sirupsen/logrus"
)

func RetryLoadManArticleDetails(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			logrus.Infof("RetryLoadManArticleDetails: to LoadManArticleDetails")
			runRecovered("RetryLoadManArticleDetails", LoadManArticleDetails)
			logrus.Infof("RetryLoadManArticleDetails: to sleep 1 min")
			if !sleepCtx(ctx, 1*time.Minute) {
				return nil
			}
		}
	}
}

func LoadManArticleDetails() (err error) {
	nextBrdname := ""
	count := 0

	for {
		boardIDs, err := schema.GetBoardIDs(nextBrdname, false, N_BOARDS+1, true)
		if err != nil {
			return err
		}

		newNextBrdname := ""
		if len(boardIDs) == N_BOARDS+1 {
			newNextBoardID := boardIDs[N_BOARDS]
			newNextBrdname = newNextBoardID.Brdname
			boardIDs = boardIDs[:N_BOARDS]
		}

		for _, each := range boardIDs {
			err = loadManArticleDetails(each.BBoardID, "")
			if err == nil {
				count++
			}
		}

		if newNextBrdname == "" {
			logrus.Infof("cron.LoadMainArticleDetails: load %v boards", count)
			return nil

		}

		nextBrdname = newNextBrdname
	}
}

func loadManArticleDetails(boardID bbs.BBoardID, levelIdx types.ManArticleID) (err error) {
	manArticleSummaries, err := schema.GetManArticleDetailSummaries(boardID, levelIdx)
	if err != nil {
		return err
	}

	for _, each := range manArticleSummaries {
		if each.IsDir {
			continue
		}

		if each.MTime <= each.ContentMTime && each.MTime < each.ContentUpdateNanoTS {
			continue
		}

		_, _, _, _ = api.TryGetManArticleContentInfo("SYSOP", each.BBoardID, each.ArticleID, nil, true, false)
	}

	for _, each := range manArticleSummaries {
		if !each.IsDir {
			continue
		}

		_ = loadManArticleDetails(boardID, each.ArticleID)
	}

	return nil
}
