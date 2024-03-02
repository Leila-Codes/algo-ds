package tree

import "algo-ds/queue"

// BreadthFirstSearch implements breadth-first-search (BFS)
// Starts at the given head/root and attempts to find the given needle.
func BreadthFirstSearch[T comparable](root BinaryNode[T], needle T) bool {
	q := queue.New(root)

	for q.Length() > 0 {
		curr, _ := q.Dequeue()

		if curr.Value == needle {
			return true
		}

		if curr.Left != nil {
			q.Enqueue(*curr.Left)
		}

		if curr.Right != nil {
			q.Enqueue(*curr.Right)
		}
	}

	return false
}
