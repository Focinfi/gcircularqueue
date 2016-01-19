package gcircularqueue

import (
	"sync"
)

type CircularQueueThreadSafe struct {
	sync.RWMutex
	*CircularQueue
}

func NewCircularQueueThreadSafe(size int) *CircularQueueThreadSafe {
	return &CircularQueueThreadSafe{CircularQueue: NewCircularQueue(size)}
}

func (cq *CircularQueueThreadSafe) IsEmpty() bool {
	cq.RLock()
	defer cq.RUnlock()
	return cq.CircularQueue.IsEmpty()
}

func (cq *CircularQueueThreadSafe) IsFull() bool {
	cq.RLock()
	defer cq.RUnlock()
	return cq.CircularQueue.IsFull()
}

func (cq *CircularQueueThreadSafe) Push(element interface{}) {
	cq.Lock()
	defer cq.Unlock()
	cq.CircularQueue.Push(element)
}

func (cq *CircularQueueThreadSafe) Shift() interface{} {
	cq.Lock()
	defer cq.Unlock()
	return cq.CircularQueue.Shift()
}
