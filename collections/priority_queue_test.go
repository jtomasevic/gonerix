package collections

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPriorityQueue_Enqueue(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		orders := getOrders()
		queue := PriorityQueue[Order, int]()
		for _, order := range orders {
			queue.Enqueue(order, order.Priority)
		}
		require.Equal(t, List[int]{
			3, 3, 4, 5, 7, 10, 10,
		}, queue.priorities.ToList())
		require.Equal(t, Stack[Order]{
			{
				Customer: "Tesla",
				OrderNo:  "TSL-s2",
				Priority: 3,
			},
			{
				Customer: "Google",
				OrderNo:  "G-s2",
				Priority: 3,
			},
		}, queue.values[3])
		require.Equal(t, Stack[Order]{
			{
				Customer: "IBM",
				OrderNo:  "IBM-s1",
				Priority: 4,
			},
		}, queue.values[4])
		require.Equal(t, Stack[Order]{
			{
				Customer: "Google",
				OrderNo:  "G-s1",
				Priority: 5,
			},
		}, queue.values[5])
		require.Equal(t, Stack[Order]{
			{
				Customer: "Microsoft",
				OrderNo:  "TSL-s2",
				Priority: 7,
			},
		}, queue.values[7])
		require.Equal(t, Stack[Order]{
			{
				Customer: "Apple",
				OrderNo:  "Apple-s1",
				Priority: 10,
			},
			{
				Customer: "Tesla",
				OrderNo:  "TSL-s1",
				Priority: 10,
			},
		}, queue.values[10])
	})

	t.Run("Remove", func(t *testing.T) {
		orders := getOrders()
		queue := PriorityQueue[Order, int]()
		for _, order := range orders {
			queue.Enqueue(order, order.Priority)
		}
		require.Equal(t, Order{
			Customer: "Tesla",
			OrderNo:  "TSL-s2",
			Priority: 3,
		}, *queue.Peek())
		queue.Dequeue()
		require.Equal(t, Order{
			Customer: "Google",
			OrderNo:  "G-s2",
			Priority: 3,
		}, *queue.Peek())

		require.Equal(t, List[int]{
			3, 4, 5, 7, 10, 10,
		}, queue.priorities.ToList())
		require.Equal(t, Stack[Order]{
			{
				Customer: "Google",
				OrderNo:  "G-s2",
				Priority: 3,
			},
		}, queue.values[3])
		//
		queue.Dequeue()
		require.Equal(t, List[int]{
			4, 5, 7, 10, 10,
		}, queue.priorities.ToList())
		require.Equal(t, Stack[Order]{
			{
				Customer: "IBM",
				OrderNo:  "IBM-s1",
				Priority: 4,
			},
		}, queue.values[4])
		require.Equal(t, Order{
			Customer: "IBM",
			OrderNo:  "IBM-s1",
			Priority: 4,
		}, *queue.Peek())
		queue.Dequeue()
		queue.Dequeue()
		queue.Dequeue()
		queue.Dequeue()
		queue.Dequeue()
		result := queue.Dequeue()
		require.Nil(t, result)
		obj := queue.Peek()
		require.Nil(t, obj)

	})
}

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
