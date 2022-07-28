package collections

func SortedList[K simpleTypes, V comparable](direction SortDirection) sortedList[K, V] {
	return sortedList[K, V]{
		keys: simpleSortedList[K]{
			direction: direction,
		},
	}
}

// sortedList Represents a collection of key/value pairs that are sorted by the keys and are accessible by key and by index.
type sortedList[K simpleTypes, V comparable] struct {
	keys     simpleSortedList[K]
	elements List[V]
}

// Add an element with the specified key and value to a sortedList object.
// If element with existing key exust it will overried it.
// Under one key it could be only one value.
func (list *sortedList[K, V]) Add(key K, value V) {
	index := list.keys.Add(key)
	list.elements.Insert(index, value)
}

// Remove the element with the specified key from a sortedList object.
func (list *sortedList[K, V]) Remove(key K) bool {
	index := list.keys.RemoveAndReturnIndex(key)
	if index != -1 {
		list.elements.RemoveAt(index)
		return true
	}
	return false
}

func (list *sortedList[K, V]) RemoveAt(position int) bool {
	if position >= list.keys.Count() || position < 0 {
		return false
	}
	// TODO: jt: no need to check result of list.keys.RemoveAt, because condition above is already included in this method.
	//  however, this need refactoring, and probably removeAt will return removed item + if remove was OK.
	list.keys.RemoveAt(position)
	list.elements.RemoveAt(position)
	return true
}

// ToList Return simple List from sorted list elements. 0(1) time complexity.
func (list *sortedList[K, V]) ToList() List[V] {
	return list.elements
}
