package gcircularqueue

// CirCircularQueue
//   capacity is the numbers of this queue's capacity(size - 1)
type CircularQueue struct {
	capacity int
	elements []interface{}
	first    int
	end      int
}

func NewCircularQueue(size int) *CircularQueue {
	cq := CircularQueue{capacity: size + 1, first: 0, end: 0}
	cq.elements = make([]interface{}, cq.capacity)
	return &cq
}

func (c CircularQueue) IsEmpty() bool {
	return c.first == c.end
}

func (c CircularQueue) IsFull() bool {
	return c.first == (c.end+1)%c.capacity
}

func (c *CircularQueue) Push(e interface{}) {
	if c.IsFull() {
		panic("Queue is full")
	}
	c.elements[c.end] = e
	c.end = (c.end + 1) % c.capacity
}

func (c *CircularQueue) Shift() (e interface{}) {
	if c.IsEmpty() {
		return nil
	}
	e = c.elements[c.first]
	c.first = (c.first + 1) % c.capacity
	return
}
