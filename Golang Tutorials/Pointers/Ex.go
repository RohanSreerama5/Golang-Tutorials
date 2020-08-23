package main

import (
	"fmt"
)

func main() {
	a := [3]int{1, 2, 3}
	b := a //data that gets copied to b is a pointer, not the underlying data itself.
	fmt.Println(a, b)
	a[1] = 42
	fmt.Println(a, b)

	//A slice is actually a projection of an underlying array. Slice DOES NOT contain the data itself. It contains a pointer to the first
	//element that the slice is pointint to on the underlying array. slices and maps are like this , not prim ttpes n structs.
}
