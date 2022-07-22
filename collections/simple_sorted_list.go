package collections

// ComperableFunction returns true if left paramater is higer than right paramater
type ComperableFunction[T comparable] func(left T, right T) bool

// Types supported by Sorted List
type simpleTypes interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

// type SortedList[T simpleTypes] struct {
// 	elements       []T
// 	// comparableFunc ComperableFunction[T]
// }

// SimpleSortedList Support sorted list for simple types like int, float and string.
type SimpleSortedList[T simpleTypes] []T

// Add element to sorted list. It is duplicate tolerante.
func (list *SimpleSortedList[T]) Add(element T) {

	newList := *list
	low := 0
	high := len(newList) - 1
	for low <= high {
		median := (low + high) / 2
		if newList[median] < element {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	elems := append([]T{element}, newList[low:]...)
	elems = append(newList[:low], elems...)
	*list = elems
}

// Remove element from sorted list. If there is multiple values, removes only first one. 
//  If cannot find element return false. 
func (list *SimpleSortedList[T]) Remove(element T) bool {
	if len(*list) == 0 {
		return false
	}
	newList := *list
	low := 0
	high := len(newList) - 1
	for low <= high {
		median := (low + high) / 2
		if newList[median] < element {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	if low < len(newList) && newList[low] == element {
		*list = append(newList[:low], newList[low+1:]...)
		return true
	}
	return false
}

// ToList Return simple List from sorted list elements. 
func (list *SimpleSortedList[T]) ToList() List[T] {
	res := List[T]{}
	for _, el := range *list {
		res.Add(el)
	}
	return res
}
