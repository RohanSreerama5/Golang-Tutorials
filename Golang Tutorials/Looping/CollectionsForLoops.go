package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3} //Here we have a slice, which can dynamically change size during run-time. What if we didn't literally define it?
	fmt.Println(s)

	//To loop thru a slice, we use a special for range loop . Helps deal with an aribitrary size slice that changes in run-time
	//IMPORTANT
	for k, v := range s { //k is key and v is value. k is the index of each item. v is the value. s is the collection I'm gonna range over
		//Get the key and value for the current item in the collection
		fmt.Println(k, v) //This concept works for arrays and maps
	}

	//Also watch this:

	d := "hello Go!"
	for k1, v1 := range d { //Using just v1 to print would give an integer, or the Unicode representation of the letter (char)
		fmt.Println(k1, string(v1)) //print out each letter of the string
	}

	//If you don't need the key, just use the write-only var _ instead of k1.
}
