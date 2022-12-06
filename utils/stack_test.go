package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := Stack{}

	assert.NotNil(t, stack)
	assert.Equal(t, 0, stack.Size())
}

func TestPush(t *testing.T) {
	stack := Stack{}
	stack.Push(1)

	assert.Equal(t, 1, stack.Size())
}

func TestPop(t *testing.T) {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	assert.Equal(t, 4, stack.Size())
	assert.Equal(t, 4, stack.Pop())
	assert.Equal(t, 3, stack.Size())
	assert.Equal(t, 3, stack.Pop())
	assert.Equal(t, 2, stack.Size())
	assert.Equal(t, 2, stack.Pop())
	assert.Equal(t, 1, stack.Size())
	assert.Equal(t, 1, stack.Pop())
	assert.Equal(t, 0, stack.Size())
	assert.Equal(t, nil, stack.Pop())
}

func TestPeek(t *testing.T) {
	stack := Stack{}

	assert.Equal(t, nil, stack.Peek())
	stack.Push(1)
	stack.Push(2)
	assert.Equal(t, 2, stack.Size())
	assert.Equal(t, 2, stack.Peek())
	assert.Equal(t, 2, stack.Size())
}
