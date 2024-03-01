package algo_ds

// LinkedList implements a linked-list data structure of the given type.
type LinkedList[T interface{}] struct {
	Head *Node[T]
}

// NewLinkedList constructs a new (empty) linked list
func NewLinkedList[T interface{}]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Push adds the specified item to the end of the linked list.
func (l *LinkedList[T]) Push(value T) {
	node := Node[T]{Value: value}

	if l.Head == nil {
		l.Head = &node
	}

	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = &node
}
