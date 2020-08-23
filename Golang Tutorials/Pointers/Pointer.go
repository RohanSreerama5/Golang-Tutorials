package main

import (
	"fmt"
)

func main() {

	//Cannot do pointer arthimetic in Go. So no subracting or adding to memory locations
	var a int = 42
	var b *int = &a //b is a pointer to an integer and b points to a.
	a = 27
	fmt.Println(a, *b)
	*b = 14
	fmt.Println(a, *b)
}
