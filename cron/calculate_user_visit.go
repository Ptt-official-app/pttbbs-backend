package cron

import (
	"context"
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/schema"

	"github.com/sirupsen/logrus"
)

// RetryCalculateUserVisit make loop job to call CalculateUserVisit per 10 mins
func RetryCalculateUserVisit(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			logrus.Infof("RetryCalculateUserVisit: to calculate user visit")
			CalculateUserVisit()
			logrus.Infof("RetryCalculateUserVisit: to sleep 10 mins")
			if !sleepCtx(ctx, 10*time.Minute) {
				return nil
			}
		}
	}
}

// CalculateUserVisit get user visit count from db
// and set to redis
func CalculateUserVisit() {
	count, err := schema.CalculateAllUserVisitCounts()
	if err != nil {
		logrus.Printf("get error in calculate user visit count %v", err)
		return
	}
	// set to redis
	err = schema.SetUserVisitCount(count)
	if err != nil {
		logrus.Printf("set to redis  %v", err)
	}
}
