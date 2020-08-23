package main

//Defer: idea of invoking a func, but delay its execution to some future time
//Panic: how app can panic, how GO app enters state where it cannot run anymore and how Go run-time can trigger that as well as how we can
//manually trigger that
//Recover: how to save the program when the program panics
import (
	"fmt"
)

func main() {
	defer fmt.Println("start")
	defer fmt.Println("middle") //Defers execution of this statement till rest of function is executed. Executes this last
	defer fmt.Println("end")
	//Essentially Go executes line 11. Then it recognizes there is a deferred func and executes line 13. THen Go exits main() and checks
	//if there are any deferred functions. There is on line 12, so Go will execute line 12 before returning back to its caller func.
	//Essentially, line 12 will execute after the function is done executing and before the function returns to its caller.
}

//Defers are executed in LIFO (Last in First Out) order. So if I defer all of the above, then u print out end, middle, start.
//IF you defer all of em, the statements are executing still after the main() is done but before it returns to its caller func
