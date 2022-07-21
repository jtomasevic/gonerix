package collections

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestList_Integer(t *testing.T) {
	t.Run("Collection of integers: Add, Remove", func(t *testing.T) {
		list := initIntList()
		require.Equal(t, len(list), 5)
		for i, el := range list {
			require.Equal(t, el, i+1)
		}
		// remove from middle
		res := list.Remove(2)
		require.Equal(t, len(list), 4)
		require.True(t, res)
		// remove first element
		list = initIntList()
		res = list.Remove(1)
		require.Equal(t, len(list), 4)
		require.Equal(t, res, true)
		// remove last
		list = initIntList()
		res = list.Remove(5)
		require.Equal(t, len(list), 4)
		require.Equal(t, res, true)
		// remove no-existing element, expect false
		res = list.Remove(12)
		list = initIntList()
		require.Equal(t, len(list), 5)
		require.Equal(t, res, false)
	})

	t.Run("Collection of integers: Remove at ", func(t *testing.T) {
		list := initIntList()
		require.Equal(t, len(list), 5)
		// remove first element
		removeIndex := 0
		list.RemoveAt(removeIndex)
		require.Equal(t, len(list), 4)
		require.Equal(t, list[0], 2)
		fmt.Printf("Remove first element: %v\n", list)

		// Remove last element
		list = initIntList()
		removeIndex = 4
		list.RemoveAt(removeIndex)
		require.Equal(t, len(list), 4)
		require.Equal(t, list[len(list)-1], 4)
		fmt.Printf("Remove last elementt: %v\n", list)

		// Remove from the middle
		list = initIntList()
		removeIndex = 2
		list.RemoveAt(removeIndex)
		require.Equal(t, len(list), 4)
		require.Equal(t, list[2], 4)
		fmt.Printf("Remove at index %v, %v\n", removeIndex, list)

		// Remove at, index out of range
		list = initIntList()
		removeIndex = 12
		res := list.RemoveAt(removeIndex)
		require.Equal(t, len(list), 5)
		require.Equal(t, res, false)
	})
	t.Run("Collection of integers: Insert ", func(t *testing.T) {
		list := initIntList()
		require.Equal(t, len(list), 5)
		// insert as first element
		forInsert := 0
		list.Insert(0, forInsert)
		require.Equal(t, len(list), 6)
		require.Equal(t, list[0], forInsert)
		// insert as last element
		list = initIntList()
		forInsert = 6
		list.Insert(5, forInsert)
		require.Equal(t, len(list), 6)
		require.Equal(t, list[5], forInsert)
		// insert in the middle
		list = initIntList()
		forInsert = 12
		list.Insert(2, forInsert)
		require.Equal(t, len(list), 6)
		require.Equal(t, list[2], forInsert)
		fmt.Println(list)

		// insert out of index
		list = initIntList()
		forInsert = 12
		res := list.Insert(8, forInsert)
		require.Equal(t, len(list), 5)
		require.Equal(t, res, false)
	})
	t.Run("Reverse", func(t *testing.T) {
		list := initIntList()
		list.Reverse()
		require.Equal(t, list[0], 5)
		require.Equal(t, list[1], 4)
		require.Equal(t, list[2], 3)
		require.Equal(t, list[3], 2)
		require.Equal(t, list[4], 1)

	})

	t.Run("Find", func(t *testing.T) {
		list := initIntList()
		result := list.Find(func (el int) bool {return el == -1})
		require.Equal(t, len(result), 0)
		result = list.Find(func (el int) bool {return el == 5})
		require.Equal(t, len(result), 1)
		list.Add(3)
		list.Add(3)
		list.Add(3)
		result = list.Find(func (el int) bool {return el == 3})
		require.Equal(t, len(result), 4)
	})
}

// initIntList return list with 5 members where values are going from 1 to 5.
func initIntList() List[int] {
	list := List[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	return list
}

// func initPlayerList() List[Player] {
// 	list := List[Player]{}
// 	list.Add(Player{})
// 	return list
// }

type Player struct {
	Name string
}
