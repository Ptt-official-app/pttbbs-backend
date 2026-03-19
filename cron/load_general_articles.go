package cron

import (
	"context"
	"time"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/api"
	"github.com/Ptt-official-app/pttbbs-backend/boardd"
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
				logrus.Infof("RetryLoadGeneralArticles: to sleep 10 min")
				time.Sleep(10 * time.Minute)
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
			if !types.IS_ALL_GUEST {
				err = loadGeneralArticlesBoardd(each.BBoardID)
			} else {
				err = loadGeneralArticlesPtt(each.BBoardID)
			}

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

func loadGeneralArticlesBoardd(boardID bbs.BBoardID) (err error) {
	nextIdx := int32(0)
	count := 0

	for {
		articleSummaries, newNextIdx, err := loadGeneralArticlesCoreBoardd(boardID, nextIdx)
		if err != nil {
			logrus.Errorf("cron.LoadGeneralArticles: unable to loadGeneralArticles: nextIdx: %v e: %v", nextIdx, err)
			return err
		}
		count += len(articleSummaries)

		if newNextIdx == INVALID_LOAD_GENERAL_ARTICLES_NEXT_IDX_BOARDD {
			break
		}

		nextIdx = newNextIdx
	}

	err = loadBottomArticlesBoardd(boardID)
	if err != nil {
		logrus.Errorf("loadGeneralArticles: unable to loadBottomArticles: e: %v", err)
		return err
	}

	return nil
}

func loadGeneralArticlesCoreBoardd(boardID bbs.BBoardID, startIdx int32) (articleSummaries []*schema.ArticleSummaryWithRegex, nextIdx int32, err error) {
	nextIdx = INVALID_LOAD_GENERAL_ARTICLES_NEXT_IDX_BOARDD
	brdnameStr := boardID.ToBrdname()
	// backend load-general-articles
	ctx := context.Background()
	brdname := &boardd.BoardRef_Name{Name: brdnameStr}
	req := &boardd.ListRequest{
		Ref:          &boardd.BoardRef{Ref: brdname},
		IncludePosts: true,
		Offset:       startIdx,
		Length:       N_ARTICLES + 1,
	}
	resp, err := boardd.Cli.List(ctx, req)
	if err != nil {
		return nil, INVALID_LOAD_GENERAL_ARTICLES_NEXT_IDX_BOARDD, err
	}

	posts := resp.Posts
	if len(posts) == N_ARTICLES+1 {
		nextIdx = startIdx + N_ARTICLES
		posts = posts[:N_ARTICLES]
	}

	// update to db
	updateNanoTS := types.NowNanoTS()
	articleSummaries, err = api.DeserializePBArticlesAndUpdateDB(boardID, posts, updateNanoTS, false)
	if err != nil {
		return nil, INVALID_LOAD_GENERAL_ARTICLES_NEXT_IDX_BOARDD, err
	}

	return articleSummaries, nextIdx, nil
}

func loadBottomArticlesBoardd(boardID bbs.BBoardID) (err error) {
	brdnameStr := boardID.ToBrdname()
	// backend load-general-articles

	ctx := context.Background()
	brdname := &boardd.BoardRef_Name{Name: brdnameStr}
	req := &boardd.ListRequest{
		Ref:            &boardd.BoardRef{Ref: brdname},
		IncludeBottoms: true,
		Offset:         0,
		Length:         N_ARTICLES + 1,
	}
	resp, err := boardd.Cli.List(ctx, req)
	if err != nil {
		return err
	}

	err = schema.ResetArticleIsBottom(boardID)
	if err != nil {
		return err
	}

	updateNanoTS := types.NowNanoTS()
	_, err = api.DeserializePBArticlesAndUpdateDB(boardID, resp.Bottoms, updateNanoTS, true)
	if err != nil {
		return err
	}

	return nil
}

func loadGeneralArticlesPtt(boardID bbs.BBoardID) (err error) {
	nextIdx := ""
	count := 0

	updateNanoTS := types.NowNanoTS()
	for {
		articleSummaries, newNextIdx, err := loadGeneralArticlesCorePtt(boardID, nextIdx, updateNanoTS)
		if err != nil {
			logrus.Errorf("cron.loadGeneralArticlesPtt: unable to loadGeneralArticlesCorePtt: nextIdx: %v e: %v", nextIdx, err)
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
