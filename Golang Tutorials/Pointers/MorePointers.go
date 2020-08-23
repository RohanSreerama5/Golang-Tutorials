package main

import (
	"fmt"
)

func main() {
	var ms *myStruct
	fmt.Println(ms)         //unlike C, every var u declare is preinitalized to a 0 value. a pointer is intialized to <nil>
	ms = &myStruct{foo: 42} //ms is holding an address of an object that has a field with value 42 in it
	//can also od ms = new(myStruct) //however, this initalies an empty object.
	fmt.Println(ms) //prints &{42}. ms is holding address of an object that has a field w value 42 in it
	//How to get at the foo field?
	(*ms).foo = 16
	fmt.Println((*ms).foo) //ms prints out &{42} (a memory locaiton.) *ms prints out the entire struct: {42}. *ms.foo prints out the foo field

	//can also do this for the same result in GO:

	ms.foo = 81
	fmt.Println(ms.foo) //This does not make sense actually speaking because the pointer ms does not have a field foo. ms is a pointer that points
	//to a struct that has a field foo. However, the compiler is smart so it will see ms.foo as (*ms).foo to help u out
}

type myStruct struct {
	foo int
}
