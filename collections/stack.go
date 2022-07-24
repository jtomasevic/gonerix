package collections

// Stack Represents a simple last-in-first-out (LIFO) collection of objects.
type Stack[T any] []T

// Push Inserts an object at the top of the Stack.
func (stack *Stack[T]) Push(element T) {
	*stack = append([]T{element}, *stack...)
}

// Pop Removes and returns the object at the top of the Stack.
// - if stack is empty returns nil
func (stack *Stack[T]) Pop() *T {
	if stack.IsEmpty() {
		return nil
	}
	forPop := (*stack)[0]
	*stack = (*stack)[1:len(*stack)]
	return &forPop
}

// Peek Returns the object at the top of the Stack without removing it.
// If stack is empty return nil
func (stack *Stack[T]) Peek() *T {
	if stack.IsEmpty() {
		return nil
	}
	return &(*stack)[0]
}

// IsEmpty returns true if stack is empty.
func (stack *Stack[T]) IsEmpty() bool {
	return len(*stack) == 0
}
