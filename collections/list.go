package collections

import "fmt"

type List[T comparable] []T

// Add an object to the end of the list
func (list *List[T]) Add(element T) {
	newList := append(*list, element)
	*list = newList
}

// Insert an element into the List[T] at the specified index.
func (list *List[T]) Insert(position int, element ...T) bool {
	if position > len(*list) || position < 0 {
		return false
	}
	newList := *list
	rightPart := append(element, newList[position:]...)
	newList = append(newList[:position], rightPart...)
	*list = newList
	return true
}

// Remove the first occurrence of a specific object from the list
func (list *List[T]) Remove(element T) bool {
	exist, index := list.Exist(element)
	if !exist {
		return false
	}
	newList := *list
	newList = append(newList[:index], newList[index+1:]...)
	*list = newList
	return true
}

// RemoveAt Removes the element at the specified index of the list
func (list *List[T]) RemoveAt(position int) bool {
	newList, removed := RemoveAnyAt(*list, position)
	if removed {
		*list = newList
		return true
	}
	return false
}

// Reverse the order of the elements in the entire
func (list *List[T]) Reverse() {
	newList := *list
	for i := len(newList)/2 - 1; i >= 0; i-- {
		opp := len(newList) - 1 - i
		newList[i], newList[opp] = newList[opp], newList[i]
	}
	*list = newList
}

// Exist Determines whether the specified object is equal to the current object.
func (list *List[T]) Exist(element T) (bool, int) {
	index := -1
	for i, el := range *list {
		if element == el {
			index = i
		}
	}
	if index == -1 {
		return false, -1
	}
	return true, index
}

// Find Return list of elements that satisfy function check provided by expression parameter.
func (list *List[T]) Find(expression func(element T) bool) List[T] {
	result := List[T]{}
	for _, el := range *list {
		if expression(el) {
			result.Add(el)
		}
	}
	return result
}

// AddOrReplace element in the list under some condition. Applies only to first element found under this condition
//  - Try to find element that satisfied condition in isEqual function
//  - If found replace with new provided element.
//  - If not found add it to list.
func (list *List[T]) AddOrReplace(element T, isEqual IsEqual[T]) {
	result := AddOrReplace(*list, element, isEqual)
	fmt.Printf("*** results: %v\n", result)
	*list = result
}
