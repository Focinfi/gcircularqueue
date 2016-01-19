gcircularqueue is Circular Queue in Golang

#### Install 

`go get github.com/Focinfi/gcircularqueue`

### Usage

This package is a typical circular with interfaces as following:

##### New

  Normal queue
  ```go
  queue := gcircularqueue.NewCircularQueue(size int)
  ```

  Thread-safe queue via `sync.RWMutex`
  ```go
  queue := gcircularqueue.NewCircularQueueThreadSafe(size int)
  ```

##### IsEmpty/IsFull

  ```go
  isEmpty := queue.IsEmpty()
  isFull := queue.IsFull()
  ```

##### Push

  ```go
  queue.Push("kitty")
  // note: it will panic if pushing a thing into a full queue
  ```

##### Shift

  ```go
  element := queue.Shift()
  // note: `element will be nil when shifting from a empty queue
  ```
