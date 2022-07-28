package collections

import "fmt"

func RemoveAt[T comparable](list []T, position int) ([]T, bool) {
	if position >= len(list) || position < 0 {
		return list, false
	}
	return append(list[:position], list[position+1:]...), true
}

func RemoveAnyAt[T any](list []T, position int) ([]T, bool) {
	if position >= len(list) || position < 0 {
		return list, false
	}
	return append(list[:position], list[position+1:]...), true
}

// IsEqual returns true if left parameter is considered equal to parameter right considering function output.
//  NOTE: This doesn't mean that two parameters are really equal, but this is defined by function (IsEqual) which can
//  be by for example field named Id or some other complex condition.
type IsEqual[T any] func(left T, right T) bool

// AddOrReplace element in the list under some condition. Applies only to first element found under this condition.
//  - Try to find element that satisfied condition in isEqual function
//  - If found replace with new provided element.
func AddOrReplace[T any](list []T,
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

// FindPosition Return index of element in asc sorted list or:
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

// FindPositionDesc Return index of element in desc sorted list or:
// - If element is not in elements, then return index where element can be added considering order.
func FindPositionDesc[T simpleTypes](elements []T, element T) int {
	// code is duplicate of FindPosition with one simple change.
	// we didn't want to put additional if to check if it's asc or des sorted list, so this is a way to keep optimal performance.
	left := 0
	right := len(elements) - 1
	for left <= right {
		median := (left + right) / 2
		if elements[median] > element {
			left = median + 1
		} else {
			right = median - 1
		}
	}
	return left
}
