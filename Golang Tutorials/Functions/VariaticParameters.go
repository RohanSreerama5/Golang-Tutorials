package main

import (
	"fmt"
)

func main() {
	s := sum("The sum is", 1, 2, 3, 4, 5) //s is now a pointer to an integer
	fmt.Println("The sum is ", *s)
}

func sum(msg string, values ...int) *int { //I've got 1 var here called values. It's type (int) is preceded by ... This tells Go run-time to take in all of
	//the last arguments that are passed in and wrap them up into a slice by the name values.

	//You can only have 1 variatic parameter and it must be the last one. THis is because if the VP is first, then Go run-time can't figure out
	//where the variatic arguments end and whre other arguments begin in the function call
	fmt.Println(values)
	result := 0
	for _, v := range values { //looping over values in slice and summing them up
		result += v
	}
	return &result //now returning interger pointer. We are returning the address of the result
	//This shouldn't typically make sense bc after this function is removed from the stack frame memory location &result wudve been
	//freed. You're returning now a pointer to a memory location that got freed. You don't rlly know whats there. But in Go, it understands
	//you're returning the address of a local variable. So it will put that on the heap (shared memory) automatically. Making this work
}

//GO also allows u to return a local variable as a pointer
//return result: Go will copy result to another var and thats what got assigned
