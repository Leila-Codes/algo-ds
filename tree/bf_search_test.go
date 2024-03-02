package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	tree := Binary[int]{
		Root: BinaryNode[int]{
			Value: 7,
			Left: &BinaryNode[int]{
				Value: 23,
				Left:  &BinaryNode[int]{Value: 5},
				Right: &BinaryNode[int]{Value: 4},
			},
			Right: &BinaryNode[int]{
				Value: 8,
				Left:  &BinaryNode[int]{Value: 21},
				Right: &BinaryNode[int]{Value: 15},
			},
		},
		height: 3,
	}

	assert.True(t, BreadthFirstSearch(tree.Root, 5))
	assert.True(t, BreadthFirstSearch(tree.Root, 4))
	assert.True(t, BreadthFirstSearch(tree.Root, 21))
	assert.False(t, BreadthFirstSearch(tree.Root, 50))
}
