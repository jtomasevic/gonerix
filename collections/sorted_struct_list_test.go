package collections

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortedStructList_AddRemove(t *testing.T) {
	products := getProduct()
	sorted := SortedStructList(func(left Product, right Product) bool {
		return left.Name < right.Name
	})
	for _, p := range products {
		sorted.Add(p)
	}
	require.Equal(t, sorted.elements, []Product{
		{
			Name: "Addidas",
		},
		{
			Name: "Diadora",
		},
		{
			Name: "Nike",
		},
		{
			Name: "Nokka",
		},
		{
			Name: "Puma",
		},
	})
}
