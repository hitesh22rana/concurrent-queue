package queue

import (
	"errors"
	"sync"
)

type ConcurrentQueue struct {
	queue []int32
	mu    sync.Mutex
}

func (cq *ConcurrentQueue) Enqueue(item int32) {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	cq.queue = append(cq.queue, item)
}

func (cq *ConcurrentQueue) Dequeue() (int32, error) {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	if len(cq.queue) == 0 {
		return 0, errors.New("cannot deque from an empty queue")
	}

	item := cq.queue[0]
	cq.queue = cq.queue[1:]
	return item, nil
}

func (cq *ConcurrentQueue) Size() int {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	return len(cq.queue)
}
