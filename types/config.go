package types

import (
	"os"
	"strings"
	"time"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

func config() {
	SERVICE_MODE = ServiceMode(setStringConfig("SERVICE_MODE", string(SERVICE_MODE)))

	HTTP_SCHEME = setStringConfig("HTTP_SCHEME", HTTP_SCHEME)
	HTTP_HOST = setStringConfig("HTTP_HOST", HTTP_HOST)
	URL_PREFIX = setStringConfig("URL_PREFIX", URL_PREFIX)
	GO_PTTBBS_PREFIX = setStringConfig("GO_PTTBBS_PREFIX", GO_PTTBBS_PREFIX)

	FRONTEND_PREFIX = setStringConfig("FRONTEND_PREFIX", FRONTEND_PREFIX)
	FRONTEND_LOGIN_URL = setStringConfig("FRONTEND_LOGIN_URL", FRONTEND_LOGIN_URL)
	FRONTEND_INIT_URL = setStringConfig("FRONTEND_INIT_URL", FRONTEND_INIT_URL)
	FRONTEND_ERR_URL = setStringConfig("FRONTEND_ERR_URL", FRONTEND_ERR_URL)

	API_PREFIX = setStringConfig("API_PREFIX", API_PREFIX)
	REGISTER_USER_URL = setStringConfig("REGISTER_USER_URL", REGISTER_USER_URL)
	ZK_PREFIX = setStringConfig("ZK_PREFIX", ZK_PREFIX)
	EMAIL_URL = setStringConfig("EMAIL_URL", EMAIL_URL)

	OIDC_OP_KEY = setStringConfig("OIDC_OP_KEY", OIDC_OP_KEY)
	OIDC_OP_KEY_ID = setStringConfig("OIDC_OP_KEY_ID", OIDC_OP_KEY_ID)
	OIDC_OP_POST_LOGOUT_URL = setStringConfig("OIDC_OP_POST_LOGOUT_URL", OIDC_OP_POST_LOGOUT_URL)
	OIDC_OP_ISSUER = setStringConfig("OIDC_OP_ISSUER", OIDC_OP_ISSUER)
	OIDC_OP_IS_ALLOW_HTTP = setBoolConfig("OIDC_OP_IS_ALLOW_HTTP", OIDC_OP_IS_ALLOW_HTTP)

	PTTSYSOP = bbs.UUserID(setStringConfig("PTTSYSOP", string(PTTSYSOP)))

	BBSNAME = setStringConfig("BBSNAME", BBSNAME)
	BBSNAME_EN = setStringConfig("BBSNAME_EN", BBSNAME_EN)
	SENDER_SUFFIX = setStringConfig("SENDER_SUFFIX", SENDER_SUFFIX)

	// web
	STATIC_DIR = setStringConfig("STATIC_DIR", STATIC_DIR)

	ALLOW_ORIGINS = setListStringConfig("ALLOW_ORIGINS", ALLOW_ORIGINS)
	BLOCKED_REFERERS = setListStringConfig("BLOCKED_REFERERS", BLOCKED_REFERERS)
	IS_ALLOW_CROSSDOMAIN = setBoolConfig("IS_ALLOW_CROSSDOMAIN", IS_ALLOW_CROSSDOMAIN)

	COOKIE_DOMAIN = setStringConfig("COOKIE_DOMAIN", COOKIE_DOMAIN)
	TOKEN_COOKIE_SUFFIX = setStringConfig("TOKEN_COOKIE_SUFFIX", TOKEN_COOKIE_SUFFIX)

	CSRF_SECRET = setBytesConfig("CSRF_SECRET", CSRF_SECRET)
	CSRF_TOKEN = setStringConfig("CSRF_TOKEN", CSRF_TOKEN)
	CSRF_TOKEN_TS = setIntConfig("CSRF_TOKEN_TS", CSRF_TOKEN_TS)
	CSRF_COOKIE_DOMAIN = setStringConfig("CSRF_COOKIE_DOMAIN", CSRF_COOKIE_DOMAIN)

	ACCESS_TOKEN_NAME = setStringConfig("ACCESS_TOKEN_NAME", ACCESS_TOKEN_NAME)
	ACCESS_TOKEN_EXPIRE_TS = setIntConfig("ACCESS_TOKEN_EXPIRE_TS", ACCESS_TOKEN_EXPIRE_TS)
	ACCESS_TOKEN_SECRET = setBytesConfig("ACCESS_TOKEN_SECRET", ACCESS_TOKEN_SECRET)

	REFRESH_TOKEN_EXPIRE_TS = setIntConfig("REFRESH_TOKEN_EXPIRE_TS", REFRESH_TOKEN_EXPIRE_TS)
	REFRESH_TOKEN_SECRET = setBytesConfig("REFRESH_TOKEN_SECRET", REFRESH_TOKEN_SECRET)

	IS_OVER_18_NAME = setStringConfig("IS_OVER_18_NAME", IS_OVER_18_NAME)
	IS_OVER_18_VALUE = setStringConfig("IS_OVER_18_VALUE", IS_OVER_18_VALUE)

	// email
	EMAIL_TOKEN_NAME = setStringConfig("EMAIL_TOKEN_NAME", EMAIL_TOKEN_NAME)

	EMAIL_FROM = setStringConfig("EMAIL_FROM", EMAIL_FROM)
	EMAIL_SERVER = setStringConfig("EMAIL_SERVER", EMAIL_SERVER)

	EMAILTOKEN_TITLE_TEMPLATE = setStringConfig("EMAILTOKEN_TITLE_TEMPLATE", EMAILTOKEN_TITLE_TEMPLATE)
	EMAILTOKEN_TEMPLATE = setStringConfig("EMAILTOKEN_TEMPLATE", EMAILTOKEN_TEMPLATE)
	IDEMAILTOKEN_TITLE_TEMPLATE = setStringConfig("IDEMAILTOKEN_TITLE_TEMPLATE", IDEMAILTOKEN_TITLE_TEMPLATE)
	IDEMAILTOKEN_TEMPLATE = setStringConfig("IDEMAILTOKEN_TEMPLATE", IDEMAILTOKEN_TEMPLATE)
	ATTEMPT_REGISTER_USER_TITLE_TEMPLATE = setStringConfig("ATTEMPT_REGISTER_USER_TITLE_TEMPLATE", ATTEMPT_REGISTER_USER_TITLE_TEMPLATE)
	ATTEMPT_REGISTER_USER_TEMPLATE = setStringConfig("ATTEMPT_REGISTER_USER_TEMPLATE", ATTEMPT_REGISTER_USER_TEMPLATE)
	ATTEMPT_LOGIN_TITLE_TEMPLATE = setStringConfig("ATTEMPT_LOGIN_TITLE_TEMPLATE", ATTEMPT_LOGIN_TITLE_TEMPLATE)
	ATTEMPT_LOGIN_TEMPLATE = setStringConfig("ATTEMPT_LOGIN_TEMPLATE", ATTEMPT_LOGIN_TEMPLATE)

	EXPIRE_USER_ID_EMAIL_IS_SET_NANO_TS = NanoTS(setInt64Config("EXPIRE_USER_ID_EMAIL_IS_SET_NANO_TS", int64(EXPIRE_USER_ID_EMAIL_IS_SET_NANO_TS)))
	EXPIRE_USER_EMAIL_IS_SET_NANO_TS = NanoTS(setInt64Config("EXPIRE_USER_EMAIL_IS_SET_NANO_TS", int64(EXPIRE_USER_EMAIL_IS_SET_NANO_TS)))

	EXPIRE_USER_ID_EMAIL_IS_NOT_SET_NANO_TS = NanoTS(setInt64Config("EXPIRE_USER_ID_EMAIL_IS_NOT_SET_NANO_TS", int64(EXPIRE_USER_ID_EMAIL_IS_NOT_SET_NANO_TS)))
	EXPIRE_USER_EMAIL_IS_NOT_SET_NANO_TS = NanoTS(setInt64Config("EXPIRE_USER_EMAIL_IS_NOT_SET_NANO_TS", int64(EXPIRE_USER_EMAIL_IS_NOT_SET_NANO_TS)))

	EXPIRE_ATTEMPT_REGISTER_USER_EMAIL_TS = setIntConfig("EXPIRE_ATTEMPT_REGISTER_USER_EMAIL_TS", EXPIRE_ATTEMPT_REGISTER_USER_EMAIL_TS)

	IS_2FA = setBoolConfig("IS_2FA", IS_2FA)
	MAX_2FA_TOKEN = setInt64Config("MAX_2FA_TOKEN", MAX_2FA_TOKEN)
	MAX_2FA_TOKEN_STR_PROMPT = setStringConfig("MAX_2FA_TOKEN_STR_PROMPT", MAX_2FA_TOKEN_STR_PROMPT)

	// big5
	BIG5_TO_UTF8 = setStringConfig("BIG5_TO_UTF8", BIG5_TO_UTF8)
	UTF8_TO_BIG5 = setStringConfig("UTF8_TO_BIG5", UTF8_TO_BIG5)
	AMBCJK = setStringConfig("AMBCJK", AMBCJK)

	// time-location
	TIME_LOCATION = setStringConfig("TIME_LOCATION", TIME_LOCATION)

	// carriage-return
	IS_CARRIAGE_RETURN = setBoolConfig("IS_CARRIAGE_RETURN", IS_CARRIAGE_RETURN)

	// is-all-guest
	IS_ALL_GUEST = setBoolConfig("IS_ALL_GUEST", IS_ALL_GUEST)

	// pttweb-hotboard-url
	PTTWEB_HOTBOARD_URL = setStringConfig("PTTWEB_HOTBOARD_URL", PTTWEB_HOTBOARD_URL)

	// pttweb-hotboard-url
	PTTWEB_BASE_URL = setStringConfig("PTTWEB_BASE_URL", PTTWEB_BASE_URL)

	// expire-http-request-ts
	EXPIRE_HTTP_REQUEST_TS = setIntConfig("EXPIRE_HTTP_REQUEST_TS", EXPIRE_HTTP_REQUEST_TS)

	MAX_POPULAR_BOARDS = setIntConfig("MAX_POPULAR_BOARDS", MAX_POPULAR_BOARDS)

	// brdname-white-list-map
	BRDNAME_WHITE_LIST_MAP_FILENAME = setStringConfig("BRDNAME_WHITE_LIST_MAP_FILENAME", BRDNAME_WHITE_LIST_MAP_FILENAME)

	BRDNAME_BLACK_LIST_MAP_FILENAME = setStringConfig("BRDNAME_BLACK_LIST_MAP_FILENAME", BRDNAME_BLACK_LIST_MAP_FILENAME)

	SLEEP_RETRY_LOAD_POPULAR_BOARDS_TS = setIntConfig("SLEEP_RETRY_LOAD_POPULAR_BOARDS_TS", SLEEP_RETRY_LOAD_POPULAR_BOARDS_TS)

	SLEEP_RETRY_LOAD_GENERAL_ARTICLES_TS = setIntConfig("SLEEP_RETRY_LOAD_GENERAL_ARTICLES_TS", SLEEP_RETRY_LOAD_GENERAL_ARTICLES_TS)

	SLEEP_RETRY_LOAD_ARTICLE_DETAILS_TS = setIntConfig("SLEEP_RETRY_LOAD_ARTICLE_DETAILS_TS", SLEEP_RETRY_LOAD_ARTICLE_DETAILS_TS)
}

func postConfig() (err error) {
	if _, err = setTimeLocation(TIME_LOCATION); err != nil {
		return err
	}
	if _, err = setAllowOrigins(ALLOW_ORIGINS); err != nil {
		return err
	}
	if _, err = setBlockedReferers(BLOCKED_REFERERS); err != nil {
		return err
	}
	if _, err = setCSRFTokenTS(CSRF_TOKEN_TS); err != nil {
		return err
	}
	if _, err = setAccessTokenExpireTS(ACCESS_TOKEN_EXPIRE_TS); err != nil {
		return err
	}
	if _, err = setRefreshTokenExpireTS(REFRESH_TOKEN_EXPIRE_TS); err != nil {
		return err
	}

	if _, err = setBBSName(BBSNAME); err != nil {
		return err
	}

	if _, err = setBBSNameEN(BBSNAME_EN); err != nil {
		return err
	}

	if err = replaceTemplates(); err != nil {
		return err
	}

	if _, err = setAttemptRegisterUserEmailTS(EXPIRE_ATTEMPT_REGISTER_USER_EMAIL_TS); err != nil {
		return err
	}

	if _, err = setSleepRetryLoadPopularBoardsTS(SLEEP_RETRY_LOAD_POPULAR_BOARDS_TS); err != nil {
		return err
	}

	if _, err = setSleepRetryLoadGeneralArticlesTS(SLEEP_RETRY_LOAD_GENERAL_ARTICLES_TS); err != nil {
		return err
	}

	if _, err = setSleepRetryLoadArticleDetailsTS(SLEEP_RETRY_LOAD_ARTICLE_DETAILS_TS); err != nil {
		return err
	}

	if err = initBig5(); err != nil {
		return err
	}

	return nil
}

// setTimeLocation
func setTimeLocation(timeLocation string) (origTimeLocation string, err error) {
	origTimeLocation = TIME_LOCATION
	TIME_LOCATION = timeLocation

	TIMEZONE, err = time.LoadLocation(TIME_LOCATION)

	return origTimeLocation, err
}

func setAllowOrigins(allowOrigins []string) (origAllowOrigins []string, err error) {
	origAllowOrigins = ALLOW_ORIGINS

	ALLOW_ORIGINS = allowOrigins

	return origAllowOrigins, nil
}

func setBlockedReferers(blockedReferers []string) (origBlockedReferers []string, err error) {
	origBlockedReferers = BLOCKED_REFERERS

	BLOCKED_REFERERS = blockedReferers

	return origBlockedReferers, nil
}

func setCSRFTokenTS(csrfTokenTS int) (origCSRFTokenTS int, err error) {
	origCSRFTokenTS = CSRF_TOKEN_TS

	CSRF_TOKEN_TS = csrfTokenTS

	CSRF_TOKEN_TS_DURATION = time.Duration(CSRF_TOKEN_TS) * time.Second

	return origCSRFTokenTS, nil
}

func setAccessTokenExpireTS(accessTokenExpireTS int) (origAccessTokenExpireTS int, err error) {
	origAccessTokenExpireTS = ACCESS_TOKEN_EXPIRE_TS

	ACCESS_TOKEN_EXPIRE_TS = accessTokenExpireTS

	ACCESS_TOKEN_EXPIRE_TS_DURATION = time.Duration(ACCESS_TOKEN_EXPIRE_TS) * time.Second

	return origAccessTokenExpireTS, nil
}

func setRefreshTokenExpireTS(refreshTokenExpireTS int) (origRefreshTokenExpireTS int, err error) {
	origRefreshTokenExpireTS = REFRESH_TOKEN_EXPIRE_TS

	REFRESH_TOKEN_EXPIRE_TS = refreshTokenExpireTS

	REFRESH_TOKEN_EXPIRE_TS_DURATION = time.Duration(REFRESH_TOKEN_EXPIRE_TS) * time.Second

	return origRefreshTokenExpireTS, nil
}

func setAttemptRegisterUserEmailTS(expireAttemptRegisterUserEmailTS int) (origExpireAttemptRegisterUserEmailTS int, err error) {
	origExpireAttemptRegisterUserEmailTS = EXPIRE_ATTEMPT_REGISTER_USER_EMAIL_TS
	EXPIRE_ATTEMPT_REGISTER_USER_EMAIL_TS = expireAttemptRegisterUserEmailTS

	EXPIRE_ATTEMPT_REGISTER_USER_EMAIL_TS_DURATION = time.Duration(EXPIRE_ATTEMPT_REGISTER_USER_EMAIL_TS) * time.Second

	return origExpireAttemptRegisterUserEmailTS, nil
}

func setEmailTokenTitleTemplate(emailTokenTitleTemplate string) (origEmailTokenTitleTemplate string, err error) {
	origEmailTokenTitleTemplate = EMAILTOKEN_TITLE_TEMPLATE
	EMAILTOKEN_TITLE_TEMPLATE = emailTokenTitleTemplate

	EMAILTOKEN_TITLE = replaceTemplate(EMAILTOKEN_TITLE_TEMPLATE)

	return origEmailTokenTitleTemplate, nil
}

func setEmailTokenTemplate(emailTokenTemplate string) (origEmailTokenTemplate string, err error) {
	origEmailTokenTemplate = EMAILTOKEN_TEMPLATE
	EMAILTOKEN_TEMPLATE = emailTokenTemplate

	contentBytes, err := os.ReadFile(EMAILTOKEN_TEMPLATE)
	if err != nil {
		return "", err
	}

	EMAILTOKEN_TEMPLATE_CONTENT = replaceTemplate(string(contentBytes))

	return origEmailTokenTemplate, nil
}

func setIDEmailTokenTitleTemplate(idEmailTokenTitleTemplate string) (origIDEmailTokenTitleTemplate string, err error) {
	origIDEmailTokenTitleTemplate = IDEMAILTOKEN_TITLE_TEMPLATE
	IDEMAILTOKEN_TITLE_TEMPLATE = idEmailTokenTitleTemplate

	IDEMAILTOKEN_TITLE = replaceTemplate(IDEMAILTOKEN_TITLE_TEMPLATE)

	return origIDEmailTokenTitleTemplate, nil
}

func setIDEmailTokenTemplate(idEmailTokenTemplate string) (origIDEmailTokenTemplate string, err error) {
	origIDEmailTokenTemplate = IDEMAILTOKEN_TEMPLATE
	IDEMAILTOKEN_TEMPLATE = idEmailTokenTemplate

	contentBytes, err := os.ReadFile(IDEMAILTOKEN_TEMPLATE)
	if err != nil {
		return "", err
	}

	IDEMAILTOKEN_TEMPLATE_CONTENT = replaceTemplate(string(contentBytes))

	return origIDEmailTokenTemplate, nil
}

func setAttemptRegisterUserTitleTemplate(attemptRegisterUserTitleTemplate string) (origAttemptRegisterUserTitleTemplate string, err error) {
	origAttemptRegisterUserTitleTemplate = ATTEMPT_REGISTER_USER_TITLE_TEMPLATE
	ATTEMPT_REGISTER_USER_TITLE_TEMPLATE = attemptRegisterUserTitleTemplate

	ATTEMPT_REGISTER_USER_TITLE = replaceTemplate(ATTEMPT_REGISTER_USER_TITLE_TEMPLATE)

	return origAttemptRegisterUserTitleTemplate, nil
}

func setAttemptRegisterUserTemplate(attemptRegisterUserTemplate string) (origAttemptRegisterUserTemplate string, err error) {
	origAttemptRegisterUserTemplate = ATTEMPT_REGISTER_USER_TEMPLATE
	ATTEMPT_REGISTER_USER_TEMPLATE = attemptRegisterUserTemplate

	contentBytes, err := os.ReadFile(ATTEMPT_REGISTER_USER_TEMPLATE)
	if err != nil {
		return "", err
	}

	ATTEMPT_REGISTER_USER_TEMPLATE_CONTENT = replaceTemplate(string(contentBytes))

	return origAttemptRegisterUserTemplate, nil
}

func setAttemptLoginTitleTemplate(attemptLoginTitleTemplate string) (origAttemptLoginTitleTemplate string, err error) {
	origAttemptLoginTitleTemplate = ATTEMPT_LOGIN_TITLE_TEMPLATE
	ATTEMPT_LOGIN_TITLE_TEMPLATE = attemptLoginTitleTemplate

	ATTEMPT_LOGIN_TITLE = replaceTemplate(ATTEMPT_LOGIN_TITLE_TEMPLATE)

	return origAttemptLoginTitleTemplate, nil
}

func setAttemptLoginTemplate(attemptLoginTemplate string) (origAttemptLoginTemplate string, err error) {
	origAttemptLoginTemplate = ATTEMPT_LOGIN_TEMPLATE
	ATTEMPT_LOGIN_TEMPLATE = attemptLoginTemplate

	contentBytes, err := os.ReadFile(ATTEMPT_LOGIN_TEMPLATE)
	if err != nil {
		return "", err
	}

	ATTEMPT_LOGIN_TEMPLATE_CONTENT = replaceTemplate(string(contentBytes))

	return origAttemptLoginTemplate, nil
}

func setBBSName(bbsname string) (origBBSName string, err error) {
	origBBSName = BBSNAME
	BBSNAME = bbsname

	err = replaceTemplates()
	if err != nil {
		return "", err
	}

	return origBBSName, nil
}

func setBBSNameEN(bbsnameEN string) (origBBSNameEN string, err error) {
	origBBSNameEN = BBSNAME_EN
	BBSNAME_EN = bbsnameEN

	err = replaceTemplates()
	if err != nil {
		return "", err
	}

	return origBBSNameEN, nil
}

func setSleepRetryLoadPopularBoardsTS(sleepRetryLoadPoluarBoardsTS int) (origSleepRetryLoadPoluarBoardsTS int, err error) {
	origSleepRetryLoadPoluarBoardsTS = SLEEP_RETRY_LOAD_POPULAR_BOARDS_TS
	SLEEP_RETRY_LOAD_POPULAR_BOARDS_TS = sleepRetryLoadPoluarBoardsTS

	SLEEP_RETRY_LOAD_POPULAR_BOARDS_TS_DURATION = time.Duration(SLEEP_RETRY_LOAD_POPULAR_BOARDS_TS) * time.Second

	return origSleepRetryLoadPoluarBoardsTS, nil
}

func setSleepRetryLoadGeneralArticlesTS(sleepRetryLoadGeneralArticlesTS int) (origSleepRetryLoadGeneralArticlesTS int, err error) {
	origSleepRetryLoadGeneralArticlesTS = SLEEP_RETRY_LOAD_GENERAL_ARTICLES_TS
	SLEEP_RETRY_LOAD_GENERAL_ARTICLES_TS = sleepRetryLoadGeneralArticlesTS

	SLEEP_RETRY_LOAD_GENERAL_ARTICLES_TS_DURATION = time.Duration(SLEEP_RETRY_LOAD_GENERAL_ARTICLES_TS) * time.Second

	return origSleepRetryLoadGeneralArticlesTS, nil
}

func setSleepRetryLoadArticleDetailsTS(sleepRetryLoadArticleDetailsTS int) (origSleepRetryLoadArticleDetailsTS int, err error) {
	origSleepRetryLoadArticleDetailsTS = SLEEP_RETRY_LOAD_ARTICLE_DETAILS_TS
	SLEEP_RETRY_LOAD_ARTICLE_DETAILS_TS = sleepRetryLoadArticleDetailsTS

	SLEEP_RETRY_LOAD_ARTICLE_DETAILS_TS_DURATION = time.Duration(SLEEP_RETRY_LOAD_ARTICLE_DETAILS_TS) * time.Second

	return origSleepRetryLoadArticleDetailsTS, nil
}

func replaceTemplates() (err error) {
	if _, err = setEmailTokenTitleTemplate(EMAILTOKEN_TITLE_TEMPLATE); err != nil {
		return err
	}

	if _, err = setEmailTokenTemplate(EMAILTOKEN_TEMPLATE); err != nil {
		return err
	}

	if _, err = setIDEmailTokenTitleTemplate(EMAILTOKEN_TITLE_TEMPLATE); err != nil {
		return err
	}

	_, err = setIDEmailTokenTemplate(IDEMAILTOKEN_TEMPLATE)
	if err != nil {
		return err
	}

	_, err = setAttemptRegisterUserTitleTemplate(ATTEMPT_REGISTER_USER_TITLE_TEMPLATE)
	if err != nil {
		return err
	}
	_, err = setAttemptRegisterUserTemplate(ATTEMPT_REGISTER_USER_TEMPLATE)
	if err != nil {
		return err
	}

	_, err = setAttemptLoginTitleTemplate(ATTEMPT_LOGIN_TITLE_TEMPLATE)
	if err != nil {
		return err
	}

	_, err = setAttemptLoginTemplate(ATTEMPT_LOGIN_TEMPLATE)
	if err != nil {
		return err
	}

	return nil
}

func replaceTemplate(template string) (content string) {
	return strings.ReplaceAll(
		strings.ReplaceAll(
			template, TEMPLATE_BBSNAME, BBSNAME,
		), TEMPLATE_BBSNAME_EN, BBSNAME_EN,
	)
}
