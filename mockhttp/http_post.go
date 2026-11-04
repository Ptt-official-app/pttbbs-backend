package mockhttp

import (
	"reflect"

	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/sirupsen/logrus"
)

func HTTPPost(url string, data any, result any) (statusCode int, err error) {
	logrus.Infof("HTTPPost: url: %v", url)
	switch url {
	case api.LOGIN_R:
		return parseResult(Login(data.(*api.LoginParams)), result)
	case api.REGISTER_R:
		return parseResult(Register(data.(*api.RegisterParams)), result)
	case api.LOAD_GENERAL_BOARDS_R:
		return parseResult(LoadGeneralBoards(data.(*api.LoadGeneralBoardsParams)), result)
	case api.LOAD_GENERAL_BOARD_DETAILS_R:
		return parseResult(LoadGeneralBoardDetails(data.(*api.LoadGeneralBoardDetailsParams)), result)
	case api.LOAD_AUTO_COMPLETE_BOARDS_R:
		return parseResult(LoadAutoCompleteBoards(data.(*api.LoadAutoCompleteBoardsParams)), result)
	case api.LOAD_GENERAL_BOARDS_BY_CLASS_R:
		return parseResult(LoadGeneralBoardsByClass(data.(*api.LoadGeneralBoardsParams)), result)
	case api.LOAD_HOT_BOARDS_R:
		return parseResult(LoadHotBoards(), result)
	case "/board/10_WhoAmI/article/1VtWRel9":
		return parseResult(GetArticleDetail(data.(*api.GetArticleParams)), result)
	case "/board/10_WhoAmI/article/1VrooM21":
		return parseResult(GetArticleDetail2(data.(*api.GetArticleParams)), result)
	case "/board/10_WhoAmI/article/1VrooM21/edit":
		return parseResult(EditArticleDetail2(data.(*api.EditArticleParams)), result)
	case "/board/10_WhoAmI/article/1VrooM22":
		return parseResult(GetArticleDetail2(data.(*api.GetArticleParams)), result)
	case "/board/10_WhoAmI/article/1VrooM22/edit":
		return parseResult(EditArticleDetail2(data.(*api.EditArticleParams)), result)
	case "/board/10_WhoAmI/article/1VrooM22/comment":
		return parseResult(CreateComment(data.(*api.CreateCommentParams)), result)
	case "/board/10_WhoAmI/article/1Vo_N0CD": // M.1607202240.A.30D
		return parseResult(GetArticleDetail4(data.(*api.GetArticleParams)), result)
	case "/board/10_WhoAmI/article":
		return parseResult(CreateArticle(data.(*api.CreateArticleParams)), result)
	case "/board/10_WhoAmI/articles":
		return parseResult(LoadGeneralArticles2(data.(*api.LoadGeneralArticlesParams)), result)
	case "/board/1_SYSOP/article":
		return parseResult(CreateArticle2(data.(*api.CreateArticleParams)), result)
	case "/board/1_SYSOP/article/1VrooM21":
		return parseResult(GetArticleDetail3(data.(*api.GetArticleParams)), result)
	case "/board/1_SYSOP/article/1VrooM21/edit":
		return parseResult(EditArticleDetail3(data.(*api.EditArticleParams)), result)
	case "/board/1_SYSOP/article/1VrooM21/comment":
		return parseResult(CreateComment2(data.(*api.CreateCommentParams)), result)
	case "/board/1_SYSOP/article/1VrooM23":
		return parseResult(GetArticleDetail3(data.(*api.GetArticleParams)), result)
	case "/board/1_SYSOP/article/1VrooM23/edit":
		return parseResult(EditArticleDetail3(data.(*api.EditArticleParams)), result)
	case "/board/1_SYSOP/article/1VrooM23/comment":
		return parseResult(CreateComment2(data.(*api.CreateCommentParams)), result)
	case "/board/1_SYSOP/articles":
		return parseResult(LoadGeneralArticles3(data.(*api.LoadGeneralArticlesParams)), result)
	case "/board/1_SYSOP/article/1VtWRel9/crosspost":
		return parseResult(CrossPost(data.(*api.CrossPostParams)), result)
	case "/board/10_WhoAmI/deletearticles":
		return parseResult(DeleteArticles(data.(*api.DeleteArticlesParams)), result)
	case "/board/10_WhoAmI/articles/bottom":
		return parseResult(LoadBottomArticles(nil), result)
	case "/user/SYSOP/information":
		return parseResult(GetUser(), result)
	case "/user/SYSOP/changepasswd":
		return parseResult(ChangePasswd(data.(*api.ChangePasswdParams)), result)
	case "/user/SYSOP/attemptchangeemail":
		return parseResult(AttemptChangeEmail(data.(*api.AttemptChangeEmailParams)), result)
	case "/user/SYSOP/changeemail":
		return parseResult(ChangeEmail(data.(*api.ChangeEmailParams)), result)
	case "/user/SYSOP/attemptsetidemail":
		return parseResult(AttemptSetIDEmail(data.(*api.AttemptSetIDEmailParams)), result)
	case "/user/SYSOP/setidemail":
		return parseResult(SetIDEmail(data.(*api.SetIDEmailParams)), result)
	case "/emailtoken/info":
		return parseResult(GetEmailTokenInfo(data.(*api.GetEmailTokenInfoParams)), result)
	case "/user/SYSOP/favorites":
		return parseResult(GetFavorites(data.(*api.GetFavoritesParams)), result)
	case "/user/SYSOP/favorites/post":
		return parseResult(WriteFavorites(data.(*api.WriteFavoritesParams)), result)
	case api.LOAD_BOARDS_BY_BIDS_R:
		return parseResult(LoadBoardsByBids(data.(*api.LoadBoardsByBidsParams)), result)
	case api.CHECK_EXISTS_USER_R:
		return parseResult(CheckExistsUser(data.(*api.CheckExistsUserParams)), result)
	case "/board/10_WhoAmI/isvalid":
		return parseResult(IsBoardValidUser(), result)
	case "/board/1_test1/isvalid":
		return parseResult(IsBoardValidUser(), result)
	case "/board/1_test1/summary":
		return parseResult(GetBoardSummary(data.(*api.LoadBoardSummaryParams)), result)
	case "/class/2/board":
		return parseResult(CreateBoard(data.(*api.CreateBoardParams)), result)
	case "/board/2_test2/isvalid":
		return parseResult(IsBoardValidUser(), result)
	case "/board/3_3........../isvalid":
		return parseResult(IsBoardValidUser(), result)
	case "/board/10_WhoAmI/article/1VrooM21/comment":
		return parseResult(CreateComment(data.(*api.CreateCommentParams)), result)
	case "/boards/isvalid":
		return parseResult(IsBoardsValidUser(data.(*api.IsBoardsValidUserParams)), result)
	case api.LOAD_FULL_CLASS_BOARDS_R:
		return parseResult(LoadFullClassBoards(data.(*api.LoadFullClassBoardsParams)), result)
	case api.GET_USER_VISIT_COUNT_R:
		return parseResult(GetUserVisitCount(), result)
	default:
		return 500, ErrURL
	}
}

func parseResult(backendResult any, httpResult any) (statusCode int, err error) {
	convert(backendResult, httpResult)

	return 200, nil
}

func convert(dataBackend any, dataResult any) {
	valueBackend := reflect.ValueOf(dataResult)
	valuePttbbs := reflect.ValueOf(dataBackend)
	valueBackend.Elem().Set(valuePttbbs)
}
