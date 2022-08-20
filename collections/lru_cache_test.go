package collections

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLruCache_Set(t *testing.T) {
	t.Run("Just set", func(t *testing.T) {
		cache := LRUCache[int, string](10)
		cache.Set(1, "1")
		cache.Set(2, "2")
		cache.Set(3, "3")
		require.Equal(t, []string{"1", "2", "3"}, cache.GetCacheValues())
	})
	t.Run("Set and override", func(t *testing.T) {
		cache := LRUCache[int, string](10)
		cache.Set(1, "1")
		cache.Set(2, "2")
		cache.Set(3, "3")
		cache.Set(1, "1")
		require.Equal(t, []string{"2", "3", "1"}, cache.GetCacheValues())
	})
	t.Run("Set over the limit", func(t *testing.T) {
		cache := LRUCache[int, string](3)
		cache.Set(1, "1")
		cache.Set(2, "2")
		cache.Set(3, "3")
		cache.Set(4, "4")
		require.Equal(t, []string{"2", "3", "4"}, cache.GetCacheValues())
	})
}

func TestLruCache_Get(t *testing.T) {
	t.Run("Set and get from middle", func(t *testing.T) {
		cache := LRUCache[int, string](10)
		cache.Set(1, "1")
		cache.Set(2, "2")
		cache.Set(3, "3")
		value := cache.Get(2)
		valueNotExist := cache.Get(10)
		require.Equal(t, "2", *value)
		require.Nil(t, valueNotExist)
		require.Equal(t, []string{"1", "3", "2"}, cache.GetCacheValues())
	})
	t.Run("Set and get first added", func(t *testing.T) {
		cache := LRUCache[int, string](10)
		cache.Set(1, "1")
		cache.Set(2, "2")
		cache.Set(3, "3")
		value := cache.Get(1)
		require.Equal(t, "1", *value)
		require.Equal(t, []string{"2", "3", "1"}, cache.GetCacheValues())
	})
	t.Run("Set and get last one", func(t *testing.T) {
		cache := LRUCache[int, string](3)
		cache.Set(1, "1")
		cache.Set(2, "2")
		cache.Set(3, "3")
		value := cache.Get(3)
		require.Equal(t, "3", *value)
		require.Equal(t, []string{"1", "2", "3"}, cache.GetCacheValues())
	})
}

func TestLRUCache_Capacity(t *testing.T) {
	t.Run("Just adding", func(t *testing.T) {
		cache := LRUCache[int, string](3)
		cache.Set(1, "1")
		cache.Set(2, "2")
		cache.Set(3, "3")
		cache.Set(4, "4")
		require.Equal(t, []string{"2", "3", "4"}, cache.GetCacheValues())
	})
	t.Run("Random set with override", func(t *testing.T) {
		cache := LRUCache[int, string](3)
		cache.Set(1, "1")
		cache.Set(2, "2")
		cache.Set(3, "3")
		cache.Set(4, "4")
		require.Equal(t, []string{"2", "3", "4"}, cache.GetCacheValues())
		cache.Set(2, "2")
		require.Equal(t, []string{"3", "4", "2"}, cache.GetCacheValues())
		cache.Set(5, "5")
		require.Equal(t, []string{"4", "2", "5"}, cache.GetCacheValues())

	})
}
