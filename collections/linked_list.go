package collections

// Node Represent node of linked list.
type Node[T comparable] struct {
	Value T
	Prev  *Node[T]
	Next  *Node[T]
}

type linkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
}

// LinkedList returns instance of linkedList[T] which a general-purpose linked list (double linked)
func LinkedList[T comparable]() linkedList[T] {
	return linkedList[T]{}
}

// Add to the end of list
//  - runtime 0(1)
//  - return pointer to Node which is part of list.
func (list *linkedList[T]) Add(value T) *Node[T] {
	newNode := &Node[T]{
		Value: value,
	}
	if list.head == nil {
		list.head = newNode
		list.tail = list.head
	} else {
		temp := list.tail
		newNode.Prev = temp
		temp.Next = newNode
		list.tail = newNode
	}
	return newNode
}

// Remove node from list where element has provided Value.
//  - runtime: this method performs a linear search; therefore, this method is an O(n) operation
//  - return pointer to the removed node
func (list *linkedList[T]) Remove(value T) (*Node[T], bool) {
	if list.head == nil {
		return nil, false
	}
	temp := list.head
	for temp != nil {
		if temp.Value == value {
			ok := list.RemoveNode(temp)
			if ok {
				return temp, true
			} else {
				return nil, false
			}
		} else {
			temp = temp.Next
		}
	}
	return nil, false
}

// RemoveNode works only if node is really node from this list.
//  - runtime: O(1)
func (list *linkedList[T]) RemoveNode(node *Node[T]) bool {
	if node == nil {
		return false
	}
	if node == list.head && node == list.tail {
		list.head = nil
		list.tail = nil
	} else if node == list.head {
		list.head = node.Next
		list.head.Prev = nil
	} else if node == list.tail {
		list.tail = node.Prev
		list.tail.Next = nil
	} else {
		prev := node.Prev
		next := node.Next
		if prev != nil {
			prev.Next = next
		}
		if next != nil {
			next.Prev = prev
		}
	}
	return true
}

func (list *linkedList[T]) Values() []T {
	printList := List[T]{}
	temp := list.head
	for temp != nil {
		printList.Add(temp.Value)
		temp = temp.Next
	}
	return printList
}
func (list *linkedList[T]) Head() *T {
	if list.head != nil {
		return &list.head.Value
	}
	return nil
}
func (list *linkedList[T]) Tail() *T {
	if list.tail != nil {
		return &list.tail.Value
	}
	return nil
}

func (list *linkedList[T]) HeadNode() *Node[T] {
	return list.head

}
func (list *linkedList[T]) TailNode() *Node[T] {
	return list.tail
}
