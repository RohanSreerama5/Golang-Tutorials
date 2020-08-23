package main

import (
	"fmt"
)

//This is for pushing and popping off slice (as used like a stack )
func main() {
	a := []int{1, 2, 3, 4, 5} //We don't use ... to literally declare size of slice bc size of slice is dynamic (det at run-time)
	b := a[1:]                //append() can be used for pushing. This can be used for popping. Called a shift operation Takes index 1 value and everything else into new array (reference tpye tho)
	//This removes 1st element from the slice
	fmt.Println(b)

	//TO trim element off end:

	c := a[:len(a)-1]
	fmt.Println(c)

	//TO remove element from middle?

	t := append(a[:2], a[3:]...) //first takes slice that inclduds 1,2. (2 is inclusive.).
	//Then adds to that elemets from 3 to end not inclusive of 3rd element

	fmt.Println(t) //note 3 is removed.

	//NOTE The problem this creates.

	//we have taken a slice of a (1,2) and just added (4,5) to it. 3 is gone. THe old 5 from the origjal a array is still there

	//So if we print a, we get (1,2,4,5,5) instead of the original array.

	fmt.Println(a)

	//How to fix this?

	//We need to use a loop to deep copy the origjnal a array, so that b no longer points to the same underlyoing array as a.
	//This ensures a stays the same

	//NOTE: size of slice is dynamic bc u keep appending to a slice. So if the underlying array is too small, Go
	//will have to create another array in memmory that is bigger to hold the slice. This is done in run-time.

	//Go having to keep resizing and creating new copies can cause app performance loss, so its best ot set a high capacity you know.

}
