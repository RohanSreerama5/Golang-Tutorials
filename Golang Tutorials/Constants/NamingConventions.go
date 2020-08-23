package main

import (
	"fmt"
)

func main() {
	//const myConst   //Can't use uppercase bc that wud tell GO to export this constatn and make it accessible outside the pacakge

	//TYPED CONSTANT - means after variable name we list its type in declaration

	const myConst int = 42 //Must be defined at compile time, cannot set const equal to some function return value that is determined in run-time
	fmt.Printf("%v, %T\n", myConst, myConst)

	//can also do shadowing of constatns ie. declaring at package level and then redeclaring again inside main func for ex. inner scope wins

	//UNTYPED CONSTANT

	const b = 42 //Go compiler infers the type of this number, an integer
	var c int16 = 27
	fmt.Printf("%v, %T\n", c+b, c+b) //prints 69, int16 without error because compiler replaces every b instance with a literal 42 so type
	//no longer becomes an issue

	//Thus we have this sort of idea of implicit type conversion when working with constants

}
