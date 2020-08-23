package main

//Go does not have exceptions. Because a lot of cases that are considered exceptional in other languages are considered normal
//in Go. For ex, if you try to open a file and that file does not exist, that's pretty normal. It's reasonable to assume
//we might open a DNE file, so we return an error value, NOT throw an exception bc that's not considered exceptional in Go
//There are some things that can cause a Go application to no longer continue. This is considered exceptional. Instead of the word
//exception, Go uses the word panic. Essentially, our app cannot continue functioning and so its panicking as it does not know what to do.
import (
	"fmt"
)

func main() {

	/*a, b := 1, 0
	ans := a / b //invalid
	fmt.Println(ans) //This obviously won't work, so the run-time itself prints out a panic as so: panic: runtime error: integer divide by zero */

	fmt.Println("start")
	panic("something bad happend") //You can make the program panic manually by using the panic keyword and passing it some string
	fmt.Println("end")             //end wont print out cuz the pgm panicked before this could execute.

}
