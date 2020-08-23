package main

import (
	"fmt"
)

func main() {
	greeting := "Hello"
	name := "Stacey"
	sayGreeting(&greeting, &name) //Go run-time makes a copy of name and puts into into name string in sayGreeting. Changing one won't change the other
	fmt.Println(name)             //by passing address of name, we changed name to Ted
}

func sayGreeting(greeting, name *string) { //useful to pass in pointers instead of full data structures bc that can result in huge copying
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name) //maps and slices are default reference types, so when they're passed in, you will edit them everywhere. THis
	//is bc those 2 types have (are) internal pointers to their underlying data. So they're always gonna act like you're passing pointers in
}
