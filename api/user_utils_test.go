package api

import (
	"fmt"
	"sync"
	"testing"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

func Test_serializeEmailToken(t *testing.T) {
	setupTest()
	defer teardownTest()

	jwt0, _ := pttbbsapi.CreateEmailToken("SYSOP", "", "test@ptt.test", pttbbsapi.CONTEXT_CHANGE_EMAIL)

	content0 := fmt.Sprintf("test@ptt.test, SYSOP, http://localhost:3457/bbs/user/SYSOP/changeemail?%v=%v", types.EMAIL_TOKEN_NAME, jwt0)

	type args struct {
		email           string
		userID          bbs.UUserID
		token           string
		urlTemplate     string
		contentTemplate string
	}
	tests := []struct {
		name            string
		args            args
		expectedContent string
	}{
		// TODO: Add test cases.
		{
			args:            args{email: "test@ptt.test", userID: "SYSOP", token: jwt0, urlTemplate: CHANGE_EMAIL_R, contentTemplate: types.EMAILTOKEN_TEMPLATE_CONTENT},
			expectedContent: content0,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if gotContent := serializeEmailToken(tt.args.email, tt.args.userID, tt.args.token, tt.args.urlTemplate, tt.args.contentTemplate); gotContent != tt.expectedContent {
				t.Errorf("serializeEmailToken() = %v, want %v", gotContent, tt.expectedContent)
			}
		})
	}
	wg.Wait()
}
