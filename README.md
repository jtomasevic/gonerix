# gonerix
Generic collections for golang

### Documentation still in development 

For still not documented collections check out code and test/example, to understand more. 

### Expected release + more explainational examples: 
1st September 2022.
## Usage

We'll assume that in each file you'll have import of collection package with qx prefix, to make examples more readable:
```
import (
	"fmt"
	gx "github.com/jtomasevic/gonerix/collections"
)
```
# API
- [List[T]](https://github.com/jtomasevic/gonerix#listt)
- [SimpleSortedList[T]](https://github.com/jtomasevic/gonerix/edit/main/README.md#simplesortedlist)

# List[T]

### Add
```
list := gx.List[int]{}

list.Add(1)
list.Add(2)
list.Add(4)
list.Add(8)
fmt.Println(list)
// [1 2 4 8]
```
### Insert
First parameter is index where to insert an element of more of them.
Second parameter is one or vmore alues to be added to list.
#### Insert one value
```
list.Insert(2, 3)
fmt.Println(list)
// [1 2 3 4 8]
```
#### Insert multiple value
```
list.Insert(4, 5, 6, 7)
fmt.Println(list)
// [1 2 3 4 5 6 7 8]
````
### Reverse
```
list.Reverse()
fmt.Println(list)
// [8 7 6 5 4 3 2 1]
```
### Remove element from list
```
list.Remove(7)
fmt.Println(list)
// [8 6 5 4 3 2 1]
```
### Remove element with particular index
```
list.RemoveAt(2)
fmt.Println(list)
// [8 6 4 3 2 1]

list.RemoveAt(len(list) - 1)
fmt.Println(list)
// [8 6 4 3 2]
```
### Exist
```
res, index := list.Exist(4)
fmt.Printf("element: %v, index: %v\n", res, index)
// element: true, index: 2

res, index = list.Exist(14)
fmt.Printf("element: %v, index: %v\n", res, index)
// element: false, index: -1
```

### AddOrReplace
This operation applies only to first element found under certain condition.
- Try to find element that satisfied condition in isEqual function (second parameter)
- If found replace with new provided element (first parameter).
- If not found, add it to list.
```
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
```
### Find
Return list of elements that satisfy function check, provided by expression parameter
```
cheaperThan200 := products.Find(func(left BlackCoffeeOffer) bool {
	return left.Price < 200
})
fmt.Println(cheaperThan200)
// [{95 small} {150 medium}]
```
# SimpleSortedList

### Add
```
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
```

### Remove
```
list.Remove(10)
fmt.Println(list.ToList())
// [20 30 40 50]
```

### First
```
first := list.First()
fmt.Println(*first)
// 20
```

### Last
```
last := list.Last()
fmt.Println(*last)
// 50
```

### RemoveAt
```
list.RemoveAt(2)
// before [20 30 40 50]
fmt.Println(list.ToList())
// [20 30 50]
```

### IsEmpty
```
res:= list.IsEmpty()
fmt.Println(list.res)
// false
```

### Count
```
count := list.Count()
// currently [20 30 50]
fmt.Println(count)
// 3
```


