package main

import (
	"fmt"
)

func main() {
	aDoctor := struct{ name string }{name: "John Pertwee"} //Can't really use this struct anywhere else or make new instances of it
	//bc this struct doesn't have a type name that can be used to refer to it. It's just fast
	fmt.Println(aDoctor) //this way is used very less, just for a short-lived struct that you don't need to formally define the structs type, just
	//needed some DS to hold it .

	//Now structs, unlime maps and slices, are value types NOT reference types. So a copy of a struct doesn't point to the same underlying data ie.

	anotherDoctor := aDoctor
	anotherDoctor.name = "Tom Baker"
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)

	//So when we pass around structs, we are passing around copies which can be expsnive if the structs are huge. We are not just passing around
	//references to the same underlying data.

	//How to make one structs point to the same underlying data? Just like arrays, use & operator

	newDoc := &aDoctor
	newDoc.name = "Rohan"
	fmt.Println(newDoc)
	fmt.Println(aDoctor)

	//aDoctor is a struct. newDoc is a pointer to the aDoctor struct. So manipulating newDoc alters aDoctor's underlying data.
}
