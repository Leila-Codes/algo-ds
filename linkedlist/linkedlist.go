package linkedlist

// LinkedList implements a linked-list data structure of the given type.
type LinkedList[T comparable] struct {
	Head *Node[T]
}

// New constructs a new linked list
// optionally can take initValues list of values to fill the list with.
func New[T comparable](initValues ...T) *LinkedList[T] {
	ll := &LinkedList[T]{}

	for _, v := range initValues {
		ll.Push(v)
	}

	return ll
}

// Push adds the specified item to the end of the linked list.
// Complexity: O(n) (to create the inputs) + O(n) to find next empty cell = O(n + 2) = O(2n) = O(n)
func (l *LinkedList[T]) Push(values ...T) {
	// exit early if no values
	if len(values) == 0 {
		return
	}

	// construct the node(s)
	first := &Node[T]{Value: values[0]}
	node := first
	for i := 1; i < len(values); i++ {
		node.Next = &Node[T]{Value: values[i]}
		node = node.Next
	}

	// is the head assigned?
	if l.Head == nil {
		l.Head = first
		return
	}

	// otherwise, find the next nil position in the list
	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}

	// place this item at this position
	curr.Next = first
}

// Contains returns whether the given value is inside the list.
// Complexity - Time: O(n)
func (l *LinkedList[T]) Contains(value T) (found bool) {
	curr := l.Head

	for !found && curr != nil {
		found = curr.Value == value
		curr = curr.Next
	}

	return found
}
