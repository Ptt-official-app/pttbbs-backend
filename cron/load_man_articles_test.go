package cron

import (
	"context"
	"testing"
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/mand"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

// hangingManClient blocks List until the request context is done.
type hangingManClient struct {
	mand.ManServiceClient
}

func (c *hangingManClient) List(ctx context.Context, _ *mand.ListRequest, _ ...grpc.CallOption) (*mand.ListReply, error) {
	<-ctx.Done()
	return nil, ctx.Err()
}

func Test_loadManArticlesCore_listTimeout(t *testing.T) {
	setupTest()
	defer teardownTest()

	oldCli := mand.Cli
	oldTS := types.EXPIRE_HTTP_REQUEST_TS
	defer func() {
		mand.Cli = oldCli
		types.EXPIRE_HTTP_REQUEST_TS = oldTS
	}()
	mand.Cli = &hangingManClient{}
	types.EXPIRE_HTTP_REQUEST_TS = 1

	done := make(chan error, 1)
	go func() {
		_, err := loadManArticlesCore("10_WhoAmI", "")
		done <- err
	}()

	select {
	case err := <-done:
		assert.ErrorIs(t, err, context.DeadlineExceeded)
	case <-time.After(5 * time.Second):
		t.Fatal("mand.Cli.List has no deadline")
	}
}
