package tree

// Binary represents a binary tree of items with a root. Each child has at most 2 children.
type Binary[T comparable] struct {
	Root   BinaryNode[T]
	height int
}

// Height returns the current height of the tree.
func (bt *Binary[T]) Height() int {
	return bt.height
}
