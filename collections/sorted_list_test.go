package collections

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortedList_AddRemove(t *testing.T) {
	products := getProduct()
	sorted := SortedList[string, Product]{}
	for _, p := range products {
		sorted.Add(p.Name, p)
	}
	require.Equal(t, sorted.keys.ToList(), List[string]{
		"Addidas",
		"Diadora",
		"Nike",
		"Nokka",
		"Puma",
	})

	products = getProduct()
	sorted = SortedList[string, Product]{}
	for _, p := range products {
		sorted.Add(p.Name, p)
	}
	sorted.Remove("Nike")
	require.Equal(t, sorted.keys.ToList(), List[string]{
		"Addidas",
		"Diadora",
		"Nokka",
		"Puma",
	})
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

	products = getProduct()
	sorted = SortedList[string, Product]{}
	for _, p := range products {
		sorted.Add(p.Name, p)
	}
	require.Equal(t, sorted.Values(), []Product{
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
	}, sorted)

}
