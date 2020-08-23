package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	defer fmt.Println("This was deferred")
	panic("something bad happend")
	fmt.Println("end")
}

//NOTE: Panic execute after deferred statements os line 8-9 work, then panic is printed.
//Order of execution: execute main function, execute defers, execute panics, then handle return value

//Why is that note important?
//Deferred statements that are going to close resources will succeed if the app panics
//So if somewhere up the call stack, you recover from the panic, you don't have to worry about the resources being left open. THey will get closed off.
