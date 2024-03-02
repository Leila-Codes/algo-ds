package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tree := Binary[int]{
		Root: BinaryNode[int]{
			Value: 7,
			Left: &BinaryNode[int]{
				Value: 5,
				Left:  &BinaryNode[int]{Value: 3},
				Right: &BinaryNode[int]{Value: 15},
			},
			Right: &BinaryNode[int]{
				Value: 23,
				Left:  &BinaryNode[int]{Value: 21},
				Right: &BinaryNode[int]{Value: 41},
			},
		},
		height: 3,
	}

	assert.False(t, BinarySearch(&tree.Root, 4))
	assert.True(t, BinarySearch(&tree.Root, 21))
}
