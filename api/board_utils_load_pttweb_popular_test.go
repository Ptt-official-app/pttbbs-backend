package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTryLoadPttWebPopularBoards_fetchFailureKeepsPopularBoards(t *testing.T) {
	tests := []struct {
		name    string
		handler http.HandlerFunc
	}{
		{
			name: "server error",
			handler: func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "boom", http.StatusInternalServerError)
			},
		},
		{
			name: "page without boards",
			handler: func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write([]byte("<html><body>maintenance</body></html>"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupTest()
			defer teardownTest()

			defer schema.Board_c.Drop()

			server := httptest.NewServer(tt.handler)
			defer server.Close()

			origURL := types.PTTWEB_HOTBOARD_URL
			types.PTTWEB_HOTBOARD_URL = server.URL
			defer func() { types.PTTWEB_HOTBOARD_URL = origURL }()

			now := types.NowNanoTS()
			board, err := schema.NewBoardSummaryFromPttWeb("test1", "title", "◎", "class", 1, 1, 1, false, now)
			require.NoError(t, err)
			err = schema.UpdateBoardSummaries([]*schema.BoardSummary{board}, now)
			require.NoError(t, err)

			_, err = TryLoadPttWebPopularBoards(nil)
			require.Error(t, err)

			got, err := schema.GetPopularBoardSummaries()
			require.NoError(t, err)
			assert.Len(t, got, 1)
		})
	}
}
