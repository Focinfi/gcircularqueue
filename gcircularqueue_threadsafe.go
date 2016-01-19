package gcircularqueue

import (
	"sync"
)

// CircularQueueThreadSafe
// composing sync.RWMutex and pointer of CircularQueue
type CircularQueueThreadSafe struct {
	sync.RWMutex
	*CircularQueue
}

// NewCircularQueueThreadSafe return a new NewCircularQueueThreadSafe
func NewCircularQueueThreadSafe(size int) *CircularQueueThreadSafe {
	return &CircularQueueThreadSafe{CircularQueue: NewCircularQueue(size)}
}

// IsEmpty return cq.CircularQueue.IsEmpty() wrapped by RLock
func (cq *CircularQueueThreadSafe) IsEmpty() bool {
	cq.RLock()
	defer cq.RUnlock()
	return cq.CircularQueue.IsEmpty()
}

// IsFull return cq.CircularQueue.IsFull() wrapped by RLock
func (cq *CircularQueueThreadSafe) IsFull() bool {
	cq.RLock()
	defer cq.RUnlock()
	return cq.CircularQueue.IsFull()
}

// Push push a element into cq.CircularQueue wrapped by Lock
func (cq *CircularQueueThreadSafe) Push(element interface{}) {
	cq.Lock()
	defer cq.Unlock()
	cq.CircularQueue.Push(element)
}

// Shift shift a element from cq.CircularQueue wrapped by Lock
func (cq *CircularQueueThreadSafe) Shift() interface{} {
	cq.Lock()
	defer cq.Unlock()
	return cq.CircularQueue.Shift()
}
