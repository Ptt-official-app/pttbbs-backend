package cron

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_sleepCtx(t *testing.T) {
	assert.True(t, sleepCtx(context.Background(), time.Millisecond))

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	start := time.Now()
	assert.False(t, sleepCtx(ctx, time.Hour))
	assert.Less(t, time.Since(start), time.Second)
}

func Test_runRecovered(t *testing.T) {
	assert.NotPanics(t, func() {
		runRecovered("test", func() error { panic("boom") })
	})

	called := false
	runRecovered("test", func() error { called = true; return errors.New("e") })
	assert.True(t, called)
}

func TestRetryLoops_shutdown(t *testing.T) {
	setupTest()
	defer teardownTest()

	// RetryLoadGeneralBoards is not covered: the mocked backend always returns a next idx, so one round never ends.
	tests := []struct {
		name string
		fn   func(ctx context.Context) error
	}{
		{name: "RetryCalculateUserVisit", fn: RetryCalculateUserVisit},
		{name: "RetryLoadGeneralArticles", fn: RetryLoadGeneralArticles},
		{name: "RetryLoadArticleDetails", fn: RetryLoadArticleDetails},
		{name: "RetryLoadManArticles", fn: RetryLoadManArticles},
		{name: "RetryLoadManArticleDetails", fn: RetryLoadManArticleDetails},
		{name: "RetryLoadPopularBoards", fn: RetryLoadPopularBoards},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			done := make(chan error, 1)
			go func() { done <- tt.fn(ctx) }()

			// let the loop enter its sleep, then shut down.
			time.Sleep(500 * time.Millisecond)
			cancel()

			select {
			case err := <-done:
				assert.NoError(t, err)
			case <-time.After(3 * time.Second):
				t.Fatal("cron loop did not stop after ctx is canceled")
			}
		})
	}
}
