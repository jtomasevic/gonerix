package collections

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStack_Integer(t *testing.T) {

	t.Run("Pop", func(t *testing.T) {
		stack:= Stack[int]{}
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)
		popped := stack.Pop()
		require.Equal(t, *popped, 3)
		popped = stack.Pop()
		require.Equal(t, *popped, 2)
		popped = stack.Pop()
		require.Equal(t, *popped, 1)
		popped = stack.Pop()
		require.Nil(t, popped)
	})

	t.Run("Pop random", func(t *testing.T) {
		stack:= Stack[int]{}
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)
		stack.Push(4)
		
		popped := stack.Pop()
		require.Equal(t, *popped, 4)
		popped = stack.Pop()
		require.Equal(t, *popped, 3)

		stack.Push(15)
		popped = stack.Pop()
		require.Equal(t, *popped, 15)

		popped = stack.Pop()
		require.Equal(t, *popped, 2)
		popped = stack.Pop()
		require.Equal(t, *popped, 1)

		popped = stack.Pop()
		require.Nil(t, popped)
	})
	t.Run("Pop", func(t *testing.T) {
		stack:= Stack[int]{}
		
		stack.Push(1)
		pik := stack.Peek()
		require.Equal(t, *pik, 1)

		stack.Push(2)
		pik = stack.Peek()
		require.Equal(t, *pik, 2)

		stack.Push(3)
		pik = stack.Peek()
		require.Equal(t, *pik, 3)

		popped := stack.Pop()
		pik = stack.Peek()
		require.Equal(t, *pik, 2)
		require.Equal(t, *popped, 3)

		popped = stack.Pop()
		pik = stack.Peek()
		require.Equal(t, *pik, 1)
		require.Equal(t, *popped, 2)

		popped = stack.Pop()
		require.Equal(t, *popped, 1)
		pik = stack.Peek()
		require.Nil(t, pik)

		popped = stack.Pop()
		pik = stack.Peek()
		require.Nil(t, pik)
		require.Nil(t, popped)
	})
}
