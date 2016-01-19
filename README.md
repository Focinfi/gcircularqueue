gcircularqueue is Circular Queue in Golang

#### Install 

`go get github.com/Focinfi/gcircularqueue`

#### Usage

This package is a typical circular with interfaces as following:

1. Create new cirular queue passing a integer as its size

```go
queue := gcircularqueue.NewCircularQueue(size int)
```

2. Check if queue is empty or full

```go
isEmpty := queue.IsEmpty()
isFull := queue.IsFull()
```

3. Push

```go
queue.Push("kitty")
```

**note**: it will panic if pushing a thing into a full queue

4. Shift

```go
element := queue.Shift()
```

**note**: `element` will be nil when shifting from a empty queue
