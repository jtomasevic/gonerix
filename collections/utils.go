package collections

func RemoveAt[T comparable](list []T, position int) ([]T, bool) {
	if position >= len(list) || position < 0 {
		return list, false
	}
	return append(list[:position], list[position+1:]...), true
}
