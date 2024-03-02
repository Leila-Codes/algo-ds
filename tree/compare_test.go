package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompare(t *testing.T) {
	treeOne := Binary[int]{
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

	treeTwo := Binary[int]{
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

	treeThree := Binary[int]{
		Root: BinaryNode[int]{
			Value: 7,
			Left: &BinaryNode[int]{
				Value: 18,
				Left:  &BinaryNode[int]{Value: 5},
			},
			Right: &BinaryNode[int]{
				Value: 8,
				Left:  &BinaryNode[int]{Value: 21},
				Right: &BinaryNode[int]{Value: 15},
			},
		},
		height: 3,
	}

	treeFour := Binary[int]{
		Root: BinaryNode[int]{
			Value: 7,
			Left: &BinaryNode[int]{
				Value: 23,
				Left: &BinaryNode[int]{
					Value: 8,
					Left:  &BinaryNode[int]{Value: 21},
					Right: &BinaryNode[int]{Value: 15},
				},
				Right: &BinaryNode[int]{Value: 4},
			},
		},
		height: 4,
	}

	assert.True(t, Compare(&treeOne.Root, &treeTwo.Root))
	assert.False(t, Compare(&treeOne.Root, &treeThree.Root))
	assert.False(t, Compare(&treeOne.Root, &treeFour.Root))
}
