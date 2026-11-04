package queue

import (
	"cmp"
	"slices"
	"sync"
	"testing"
	"time"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/dbcs"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func TestQueueCommentDBCSStr(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		bboardID          bbs.BBoardID
		articleID         bbs.ArticleID
		ownerID           bbs.UUserID
		commentDBCS       string
		articleCreateTime types.NanoTS
		articleMTime      types.NanoTS
		updateNanoTS      types.NanoTS
	}
	tests := []struct {
		name     string
		args     args
		expected []*schema.Comment
		sleepTS  int
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				bboardID:          "test",
				articleID:         "test7",
				ownerID:           "testOwner",
				commentDBCS:       string(testUtf8Comment7),
				articleCreateTime: types.NanoTS(1516958302000000000),
				articleMTime:      types.NanoTS(1608567097000000000),
				updateNanoTS:      types.NanoTS(1688888888000000001),
			},
			expected: testUtf8FullComments7,
			sleepTS:  1,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()

			firstComments, _, _, _ := dbcs.ParseFirstCommentsStr(tt.args.bboardID, tt.args.articleID, tt.args.ownerID, tt.args.articleCreateTime, tt.args.articleMTime, tt.args.commentDBCS, "")

			logrus.Infof("TestQueueCommentDBCSStr: to update comments: firstComments: %v", firstComments)
			schema.UpdateComments(firstComments, 1688888888000000000)

			if err := QueueCommentDBCSStr(tt.args.bboardID, tt.args.articleID, tt.args.ownerID, tt.args.commentDBCS, tt.args.articleCreateTime, tt.args.articleMTime, tt.args.updateNanoTS); (err != nil) != tt.wantErr {
				t.Errorf("QueueCommentDBCSStr() error = %v, wantErr %v", err, tt.wantErr)
			}

			time.Sleep(time.Duration(tt.sleepTS * int(time.Second)))

			query := bson.M{
				schema.COMMENT_BBOARD_ID_b:  tt.args.bboardID,
				schema.COMMENT_ARTICLE_ID_b: tt.args.articleID,
			}
			var got []*schema.Comment
			_ = schema.Comment_c.Find(query, 0, &got, nil, nil)

			slices.SortStableFunc(got, func(a, b *schema.Comment) int {
				return cmp.Compare(a.SortTime, b.SortTime)
			})
			for _, each := range got {
				each.UpdateNanoTS = types.NanoTS(0)
			}
			testutil.TDeepEqual(t, "got", got, tt.expected)
		})
		wg.Wait()
	}
}
