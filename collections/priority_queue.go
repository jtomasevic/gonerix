package collections

import "fmt"

func PriorityQueue[T any, P simpleTypes]() priorityQueue[T, P] {
	return priorityQueue[T, P]{
		priorities: SimpleSortedList[P]{},
		values:     Dictionary[P, Stack[T]]{},
	}
}

// priorityQueue Represents a collection of items that have a value and a priority. On dequeue, the item with the lowest priority value is removed.
//  - T Specifies the type of elements in the queue.
//  - P Specifies the type of priority associated with enqueued elements.
type priorityQueue[T any, P simpleTypes] struct {
	priorities SimpleSortedList[P]
	values     Dictionary[P, Stack[T]]
}

// Enqueue Adds the specified element with associated priority
func (queue *priorityQueue[T, P]) Enqueue(value T, priority P) {
	values := queue.values
	queue.priorities.Add(priority)
	// check if this is first item with provided priority
	if !values.Exist(priority) {
		values[priority] = Stack[T]{}
	}
	priorityStack := values[priority]
	priorityStack.Push(value)
	values[priority] = priorityStack
	queue.values = values
}

// Dequeue Removes and returns the minimal element from the PriorityQueue - that is, the element with the lowest priority value.
//  Note: If there is multiple elements with same priority last added will be removed first (Stack approach).
//  If value with priority not exit return value will be false.
func (queue *priorityQueue[T, P]) Dequeue() bool {
	if len(queue.priorities) == 0 {
		return false
	}
	priorities := queue.priorities
	values := queue.values
	minPriority := queue.priorities.First()
	priorityElements := values[*minPriority]
	// TODO JT: fix this code, looks ugly,=,
	forRemoving := priorityElements.Pop()
	values[*minPriority] = priorityElements
	// this should never happen, but for every case.
	if forRemoving == nil {
		return false
	}

	fmt.Printf("*** to rempve %v , element: %v \n", *minPriority, forRemoving)
	priorities.Remove(*minPriority)
	// if this was last element
	if priorityElements.IsEmpty() {
		values.Remove(*minPriority)
	}
	queue.priorities = priorities
	queue.values = values
	fmt.Printf("**** prios: %v\n", queue.priorities)
	return true
}

// Peek Returns the minimal element from the PriorityQueue without removing it.
//  if queue is empty it returns nil.
func (queue *priorityQueue[T, P]) Peek() *T {
	if len(queue.priorities) == 0 {
		return nil
	}
	elementStack := queue.values[*queue.priorities.First()]
	return elementStack.Peek()
}
