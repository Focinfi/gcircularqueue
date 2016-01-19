package gcircularqueue

import (
	"testing"
)

func TestQueueInit(t *testing.T) {
	cq := NewCircularQueue(5)
	if cq.capacity != 6 {
		t.Error("It has wrong capacity:", cq.capacity)
	}

	if !cq.IsEmpty() {
		t.Error("It is not empty")
	}
}

func TestPush(t *testing.T) {
	cq := NewCircularQueue(5)
	firstElement := "First"
	cq.Push(firstElement)
	if cq.elements[0] != firstElement {
		t.Error("It's Push Func is wrong, the first elment is:", cq.elements[0])
	}
}

func TestShift(t *testing.T) {
	cq := NewCircularQueue(5)
	firstElement := "First"
	cq.Push(firstElement)
	e := cq.Shift()
	if e != firstElement {
		t.Error("It can not shift the first element")
	}
	// shift from a empty queue
	e = cq.Shift()
	if e != nil {
		t.Errorf("Can not shift from a empty queue, it is:%v", e)
	}
}

func TestIsEmpty(t *testing.T) {
	cq := NewCircularQueue(5)
	if !cq.IsEmpty() {
		t.Error("It's IsEmpty Func is wrong")
	}
}

func TestIsFull(t *testing.T) {
	cq := NewCircularQueue(1)
	cq.Push("First")
	if !cq.IsFull() {
		t.Error("It's IsFull is wrong")
	}
}

func TestFIFO(t *testing.T) {
	cq := NewCircularQueue(3)
	cq.Push(1)
	cq.Push(2)
	firstElement := cq.Shift()
	if firstElement != 1 {
		t.Error("It doesn't support FIFO")
	}
}

func TestCirculartAility(t *testing.T) {
	cq := NewCircularQueue(3)
	cq.Push(1)
	cq.Push(2)
	cq.Push(3)
	cq.Shift()
	cq.Push(3)
}
