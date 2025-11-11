//go:build queue

package queue

import (
	"sort"
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

func TestQueueCommentDBCS(t *testing.T) {
	setupTest()
	defer teardownTest()

	schema.Comment_c.Drop()
	defer schema.Comment_c.Drop()

	// move setupTest inside
	type args struct {
		bboardID          bbs.BBoardID
		articleID         bbs.ArticleID
		ownerID           bbs.UUserID
		commentDBCS       []byte
		articleCreateTime types.NanoTS
		articleMTime      types.NanoTS
		updateNanoTS      types.NanoTS
	}
	tests := []struct {
		name     string
		args     args
		expected []*schema.Comment
		wantErr  bool
		sleepTS  int
	}{
		// TODO: Add test cases.
		{
			args: args{
				bboardID:          "test",
				articleID:         "test11",
				ownerID:           "testOwner",
				commentDBCS:       testComment11,
				articleCreateTime: types.NanoTS(1608551120000000000),
				articleMTime:      types.NanoTS(1608567097000000000),
				updateNanoTS:      types.NanoTS(1688888888000000001),
			},
			expected: testFullComments11,
		},
		/*
			{
				args: args{
					bboardID:          "test",
					articleID:         "test11",
					ownerID:           "testOwner",
					commentDBCS:       testComment11,
					articleCreateTime: types.NanoTS(1608551120000000000),
					articleMTime:      types.NanoTS(1608567097000000000),
					updateNanoTS:      types.NanoTS(1688888888000000002),
				},
				expected: testFullTheRestComments11,
			},
		*/
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()

			firstComments, _, _, _ := dbcs.ParseFirstComments(tt.args.bboardID, tt.args.articleID, "testOwner", tt.args.articleCreateTime, tt.args.articleMTime, tt.args.commentDBCS, "")

			logrus.Infof("QueueCommentDBCS: firstComments: %v", len(firstComments))

			schema.UpdateComments(firstComments, 1688888888000000000)

			if err := QueueCommentDBCS(tt.args.bboardID, tt.args.articleID, tt.args.ownerID, tt.args.commentDBCS, tt.args.articleCreateTime, tt.args.articleMTime, tt.args.updateNanoTS); (err != nil) != tt.wantErr {
				t.Errorf("QueueCommentDBCS() error = %v, wantErr %v", err, tt.wantErr)
			}

			time.Sleep(time.Duration(5) * time.Second)

			query := bson.M{
				schema.COMMENT_BBOARD_ID_b:  tt.args.bboardID,
				schema.COMMENT_ARTICLE_ID_b: tt.args.articleID,
			}
			var got []*schema.Comment
			_ = schema.Comment_c.Find(query, 0, &got, nil, nil)

			sort.SliceStable(got, func(i, j int) bool {
				return got[i].SortTime <= got[j].SortTime
			})
			for _, each := range got {
				each.UpdateNanoTS = types.NanoTS(0)
			}
			testutil.TDeepEqual(t, "got", got, tt.expected)
		})
		wg.Wait()
	}
}
