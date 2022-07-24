package collections

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortedList_Integer(t *testing.T) {
	t.Run("Add / Remove", func(t *testing.T) {
		list := SimpleSortedList[int]{}
		list.Add(3)
		list.Add(1)
		list.Add(2)
		list.Add(2)
		list.Add(10)
		list.Add(9)
		list.Add(5)
		list.Add(5)
		list.Add(3)
		list.Add(7)
		list.Add(4)
		list.Add(6)
		list.Add(8)
		list.Add(2)

		currentList := List[int]{
			1, 2, 2, 2, 3, 3, 4, 5, 5, 6, 7, 8, 9, 10,
		}
		require.Equal(t, currentList, list.ToList())
		res := list.Remove(9)
		require.True(t, res)
		require.Equal(
			t,
			list.ToList(),
			List[int]{
				1, 2, 2, 2, 3, 3, 4, 5, 5, 6, 7, 8, 10,
			},
		)

		res = list.Remove(2)
		require.True(t, res)
		require.Equal(
			t,
			list.ToList(),
			List[int]{
				1, 2, 2, 3, 3, 4, 5, 5, 6, 7, 8, 10,
			},
		)

		fmt.Println("remove 5")
		res = list.Remove(5)
		require.True(t, res)
		require.Equal(
			t,
			list.ToList(),
			List[int]{
				1, 2, 2, 3, 3, 4, 5, 6, 7, 8, 10,
			},
		)

		fmt.Println("remove 10")
		list.Remove(10)
		require.True(t, res)
		require.Equal(
			t,
			list.ToList(),
			List[int]{
				1, 2, 2, 3, 3, 4, 5, 6, 7, 8,
			},
		)

		fmt.Println("remove 1")
		res = list.Remove(1)
		require.True(t, res)
		require.Equal(
			t,
			list.ToList(),
			List[int]{
				2, 2, 3, 3, 4, 5, 6, 7, 8,
			},
		)

		fmt.Println("remove 8")
		res = list.Remove(8)
		require.True(t, res)
		require.Equal(
			t,
			list.ToList(),
			List[int]{
				2, 2, 3, 3, 4, 5, 6, 7,
			},
		)

		fmt.Println("remove 4 ")
		res = list.Remove(4)
		require.True(t, res)
		require.Equal(
			t,
			list.ToList(),
			List[int]{
				2, 2, 3, 3, 5, 6, 7,
			},
		)
		fmt.Println("remove 4 not existing")
		res = list.Remove(4)
		require.False(t, res)
		require.Equal(
			t,
			list.ToList(),
			List[int]{
				2, 2, 3, 3, 5, 6, 7,
			},
		)
		fmt.Println("remove 3")
		res = list.Remove(3)
		require.True(t, res)
		require.Equal(
			t,
			list.ToList(),
			List[int]{
				2, 2, 3, 5, 6, 7,
			},
		)
		fmt.Println("remove 7")

		res = list.Remove(7)
		require.True(t, res)
		require.Equal(
			t,
			list.ToList(),
			List[int]{
				2, 2, 3, 5, 6,
			},
		)

		fmt.Println("remove 6")
		res = list.Remove(6)
		require.True(t, res)
		require.Equal(
			t,
			list.ToList(),
			List[int]{
				2, 2, 3, 5,
			},
		)

		fmt.Println("remove 5")
		res = list.Remove(5)
		require.True(t, res)
		require.Equal(
			t,
			list.ToList(),
			List[int]{
				2, 2, 3,
			},
		)
		fmt.Println("remove non existing 5")
		res = list.Remove(5)
		require.False(t, res)
		require.Equal(
			t,
			list.ToList(),
			List[int]{
				2, 2, 3,
			},
		)
		res = list.Remove(2)
		require.True(t, res)
		require.Equal(
			t,
			list.ToList(),
			List[int]{
				2, 3,
			},
		)

		res = list.Remove(2)
		require.True(t, res)
		require.Equal(
			t,
			list.ToList(),
			List[int]{
				3,
			},
		)
		res = list.Remove(3)
		require.True(t, res)
		require.Equal(
			t,
			list.ToList(),
			List[int]{},
		)
		res = list.Remove(2)
		require.False(t, res)
		require.Equal(
			t,
			list.ToList(),
			List[int]{},
		)
	})
	t.Run("Last member", func(t *testing.T) {
		list := SimpleSortedList[int]{}
		list.Add(3)
		list.Add(1)
		list.Add(2)
		list.Add(10)
		list.Add(9)
		list.Add(5)
		list.Add(7)
		list.Add(4)
		list.Add(6)
		list.Add(8)
		currentList := List[int]{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		}
		require.Equal(t, currentList, list.ToList())

		last := list.Last()
		first := list.First()
		require.Equal(t, 10, *last)
		require.Equal(t, 1, *first)
		list = SimpleSortedList[int]{}
		last = list.Last()
		require.Nil(t, last)
		first = list.First()
		require.Nil(t, first)
	})
}

func TestSimpleSortedList_RemoveAt(t *testing.T) {
	list := SimpleSortedList[int]{}
	list.Add(3)
	list.Add(1)
	list.Add(2)
	list.Add(10)
	list.Add(9)
	list.Add(5)
	list.Add(7)
	list.Add(4)
	list.Add(6)
	list.Add(8)
	currentList := List[int]{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	require.Equal(t, currentList, list.ToList())

	response := list.RemoveAt(0)
	require.True(t, response)
	require.Equal(t, List[int]{
		2, 3, 4, 5, 6, 7, 8, 9, 10,
	}, list.ToList())

	response = list.RemoveAt(9)
	require.False(t, response)
	require.Equal(t, List[int]{
		2, 3, 4, 5, 6, 7, 8, 9, 10,
	}, list.ToList())

	response = list.RemoveAt(-2)
	require.False(t, response)
	require.Equal(t, List[int]{
		2, 3, 4, 5, 6, 7, 8, 9, 10,
	}, list.ToList())

	response = list.RemoveAt(8)
	require.True(t, response)
	require.Equal(t, List[int]{
		2, 3, 4, 5, 6, 7, 8, 9,
	}, list.ToList())

	response = list.RemoveAt(2)
	require.True(t, response)
	require.Equal(t, List[int]{
		2, 3, 5, 6, 7, 8, 9,
	}, list.ToList())

	response = list.RemoveAt(2)
	require.True(t, response)
	require.Equal(t, List[int]{
		2, 3, 6, 7, 8, 9,
	}, list.ToList())

}
