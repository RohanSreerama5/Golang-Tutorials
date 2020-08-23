package main

import (
	"fmt"
)

func main() {

	//or write: var f func() = func() {  }
	f := func() { //Assigned anonymous func to f
		fmt.Println("Hello Go!")
	} //this func is defined as a variable ow, so its free to pass around the application
	f() //invoking func
}
