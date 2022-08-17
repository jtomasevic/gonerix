package collections

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func getListOfPriorityQueueValues[T comparable, P simpleTypes](queue priorityQueue[T, P]) []T {
	var result []T = make([]T, 0)
	current := queue.elements.head
	for current != nil {
		result = append(result, current.Value.Value)
		current = current.Next
	}
	return result
}

func TestPriorityQueue_Enqueue(t *testing.T) {
	t.Run("Random 1 - ASC", func(t *testing.T) {
		queue := PriorityQueue[string, int](ASC, 5)
		queue.Enqueue("1", 1)
		queue.Enqueue("2", 2)
		queue.Enqueue("4", 4)
		queue.Enqueue("5", 5)
		queue.Enqueue("3", 3)
		queue.Enqueue("8", 8)
		queue.Enqueue("6", 6)
		queue.Enqueue("7", 7)

		actual := getListOfPriorityQueueValues(queue)
		require.Equal(t, []string{
			"1", "2", "3", "4", "5", "6", "7", "8",
		}, actual)
	})
	t.Run("Random 1 - DESC", func(t *testing.T) {
		queue := PriorityQueue[string, int](DESC, 5)
		queue.Enqueue("1", 1)
		queue.Enqueue("2", 2)
		queue.Enqueue("4", 4)
		queue.Enqueue("5", 5)
		queue.Enqueue("3", 3)
		queue.Enqueue("8", 8)
		queue.Enqueue("6", 6)
		queue.Enqueue("7", 7)

		actual := getListOfPriorityQueueValues(queue)
		require.Equal(t, []string{
			"8", "7", "6", "5", "4", "3", "2", "1",
		}, actual)
	})

	t.Run("Random 2 - ASC", func(t *testing.T) {
		queue := PriorityQueue[string, int](ASC, 5)
		queue.Enqueue("4", 4)
		queue.Enqueue("3", 3)
		queue.Enqueue("2", 2)
		queue.Enqueue("1", 1)
		actual := getListOfPriorityQueueValues(queue)
		require.Equal(t, []string{
			"1", "2", "3", "4",
		}, actual)
	})
	t.Run("Random 2 - DESC", func(t *testing.T) {
		queue := PriorityQueue[string, int](DESC, 5)
		queue.Enqueue("4", 4)
		queue.Enqueue("3", 3)
		queue.Enqueue("2", 2)
		queue.Enqueue("1", 1)
		actual := getListOfPriorityQueueValues(queue)
		require.Equal(t, []string{
			"4", "3", "2", "1",
		}, actual)
	})

	t.Run("Random 3 - ASC", func(t *testing.T) {
		queue := PriorityQueue[string, int](ASC, 5)
		queue.Enqueue("1", 1)
		queue.Enqueue("2", 2)
		queue.Enqueue("3", 3)
		queue.Enqueue("4", 4)
		actual := getListOfPriorityQueueValues(queue)
		require.Equal(t, []string{
			"1", "2", "3", "4",
		}, actual)
	})
	t.Run("Random 3 - DESC", func(t *testing.T) {
		queue := PriorityQueue[string, int](DESC, 5)
		queue.Enqueue("1", 1)
		queue.Enqueue("2", 2)
		queue.Enqueue("3", 3)
		queue.Enqueue("4", 4)
		actual := getListOfPriorityQueueValues(queue)
		require.Equal(t, []string{
			"4", "3", "2", "1",
		}, actual)
	})

	t.Run("Random 4 - ASC", func(t *testing.T) {
		queue := PriorityQueue[string, int](ASC, 5)
		queue.Enqueue("5", 5)
		queue.Enqueue("7", 7)
		queue.Enqueue("2", 2)
		queue.Enqueue("3", 3)
		queue.Enqueue("8", 8)
		queue.Enqueue("6", 6)
		queue.Enqueue("4", 4)
		queue.Enqueue("1", 1)

		actual := getListOfPriorityQueueValues(queue)
		require.Equal(t, []string{
			"1", "2", "3", "4", "5", "6", "7", "8",
		}, actual)
	})
	t.Run("Random 4 - DESC", func(t *testing.T) {
		queue := PriorityQueue[string, int](DESC, 5)
		queue.Enqueue("5", 5)
		queue.Enqueue("7", 7)
		queue.Enqueue("2", 2)
		queue.Enqueue("3", 3)
		queue.Enqueue("8", 8)
		queue.Enqueue("6", 6)
		queue.Enqueue("4", 4)
		queue.Enqueue("1", 1)

		actual := getListOfPriorityQueueValues(queue)
		require.Equal(t, []string{
			"8", "7", "6", "5", "4", "3", "2", "1",
		}, actual)
	})

	t.Run("Random 5 - ASC", func(t *testing.T) {
		queue := PriorityQueue[string, int](ASC, 5)
		queue.Enqueue("1", 1)
		queue.Enqueue("2", 2)
		queue.Enqueue("3", 3)
		queue.Enqueue("6", 6)
		queue.Enqueue("5", 5)
		queue.Enqueue("4", 4)
		queue.Enqueue("7", 7)
		queue.Enqueue("8", 8)

		actual := getListOfPriorityQueueValues(queue)
		require.Equal(t, []string{
			"1", "2", "3", "4", "5", "6", "7", "8",
		}, actual)
	})
	t.Run("Random 5 - DESC", func(t *testing.T) {
		queue := PriorityQueue[string, int](DESC, 5)
		queue.Enqueue("1", 1)
		queue.Enqueue("2", 2)
		queue.Enqueue("3", 3)
		queue.Enqueue("6", 6)
		queue.Enqueue("5", 5)
		queue.Enqueue("4", 4)
		queue.Enqueue("7", 7)
		queue.Enqueue("8", 8)

		actual := getListOfPriorityQueueValues(queue)
		require.Equal(t, []string{
			"8", "7", "6", "5", "4", "3", "2", "1",
		}, actual)
	})

	t.Run("Random with repeated values - ASC", func(t *testing.T) {
		queue := PriorityQueue[string, int](ASC, 5)
		queue.Enqueue("5", 5)
		queue.Enqueue("7", 7)
		queue.Enqueue("2", 2)
		queue.Enqueue("3", 3)
		queue.Enqueue("8", 8)
		queue.Enqueue("6", 6)
		queue.Enqueue("4", 4)
		queue.Enqueue("1", 1)
		queue.Enqueue("5.1", 5)
		queue.Enqueue("7.1", 7)
		queue.Enqueue("2.1", 2)
		queue.Enqueue("3.1", 3)
		queue.Enqueue("1.1", 1)
		queue.Enqueue("8.1", 8)
		queue.Enqueue("6.1", 6)
		queue.Enqueue("4.1", 4)
		queue.Enqueue("1.2", 1)
		queue.Enqueue("6.2", 6)
		queue.Enqueue("8.2", 8)
		actual := getListOfPriorityQueueValues(queue)
		require.Equal(t, []string{
			"1", "1.1", "1.2",
			"2", "2.1",
			"3", "3.1",
			"4", "4.1",
			"5", "5.1",
			"6", "6.1", "6.2",
			"7", "7.1",
			"8", "8.1", "8.2",
		}, actual)
	})

	t.Run("Random with repeated values - DESC", func(t *testing.T) {
		queue := PriorityQueue[string, int](DESC, 5)
		queue.Enqueue("5", 5)
		queue.Enqueue("7", 7)
		queue.Enqueue("2", 2)
		queue.Enqueue("3", 3)
		queue.Enqueue("8", 8)
		queue.Enqueue("6", 6)
		queue.Enqueue("4", 4)
		queue.Enqueue("1", 1)
		queue.Enqueue("5.1", 5)
		queue.Enqueue("7.1", 7)
		queue.Enqueue("2.1", 2)
		queue.Enqueue("3.1", 3)
		queue.Enqueue("1.1", 1)
		queue.Enqueue("8.1", 8)
		queue.Enqueue("6.1", 6)
		queue.Enqueue("4.1", 4)
		queue.Enqueue("1.2", 1)
		queue.Enqueue("6.2", 6)
		queue.Enqueue("8.2", 8)
		actual := getListOfPriorityQueueValues(queue)
		require.Equal(t, []string{
			"8", "8.1", "8.2",
			"7", "7.1",
			"6", "6.1", "6.2",
			"5", "5.1",
			"4", "4.1",
			"3", "3.1",
			"2", "2.1",
			"1", "1.1", "1.2",
		}, actual)
	})

}

func TestPriorityQueue_Dequeue(t *testing.T) {
	t.Run("ASC", func(t *testing.T) {
		queue := PriorityQueue[string, int](ASC, 5)
		queue.Enqueue("1", 1)
		queue.Enqueue("2", 2)
		queue.Enqueue("3", 3)
		queue.Enqueue("4", 4)

		dequeued := queue.Dequeue()
		peek := queue.Peek()
		require.Equal(t, "1", *dequeued)
		require.Equal(t, "2", *peek)
		actual := getListOfPriorityQueueValues(queue)
		require.Equal(t, []string{
			"2", "3", "4",
		}, actual)

		dequeued = queue.Dequeue()
		peek = queue.Peek()
		require.Equal(t, "2", *dequeued)
		require.Equal(t, "3", *peek)
		actual = getListOfPriorityQueueValues(queue)
		require.Equal(t, []string{
			"3", "4",
		}, actual)

		dequeued = queue.Dequeue()
		peek = queue.Peek()
		require.Equal(t, "3", *dequeued)
		require.Equal(t, "4", *peek)
		actual = getListOfPriorityQueueValues(queue)
		require.Equal(t, []string{
			"4",
		}, actual)

		dequeued = queue.Dequeue()
		peek = queue.Peek()
		require.Equal(t, "4", *dequeued)
		require.Nil(t, peek)
		actual = getListOfPriorityQueueValues(queue)
		require.Equal(t, []string{}, actual)

		dequeued = queue.Dequeue()
		peek = queue.Peek()
		require.Nil(t, dequeued)
		require.Nil(t, peek)
	})
	t.Run("DESC", func(t *testing.T) {
		queue := PriorityQueue[string, int](DESC, 5)
		queue.Enqueue("1", 1)
		queue.Enqueue("2", 2)
		queue.Enqueue("3", 3)
		queue.Enqueue("4", 4)

		dequeued := queue.Dequeue()
		require.Equal(t, "4", *dequeued)
		actual := getListOfPriorityQueueValues(queue)
		require.Equal(t, []string{
			"3", "2", "1",
		}, actual)
	})
}

//func TestPriorityQueue_Enqueue(t *testing.T) {
//	t.Run("Add", func(t *testing.T) {
//		orders := getOrders()
//		// it means lower number means higher priority. So priority with value 1 is higher than priority with value 2.
//		queue := PriorityQueue[Order, int](ASC, 5)
//		for _, order := range orders {
//			queue.Enqueue(order, order.Priority)
//		}
//		expected := getListOfPriorityQueueValues(queue)
//		for _, exp := range expected {
//			fmt.Printf("!!! %v\n", exp)
//		}
//		fmt.Printf("!!! elements.values %v", queue.elements.Values())
//		//require.Equal(t, List[int]{
//		//	expected,
//		//}, queue.elements.Values())
//
//		require.Equal(t, List[int]{
//			3, 3, 4, 5, 7, 10, 10,
//		}, queue.priorities.ToList())
//		require.Equal(t, Stack[Order]{
//			{
//				Customer: "Tesla",
//				OrderNo:  "TSL-s2",
//				Priority: 3,
//			},
//			{
//				Customer: "Google",
//				OrderNo:  "G-s2",
//				Priority: 3,
//			},
//		}, queue.values[3])
//		require.Equal(t, Stack[Order]{
//			{
//				Customer: "IBM",
//				OrderNo:  "IBM-s1",
//				Priority: 4,
//			},
//		}, queue.values[4])
//		require.Equal(t, Stack[Order]{
//			{
//				Customer: "Google",
//				OrderNo:  "G-s1",
//				Priority: 5,
//			},
//		}, queue.values[5])
//		require.Equal(t, Stack[Order]{
//			{
//				Customer: "Microsoft",
//				OrderNo:  "TSL-s2",
//				Priority: 7,
//			},
//		}, queue.values[7])
//		require.Equal(t, Stack[Order]{
//			{
//				Customer: "Apple",
//				OrderNo:  "Apple-s1",
//				Priority: 10,
//			},
//			{
//				Customer: "Tesla",
//				OrderNo:  "TSL-s1",
//				Priority: 10,
//			},
//		}, queue.values[10])
//	})
//
//	t.Run("Remove", func(t *testing.T) {
//		orders := getOrders()
//		queue := PriorityQueue[Order, int](ASC, 10)
//		for _, order := range orders {
//			queue.Enqueue(order, order.Priority)
//		}
//		require.Equal(t, Order{
//			Customer: "Tesla",
//			OrderNo:  "TSL-s2",
//			Priority: 3,
//		}, *queue.Peek())
//		queue.Dequeue()
//		require.Equal(t, Order{
//			Customer: "Google",
//			OrderNo:  "G-s2",
//			Priority: 3,
//		}, *queue.Peek())
//
//		require.Equal(t, List[int]{
//			3, 4, 5, 7, 10, 10,
//		}, queue.priorities.ToList())
//		require.Equal(t, Stack[Order]{
//			{
//				Customer: "Google",
//				OrderNo:  "G-s2",
//				Priority: 3,
//			},
//		}, queue.values[3])
//		//
//		queue.Dequeue()
//		require.Equal(t, List[int]{
//			4, 5, 7, 10, 10,
//		}, queue.priorities.ToList())
//		require.Equal(t, Stack[Order]{
//			{
//				Customer: "IBM",
//				OrderNo:  "IBM-s1",
//				Priority: 4,
//			},
//		}, queue.values[4])
//		require.Equal(t, Order{
//			Customer: "IBM",
//			OrderNo:  "IBM-s1",
//			Priority: 4,
//		}, *queue.Peek())
//		queue.Dequeue()
//		queue.Dequeue()
//		queue.Dequeue()
//		queue.Dequeue()
//		queue.Dequeue()
//		result := queue.Dequeue()
//		require.Nil(t, result)
//		obj := queue.Peek()
//		require.Nil(t, obj)
//
//	})
//}

type Order struct {
	Customer string
	OrderNo  string
	Priority int
}

func getOrders() []Order {
	return []Order{
		{
			Customer: "Google",
			OrderNo:  "G-s1",
			Priority: 5,
		},
		{
			Customer: "IBM",
			OrderNo:  "IBM-s1",
			Priority: 4,
		},
		{
			Customer: "Google",
			OrderNo:  "G-s2",
			Priority: 3,
		},
		{
			Customer: "Tesla",
			OrderNo:  "TSL-s1",
			Priority: 10,
		},
		{
			Customer: "Tesla",
			OrderNo:  "TSL-s2",
			Priority: 3,
		},
		{
			Customer: "Microsoft",
			OrderNo:  "TSL-s2",
			Priority: 7,
		},
		{
			Customer: "Apple",
			OrderNo:  "Apple-s1",
			Priority: 10,
		},
	}
}
