package collections

// IsLower returns true if left parameter is lower than right parameter
type IsLower[T comparable] func(left T, right T) bool

// SortedStructList create
func SortedStructList[T comparable](comparableFunction IsLower[T]) sortedStructList[T] {
	return sortedStructList[T]{
		isLower: comparableFunction,
	}
}

type sortedStructList[T comparable] struct {
	elements []T
	isLower  IsLower[T]
}

// Add element to sorted list. It is duplicate tolerant.
func (list *sortedStructList[T]) Add(element T) {

	newList := list.elements
	low := 0
	high := len(newList) - 1
	for low <= high {
		median := (low + high) / 2
		if list.isLower(newList[median], element) {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	elems := append([]T{element}, newList[low:]...)
	elems = append(newList[:low], elems...)
	list.elements = elems
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
		if list.isLower(newList[median], element) {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	if low < len(newList) && newList[low] == element {
		newList = append(newList[:low], newList[low+1:]...)
		list.elements = newList
		return true
	}
	return false
}

// RemoveAt Removes the element at the specified index of the list
func (list *sortedStructList[T]) RemoveAt(position int) bool {
	if position >= len(list.elements) || position < 0 {
		return false
	}
	newList := list.elements
	newList = append(newList[:position], newList[position+1:]...)
	(*list).elements = newList
	return true
}

// Values return list of sorted structs.
func (list *sortedStructList[T]) Values() []T {
	return (*list).elements
}

// Size number of elements in collection
func (list *sortedStructList[T]) Size() int {
	return len(list.elements)
}
