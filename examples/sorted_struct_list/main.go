package main

import (
	"fmt"
	gx "github.com/jtomasevic/gonerix/collections"
)

func main() {
	// Add
	type Brand struct {
		Name string
	}
	brands := []Brand{
		{
			Name: "Diadora",
		},
		{
			Name: "Nike",
		},
		{
			Name: "Addidas",
		},
		{
			Name: "Puma",
		},
		{
			Name: "Nokka",
		},
	}
	sorted := gx.SortedStructList(func(left Brand, right Brand) bool {
		return left.Name > right.Name
	})
	for _, brand := range brands {
		sorted.Add(brand)
	}
	// this is O(1) operation.
	fmt.Println(sorted.Values())
	// [{Puma} {Nokka} {Nike} {Diadora} {Addidas}]

	// Remove
	sorted.Remove(Brand{
		Name: "Puma",
	})
	fmt.Println(sorted.Values())
	// [{Nokka} {Nike} {Diadora} {Addidas}]
	// non existing
	result := sorted.Remove(Brand{
		Name: "BOSS",
	})
	fmt.Println(result)
	// false

	// Remove At
	sorted.RemoveAt(0)
	fmt.Println(sorted.Values())
	// [{Nike} {Diadora} {Addidas}]

	sorted.Size()
	// return size of element.
}
