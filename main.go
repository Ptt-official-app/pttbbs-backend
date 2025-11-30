package main

import (
	"context"
	"net/http"
	"path/filepath"
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/api"
	"github.com/Ptt-official-app/pttbbs-backend/cron"
	"github.com/Ptt-official-app/pttbbs-backend/queue"
	"github.com/Ptt-official-app/pttbbs-backend/types"

	"github.com/appleboy/graceful"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func withPrefix(path string) string {
	return types.API_PREFIX + path
}

func initGin() (*gin.Engine, error) {
	if types.IS_ALL_GUEST {
		return initGinAllGuest()
	}

	return initGinCore()
}

func initGinAllGuest() (*gin.Engine, error) {
	router := gin.Default()

	// options
	router.OPTIONS("/*path", api.OptionsWrapper)

	// index
	router.GET(withPrefix(api.INDEX_R), api.IndexWrapper)

	router.GET(withPrefix(api.GET_VERSION_R), api.GetVersionWrapper)

	// board
	router.GET(withPrefix(api.LOAD_POPULAR_BOARDS_R), api.LoadPopularBoardsAllGuestWrapper)
	router.GET(withPrefix(api.GET_BOARD_DETAIL_R), api.GetBoardDetailAllGuestWrapper)
	router.GET(withPrefix(api.GET_BOARD_SUMMARY_R), api.GetBoardSummaryAllGuestWrapper)

	router.GET(withPrefix(api.REFRESH_BOARD_ID_WHITE_LIST_MAP_R), api.RefreshBoardIDWhiteListMapAllGuestWrapper)

	// article
	router.GET(withPrefix(api.LOAD_GENERAL_ARTICLES_R), api.LoadGeneralArticlesAllGuestWrapper)
	router.GET(withPrefix(api.LOAD_BOTTOM_ARTICLES_R), api.LoadBottomArticlesAllGuestWrapper)
	router.GET(withPrefix(api.GET_ARTICLE_R), api.GetArticleDetailAllGuestWrapper)
	router.GET(withPrefix(api.GET_ARTICLE_BLOCKS_R), api.GetArticleBlocksAllGuestWrapper)

	// comments
	router.GET(withPrefix(api.LOAD_ARTICLE_COMMENTS_R), api.LoadArticleCommentsAllGuestWrapper)

	// user
	router.GET(withPrefix(api.GET_USER_ID_R), api.GetUserIDAllGuestWrapper)

	// html
	router.GET(api.ROOT_HTML_R, api.IndexHTMLWrapper)
	router.GET(api.INDEX_HTML_R, api.IndexHTMLWrapper)

	router.Static("/static", filepath.Join(types.STATIC_DIR, "static"))

	staticFiles := []string{
		"asset-manifest.json",
		"favicon.ico",
		"logo192.png",
		"logo512.png",
		"manifest.json",
		"robots.txt",
	}

	for _, each := range staticFiles {
		router.StaticFile("/"+each, filepath.Join(types.STATIC_DIR, each))
	}

	router.NoRoute(api.AllHTMLWrapper)

	return router, nil
}

func initGinCore() (*gin.Engine, error) {
	router := gin.Default()

	// options
	router.OPTIONS("/*path", api.OptionsWrapper)

	// index
	router.GET(withPrefix(api.INDEX_R), api.IndexWrapper)

	router.GET(withPrefix(api.GET_USER_VISIT_COUNT_R), api.GetUserVisitCountWrapper)
	router.GET(withPrefix(api.GET_VERSION_R), api.GetVersionWrapper)

	// register/login
	router.POST(withPrefix(api.REGISTER_CLIENT_R), api.RegisterClientWrapper)
	router.POST(withPrefix(api.REGISTER_USER_R), api.RegisterUserWrapper)
	router.POST(withPrefix(api.LOGIN_R), api.LoginWrapper)
	router.POST(withPrefix(api.ATTEMPT_REGISTER_USER_R), api.AttemptRegisterUserWrapper)
	router.POST(withPrefix(api.CHECK_EXISTS_USER_R), api.CheckExistsUserWrapper)
	router.POST(withPrefix(api.LOGOUT_R), api.LogoutWrapper)
	router.POST(withPrefix(api.REFRESH_R), api.RefreshWrapper)

	// board
	router.GET(withPrefix(api.LOAD_GENERAL_BOARDS_R), api.LoadGeneralBoardsWrapper)
	router.GET(withPrefix(api.GET_BOARD_DETAIL_R), api.GetBoardDetailWrapper)
	router.GET(withPrefix(api.GET_BOARD_SUMMARY_R), api.GetBoardSummaryWrapper)
	router.GET(withPrefix(api.LOAD_POPULAR_BOARDS_R), api.LoadPopularBoardsWrapper)
	router.GET(withPrefix(api.LOAD_GENERAL_BOARDS_BY_CLASS_R), api.LoadGeneralBoardsByClassWrapper)
	router.GET(withPrefix(api.LOAD_AUTO_COMPLETE_BOARDS_R), api.LoadAutoCompleteBoardsWrapper)
	router.POST(withPrefix(api.CREATE_BOARD_R), api.CreateBoardWrapper)
	router.GET(withPrefix(api.LOAD_CLASS_BOARDS_R), api.LoadClassBoardsWrapper)

	// article
	router.GET(withPrefix(api.LOAD_GENERAL_ARTICLES_R), api.LoadGeneralArticlesWrapper)
	router.GET(withPrefix(api.LOAD_BOTTOM_ARTICLES_R), api.LoadBottomArticlesWrapper)
	router.GET(withPrefix(api.GET_ARTICLE_R), api.GetArticleDetailWrapper)
	router.GET(withPrefix(api.GET_ARTICLE_BLOCKS_R), api.GetArticleBlocksWrapper)
	router.GET(withPrefix(api.LOAD_POPULAR_ARTICLES_R), api.LoadPopularArticlesWrapper)
	router.POST(withPrefix(api.CREATE_ARTICLE_R), api.CreateArticleWrapper)
	router.GET(withPrefix(api.CROSS_POST_R), api.CrossPostWrapper)
	router.POST(withPrefix(api.EDIT_ARTICLE_R), api.EditArticleDetailWrapper)
	router.POST(withPrefix(api.REPLY_COMMENTS_R), api.ReplyCommentsWrapper)
	router.POST(withPrefix(api.DELETE_COMMENTS_R), api.DeleteCommentsWrapper)
	router.POST(withPrefix(api.DELETE_ARTICLES_R), api.DeleteArticlesWrapper)

	// manual
	router.GET(withPrefix(api.LOAD_MAN_ARTICLES_R), api.LoadManArticlesWrapper)
	router.GET(withPrefix(api.GET_MAN_ARTICLE_R), api.GetManArticleDetailWrapper)
	router.GET(withPrefix(api.GET_MAN_ARTICLE_BLOCKS_R), api.GetManArticleBlocksWrapper)

	// user
	router.GET(withPrefix(api.GET_USER_INFO_R), api.GetUserInfoWrapper)
	router.GET(withPrefix(api.LOAD_FAVORITE_BOARDS_R), api.LoadFavoriteBoardsWrapper)
	router.GET(withPrefix(api.LOAD_USER_ARTICLES_R), api.LoadUserArticlesWrapper)
	router.POST(withPrefix(api.CHANGE_PASSWD_R), api.ChangePasswdWrapper)
	router.POST(withPrefix(api.ATTEMPT_CHANGE_EMAIL_R), api.AttemptChangeEmailWrapper)
	router.POST(withPrefix(api.CHANGE_EMAIL_R), api.ChangeEmailWrapper)
	router.POST(withPrefix(api.ATTEMPT_SET_ID_EMAIL_R), api.AttemptSetIDEmailWrapper)
	router.POST(withPrefix(api.SET_ID_EMAIL_R), api.SetIDEmailWrapper)
	router.GET(withPrefix(api.GET_USER_ID_R), api.GetUserIDWrapper)
	router.POST(withPrefix(api.ADD_FAVORITE_BOARD_R), api.AddFavoriteBoardWrapper)
	router.POST(withPrefix(api.ADD_FAVORITE_FOLDER_R), api.AddFavoriteFolderWrapper)
	router.POST(withPrefix(api.ADD_FAVORITE_LINE_R), api.AddFavoriteLineWrapper)
	router.POST(withPrefix(api.DELETE_FAVORITE_R), api.DeleteFavoriteWrapper)

	// comments
	router.GET(withPrefix(api.LOAD_ARTICLE_COMMENTS_R), api.LoadArticleCommentsWrapper)
	router.GET(withPrefix(api.LOAD_USER_COMMENTS_R), api.LoadUserCommentsWrapper)
	router.POST(withPrefix(api.CREATE_COMMENT_R), api.CreateCommentWrapper)

	// ranks
	router.POST(withPrefix(api.CREATE_RANK_R), api.CreateRankWrapper)

	// html
	router.GET(api.ROOT_HTML_R, api.IndexHTMLWrapper)
	router.GET(api.INDEX_HTML_R, api.IndexHTMLWrapper)

	router.GET(api.REGISTER_HTML_R, api.RegisterHTMLWrapper)
	router.GET(api.LOGIN_HTML_R, api.LoginHTMLWrapper)

	router.GET(api.USER_HTML_R, api.UserHTMLWrapper)
	router.GET(api.USER_CHANGE_PASSWD_HTML_R, api.UserChangePasswdHTMLWrapper)
	router.GET(api.USER_ATTEMPT_CHANGE_EMAIL_HTML_R, api.UserAttemptChangeEmailHTMLWrapper)
	router.GET(api.USER_CHANGE_EMAIL_HTML_R, api.UserChangeEmailHTMLWrapper)
	router.GET(api.USER_ATTEMPT_SET_ID_EMAIL_HTML_R, api.UserAttemptSetIDEmailHTMLWrapper)
	router.GET(api.USER_SET_ID_EMAIL_HTML_R, api.UserSetIDEmailHTMLWrapper)

	router.Static("/static", filepath.Join(types.STATIC_DIR, "static"))

	staticFiles := []string{
		"asset-manifest.json",
		"favicon.ico",
		"logo192.png",
		"logo512.png",
		"manifest.json",
		"robots.txt",
	}

	for _, each := range staticFiles {
		router.StaticFile("/"+each, filepath.Join(types.STATIC_DIR, each))
	}

	router.NoRoute(api.AllHTMLWrapper)

	return router, nil
}

func main() {
	err := initMain()
	if err != nil {
		logrus.Fatalf("unable to initMain: e: %v", err)
		return
	}

	if err := queue.Start(); err != nil {
		logrus.Fatal(err)
	}

	r, err := initGin()
	if err != nil {
		logrus.Fatal(err)
	}

	s := &http.Server{
		Addr:              types.HTTP_HOST,
		Handler:           r,
		ReadHeaderTimeout: time.Duration(180 * time.Second),
		ReadTimeout:       time.Duration(180 * time.Second),
	}

	g := graceful.NewManager()
	g.AddShutdownJob(func() error {
		queue.Close()
		return nil
	})

	if !types.IS_ALL_GUEST {
		// retry load general boards
		g.AddRunningJob(cron.RetryLoadGeneralBoards)
		// retry to calculate user visit count
		g.AddRunningJob(cron.RetryCalculateUserVisit)

		// retry load man articles
		g.AddRunningJob(cron.RetryLoadManArticles)

		// retry load man article details
		g.AddRunningJob(cron.RetryLoadManArticleDetails)

	} else {
		// retry load popular boards
		g.AddRunningJob(cron.RetryLoadPopularBoards)
	}

	// retry load general articles
	g.AddRunningJob(cron.RetryLoadGeneralArticles)

	// retry load article details
	g.AddRunningJob(cron.RetryLoadArticleDetails)

	g.AddRunningJob(func(ctx context.Context) error {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return s.Shutdown(ctx)
	})

	g.AddRunningJob(func(ctx context.Context) error {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	<-g.Done()
}
