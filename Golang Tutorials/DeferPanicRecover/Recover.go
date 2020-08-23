package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err) //logging that error
		}
		//This is known as an anonymous function. It has no name and can be called once. Note the () (thats how you call it.)
		//Defer takes as argument function calls (NOT functions) so to defer a function ,we must invoke a call
		//Nothing else can call this anonymous func. It is defined and can be called one time
	}()
	panic("something bad happend")
	fmt.Println("end")

	//recover() will return nil if app is not panicking. If not nil, it will return the err that is causing the app to panic.
}
