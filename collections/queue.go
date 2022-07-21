package collections

type Queue[T comparable] struct{
	elements []T
}

// Adds an object to the end of the Queue.
func (queue *Queue[T]) Enqueue (element T) {
	queue.elements = append(queue.elements, element)
}

// Removes and returns the object at the beginning of the Queue.
// If stack is empty return nil
func (queue *Queue[T]) Dequeue() *T {
	if queue.IsEmpty() {
		return nil 
	}
	removed:=queue.elements[0]
	queue.elements = queue.elements[1:len(queue.elements)]
	return &removed
}

// Peek Returns the object at the top of the Queue without removing it.
// If stack is empty return nil
func (queue *Queue[T]) Peek() *T {
	if queue.IsEmpty() {
		return nil 
	}
	return &queue.elements[0]
}

func (queue *Queue[T]) IsEmpty() bool{
	return len(queue.elements) == 0
}
