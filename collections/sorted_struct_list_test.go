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
	t.Run("Add ", func(t *testing.T) {
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
	})

	t.Run("Remove from middle ", func(t *testing.T) {
		result := sorted.Remove(Product{
			Name: "Nike",
		})
		require.True(t, result)
		require.Equal(t, sorted.elements, []Product{
			{
				Name: "Addidas",
			},
			{
				Name: "Diadora",
			},
			{
				Name: "Nokka",
			},
			{
				Name: "Puma",
			},
		})
	})
	t.Run("Remove non existing element", func(t *testing.T) {
		result := sorted.Remove(Product{
			Name: "Nike",
		})
		require.False(t, result)
		require.Equal(t, sorted.elements, []Product{
			{
				Name: "Addidas",
			},
			{
				Name: "Diadora",
			},
			{
				Name: "Nokka",
			},
			{
				Name: "Puma",
			},
		})
	})
	t.Run("Remove First", func(t *testing.T) {
		result := sorted.Remove(Product{
			Name: "Addidas",
		})
		require.True(t, result)
		require.Equal(t, sorted.elements, []Product{
			{
				Name: "Diadora",
			},
			{
				Name: "Nokka",
			},
			{
				Name: "Puma",
			},
		})
	})
	t.Run("Remove Last", func(t *testing.T) {
		result := sorted.Remove(Product{
			Name: "Puma",
		})
		require.True(t, result)
		require.Equal(t, sorted.elements, []Product{
			{
				Name: "Diadora",
			},
			{
				Name: "Nokka",
			},
		})
	})
	t.Run("Remove From empty", func(t *testing.T) {
		result := sorted.Remove(Product{
			Name: "Diadora",
		})
		require.True(t, result)
		result = sorted.Remove(Product{
			Name: "Nokka",
		})
		require.True(t, result)
		require.Equal(t, sorted.elements, []Product{})
		result = sorted.Remove(Product{
			Name: "Some random product",
		})
		require.False(t, result)
	})
}

type Product struct {
	Name string
}

func getProduct() []Product {
	return []Product{
		{
			Name: "Diadora",
		},
		{
			Name: "Nike",
		},
		{
			Name: "Addidas",
		},
		{
			Name: "Puma",
		},
		{
			Name: "Nokka",
		},
	}
}
