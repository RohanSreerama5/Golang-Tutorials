package main

import (
	"fmt"
)

func main() {
	/*if true {
		fmt.Println("The test is true")
	} */

	statePopulations := map[string]int{ //statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	if pop, ok := statePopulations["Florida"]; ok { //ok is the boolean test being used. If value exists, then pop (the value) is printed out
		fmt.Println(pop) //pop is scoped to the if block
	}

	//pop, ok := statePopulations["Florida"] is the initalizer. IT allows us to run a statement and generate some information that sets us up to work w
	//the if block. The intializer here is generating a boolean result in ok. Which we use as a boolean test in the if block
}
