package collections

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLinkedList_Add(t *testing.T) {
	t.Run("Add ", func(t *testing.T) {
		list := LinkedList[int]()
		result := list.Add(5)
		require.Equal(t, 5, result.Value)
		require.Nil(t, result.Next)
		require.Nil(t, result.Prev)
		require.Equal(t, 5, *list.Head())
		require.Equal(t, 5, *list.Tail())

		result = list.Add(2)
		require.Equal(t, 2, result.Value)
		require.Equal(t, 5, result.Prev.Value)
		require.Nil(t, result.Next)
		require.Equal(t, 2, *list.Tail())

		result = list.Add(3)
		require.Equal(t, 3, result.Value)
		require.Equal(t, 2, result.Prev.Value)
		require.Nil(t, result.Next)
		require.Equal(t, 5, *list.Head())
		require.Equal(t, 3, *list.Tail())

		list.Add(10)
		require.Equal(t, 5, *list.Head())
		require.Equal(t, 10, *list.Tail())
		require.Equal(t, []int{
			5, 2, 3, 10,
		}, list.Values())
	})
}

func TestLinkedList_Remove(t *testing.T) {
	t.Run("Remove from middle 1", func(t *testing.T) {
		list := LinkedList[int]()
		list.Add(1)
		list.Add(2)
		list.Add(3)
		list.Add(4)
		list.Add(5)
		result, ok := list.Remove(3)
		require.Equal(t, 3, (*result).Value)
		require.Equal(t, 1, *list.Head())
		require.Equal(t, 5, *list.Tail())
		require.Equal(t, true, ok)
		require.Equal(t, []int{
			1, 2, 4, 5,
		}, list.Values())
	})

	t.Run("Remove from middle 2", func(t *testing.T) {
		list := LinkedList[int]()
		list.Add(1)
		list.Add(2)
		list.Add(3)
		list.Add(4)
		list.Add(5)
		result, ok := list.Remove(2)
		require.Equal(t, 2, (*result).Value)
		require.Equal(t, true, ok)
		require.Equal(t, []int{
			1, 3, 4, 5,
		}, list.Values())
		require.Equal(t, 1, *list.Head())
		require.Equal(t, 5, *list.Tail())
	})

	t.Run("Remove from middle 3", func(t *testing.T) {
		list := LinkedList[int]()
		list.Add(1)
		list.Add(2)
		list.Add(3)
		list.Add(4)
		list.Add(5)
		result, ok := list.Remove(4)
		require.Equal(t, 4, (*result).Value)
		require.Equal(t, true, ok)
		require.Equal(t, []int{
			1, 2, 3, 5,
		}, list.Values())
		require.Equal(t, 1, *list.Head())
		require.Equal(t, 5, *list.Tail())
	})

	t.Run("Remove first", func(t *testing.T) {
		list := LinkedList[int]()
		list.Add(1)
		list.Add(2)
		list.Add(3)
		list.Add(4)
		list.Add(5)
		result, ok := list.Remove(1)
		require.Equal(t, 1, (*result).Value)
		require.Equal(t, true, ok)
		require.Equal(t, []int{
			2, 3, 4, 5,
		}, list.Values())
		require.Equal(t, 2, *list.Head())
		require.Equal(t, 5, *list.Tail())
	})

	t.Run("Remove last", func(t *testing.T) {
		list := LinkedList[int]()
		list.Add(1)
		list.Add(2)
		list.Add(3)
		list.Add(4)
		list.Add(5)
		result, ok := list.Remove(5)
		require.Equal(t, 5, (*result).Value)
		require.Equal(t, true, ok)
		require.Equal(t, []int{
			1, 2, 3, 4,
		}, list.Values())
		require.Equal(t, 1, *list.Head())
		require.Equal(t, 4, *list.Tail())

		result, ok = list.Remove(4)
		require.Equal(t, 4, (*result).Value)
		require.Equal(t, true, ok)
		require.Equal(t, []int{
			1, 2, 3,
		}, list.Values())
		require.Equal(t, 1, *list.Head())
		require.Equal(t, 3, *list.Tail())

		result, ok = list.Remove(3)
		require.Equal(t, 3, (*result).Value)
		require.Equal(t, true, ok)
		require.Equal(t, []int{
			1, 2,
		}, list.Values())
		require.Equal(t, 1, *list.Head())
		require.Equal(t, 2, *list.Tail())

		result, ok = list.Remove(2)
		require.Equal(t, 2, (*result).Value)
		require.Equal(t, true, ok)
		require.Equal(t, []int{
			1,
		}, list.Values())
		require.Equal(t, 1, *list.Head())
		require.Equal(t, 1, *list.Tail())

		result, ok = list.Remove(1)
		require.Equal(t, 1, (*result).Value)
		require.Equal(t, true, ok)
		require.Equal(t, []int{}, list.Values())
		require.Nil(t, list.Head())
		require.Nil(t, list.Tail())

	})

	t.Run("Remove/Add random", func(t *testing.T) {
		list := LinkedList[int]()
		list.Add(1)
		list.Add(2)
		require.Equal(t, []int{
			1, 2,
		}, list.Values())

		_, ok := list.Remove(2)
		require.True(t, ok)
		require.Equal(t, []int{
			1,
		}, list.Values())

		_, ok = list.Remove(1)
		require.True(t, ok)
		require.Equal(t, []int{}, list.Values())
		require.Nil(t, list.Head())
		require.Nil(t, list.Tail())

		list.Add(1)
		list.Add(2)
		list.Add(3)
		list.Add(4)
		list.Add(5)

		_, ok = list.Remove(2)
		require.True(t, ok)
		require.Equal(t, []int{1, 3, 4, 5}, list.Values())

		list.Add(6)
		require.True(t, ok)
		require.Equal(t, []int{1, 3, 4, 5, 6}, list.Values())

		_, ok = list.Remove(2)
		require.False(t, ok)
		require.Equal(t, []int{1, 3, 4, 5, 6}, list.Values())

		_, ok = list.Remove(1)
		require.True(t, ok)
		require.Equal(t, []int{3, 4, 5, 6}, list.Values())

	})

	t.Run("Remove non existing", func(t *testing.T) {
		list := LinkedList[int]()
		list.Add(1)
		list.Add(2)
		list.Add(3)
		list.Add(4)
		list.Add(5)
		result, ok := list.Remove(8)
		require.False(t, ok)
		require.Nil(t, result)
		require.Equal(t, []int{
			1, 2, 3, 4, 5,
		}, list.Values())
	})
}

func TestLinkedList_RemoveNode(t *testing.T) {
	t.Run("Remove from middle 1", func(t *testing.T) {
		list := LinkedList[int]()
		list.Add(1)
		list.Add(2)
		node3 := list.Add(3)
		list.Add(4)
		list.Add(5)
		ok := list.RemoveNode(node3)
		require.Equal(t, 1, *list.Head())
		require.Equal(t, 5, *list.Tail())
		require.Equal(t, true, ok)
		require.Equal(t, []int{
			1, 2, 4, 5,
		}, list.Values())
	})

	t.Run("Remove from middle 2", func(t *testing.T) {
		list := LinkedList[int]()
		list.Add(1)
		node2 := list.Add(2)
		list.Add(3)
		list.Add(4)
		list.Add(5)
		ok := list.RemoveNode(node2)
		require.Equal(t, true, ok)
		require.Equal(t, []int{
			1, 3, 4, 5,
		}, list.Values())
		require.Equal(t, 1, *list.Head())
		require.Equal(t, 5, *list.Tail())
	})

	t.Run("Remove from middle 3", func(t *testing.T) {
		list := LinkedList[int]()
		list.Add(1)
		list.Add(2)
		list.Add(3)
		node4 := list.Add(4)
		list.Add(5)
		ok := list.RemoveNode(node4)
		require.Equal(t, true, ok)
		require.Equal(t, []int{
			1, 2, 3, 5,
		}, list.Values())
		require.Equal(t, 1, *list.Head())
		require.Equal(t, 5, *list.Tail())
	})

	t.Run("Remove first", func(t *testing.T) {
		list := LinkedList[int]()
		nodeFirst := list.Add(1)
		list.Add(2)
		list.Add(3)
		list.Add(4)
		list.Add(5)
		ok := list.RemoveNode(nodeFirst)
		require.Equal(t, true, ok)
		require.Equal(t, []int{
			2, 3, 4, 5,
		}, list.Values())
		require.Equal(t, 2, *list.Head())
		require.Equal(t, 5, *list.Tail())
	})

	t.Run("Remove last", func(t *testing.T) {
		list := LinkedList[int]()
		node1 := list.Add(1)
		node2 := list.Add(2)
		node3 := list.Add(3)
		lastNode := list.Add(4)

		ok := list.RemoveNode(lastNode)
		require.Equal(t, true, ok)
		require.Equal(t, []int{
			1, 2, 3,
		}, list.Values())
		require.Equal(t, 1, *list.Head())
		require.Equal(t, 3, *list.Tail())

		ok = list.RemoveNode(node3)
		require.Equal(t, true, ok)
		require.Equal(t, []int{
			1, 2,
		}, list.Values())
		require.Equal(t, 1, *list.Head())
		require.Equal(t, 2, *list.Tail())

		ok = list.RemoveNode(node2)
		require.Equal(t, true, ok)
		require.Equal(t, []int{
			1,
		}, list.Values())
		require.Equal(t, 1, *list.Head())
		require.Equal(t, 1, *list.Tail())

		ok = list.RemoveNode(node1)
		require.Equal(t, true, ok)
		require.Equal(t, []int{}, list.Values())
		require.Nil(t, list.Head())
		require.Nil(t, list.Tail())

	})
}
