package collections

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortedStructList_Add(t *testing.T) {
	sorted := getSortedProducts()

	t.Run("Add ", func(t *testing.T) {
		require.Equal(t, sorted.Values(), []Product{
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
}
func TestSortedStructList_Remove(t *testing.T) {
	sorted := getSortedProducts()
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

func TestSortedStructList_RemoveAt(t *testing.T) {
	sorted := getSortedProducts()

	t.Run("Remove from middle ", func(t *testing.T) {
		result := sorted.RemoveAt(2)
		require.True(t, result)
		require.Equal(t, []Product{
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
		}, sorted.elements)
		require.Equal(t, sorted.Size(), 4)
	})
	t.Run("Remove with non existing index", func(t *testing.T) {
		result := sorted.RemoveAt(10)
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
		result = sorted.RemoveAt(-1)
		require.False(t, result)
	})
	t.Run("Remove at First", func(t *testing.T) {
		result := sorted.RemoveAt(0)
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
	t.Run("Remove at Last", func(t *testing.T) {
		result := sorted.RemoveAt(2)
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
	t.Run("Remove at From empty", func(t *testing.T) {
		result := sorted.RemoveAt(0)
		require.True(t, result)
		result = sorted.RemoveAt(0)
		require.True(t, result)
		require.Equal(t, sorted.elements, []Product{})
		require.Equal(t, sorted.Size(), 0)

		result = sorted.RemoveAt(0)
		require.False(t, result)
		result = sorted.RemoveAt(10)
		require.False(t, result)
	})
}

func getSortedProducts() sortedStructList[Product] {
	products := getProduct()
	sorted := SortedStructList(func(left Product, right Product) bool {
		return left.Name < right.Name
	})
	for _, p := range products {
		sorted.Add(p)
	}
	return sorted
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
