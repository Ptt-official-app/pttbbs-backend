package queue

import (
	"github.com/golang-queue/queue"
	"github.com/sirupsen/logrus"
)

var client *queue.Queue

func Start() error {
	client = queue.NewPool(
		WORKER_NUM_i64,
		queue.WithQueueSize(QUEUE_SIZE),
	)

	// start the worker
	logrus.Infof("queue.Start: client start")
	client.Start()

	return nil
}

func Close() {
	// shutdown the service and notify all the worker
	// wait all jobs are complete.
	defer logrus.Infof("queue.Close: done")
	client.Release()
}
