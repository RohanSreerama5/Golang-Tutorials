package main

import (
	"fmt"
)

func main() {
	//In Go, most common use case for arrays is to back a slice
	/*a := []int{1, 2, 3} //A slice is initialized as a literal as so
	fmt.Println(a)      //looks just like an array. Everything we do w array we can do w slice
	fmt.Printf("Length: %v\n", len(a))
	//WE also have cap() for capacity in slices

	//That's bc the # elements in the slice doesn't have to be the same as the number in the backing array
	//The slice is a projection of the underlying array, so we cud have a huge array and only be looking at a slice

	fmt.Printf("Capacity: %v\n", cap(a)) //In this case cap = len, so underlying array is  same size as slice

	//UNLIKE arrays, slices are naturally reference types. They refer to the same underlying data. So you don't need b := &a, just do b := a



	//MANY different ways to make a slice: */

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:] //slice of all elements
	b[4] = 40
	c := a[3:]  //slice from 4th element to end
	d := a[:6]  //slice first 6 elements
	e := a[3:6] //slice 4th 5th and 6th elements
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	//All of these new slices still point to a's underlying memory locations so altering a value in a[] alters a value in all the others

	//Slicing operations cam have as their source an array or a slice. IMPORTANT. This means a could'ev been intialized
	//as so: like an array: a := [...]int{etc.} //Result will be the same exactly

	//Last way to make a slice is using the make() func

}
