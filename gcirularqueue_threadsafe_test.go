package gcircularqueue

import (
	"testing"
)

func TestQueueThreadSafeInit(t *testing.T) {
	cq := NewCircularQueueThreadSafe(5)
	if cq.capacity != 6 {
		t.Error("It has wrong size:", cq.capacity)
	}

	if !cq.IsEmpty() {
		t.Error("It is not empty")
	}
}

func TestThreadSafePush(t *testing.T) {
	cq := NewCircularQueueThreadSafe(1)
	firstElement := "First"
	cq.Push(firstElement)
	if cq.elements[0] != firstElement {
		t.Error("It's Push Func is wrong, the first elment is:", cq.elements[0])
	}
}

func TestThreadSafeShift(t *testing.T) {
	cq := NewCircularQueueThreadSafe(5)
	firstElement := "First"
	cq.Push(firstElement)
	e := cq.Shift()
	if e != firstElement {
		t.Error("It can not shift the first element")
	}
	// shift from a empty queueThreadSafe
	e = cq.Shift()
	if e != nil {
		t.Errorf("Can not shift from a empty queueThreadSafe, it is:%v", e)
	}
}

func TestThreadSafeIsEmpty(t *testing.T) {
	cq := NewCircularQueueThreadSafe(5)
	if !cq.IsEmpty() {
		t.Error("It's IsEmpty Func is wrong")
	}
}

func TestThreadSafeIsFull(t *testing.T) {
	cq := NewCircularQueueThreadSafe(1)
	cq.Push("First")
	if !cq.IsFull() {
		t.Error("It's IsFull is wrong")
	}

}

func TestThreadSafeFIFO(t *testing.T) {
	cq := NewCircularQueueThreadSafe(3)
	cq.Push(1)
	cq.Push(2)
	firstElement := cq.Shift()
	if firstElement != 1 {
		t.Error("It doesn't support FIFO")
	}
}

func TestThreadSafeCirculartAility(t *testing.T) {
	cq := NewCircularQueueThreadSafe(3)
	cq.Push(1)
	cq.Push(2)
	cq.Push(3)
	cq.Shift()
	cq.Push(3)
}
