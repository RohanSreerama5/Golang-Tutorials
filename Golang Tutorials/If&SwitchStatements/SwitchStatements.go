package main

import (
	"fmt"
)

func main() {

	/*switch 5 { //2 is the tag.
	case 1: //1 is compared to 2. .
		fmt.Println("one")
	case 2, 5, 10: //2 is compared to 2. //Can also test multiple nums at a time as so:
		fmt.Println("two")
	default: //if nothing works, this is the defualt
		fmt.Println("not one or two")
	}*/

	//Can also do this with an initalizer as so :

	/*switch i := 2 + 3; i { //testing agaginst the tag i
	case 1: //1 is compared to 2. .
		fmt.Println("one")
	case 2, 5, 10: //2 is compared to 2. //Can also test multiple nums at a time as so:
		fmt.Println("two")
	default: //if nothing works, this is the defualt
		fmt.Println("not one or two")
	}*/

	//WE can also have tagless switch statements as so:

	i := 10
	switch { //No tag here
	//The way we wrote this, 10 passes case 1 and 2. But it will only do case 1. If you want your case to fallthrough. ie. do case 1 and 2 use fallthrough keyword
	case i <= 10: //Now we're using the cases as the logical tests
		fmt.Println("one")
		fallthrough
	case i <= 20: //2 is compared to 2. //Can also test multiple nums at a time as so:
		//Even if this case is false, it will run if I write fallthrough above. Fallthrough is logic-less, it just makes the next case execute.
		fmt.Println("two")
		fallthrough
	default: //if nothing works, this is the defualt
		fmt.Println("not one or two")
	}

}
