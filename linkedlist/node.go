package linkedlist

// Node is an object containing a value, plus a pointer to the next value in the sequence.
type Node[T interface{}] struct {
	Value T
	Next  *Node[T]
}

// DoubleLinkedNode extends Node with an additional pointer to the previous object in the list.
type DoubleLinkedNode[T interface{}] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}
