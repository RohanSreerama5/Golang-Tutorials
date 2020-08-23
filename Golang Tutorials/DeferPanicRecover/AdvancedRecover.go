package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil { //In this case, we are handling the inital panic. REcover() told us we can go run the caller func still
			log.Println("Error:", err)
			panic(err) //Add this line if the err is something you don't know to handle and you want to repanic the app
			//With the above statement "end" won't print out bc we have no other defer func to recover from this panic. WE just panic
			//Go run-time is hit and Go run-time can't handle panics, so it will exit.
		}
	}()
	panic("Something bad happend") //start is printed, then about to panic, then app panics.
	fmt.Println("done panicking")  //cannot execute bc app panicked
	//So panicker() stops and executes any deferred functions.
}

//So ooce panicker() halts, we can still excute its caller main func. So that's why main continues to print "end" This is
//because recover() said that your app is still in a state where it can continue

//Recover only useful inside of defered funcs bc once an app panics, it terminates that func and only executes deferred funcs. So
//putting recover() inside a deferrred func can help it actually recover
//The recover() ig will look fora a panicking situation and decide what to do if the app is panicking

//THe big idea:

//Current function will not attempt to continue, but higher functions in call stack will.
//If you dont like this behavior bc you cant handle that particular err (meaning you cant let the rest of the funcs execute). then rethrow the panic

//IF you call recover(), look atht err it returns and realize you can't handle it, rethrow the panic
