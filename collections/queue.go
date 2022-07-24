package collections

type Queue[T any] []T

// Enqueue Adds an object to the end of the Queue.
func (queue *Queue[T]) Enqueue(element T) {
	*queue = append(*queue, element)
}

// Dequeue Removes and returns the object at the beginning of the Queue.
// If stack is empty return nil
func (queue *Queue[T]) Dequeue() *T {
	if queue.IsEmpty() {
		return nil
	}
	removed := (*queue)[0]
	*queue = (*queue)[1:len(*queue)]
	return &removed
}

// Peek Returns the object at the top of the Queue without removing it.
// If stack is empty return nil
func (queue *Queue[T]) Peek() *T {
	if queue.IsEmpty() {
		return nil
	}
	return &(*queue)[0]
}

// IsEmpty return bool is true, otherwise false
func (queue *Queue[T]) IsEmpty() bool {
	return len(*queue) == 0
}
