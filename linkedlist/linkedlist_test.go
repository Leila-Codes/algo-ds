package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_Contains(t *testing.T) {
	ll := New[int](5, 21, 7, 31)

	assert.True(t, ll.Contains(7))
	assert.True(t, ll.Contains(31))
	assert.False(t, ll.Contains(90))
}

func TestLinkedList_Push(t *testing.T) {
	ll := New[int](5, 89, 31, 41)

	assert.False(t, ll.Contains(13))
	ll.Push(13)
	assert.True(t, ll.Contains(13))
}
