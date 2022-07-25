# gonerix
Generic collections for golang

# Still in preparation 

Check out code and test/example, to understand more. 

#### MORE EXAMPLES SOON

### Expected release + more explainational examples: 
1st September 2022.
## List
Example: 
```
import (
	"fmt"
	gx "github.com/jtomasevic/gonerix/collections"
)

func main() {
	list := gx.List[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(4)
	list.Add(3)
	fmt.Println(list)
	// [1 2 4 3]

	list.Reverse()
	fmt.Println(list)
	// [3 4 2 1]

	list.Remove(2)
	fmt.Println(list)
	// [3 4 1]

	list.Insert(0, 10)
	fmt.Println(list)
	// [10 3 4 1]

	list.Insert(2, 5, 6, 7, 8)
	fmt.Println(list)
	// [10 3 5 6 7 8 4 1]

	list.RemoveAt(0)
	fmt.Println(list)
	// [3 5 6 7 8 4 1]

	list.RemoveAt(len(list) - 1)
	fmt.Println(list)
	// [3 5 6 7 8 4]
}
```

