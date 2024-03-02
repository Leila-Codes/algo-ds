package queue

import "errors"

var (
	ErrQueueEmpty = errors.New("queue is empty")
)

type Queue[T comparable] struct {
	Head   *Node[T]
	Tail   *Node[T]
	length int
}

// Length returns the length of the queue
func (q *Queue[T]) Length() int {
	return q.length
}

func New[T comparable](initValues ...T) *Queue[T] {
	q := &Queue[T]{}
	for _, v := range initValues {
		q.Enqueue(v)
	}
	return q
}

// Enqueue pushes the given value to the end of the queue
func (q *Queue[T]) Enqueue(values ...T) {
	for _, v := range values {
		node := &Node[T]{Value: v}
		if q.Head == nil {
			q.Head = node
			q.Tail = q.Head
		}

		q.Tail.Next = node
		q.Tail = node

		q.length++
	}
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
	q.length--

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
