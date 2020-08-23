package main

import "fmt"

// No shortcut := syntax here

//var j int = 58

//To declare a buncha new variables do:

var (
	actorName    string = "Elizabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

func main() {

	//Way 1:
	/*var i int //variable i of int type
	i = 42*/

	//Way 2:
	//var i int = 42

	//Way 3:
	//Go compiler can also figure out what your datatype is on its own

	j := 42.0 //This isn't always best if you want 42 to be a float32 for ex. Go will infer this is an integer. 42. will be inferred float64

	fmt.Println(j)

	fmt.Printf("%v, %T", j, j) //like C, you're printing by giving formats. It prints 42, int

}
