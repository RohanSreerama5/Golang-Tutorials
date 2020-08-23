package main

import (
	"fmt"
)

//make() can take 2 or 3 arguments
func main() {
	a := make([]int, 3) //making a slice using make(type of object we want to create, length of slice)

	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	//we can also do:

	//b := make([]int, 3, 100) //This means the underlying array is of size 100, so length is 3 and capacity is 100

	//Why do this?

	//Unlike arrays, slices don't have to have a fixed size. We can add and remove elements

	//ADDING TO SLICE:

	e := []int{}
	fmt.Println(e)
	fmt.Printf("Length: %v\n", len(e))
	fmt.Printf("Capacity: %v\n", cap(e))
	e = append(e, 1, 3, 6, 7, 7, 10, 11) //adds #1 to e
	// can also append multiple as so: e = append(e, 1, 2 ,4 ,5 )

	fmt.Println(e)
	fmt.Printf("Length: %v\n", len(e))
	fmt.Printf("Capacity: %v\n", cap(e))

	//When we initalized a slice to value  e, Go looks for a memory location to store the slice e. Since it didn't have to store anything, Go
	//created an underlying array of 0 elements. As soon as we append #1 to the slice, it couldn't fit this into a 0-element array, so it needed to
	//copy all of the elements (nothing in this case) to a new array of a larger size.

	//In this case Go created a new array of size 1 and put #1 into that array

	//As things get very large, copy operations become tough and expensive

	//This is where the 3-arg make() func comes in

	//If we set the capacity to some large thing like 100, we can keep appending elements to that same underlying array
	//without having to create a new one  and runing out of space

	//so how does Go resize the underlying array when a slice takes up all of the original array space?

	//Go usually just doubles the size of the original array when its filled/
	//So it goes 0 1 2 4 8 16 32 etc.

	//----------------------------------------------------------------------

	//How to combine 2 slices? Do using spread operator:

	//a = append(a, []int{1, 2 , 3}...)
	//same as a = append(a, 1, 2, ,3 )

}
