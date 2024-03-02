package tree

// Node is an object on a tree with 0 or more children.
//type Node[T comparable] struct {
//	Value    T
//	Children []*Node[T]
//}

// BinaryNode is an object containing a value, plus 2 pointers
// one (Left) and one (Right) - both default to nil.
type BinaryNode[T comparable] struct {
	Value T
	Left  *BinaryNode[T]
	Right *BinaryNode[T]
}

// IsLeaf returns whether this node is a leaf (contains no children)
func (bn *BinaryNode[T]) IsLeaf() bool {
	return bn.Left == nil && bn.Right == nil
}
