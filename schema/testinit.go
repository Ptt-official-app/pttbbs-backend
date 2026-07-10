package schema

import (
	"context"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	IsTest = false
	lock   sync.Mutex
)

func SetIsTest() {
	lock.Lock()
	IsTest = true

	logrus.Info("schema: SetIsTest")

	MONGO_DBNAME = "devptt_test"
	err := Init()
	if err != nil {
		logrus.Errorf("schema.Init: unable to init: e: %v", err)
	}
}

func UnsetIsTest() {
	testResetDB()
	Close()

	logrus.Info("schema: UnsetIsTest")

	IsTest = false
	lock.Unlock()
}

func testResetDB() {
	_ = AccessToken_c.Drop()
	_ = Article_c.Drop()
	_ = Board_c.Drop()
	_ = BoardBanuser_c.Drop()
	_ = BoardFriend_c.Drop()
	_ = Client_c.Drop()
	_ = Comment_c.Drop()
	_ = User_c.Drop()
	_ = UserAloha_c.Drop()
	_ = UserFavorites_c.Drop()
	_ = UserFavoritesMeta_c.Drop()
	_ = UserFriend_c.Drop()
	_ = UserArticle_c.Drop()
	_ = UserBoard_c.Drop()
	_ = UserReject_c.Drop()
	_ = UserIDEmail_c.Drop()
	_ = UserEmail_c.Drop()
	_ = Rank_c.Drop()
	_ = UserVisit_c.Drop()
	_ = ContentBlock_c.Drop()
	_ = ManArticle_c.Drop()
	_ = ManContentBlock_c.Drop()

	ctx, cancel := context.WithTimeout(context.Background(), REDIS_TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		cancel()
	}()

	_, _ = rdb.FlushDB(ctx).Result()
}
