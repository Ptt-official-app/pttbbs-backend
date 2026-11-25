package api

import (
	"testing"

	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_loadPttWebBoardsGetIsOver18_cacheHitKeepsTimestamp(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer schema.Board_c.Drop()

	now := types.NowNanoTS()
	storedNanoTS := now - types.NanoTS(REFRESH_IS_OVER18_NANO_TS/2)

	board, err := schema.NewBoardSummaryFromPttWeb("test1", "title", "◎", "class", 1, 1, 1, true, storedNanoTS, storedNanoTS)
	require.NoError(t, err)
	err = schema.UpdateBoardSummaries([]*schema.BoardSummary{board}, storedNanoTS)
	require.NoError(t, err)

	// within the refresh window: served from the db, no request to pttweb.
	isOver18, isOver18UpdateNanoTS, err := loadPttWebBoardsGetIsOver18("test1", now, nil)
	require.NoError(t, err)
	assert.True(t, isOver18)
	assert.Equal(t, storedNanoTS, isOver18UpdateNanoTS)

	// the stored timestamp must survive the hourly board update, otherwise the cache never expires.
	board, err = schema.NewBoardSummaryFromPttWeb("test1", "title", "◎", "class", 1, 1, 1, isOver18, isOver18UpdateNanoTS, now)
	require.NoError(t, err)
	err = schema.UpdateBoardSummaries([]*schema.BoardSummary{board}, now)
	require.NoError(t, err)

	got, err := schema.GetBoardSummary("test1")
	require.NoError(t, err)
	require.NotNil(t, got)
	assert.Equal(t, storedNanoTS, got.IsOver18UpdateNanoTS)
	assert.Equal(t, now, got.UpdateNanoTS)
}
