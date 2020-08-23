package main

import (
	"fmt"
)

const ( //This is just like an enumeration in C

	errorSpecialist = iota //If our app guards against this strategy do _ = iota
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

func main() {
	var specialistType int                                                                   //= catSpecialist
	fmt.Printf("%v, %T\n", specialistType == catSpecialist, specialistType == catSpecialist) //prints true, bool
}

//Careful:

//If you leave specialistType unitialized, it will default to 0, catSpecialist is also 0 or iota, so if you compare them , it will be true

//How to fix this? Set 0 value of iota to be error

//_ = iota

//_ is Go's only write-only variable

//For iota, we always start w 0. But if we don't care about 0, we don't need to assign memory to a variable contating 0.
//_ tells compiler Yes i know you will generate a value here, but I don't care about it just throw it away

//So if we do _ = iota, we no longer will be able to acces that 0 value of the const block

//ALSO

//Although consts have to be define at compile-time, there is an exception:
//We can still do basic primitive type operations like add, subtract, etc., bit shifting, and bitwise operations

//ie.

// _ = iota + 5 //Useful if you need some offset

//Bit shifting is useful bc we can't raise to powers in Go, would have to import math library.
//But by bit shifting we cna raise to power of 2 because every time you shift a bit one level, you are multiplying the entire number by 2
