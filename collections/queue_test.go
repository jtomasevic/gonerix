package collections

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueue_Integer(t *testing.T) {

	t.Run("Enqueue/Dequeue/Peek", func(t *testing.T) {
		stack := Queue[int]{}
		stack.Enqueue(1)
		stack.Enqueue(2)
		stack.Enqueue(3)
		el := stack.Peek()
		require.Equal(t, *el, 1)
		el = stack.Dequeue()
		require.Equal(t, *el, 1)
		el = stack.Dequeue()
		require.Equal(t, *el, 2)
		el = stack.Dequeue()
		require.Equal(t, *el, 3)
		el = stack.Dequeue()
		require.Nil(t, el)
		el = stack.Peek()
		require.Nil(t, el)
	})
}
