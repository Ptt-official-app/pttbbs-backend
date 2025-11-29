package queue

var (
	QUEUE_SIZE = 4096
	WORKER_NUM = 2

	WORKER_NUM_i64 = int64(WORKER_NUM)
)

func config() {
	QUEUE_SIZE = setIntConfig("QUEUE_SIZE", QUEUE_SIZE)
	WORKER_NUM = setIntConfig("WORKER_NUM", WORKER_NUM)
	WORKER_NUM_i64 = int64(WORKER_NUM)
}
