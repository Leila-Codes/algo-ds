package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	g := make(WeightedAdjacencyList, 4)

	g[0] = []int{0, -1, 5, -1}
	g[1] = []int{-1, -1, -1, -1}
	g[2] = []int{3, -1, -1, -1}
	g[3] = []int{-1, 4, 3, -1}

	res := BreadthFirstSearch(g, 0, 3)
	assert.True(t, res)
}
