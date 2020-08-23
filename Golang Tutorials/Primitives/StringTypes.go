package main

import (
	"fmt"
)

func main() {
	//Text in Go falls into 2 categories: string

	//String in Go is any UTF-8 character

	s := "This is a string"
	fmt.Printf("%v, %T\n", s, s)
	//string is a char array so you shud be able to index a string
	fmt.Printf("%v, %T", s[2], s[2]) //Prints 105, uint8
	//This is happening because strings in Go are actually aliases for bytes. Strings are encoded in uint8 bit streams

	fmt.Printf("%v, %T\n", string(s[2]), string(s[2])) //need to conv those bytes back into a string

	//Can't Do:

	//s[2] = "u" bc you can't assign a string to a byte, also strings are immutable, you can't change em

	//We do have string concatenation: adding strings together

	s2 := "This is also a string"

	fmt.Printf("%v, %T\n", s+s2, s+s2)

	//Can also convert a string to a collection of bytes, referred to as a slice of bytes

	b := []byte(s) //collection of bytes, collection of uint8 values type: []uint8

	fmt.Printf("%v, %T\n", b, b) //prints array of uint8 values the UTF-8 values of each character in the string

	//How is this useful?

	//If you want to send as a response to a web-service call a string back, you cna convert it to a collection of bytes

	//If you want to send a file back, a file is also just a collection of bytes. This method makes things flexible and universal
	//Don't need to worry about line formating and such if you just convert string sto bytes when sending them off

	//Called converting a string to a byte slice

	// ROON PRIMITIVE DATA TYPE

	// a string represented any UTF-8 character
	//a ROON represents any UTF-32 character

	//any character in UTF-32 can be up to 32 bits long; however, it doesn't have to be that many

	//This means any valid UTF-8 character (8 bits long) is also a valid UTF-32 character

	//Roon delcaration:

	r := 'a'

	fmt.Printf("%v, %T\n", r, r) //prints 97, int32

	//Roons are a true type alias for int32

	//Strings can be converted back forth between bytes and strings; on the other hand, roons are a true type alias for int32

	//Roons can be useful when reading from data streams that are encoded in UTF-32

	//The GO API has special functions to read roons

	//STRINGS represented as a collection of UTF-8 characters

}
