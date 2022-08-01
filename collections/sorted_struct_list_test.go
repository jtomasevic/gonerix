package collections

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortedStructList_Add(t *testing.T) {
	sorted := getSortedProductsAsc()

	t.Run("Add ", func(t *testing.T) {
		require.Equal(t, []Product{
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
		}, sorted.Values())
	})
}

func TestSortedStructList_Remove(t *testing.T) {
	sorted := getSortedProductsAsc()
	t.Run("Remove from middle ", func(t *testing.T) {

		result := sorted.Remove(Product{
			Name: "Nike",
		})
		require.True(t, result)
		require.Equal(t, List[Product]{
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
	})
	t.Run("Remove non existing element", func(t *testing.T) {
		result := sorted.Remove(Product{
			Name: "Nike",
		})
		require.False(t, result)
		require.Equal(t, List[Product]{
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
	})
	t.Run("Remove First", func(t *testing.T) {
		result := sorted.Remove(Product{
			Name: "Addidas",
		})
		require.True(t, result)
		require.Equal(t, List[Product]{
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
	})
	t.Run("Remove Last", func(t *testing.T) {
		result := sorted.Remove(Product{
			Name: "Puma",
		})
		require.True(t, result)
		require.Equal(t, List[Product]{
			{
				Name: "Diadora",
			},
			{
				Name: "Nokka",
			},
		}, sorted.elements)
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
		require.Equal(t, List[Product]{}, sorted.elements)
		result = sorted.Remove(Product{
			Name: "Some random product",
		})
		require.False(t, result)
	})
}

func TestSortedStructList_RemoveAt(t *testing.T) {
	sorted := getSortedProductsAsc()

	t.Run("Remove from middle ", func(t *testing.T) {
		result := sorted.RemoveAt(2)
		require.True(t, result)
		require.Equal(t, List[Product]{
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
		require.Equal(t, List[Product]{
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
		result = sorted.RemoveAt(-1)
		require.False(t, result)
	})
	t.Run("Remove at First", func(t *testing.T) {
		result := sorted.RemoveAt(0)
		require.True(t, result)
		require.Equal(t, List[Product]{
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
	})
	t.Run("Remove at Last", func(t *testing.T) {
		result := sorted.RemoveAt(2)
		require.True(t, result)
		require.Equal(t, List[Product]{
			{
				Name: "Diadora",
			},
			{
				Name: "Nokka",
			},
		}, sorted.elements)
	})
	t.Run("Remove at From empty", func(t *testing.T) {
		result := sorted.RemoveAt(0)
		require.True(t, result)
		result = sorted.RemoveAt(0)
		require.True(t, result)
		require.Equal(t, List[Product]{}, sorted.elements)
		require.Equal(t, sorted.Size(), 0)

		result = sorted.RemoveAt(0)
		require.False(t, result)
		result = sorted.RemoveAt(10)
		require.False(t, result)
	})
}

func TestSortedStructListDesc_Add(t *testing.T) {
	sorted := getSortedProductsDesc()

	t.Run("Add ", func(t *testing.T) {
		require.Equal(t, List[Product]{
			{
				Name: "Puma",
			},
			{
				Name: "Nokka",
			},
			{
				Name: "Nike",
			},
			{
				Name: "Diadora",
			},
			{
				Name: "Addidas",
			},
		}, sorted.elements)
	})
}

func TestSortedStructListDesc_Remove(t *testing.T) {
	sorted := getSortedProductsDesc()
	t.Run("Remove from middle ", func(t *testing.T) {

		result := sorted.Remove(Product{
			Name: "Nike",
		})
		require.True(t, result)
		require.Equal(t, List[Product]{
			{
				Name: "Puma",
			},
			{
				Name: "Nokka",
			},
			{
				Name: "Diadora",
			},
			{
				Name: "Addidas",
			},
		}, sorted.elements)
	})
	t.Run("Remove non existing element", func(t *testing.T) {
		result := sorted.Remove(Product{
			Name: "Nike",
		})
		require.False(t, result)
		require.Equal(t, List[Product]{
			{
				Name: "Puma",
			},
			{
				Name: "Nokka",
			},
			{
				Name: "Diadora",
			},
			{
				Name: "Addidas",
			},
		}, sorted.elements)
	})
	t.Run("Remove First", func(t *testing.T) {
		result := sorted.Remove(Product{
			Name: "Puma",
		})
		require.True(t, result)
		require.Equal(t, List[Product]{
			{
				Name: "Nokka",
			},
			{
				Name: "Diadora",
			},
			{
				Name: "Addidas",
			},
		}, sorted.elements)
	})
	t.Run("Remove Last", func(t *testing.T) {
		result := sorted.Remove(Product{
			Name: "Addidas",
		})
		require.True(t, result)
		require.Equal(t, List[Product]{
			{
				Name: "Nokka",
			},
			{
				Name: "Diadora",
			},
		}, sorted.elements)
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
		require.Equal(t, List[Product]{}, sorted.elements)
		result = sorted.Remove(Product{
			Name: "Some random product",
		})
		require.False(t, result)
	})
}

func TestSortedStructListDesc_RemoveAt(t *testing.T) {
	sorted := getSortedProductsDesc()

	t.Run("Remove from middle ", func(t *testing.T) {
		result := sorted.RemoveAt(2)
		require.True(t, result)
		require.Equal(t, List[Product]{
			{
				Name: "Puma",
			},
			{
				Name: "Nokka",
			},
			{
				Name: "Diadora",
			},
			{
				Name: "Addidas",
			},
		}, sorted.elements)
		require.Equal(t, sorted.Size(), 4)
	})
	t.Run("Remove with non existing index", func(t *testing.T) {
		result := sorted.RemoveAt(10)
		require.False(t, result)
		require.Equal(t, List[Product]{
			{
				Name: "Puma",
			},
			{
				Name: "Nokka",
			},
			{
				Name: "Diadora",
			},
			{
				Name: "Addidas",
			},
		}, sorted.elements)
		result = sorted.RemoveAt(-1)
		require.False(t, result)
	})
	t.Run("Remove at First", func(t *testing.T) {
		result := sorted.RemoveAt(0)
		require.True(t, result)
		require.Equal(t, List[Product]{
			{
				Name: "Nokka",
			},
			{
				Name: "Diadora",
			},
			{
				Name: "Addidas",
			},
		}, sorted.elements)
	})
	t.Run("Remove at Last", func(t *testing.T) {
		result := sorted.RemoveAt(2)
		require.True(t, result)
		require.Equal(t, List[Product]{
			{
				Name: "Nokka",
			},
			{
				Name: "Diadora",
			},
		}, sorted.elements)
	})
	t.Run("Remove at From empty", func(t *testing.T) {
		result := sorted.RemoveAt(0)
		require.True(t, result)
		result = sorted.RemoveAt(0)
		require.True(t, result)
		require.Equal(t, List[Product]{}, sorted.elements)
		require.Equal(t, sorted.Size(), 0)

		result = sorted.RemoveAt(0)
		require.False(t, result)
		result = sorted.RemoveAt(10)
		require.False(t, result)
	})
}

func getSortedProductsAsc() sortedStructList[Product] {
	products := getProduct()
	sorted := SortedStructList(func(left Product, right Product) bool {
		return left.Name < right.Name
	})
	for _, p := range products {
		sorted.Add(p)
	}
	return sorted
}

func getSortedProductsDesc() sortedStructList[Product] {
	products := getProduct()
	sorted := SortedStructList(func(left Product, right Product) bool {
		return left.Name > right.Name
	})
	for _, p := range products {
		sorted.Add(p)
	}
	return sorted
}

type Product struct {
	Name string
}

func getProduct() List[Product] {
	return List[Product]{
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
