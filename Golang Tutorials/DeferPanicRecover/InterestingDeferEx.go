package main

import (
	"fmt"
)

func main() {
	a := "start"
	defer fmt.Println(a)
	a = "end"
}

//Output is start. Reason: When you defer a  function, it takes the argument at the time the defer is called, not after the main func is done executing
