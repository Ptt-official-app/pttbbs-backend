package cron

import (
	"context"
	"runtime/debug"
	"time"

	"github.com/sirupsen/logrus"
)

// sleepCtx sleeps for d and returns true, or returns false as soon as ctx is done,
// so that the cron loops do not block graceful shutdown.
func sleepCtx(ctx context.Context, d time.Duration) bool {
	timer := time.NewTimer(d)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return false
	case <-timer.C:
		return true
	}
}

// runRecovered runs one round of a cron loop and recovers from panic,
// so that a panic only fails that round instead of ending the cron permanently.
func runRecovered(name string, fn func() error) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("%v: Recovered r: %v stack: %v", name, r, string(debug.Stack()))
		}
	}()

	_ = fn()
}
