package collections

// Stack Represents a simple last-in-first-out (LIFO) collection of objects.
type Stack[T any] struct {
	elements []T
}

// Push Inserts an object at the top of the Stack.
func (stack *Stack[T]) Push(element T) {
	stack.elements = append([]T{element}, stack.elements...)
}

// Pop Removes and returns the object at the top of the Stack.
// - if stack is empty returns nil
func (stack *Stack[T]) Pop() *T {
	if stack.IsEmpty() {
		return nil
	}
	forPop := stack.elements[0]
	stack.elements = stack.elements[1:len(stack.elements)]
	return &forPop
}

// Peek Returns the object at the top of the Stack without removing it.
// If stack is empty return nil
func (stack *Stack[T]) Peek() *T {
	if stack.IsEmpty() {
		return nil
	}
	return &stack.elements[0]
}

// IsEmpty returns true if stack is empty.
func (stack *Stack[T]) IsEmpty() bool {
	return len(stack.elements) == 0
}
