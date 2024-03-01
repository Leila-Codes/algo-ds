package algo_ds

// Node is an object containing a value, plus a pointer to the next value in the sequence.
type Node[T interface{}] struct {
	Value T
	Next  *Node[T]
}

// BinaryNode is an object containing a value, plus 2 pointers
// one (Left) and one (Right) - both default to nil.
type BinaryNode[T interface{}] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

// IsLeaf returns whether this node is a leaf (contains no children)
func (bn *BinaryNode[T]) IsLeaf() bool {
	return bn.Left == nil && bn.Right == nil
}
