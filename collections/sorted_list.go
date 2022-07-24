package collections

// SortedList Represents a collection of key/value pairs that are sorted by the keys and are accessible by key and by index.
type SortedList[K simpleTypes, V any] struct {
	keys   SimpleSortedList[K]
	values Dictionary[K, V]
}

// Add an element with the specified key and value to a SortedList object.
// If element with existing key exust it will overried it.
// Under one key it could be only one value.
func (list *SortedList[K, V]) Add(key K, value V) {
	if list.values == nil {
		list.values = Dictionary[K, V]{}
	}
	list.keys.Add(key)
	list.values[key] = value
}

// Remove the element with the specified key from a SortedList object.
func (list *SortedList[K, V]) Remove(key K) bool {
	res := list.keys.Remove(key)
	if res {
		list.values.Remove(key)
		return true
	}
	return false
}

func (list *SortedList[K, V]) RemoveAt(position int) bool {
	if position >= len(list.keys) || position < 0 {
		return false
	}
	key := list.keys[position]
	// TODO: jt: no need to check result of list.keys.RemoveAt, because condition above is already included in this method.
	//  however, this need refactoring, and probably removeAt will return removed item + if remove was OK.
	list.keys.RemoveAt(position)
	list.values.Remove(key)
	return true
}

// Get sorted values.
// TODO: this need optimisation.
func (list *SortedList[K, V]) Values() []V {
	values := []V{}
	for _, key := range list.keys {
		values = append(values, list.values[key])
	}
	return values
}
