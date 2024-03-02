package queue

// Node is an object containing a value, plus a pointer to the next value in the sequence.
type Node[T interface{}] struct {
	Value T
	Next  *Node[T]
}
