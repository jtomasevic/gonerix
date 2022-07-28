package collections

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimpleSortedListAsc_AddRemove(t *testing.T) {
	t.Run("Add / Remove", func(t *testing.T) {
		list := SimpleSortedList[int](ASC)
		index := list.Add(3)
		require.Equal(t, 0, index)

		index = list.Add(1)
		require.Equal(t, 0, index)

		index = list.Add(2)
		require.Equal(t, 1, index)

		index = list.Add(2)
		require.Equal(t, 1, index)

		index = list.Add(10)
		require.Equal(t, 4, index)

		index = list.Add(9)
		require.Equal(t, 4, index)

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
		require.Equal(t, list.Count(), len(currentList))
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
		list := SimpleSortedList[int](ASC)
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
		list = SimpleSortedList[int](ASC)
		last = list.Last()
		require.Nil(t, last)
		first = list.First()
		require.Nil(t, first)
	})
}

func TestSimpleSortedListAsc_RemoveAt(t *testing.T) {
	list := SimpleSortedList[int](ASC)
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

func TestSimpleSortedListAscDesc_AddRemove(t *testing.T) {

	t.Run("Adding", func(t *testing.T) {
		list := initIntDescList()

		currentList := List[int]{
			10, 9, 8, 7, 6, 5, 5, 4, 3, 3, 2, 2, 2, 1,
		}
		require.Equal(t, currentList, list.ToList())
	})

	t.Run("Removing ", func(t *testing.T) {
		list := initIntDescList()
		// 10, 9, 8, 7, 6, 5, 5, 4, 3, 3, 2, 2, 2, 1,

		res := list.Remove(9)
		require.True(t, res)
		require.Equal(
			t,
			List[int]{
				10, 8, 7, 6, 5, 5, 4, 3, 3, 2, 2, 2, 1,
			},
			list.ToList(),
		)

		// 10, 8, 7, 6, 5, 5, 4, 3, 3, 2, 2, 2, 1,
		res = list.Remove(2)
		require.True(t, res)
		require.Equal(
			t,
			list.ToList(),
			List[int]{
				10, 8, 7, 6, 5, 5, 4, 3, 3, 2, 2, 1,
			},
		)

		// 10, 8, 7, 6, 5, 5, 4, 3, 3, 2, 2, 1,
		res = list.Remove(5)
		require.True(t, res)
		require.Equal(
			t,
			list.ToList(),
			List[int]{
				10, 8, 7, 6, 5, 4, 3, 3, 2, 2, 1,
			},
		)

		// 10, 8, 7, 6, 5, 4, 3, 3, 2, 2, 1,
		fmt.Println("remove 10")
		list.Remove(10)
		require.True(t, res)
		require.Equal(
			t,
			list.ToList(),
			List[int]{
				8, 7, 6, 5, 4, 3, 3, 2, 2, 1,
			},
		)
		// 8, 7, 6, 5, 4, 3, 3, 2, 2, 1,
		fmt.Println("remove 1")
		res = list.Remove(1)
		require.True(t, res)
		require.Equal(
			t,
			list.ToList(),
			List[int]{
				8, 7, 6, 5, 4, 3, 3, 2, 2,
			},
		)

		// 8, 7, 6, 5, 4, 3, 3, 2, 2,
		fmt.Println("remove 8")
		res = list.Remove(8)
		require.True(t, res)
		require.Equal(
			t,
			list.ToList(),
			List[int]{
				7, 6, 5, 4, 3, 3, 2, 2,
			},
		)

		// remove non existing
		res = list.Remove(14)
		require.False(t, res)

	})

	t.Run("Add / Remove", func(t *testing.T) {

	})
	t.Run("Last member", func(t *testing.T) {
		list := SimpleSortedList[int](DESC)
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
			10, 9, 8, 7, 6, 5, 4, 3, 2, 1,
		}
		require.Equal(t, currentList, list.ToList())

		last := list.Last()
		first := list.First()
		require.Equal(t, 1, *last)
		require.Equal(t, 10, *first)
		list = SimpleSortedList[int](DESC)
		last = list.Last()
		require.Nil(t, last)
		first = list.First()
		require.Nil(t, first)
	})
}

func TestSimpleSortedListAscDesc_RemoveAt(t *testing.T) {
	list := initIntDescList()
	// 10, 9, 8, 7, 6, 5, 5, 4, 3, 3, 2, 2, 2, 1,

	response := list.RemoveAt(0)
	require.True(t, response)
	require.Equal(t, List[int]{
		9, 8, 7, 6, 5, 5, 4, 3, 3, 2, 2, 2, 1,
	}, list.ToList())

	response = list.RemoveAt(9)
	require.True(t, response)
	require.Equal(t, List[int]{
		9, 8, 7, 6, 5, 5, 4, 3, 3, 2, 2, 1,
	}, list.ToList())

	response = list.RemoveAt(-2)
	require.False(t, response)

	response = list.RemoveAt(8)
	require.True(t, response)
	require.Equal(t, List[int]{
		9, 8, 7, 6, 5, 5, 4, 3, 2, 2, 1,
	}, list.ToList())

	response = list.RemoveAt(2)
	require.True(t, response)
	require.Equal(t, List[int]{
		9, 8, 6, 5, 5, 4, 3, 2, 2, 1,
	}, list.ToList())

	response = list.RemoveAt(2)
	require.True(t, response)
	require.Equal(t, List[int]{
		9, 8, 5, 5, 4, 3, 2, 2, 1,
	}, list.ToList())

	response = list.RemoveAt(8)
	require.True(t, response)
	require.Equal(t, List[int]{
		9, 8, 5, 5, 4, 3, 2, 2,
	}, list.ToList())

}

func TestSimpleSortedList_CountFirstLast(t *testing.T) {
	list := SimpleSortedList[string](ASC)
	list.Add("Knowledge")
	list.Add("Argument")
	list.Add("Row")
	list.Add("Beep")
	list.Add("Stone")
	require.Equal(t, "Argument", *list.First())
	require.Equal(t, "Stone", *list.Last())
	require.Equal(t, 5, list.Count())
	require.False(t, list.IsEmpty())

	list = SimpleSortedList[string](ASC)
	require.Nil(t, list.First())
	require.Nil(t, list.Last())
	require.True(t, list.IsEmpty())
}

func initIntDescList() simpleSortedList[int] {
	list := SimpleSortedList[int](DESC)
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
	return list
}

func initIntAscList() simpleSortedList[int] {
	list := SimpleSortedList[int](ASC)
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
	return list
}

func TestSimpleSortedList_RemoveAndReturnIndex(t *testing.T) {
	sorted := SimpleSortedList[int](ASC)
	sorted.Add(35)
	sorted.Add(45)
	sorted.Add(15)
	sorted.Add(25)
	sorted.Add(55)

	index := sorted.RemoveAndReturnIndex(25)
	require.Equal(t, List[int]{
		15,
		35,
		45,
		55,
	}, sorted.ToList())
	require.Equal(t, index, 1)

	index = sorted.RemoveAndReturnIndex(105)
	require.Equal(t, List[int]{
		15,
		35,
		45,
		55,
	}, sorted.ToList())
	require.Equal(t, index, -1)

	index = sorted.RemoveAndReturnIndex(15)
	require.Equal(t, List[int]{
		35,
		45,
		55,
	}, sorted.ToList())
	require.Equal(t, index, 0)

	index = sorted.RemoveAndReturnIndex(10)
	require.Equal(t, List[int]{
		35,
		45,
		55,
	}, sorted.ToList())
	require.Equal(t, index, -1)

	index = sorted.RemoveAndReturnIndex(45)
	require.Equal(t, index, 1)

	index = sorted.RemoveAndReturnIndex(35)
	require.Equal(t, index, 0)

	index = sorted.RemoveAndReturnIndex(55)
	require.Equal(t, index, 0)

	require.True(t, sorted.IsEmpty())

	index = sorted.RemoveAndReturnIndex(55)
	require.Equal(t, index, -1)

}
