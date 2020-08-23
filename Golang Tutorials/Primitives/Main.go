package main

import (
	"fmt"
)

func main() {

	/*var n bool = true
	fmt.Printf("%v, %T\n", n, n)*/

	n := 1 == 1 //logical test
	m := 1 == 2
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T", m, m)

	//NOTE: In Go, all values unitialzied manually are initialized to 0. Bool default is false.

}
