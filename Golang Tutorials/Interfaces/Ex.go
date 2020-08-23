package main

import (
	"fmt"
)

func main() {
	myInt := IntCounter(0)       //have to cast an integer to an IntCounter
	var inc Incrementer = &myInt //create Incrementer and assign it as a pointer to myInt object
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

type Incrementer interface {
	Increment() int
}

type IntCounter int //type alias for int (still an int, just name this type IntCounter)

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
