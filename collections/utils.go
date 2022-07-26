package collections

import "fmt"

func RemoveAt[T comparable](list []T, position int) ([]T, bool) {
	if position >= len(list) || position < 0 {
		return list, false
	}
	return append(list[:position], list[position+1:]...), true
}

// IsEqual returns true if left parameter is considered equal to parameter right considering function output.
//  NOTE: This doesn't mean that two parameters are really equal, but this is defined by function (IsEqual) which can
//  be by for example field named Id or some other complex condition.
type IsEqual[T comparable] func(left T, right T) bool

// AddOrReplace element in the list under some condition. Applies only to first element found under this condition.
//  - Try to find element that satisfied condition in isEqual function
//  - If found replace with new provided element.
func AddOrReplace[T comparable](list []T,
	element T,
	isEqual IsEqual[T]) []T {
	for i, el := range list {
		if isEqual(el, element) {
			list[i] = element
			return list
		}
	}
	l := append(list, element)
	fmt.Printf("**** generic result: %v\n", l)
	return l
}

// FindPosition Return index of element or:
// - If element is not in elements, then return index where element can be added considering order.
func FindPosition[T simpleTypes](elements []T, element T) int {
	left := 0
	right := len(elements) - 1
	for left <= right {
		median := (left + right) / 2
		if elements[median] < element {
			left = median + 1
		} else {
			right = median - 1
		}
	}
	return left
}
