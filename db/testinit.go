package db

import (
	"sync"
	"time"
)

// The default 10 ms timeout is meant for production; a local Redis (e.g. in a docker VM)
// sometimes exceeds it, making tests fail with context deadline exceeded.
const TEST_REDIS_TIMEOUT_MILLI_TS = time.Duration(1000)

var (
	isTest = false
	lock   sync.Mutex

	origRedisTimeoutMilliTS time.Duration
)

func setupTest() {
	SetIsTest()
}

func teardownTest() {
	UnsetIsTest()
}

func SetIsTest() {
	lock.Lock()
	isTest = true

	origRedisTimeoutMilliTS = REDIS_TIMEOUT_MILLI_TS
	REDIS_TIMEOUT_MILLI_TS = TEST_REDIS_TIMEOUT_MILLI_TS
}

func UnsetIsTest() {
	REDIS_TIMEOUT_MILLI_TS = origRedisTimeoutMilliTS

	isTest = false
	lock.Unlock()
}
