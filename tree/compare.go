package tree

// Compare compares the two binary trees from the given nodes a and b. This function is recursive.
func Compare[T comparable](a *BinaryNode[T], b *BinaryNode[T]) bool {
	if a == nil || b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Value != b.Value {
		return false
	}

	return Compare(a.Left, b.Left) && Compare(a.Right, b.Right)
}
