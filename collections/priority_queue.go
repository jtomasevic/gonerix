package collections

type priorityNode[T comparable, P simpleTypes] struct {
	Priority P
	Value    T
}

// PriorityQueue  is an extension of the queue with the following properties:
//  - Every item has a priority associated with it.
//  - An element with high priority is dequeued before an element with low priority.
//  - If two elements have the same priority, they are served according to their order in the queue.
//  - Capacity: queue has capacity, so some elements could be lost if capacity is full.
func PriorityQueue[T comparable, P simpleTypes](priorityOrderType SortDirection, capacity int) priorityQueue[T, P] {
	var ascCondition = func(priority P, node *Node[priorityNode[T, P]]) bool {
		return priority >= node.Value.Priority
	}
	var descCondition = func(priority P, node *Node[priorityNode[T, P]]) bool {
		return priority <= node.Value.Priority
	}
	var condition func(priority P, node *Node[priorityNode[T, P]]) bool
	if priorityOrderType == ASC {
		condition = ascCondition
	} else {
		condition = descCondition
	}
	return priorityQueue[T, P]{
		capacity:        capacity,
		elements:        LinkedList[priorityNode[T, P]](),
		compareFunction: condition,
	}
}

// priorityQueue Represents a collection of items that have a value and a priority.
//  - T Specifies the type of elements in the queue.
//  - P Specifies the type of priority associated with enqueued elements.
type priorityQueue[T comparable, P simpleTypes] struct {
	capacity        int
	elements        linkedList[priorityNode[T, P]]
	compareFunction func(priority P, node *Node[priorityNode[T, P]]) bool
}

// Enqueue Adds the specified element with associated priority.
//  - O(n) = n
func (queue *priorityQueue[T, P]) Enqueue(value T, priority P) {
	current := queue.elements.HeadNode()
	// if list is empty
	if current == nil {
		queue.elements.Add(priorityNode[T, P]{
			Priority: priority,
			Value:    value,
		})
		return
	}
	// keep track on prev because if current is nil at the end, then we should refer to prev.
	prev := current
	// for current != nil && priority >= current.Value.Priority {
	for current != nil && queue.compareFunction(priority, current) {
		prev = current
		current = current.Next
	}
	if current == nil {
		current = prev
	}
	if queue.compareFunction(priority, current) {
		queue.elements.AddAfter(current, priorityNode[T, P]{
			Priority: priority,
			Value:    value,
		})
	} else {
		queue.elements.AddBefore(current, priorityNode[T, P]{
			Priority: priority,
			Value:    value,
		})
	}
}

// Dequeue Removes and returns the minimal element from the PriorityQueue - that is, the element with the lowest priority value.
//  Note: If there is multiple elements with same priority last added will be removed first (Stack approach).
//  If value with priority not exit return value will be false.
//  - O(n) = 1
func (queue *priorityQueue[T, P]) Dequeue() *T {
	if queue.elements.head == nil {
		return nil
	}
	forDequeue := queue.elements.head
	if queue.elements.head.Next != nil {
		queue.elements.head.Next.Prev = nil
		queue.elements.head = queue.elements.head.Next
	} else {
		queue.elements.head = nil
		queue.elements.tail = nil
	}
	return &forDequeue.Value.Value
}

// Peek Returns the minimal element from the PriorityQueue without removing it.
//  if queue is empty it returns nil.
//  - O(n) = 1
func (queue *priorityQueue[T, P]) Peek() *T {
	if queue.elements.head == nil {
		return nil
	}
	return &queue.elements.head.Value.Value
}
