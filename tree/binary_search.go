package tree

import "algo-ds/types"

// BinarySearch implements a binary search across a tree data structure.
// Note the additional limitation that values must be of type Number to perform comparisons.
// Complexity: O(log(n)) -> O(n) - dependant on balance of the tree.
func BinarySearch[T types.Number](node *BinaryNode[T], needle T) bool {
	if node == nil {
		return false
	}

	if node.Value == needle {
		return true
	}

	if needle > node.Value {
		return BinarySearch(node.Right, needle)
	} else {
		return BinarySearch(node.Left, needle)
	}
}
