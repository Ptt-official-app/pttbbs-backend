package api

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/boardd"
	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/mand"
	"github.com/Ptt-official-app/pttbbs-backend/mockhttp"
	"github.com/Ptt-official-app/pttbbs-backend/queue"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/go-querystring/query"
	"github.com/sirupsen/logrus"
)

var testIP = "127.0.0.1"

func setupTest() {
	mockhttp.SetIsTest()

	utils.SetIsTest()

	db.SetIsTest()

	types.SetIsTest("api")

	schema.SetIsTest()

	queue.SetIsTest()

	boardd.SetIsTest()

	mand.SetIsTest()

	SetIsTest()

	initTest()

	params := &RegisterClientParams{ClientID: "default_client_id", ClientType: types.CLIENT_TYPE_APP}
	user := &UserInfo{UserID: bbs.UUserID("SYSOP"), IsOver18: true}
	_, _, _ = RegisterClient("localhost", user, params, nil)
	// logrus.Infof("api.setupTest: after RegisterClient: status: %v e: %v", statusCode, err)
}

func teardownTest() {
	defer mockhttp.UnsetIsTest()

	defer utils.UnsetIsTest()

	defer db.UnsetIsTest()

	defer types.UnsetIsTest("api")

	defer schema.UnsetIsTest()

	defer queue.UnsetIsTest()

	defer boardd.UnsetIsTest()

	defer mand.UnsetIsTest()

	defer UnsetIsTest()
}

func testSetRequest(reqPath string, pathPattern string, params any, jwt string, csrfToken string, headers map[string]string, method string, f gin.HandlerFunc) (*httptest.ResponseRecorder, *gin.Context, *gin.Engine) {
	var jsonBytes []byte

	if method == "GET" {
		v, _ := query.Values(params)
		reqPath = reqPath + "?" + v.Encode()
	} else {
		jsonBytes, _ = json.Marshal(params)
	}

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	switch method {
	case "POST":
		r.POST(pathPattern, f)
	case "GET":
		r.GET(pathPattern, f)
	}

	req := httptest.NewRequest(method, reqPath, bytes.NewBuffer(jsonBytes))

	logrus.Infof("testSetRequest: method: %v reqPath: %v pathPattern: %v f: %v", method, reqPath, pathPattern, f)

	req.Header.Set("Host", "localhost:5678")
	if jwt != "" {
		req.Header.Set("Authorization", "bearer "+jwt)
	}

	if csrfToken != "" {
		req.Header.Set("X-CSRFToken", csrfToken)
		req.Header.Set("Cookie", "csrftoken="+csrfToken)
	}

	for key, val := range headers {
		req.Header.Set(key, val)
	}

	c.Request = req

	return w, c, r
}
