package collections

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortedList_AddRemove(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		products := getProduct()
		sorted := SortedList[string, Product](ASC)
		for _, p := range products {
			sorted.Add(p.Name, p)
		}
		fmt.Printf("!!!!! %v\n", sorted.keys)

		require.Equal(t, List[string]{
			"Addidas",
			"Diadora",
			"Nike",
			"Nokka",
			"Puma",
		}, sorted.keys.ToList())

		require.Equal(t, List[Product]{
			{
				Name: "Addidas",
			}, {
				Name: "Diadora",
			}, {
				Name: "Nike",
			}, {
				Name: "Nokka",
			}, {
				Name: "Puma",
			},
		}, sorted.ToList())
	})
	t.Run("Remove", func(t *testing.T) {
		products := getProduct()
		sorted := SortedList[string, Product](ASC)
		for _, p := range products {
			sorted.Add(p.Name, p)
		}
		sorted.Remove("Nike")
		require.Equal(t, List[string]{
			"Addidas",
			"Diadora",
			"Nokka",
			"Puma",
		}, sorted.keys.ToList())
		sorted.Remove("Puma")
		sorted.Remove("Addidas")
		require.Equal(t, sorted.keys.ToList(), List[string]{
			"Diadora",
			"Nokka",
		})

		res := sorted.Remove("Addidas")
		require.Equal(t, sorted.keys.ToList(), List[string]{
			"Diadora",
			"Nokka",
		})
		require.False(t, res)

		sorted.Remove("Diadora")
		sorted.Remove("Nokka")
		require.Equal(t, sorted.keys.ToList(), List[string]{})
		res = sorted.Remove("Nokka")
		require.False(t, res)
		require.Equal(t, sorted.keys.ToList(), List[string]{})
	})
	t.Run("RemoveAt", func(t *testing.T) {
		products := getProduct()
		sorted := SortedList[string, Product](ASC)
		for _, p := range products {
			sorted.Add(p.Name, p)
		}
		sorted.RemoveAt(1)
		require.Equal(t, List[string]{
			"Addidas",
			"Nike",
			"Nokka",
			"Puma",
		}, sorted.keys.ToList())
		result := sorted.RemoveAt(15)
		require.False(t, result)
		sorted = SortedList[string, Product](ASC)
		result = sorted.RemoveAt(0)
		require.False(t, result)
		result = sorted.RemoveAt(+1)
		require.False(t, result)

	})
}

func TestSortedListDesc_AddRemove(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		products := getProduct()
		sorted := SortedList[string, Product](DESC)
		for _, p := range products {
			sorted.Add(p.Name, p)
		}
		fmt.Printf("!!!!! %v\n", sorted.keys)

		require.Equal(t, List[string]{
			"Puma",
			"Nokka",
			"Nike",
			"Diadora",
			"Addidas",
		}, sorted.keys.ToList())
	})
	t.Run("Remove", func(t *testing.T) {
		products := getProduct()
		sorted := SortedList[string, Product](DESC)
		for _, p := range products {
			sorted.Add(p.Name, p)
		}
		sorted.Remove("Nike")
		require.Equal(t, List[string]{
			"Puma",
			"Nokka",
			"Diadora",
			"Addidas",
		}, sorted.keys.ToList())
		sorted.Remove("Puma")
		sorted.Remove("Addidas")
		require.Equal(t, sorted.keys.ToList(), List[string]{
			"Nokka",
			"Diadora",
		})

		res := sorted.Remove("Addidas")
		require.Equal(t, sorted.keys.ToList(), List[string]{
			"Nokka",
			"Diadora",
		})
		require.False(t, res)

		sorted.Remove("Diadora")
		sorted.Remove("Nokka")
		require.Equal(t, sorted.keys.ToList(), List[string]{})
		res = sorted.Remove("Nokka")
		require.False(t, res)
		require.Equal(t, sorted.keys.ToList(), List[string]{})
	})
	t.Run("RemoveAt", func(t *testing.T) {
		products := getProduct()
		sorted := SortedList[string, Product](DESC)
		for _, p := range products {
			sorted.Add(p.Name, p)
		}
		sorted.RemoveAt(1)
		require.Equal(t, List[string]{
			"Puma",
			"Nike",
			"Diadora",
			"Addidas",
		}, sorted.keys.ToList())
		sorted.RemoveAt(0)
		require.Equal(t, List[string]{
			"Nike",
			"Diadora",
			"Addidas",
		}, sorted.keys.ToList())
		sorted.RemoveAt(2)
		require.Equal(t, List[string]{
			"Nike",
			"Diadora",
		}, sorted.keys.ToList())

		result := sorted.RemoveAt(15)
		require.False(t, result)

		sorted = SortedList[string, Product](DESC)
		result = sorted.RemoveAt(0)
		require.False(t, result)
		result = sorted.RemoveAt(+1)
		require.False(t, result)

	})
}
