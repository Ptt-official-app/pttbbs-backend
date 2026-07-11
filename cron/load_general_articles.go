package cron

import (
	"context"
	"time"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/api"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/sirupsen/logrus"
)

func RetryLoadGeneralArticles(ctx context.Context) error {
	time.Sleep(10 * time.Second)

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			logrus.Infof("RetryLoadGeneralArticles: to LoadGeneralArticles")
			_ = LoadGeneralArticles()
			select {
			case <-ctx.Done():
				return nil
			default:
				logrus.Infof("RetryLoadGeneralArticles: to sleep %v seconds", types.SLEEP_RETRY_LOAD_GENERAL_ARTICLES_TS_DURATION.Seconds())
				time.Sleep(types.SLEEP_RETRY_LOAD_GENERAL_ARTICLES_TS_DURATION)
			}
		}
	}
}

func LoadGeneralArticles() (err error) {
	nextBrdname := ""
	count := 0

	for {
		boardIDs, err := schema.GetBoardIDs(nextBrdname, false, N_BOARDS+1, true)
		if err != nil {
			logrus.Errorf("cron.LoadGeneralArticles: unable to GetBoardIDs: e: %v", err)
			return err
		}

		newNextBrdname := ""
		if len(boardIDs) == N_BOARDS+1 {
			newNextBoardID := boardIDs[N_BOARDS]
			newNextBrdname = newNextBoardID.Brdname
			boardIDs = boardIDs[:N_BOARDS]
		}

		for _, each := range boardIDs {
			err = loadGeneralArticlesPtt(each.BBoardID)

			if err == nil {
				count++
			}
		}

		if newNextBrdname == "" {
			logrus.Infof("cron.LoadGeneralArticles: load %v boards", count)
			return nil

		}

		nextBrdname = newNextBrdname
	}
}

func loadGeneralArticlesPtt(boardID bbs.BBoardID) (err error) {
	nextIdx := ""
	count := 0

	updateNanoTS := types.NowNanoTS()
	for {
		articleSummaries, newNextIdx, err := loadGeneralArticlesCorePtt(boardID, nextIdx, updateNanoTS)
		if err != nil {
			logrus.Errorf("cron.loadGeneralArticlesPtt: unable to loadGeneralArticlesCorePtt: board: %v nextIdx: %v e: %v", boardID, nextIdx, err)
			return err
		}
		count += len(articleSummaries)

		if newNextIdx == INVALID_LOAD_GENERAL_ARTICLES_NEXT_IDX_PTT {
			break
		}

		nextIdx = newNextIdx
	}

	bottomUpdateNanoTS := types.NowNanoTS()
	err = loadBottomArticlesPtt(boardID, bottomUpdateNanoTS)
	if err != nil {
		logrus.Errorf("cron.loadGeneralArticlesPtt: unable to loadBottomArticles: boardID: %v e: %v", boardID, err)
		return err
	}

	err = schema.DeleteArticlesByBoardID(boardID, updateNanoTS)
	if err != nil {
		logrus.Errorf("cron.loadGeneralArticlesPtt: unable to delete previous articles: boardID: %v e: %v", boardID, err)
		return err
	}

	return nil
}

func loadGeneralArticlesCorePtt(boardID bbs.BBoardID, startIdx string, updateNanoTS types.NanoTS) (articleSummaries []*schema.ArticleSummaryWithRegex, nextIdx string, err error) {
	// backend load-general-articles
	theParams_b := &pttbbsapi.LoadGeneralArticlesParams{
		StartIdx:  startIdx,
		NArticles: N_ARTICLES,
		Desc:      true,
		IsSystem:  true,
	}
	var result_b *pttbbsapi.LoadGeneralArticlesResult

	urlMap := map[string]string{
		"bid": string(boardID),
	}
	url := utils.MergeURL(urlMap, pttbbsapi.LOAD_GENERAL_ARTICLES_R)
	statusCode, err := utils.BackendGet(nil, url, theParams_b, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, "", err
	}

	articleSummaries, err = api.DeserializeArticlesAndUpdateDB(result_b.Articles, updateNanoTS, false)
	if err != nil {
		return nil, "", err
	}

	return articleSummaries, result_b.NextIdx, nil
}

func loadBottomArticlesPtt(boardID bbs.BBoardID, updateNanoTS types.NanoTS) (err error) {
	// backend load-general-articles
	var result_b *pttbbsapi.LoadBottomArticlesResult

	urlMap := map[string]string{
		"bid": string(boardID),
	}
	url := utils.MergeURL(urlMap, pttbbsapi.LOAD_BOTTOM_ARTICLES_R)
	statusCode, err := utils.BackendGet(nil, url, nil, nil, &result_b)
	if err != nil || statusCode != 200 {
		return err
	}

	_, err = api.DeserializeArticlesAndUpdateDB(result_b.Articles, updateNanoTS, true)
	if err != nil {
		return err
	}

	return nil
}
