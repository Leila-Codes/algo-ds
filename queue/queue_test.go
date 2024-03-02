package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	queue := New[int]()

	assert.Equal(t, queue.Length(), 0)
	assert.Nil(t, queue.Head)
	assert.Nil(t, queue.Tail)

	queue.Enqueue(5)
	assert.Equal(t, queue.Length(), 1)
	assert.NotNil(t, queue.Head)
	assert.NotNil(t, queue.Tail)
	assert.Equal(t, queue.Head.Value, 5)
}

func TestQueue_Dequeue(t *testing.T) {
	queue := New[int](5, 19, 3, 7)

	assert.Equal(t, queue.Length(), 4)

	item, err := queue.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, *item, 5)
	assert.Equal(t, queue.Length(), 3)

	item, err = queue.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, *item, 19)
	assert.Equal(t, queue.Length(), 2)

	item, err = queue.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, *item, 3)
	assert.Equal(t, queue.Length(), 1)

	item, err = queue.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, *item, 7)
	assert.Equal(t, queue.Length(), 0)

	item, err = queue.Dequeue()
	assert.Equal(t, err, ErrQueueEmpty)
	assert.Nil(t, item)
}

func TestQueue_Peek(t *testing.T) {
	queue := New[int](5, 19, 3, 7)
	assert.Equal(t, queue.Length(), 4)
	item, err := queue.Peek()
	assert.Nil(t, err)
	assert.Equal(t, *item, 5)
	assert.Equal(t, queue.Length(), 4)

	item, err = queue.Peek()
	assert.Nil(t, err)
	assert.Equal(t, *item, 5)
	assert.Equal(t, queue.Length(), 4)
}
