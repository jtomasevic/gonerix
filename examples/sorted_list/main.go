package main

import (
	"fmt"
	gx "github.com/jtomasevic/gonerix/collections"
)

func main() {
	// Add
	list := gx.SortedList[int, Order](gx.ASC)

	list.Add(3, Order{
		Customer: "C",
		OrderNo:  52,
	})
	list.Add(1, Order{
		Customer: "A",
		OrderNo:  50,
	})
	list.Add(4, Order{
		Customer: "D",
		OrderNo:  53,
	})
	list.Add(2, Order{
		Customer: "B",
		OrderNo:  51,
	})

	fmt.Println(list.ToList())
	// [{A 50} {B 51} {C 52} {D 5
}

type Order struct {
	Customer string
	OrderNo  int
}
