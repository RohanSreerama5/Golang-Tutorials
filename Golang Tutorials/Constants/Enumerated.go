package main

import (
	"fmt"
)

//const a = iota //prints 0, int (Go inferred its type to be int)

//Iota is a counter that can be used when creatin enumerated constants

const (
	a = iota
	b = iota
	c = iota //Prints 0 1 2
)

/*const (
	a = iota //I cna also do this. Go compiler will infer the pattern used to assign the varibales. So it assigns iota to b and c as well
	b
	c //Prints 0 1 2
)*/

//Iota is scoped to its const block

const (
	a2 = iota
)

func main() {
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", a2, a2)

}
