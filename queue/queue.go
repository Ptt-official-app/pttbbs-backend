package queue

import (
	"context"
	"fmt"
	"sync"

	"github.com/golang-queue/queue"
	"github.com/golang-queue/queue/job"
	"github.com/sirupsen/logrus"
)

var (
	client   *queue.Queue
	clientMu sync.Mutex
	closing  chan struct{}
	pending  sync.WaitGroup
)

func Start() error {
	clientMu.Lock()
	defer clientMu.Unlock()
	if closing != nil {
		return queue.ErrQueueShutdown
	}
	if client != nil {
		return nil
	}
	if WORKER_NUM <= 0 || QUEUE_SIZE < 0 {
		return fmt.Errorf("queue: WORKER_NUM must be positive and QUEUE_SIZE nonnegative (got %d, %d)", WORKER_NUM, QUEUE_SIZE)
	}
	// NewPool starts its workers automatically.
	client = queue.NewPool(int64(WORKER_NUM), queue.WithQueueSize(QUEUE_SIZE))
	logrus.Infof("queue.Start: client start")
	return nil
}

// queueTask owns admission and tracks actual task completion, even if the
// library's timeout expires before a task that ignores context returns.
// Tasks must not call Start or Close. Automatic retries are intentionally off:
// comment processing performs multiple database writes and is not atomic.
func queueTask(task job.TaskFunc, opts ...job.AllowOption) error {
	clientMu.Lock()
	defer clientMu.Unlock()
	if client == nil || closing != nil {
		return queue.ErrQueueShutdown
	}
	// Force a single attempt so completion accounting also holds if options
	// are supplied. Retrying partial comment writes needs a separate policy.
	if len(opts) == 0 {
		opts = []job.AllowOption{{}}
	}
	option := opts[0]
	option.RetryCount = job.Int64(0)
	pending.Add(1)
	err := client.QueueTask(func(ctx context.Context) error {
		defer pending.Done()
		return task(ctx)
	}, option)
	if err != nil {
		pending.Done()
	}
	return err
}

// Close rejects new work and drains accepted tasks before releasing the pool.
// It is safe before Start and for concurrent or repeated callers.
func Close() {
	clientMu.Lock()
	if done := closing; done != nil {
		clientMu.Unlock()
		<-done
		return
	}
	if client == nil {
		clientMu.Unlock()
		return
	}
	done := make(chan struct{})
	closing = done
	q := client
	clientMu.Unlock()

	pending.Wait()
	q.Release()

	clientMu.Lock()
	client = nil
	closing = nil
	close(done)
	clientMu.Unlock()
	logrus.Infof("queue.Close: done")
}
