package main

import (
	"fmt"
	gx "github.com/jtomasevic/gonerix/collections"
)

func main() {
	// Add
	list := gx.SimpleSortedList[int](gx.ASC)
	list.Add(30)
	list.Add(50)
	list.Add(20)
	list.Add(10)
	list.Add(40)
	fmt.Println(list)
	// {[10 20 30 40 50] asc}
	fmt.Println(list.ToList())
	// NOTE: ToList is O(1), no transformation just pass internal reference.
	// [10 20 30 40 50]

	// Remove
	list.Remove(10)
	fmt.Println(list.ToList())
	// [20 30 40 50]

	// First
	first := list.First()
	fmt.Println(*first)
	// 20

	// Last
	last := list.Last()
	fmt.Println(*last)
	// 50

	// RemoveAt
	list.RemoveAt(2)
	fmt.Println(list.ToList())
	// [20 30 50]

	// IsEmpty
	list.IsEmpty()
	fmt.Println(list.IsEmpty())
	// false

	// Count
	count := list.Count()
	fmt.Println(count)
	// 3
}
