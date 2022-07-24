package collections

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDictionary(t *testing.T) {
	t.Run("Dictionary basic + remove", func(t *testing.T) {
		dic := Dictionary[int, string]{}
		dic[1] = "one"
		dic[2] = "two"
		dic[3] = "three"
		require.Equal(t, dic.Length(), 3)
		require.Equal(t, dic[1], "one")
		require.Equal(t, dic[2], "two")
		require.Equal(t, dic[3], "three")

		exist := dic.Exist(1)
		require.True(t, exist)

		exist = dic.Exist(10)
		require.False(t, exist)

		isOk, result := dic.Remove(1)
		require.True(t, isOk)
		require.Equal(t, result, "one")

		isOk, result = dic.Remove(90)
		require.False(t, isOk)
		require.Equal(t, result, "")
	})
}