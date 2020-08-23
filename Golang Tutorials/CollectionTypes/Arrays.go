package main

import (
	"fmt"
)

//2 collection types: arrays and slices

func main() {
	grades := [...]int{97, 85, 93} //When declaring literally with elements, just use ... instead of typing 3 bc that declares size twice

	fmt.Printf("Grades: %v", grades)

	var students [3]string
	students[0] = "Lisa"
	fmt.Printf("Students: %v", students)

	//Use len(nameOfArray) to get length of array

	//Array of arrays: 2D array

	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}

	fmt.Println(identityMatrix)

	//In GO, arrays are considered values  NOT POINTERS
	//Usually, when you create an array its actually pointing to the values in that array
	//So when you manipulate, its just passing around the same underlying data

	//In Go, when you copy an array you are creating an entirely new array with new data (same number values ofc).
	//This means the new array is not pointing to the same underlying values that the original array pointed.
	//When you create a new array, it is not poitning to the same underlying data, its pointing to a new diff set of data

	//Ie.

	a := [...]int{1, 2, 3}
	b := a //Go is literally copying the entire array (this cud slow program if array is huge )
	b[1] = 5
	fmt.Println(a) //prints [1 2 3 ]
	fmt.Println(b) //prints [1 5 3 ]

	//a and b are diff bc b[1] does not point to the same underlying memory location. In Go, a whole new array is created

	//ie. If you pass an array into a function, Go will copy that entire array over. So you can't rlly edit that original array, it's
	//gonna edit the copy Go created

	//HOW CAN I FIX THIS so that a and b are the same always?
	// do b := &a. b is now equal to address of a. So b is poitning to same underlying data as 'a' array.
	//it'll print oit a and b as the same

}
