package algo_ds

import "errors"

var (
	ErrQueueEmpty = errors.New("queue is empty")
)

type Queue[T interface{}] struct {
	Head *Node[T]
	Tail *Node[T]
}

// Enqueue pushes the given value to the end of the queue
func (q *Queue[T]) Enqueue(value T) {
	node := &Node[T]{Value: value}
	if q.Head == nil {
		q.Head = node
		q.Tail = q.Head
	}

	q.Tail.Next = node
	q.Tail = node
}

// Dequeue retrieves the next item from the queue.
//
// If the queue is empty an ErrQueueEmpty is returned
func (q *Queue[T]) Dequeue() (*T, error) {
	node := q.Head
	if node == nil {
		return nil, ErrQueueEmpty
	}

	q.Head = q.Head.Next
	return &node.Value, nil
}

// Peek looks up the current head of the queue without removing it.
//
// If the queue is empty an ErrQueueEmpty is returned
func (q *Queue[T]) Peek() (*T, error) {
	if q.Head == nil {
		return nil, ErrQueueEmpty
	}

	return &q.Head.Value, nil
}
