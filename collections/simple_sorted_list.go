package collections

import "fmt"

// Types supported by Sorted List
type simpleTypes interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

type SortDirection string

const (
	ASC  SortDirection = "asc"
	DESC SortDirection = "desc"
)

func SimpleSortedList[T simpleTypes](direction SortDirection) simpleSortedList[T] {
	return simpleSortedList[T]{
		elements:  List[T]{},
		direction: direction,
	}
}

// simpleSortedList Support sorted list for simple types like int, float and string.
type simpleSortedList[T simpleTypes] struct {
	elements  List[T]
	direction SortDirection
}

// Add element to sorted list. It is duplicate tolerant.
//  - Return index of element in the list.
func (list *simpleSortedList[T]) Add(element T) int {
	left := list.findPosition(element)
	elems := append([]T{element}, (list.elements)[left:]...)
	elems = append(list.elements[:left], elems...)
	(*list).elements = elems
	return left
}

// Remove element from sorted list. If there is multiple values, removes only first one.
//  If cannot find element return false.
func (list *simpleSortedList[T]) Remove(element T) bool {
	if len(list.elements) == 0 {
		return false
	}
	left := list.findPosition(element)
	if left < len(list.elements) && (list.elements)[left] == element {
		(*list).elements = append(list.elements[:left], (list.elements)[left+1:]...)
		return true
	}
	return false
}

// RemoveAndReturnIndex element from sorted list. If there is multiple values, removes only first one.
//  If cannot find element return -1, if found return its index.
func (list *simpleSortedList[T]) RemoveAndReturnIndex(element T) int {
	if len(list.elements) == 0 {
		return -1
	}
	left := list.findPosition(element)
	if left < len(list.elements) && (list.elements)[left] == element {
		(*list).elements = append(list.elements[:left], (list.elements)[left+1:]...)
		return left
	}
	return -1
}

// RemoveAt If cannot find element return false.
func (list *simpleSortedList[T]) RemoveAt(position int) bool {
	newList, removed := RemoveAt(list.elements, position)
	if removed {
		(*list).elements = newList
		return true
	}
	return false
}

// ToList Return simple List from sorted list elements. 0(1) time complexity.
func (list *simpleSortedList[T]) ToList() List[T] {
	return list.elements
}

// Last Returns element from the end of the list
//  - nil if list is empty.
func (list *simpleSortedList[T]) Last() *T {
	if len(list.elements) == 0 {
		return nil
	}
	last := list.elements[len(list.elements)-1]
	return &last
}

// First Returns first element of the kust
//  - nil if list is empty.
func (list *simpleSortedList[T]) First() *T {
	if len(list.elements) == 0 {
		return nil
	}
	first := (list.elements)[0]
	return &first
}

// Count return number of elements.
func (list *simpleSortedList[T]) Count() int {
	return len(list.elements)
}

// IsEmpty returns true if list has no elements.
func (list *simpleSortedList[T]) IsEmpty() bool {
	return len(list.elements) == 0
}

func (list *simpleSortedList[T]) findPosition(element T) int {
	if list.direction == ASC {
		return FindPosition(list.elements, element)
	} else if list.direction == DESC {
		return FindPositionDesc(list.elements, element)
	}
	panic(fmt.Sprintf("Wrong direction to find position in simpleSortedList: %v", list.direction))
}
