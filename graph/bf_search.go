package graph

import "algo-ds/queue"

func BreadthFirstSearch(
	graph WeightedAdjacencyList,
	start int,
	needle int,
) bool {
	q := queue.New(start)
	seen := make([]bool, len(graph))
	prev := make([]int, len(graph))

	for q.Length() > 0 {
		curr, _ := q.Dequeue()

		if *curr == needle {
			return true
		}

		for _, edge := range graph[*curr] {

			// edge of -1 means no link.
			if edge == -1 {
				continue
			}

			if seen[edge] {
				continue
			}

			seen[edge] = true
			prev[edge] = *curr

			q.Enqueue(edge)
		}
	}

	return prev[needle] == -1
}
