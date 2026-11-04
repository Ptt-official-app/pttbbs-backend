package queue

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	pool "github.com/golang-queue/queue"
	"github.com/golang-queue/queue/job"
)

func startTestPool(t *testing.T, size int) {
	t.Helper()
	Close()
	workers, capacity := WORKER_NUM, QUEUE_SIZE
	WORKER_NUM, QUEUE_SIZE = 1, size
	t.Cleanup(func() { Close(); WORKER_NUM, QUEUE_SIZE = workers, capacity })
	if err := Start(); err != nil {
		t.Fatal(err)
	}
}

func await(t *testing.T, ch <-chan struct{}) {
	t.Helper()
	select {
	case <-ch:
	case <-time.After(3 * time.Second):
		t.Fatal("timed out waiting for queue")
	}
}

func TestQueueLifecycle(t *testing.T) {
	Close()
	Close()
	task := func(context.Context) error { return nil }
	if err := queueTask(task); !errors.Is(err, pool.ErrQueueShutdown) {
		t.Fatalf("before Start: %v", err)
	}
	startTestPool(t, 8)
	original := client
	if err := Start(); err != nil {
		t.Fatal(err)
	}
	if client != original {
		t.Fatal("Start replaced a running pool")
	}
	var completed atomic.Int32
	for range 8 {
		if err := queueTask(func(context.Context) error { completed.Add(1); return nil }); err != nil {
			t.Fatal(err)
		}
	}
	Close()
	if got := completed.Load(); got != 8 {
		t.Fatalf("completed %d/8 tasks", got)
	}
	if err := queueTask(task); !errors.Is(err, pool.ErrQueueShutdown) {
		t.Fatalf("after Close: %v", err)
	}
	if err := Start(); err != nil {
		t.Fatal(err)
	}
	if err := queueTask(task); err != nil {
		t.Fatal(err)
	}
}

func TestQueueCapacity(t *testing.T) {
	startTestPool(t, 1)
	started, unblock := make(chan struct{}), make(chan struct{})
	defer close(unblock)
	if err := queueTask(func(context.Context) error { close(started); <-unblock; return nil }); err != nil {
		t.Fatal(err)
	}
	await(t, started)
	if err := queueTask(func(context.Context) error { return nil }); err != nil {
		t.Fatal(err)
	}
	if err := queueTask(func(context.Context) error { t.Error("rejected task executed"); return nil }); !errors.Is(err, pool.ErrMaxCapacity) {
		t.Fatalf("full queue: %v", err)
	}
}

func TestQueueCloseWaitsForTimedOutTask(t *testing.T) {
	startTestPool(t, 1)
	expired, unblock := make(chan struct{}), make(chan struct{})
	var once sync.Once
	defer once.Do(func() { close(unblock) })
	timeout := 10 * time.Millisecond
	if err := queueTask(func(ctx context.Context) error {
		<-ctx.Done()
		close(expired)
		<-unblock // Simulate the existing database work that ignores context.
		return nil
	}, job.AllowOption{Timeout: &timeout}); err != nil {
		t.Fatal(err)
	}
	await(t, expired)
	done := make(chan struct{})
	go func() { Close(); close(done) }()
	// Wait until Close has stopped admission, without assuming scheduling order.
	deadline := time.After(3 * time.Second)
	for {
		clientMu.Lock()
		stopping := closing != nil
		clientMu.Unlock()
		if stopping {
			break
		}
		select {
		case <-deadline:
			t.Fatal("Close did not stop admission")
		default:
			time.Sleep(time.Millisecond)
		}
	}
	if err := Start(); !errors.Is(err, pool.ErrQueueShutdown) {
		t.Fatalf("Start during Close: %v", err)
	}
	if err := queueTask(func(context.Context) error { return nil }); !errors.Is(err, pool.ErrQueueShutdown) {
		t.Fatalf("enqueue during Close: %v", err)
	}
	second := make(chan struct{})
	go func() { Close(); close(second) }()
	select {
	case <-done:
		t.Fatal("Close returned before task finished")
	default:
	}
	select {
	case <-second:
		t.Fatal("second Close returned before task finished")
	default:
	}
	once.Do(func() { close(unblock) })
	await(t, done)
	await(t, second)
}

func TestQueueConcurrentAdmissionAndClose(t *testing.T) {
	startTestPool(t, 256)
	var accepted, completed atomic.Int32
	var submitters sync.WaitGroup
	for range 100 {
		submitters.Go(func() {
			err := queueTask(func(context.Context) error { completed.Add(1); return nil })
			if err == nil {
				accepted.Add(1)
			} else if !errors.Is(err, pool.ErrQueueShutdown) {
				t.Errorf("enqueue: %v", err)
			}
		})
	}
	Close()
	submitters.Wait()
	if accepted.Load() != completed.Load() {
		t.Fatalf("accepted %d, completed %d", accepted.Load(), completed.Load())
	}
}

func TestQueueInvalidConfig(t *testing.T) {
	Close()
	workers, capacity := WORKER_NUM, QUEUE_SIZE
	defer func() { WORKER_NUM, QUEUE_SIZE = workers, capacity }()
	for _, config := range [][2]int{{0, 1}, {-1, 1}, {1, -1}} {
		WORKER_NUM, QUEUE_SIZE = config[0], config[1]
		if err := Start(); err == nil {
			Close()
			t.Fatalf("accepted invalid config %v", config)
		}
	}
}
