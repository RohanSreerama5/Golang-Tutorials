package main

import (
	"fmt"
)

func main() {
	var i interface{} = 1 //i is typed to the interface type, which can take any type of Go data type. So we can assign to an interface anyhting rlly
	//like a struct, function bool. data type, etc. In this case, we assigned an int to the interface. Interface type can be type integer
	switch i.(type) { //we are still switching based on a tag
	//i.(type) pulls the underlying type of interface.
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("I is a float64")
		break //this will exit this case early and not do the next print statement
		fmt.Println("Print this")
	case string:
		fmt.Println("i is a string")
	case [2]int:
		fmt.Println("i is [2]int")
	default:
		fmt.Println("i is another type")
	}
}
