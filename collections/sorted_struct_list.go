package collections

// CompareFunction
// - If left parameter is lower than right parameter, this is ASC ordered list.
// - If left parameter is higher than right parameter, this is DESC ordered list.
type CompareFunction[T comparable] func(left T, right T) bool

// SortedStructList create
func SortedStructList[T comparable](comparableFunction CompareFunction[T]) sortedStructList[T] {
	return sortedStructList[T]{
		compare: comparableFunction,
	}
}

type sortedStructList[T comparable] struct {
	elements List[T]
	compare  CompareFunction[T]
}

// Add element to sorted list. It is duplicate tolerant.
func (list *sortedStructList[T]) Add(element T) {
	newList := list.elements
	low := 0
	high := len(newList) - 1
	for low <= high {
		median := (low + high) / 2
		if list.compare(newList[median], element) {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	list.elements.Insert(low, element)
}

// Remove element from sorted list. If there is multiple values, removes only first one.
//  If element cannot be find return false.
func (list *sortedStructList[T]) Remove(element T) bool {
	if len(list.elements) == 0 {
		return false
	}
	newList := list.elements
	low := 0
	high := len(newList) - 1
	for low <= high {
		median := (low + high) / 2
		if list.compare(newList[median], element) {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	if low < len(newList) && newList[low] == element {
		list.elements.RemoveAt(low)
		//newList = append(newList[:low], newList[low+1:]...)
		//list.elements = newList
		return true
	}
	return false
}

// RemoveAt Removes the element at the specified index of the list
func (list *sortedStructList[T]) RemoveAt(position int) bool {
	newList, removed := RemoveAt(list.elements, position)
	if removed {
		(*list).elements = newList
		return true
	}
	return false
}

// Values return list of sorted structs.
func (list *sortedStructList[T]) Values() []T {
	return (*list).elements
}

// Size number of elements in collection
func (list *sortedStructList[T]) Size() int {
	return len(list.elements)
}
