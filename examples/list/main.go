package main

import (
	"fmt"
	gx "github.com/jtomasevic/gonerix/collections"
)

func main() {
	list := gx.List[int]{}
	// Add
	list.Add(1)
	list.Add(2)
	list.Add(4)
	list.Add(8)
	fmt.Println(list)
	// [1 2 4 8]

	// Insert one value
	list.Insert(2, 3)
	fmt.Println(list)
	// [1 2 3 4 8]

	// Insert multiple value
	list.Insert(4, 5, 6, 7)
	fmt.Println(list)
	// [1 2 3 4 5 6 7 8]

	// Reverse
	list.Reverse()
	fmt.Println(list)
	// [8 7 6 5 4 3 2 1]

	list.Remove(7)
	fmt.Println(list)
	// [8 6 5 4 3 2 1]

	list.RemoveAt(2)
	fmt.Println(list)
	// [8 6 4 3 2 1]

	list.RemoveAt(len(list) - 1)
	fmt.Println(list)
	// [8 6 4 3 2]

	res, index := list.Exist(4)
	fmt.Printf("element: %v, index: %v\n", res, index)
	// element: true, index: 2

	res, index = list.Exist(14)
	fmt.Printf("element: %v, index: %v\n", res, index)
	// element: false, index: -1

	// Create some structs to be part of List
	type CupSize string
	const (
		Small      CupSize = "small"
		Medium     CupSize = "medium"
		Large      CupSize = "large"
		ExtraLarge CupSize = "extraLarge"
	)
	type BlackCoffeeOffer struct {
		Price int
		Size  CupSize
	}
	// Initialise list
	products := gx.List[BlackCoffeeOffer]{
		{
			Price: 100,
			Size:  Small,
		},
		{
			Price: 150,
			Size:  Medium,
		},
		{
			Price: 200,
			Size:  Large,
		},
	}
	// This operation applies only to first element found under certain condition
	//  - Try to find element that satisfied condition in isEqual function (second parameter)
	//  - If found replace with new provided element.
	//  - If not found add it to list.
	products.AddOrReplace(BlackCoffeeOffer{
		Price: 250,
		Size:  ExtraLarge,
	}, func(left BlackCoffeeOffer, right BlackCoffeeOffer) bool {
		return left.Size == right.Size
	})
	fmt.Println(products)
	// [{100 small} {150 medium} {200 large} {250 extraLarge}]

	// define isEqual function as variable
	var isEqual gx.IsEqual[BlackCoffeeOffer] = func(left BlackCoffeeOffer, right BlackCoffeeOffer) bool {
		return left.Size == right.Size
	}
	products.AddOrReplace(BlackCoffeeOffer{
		Price: 95,
		Size:  Small,
	}, isEqual)
	fmt.Println(products)
	// [{95 small} {150 medium} {200 large} {250 extraLarge}]

	cheaperThan200 := products.Find(func(left BlackCoffeeOffer) bool {
		return left.Price < 200
	})
	fmt.Println(cheaperThan200)
	// [{95 small} {150 medium}]
}
