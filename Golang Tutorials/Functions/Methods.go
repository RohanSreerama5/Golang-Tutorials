package main

import (
	"fmt"
)

func main() {

	g := greeter{ //Declaring a greeter struct
		greeting: "Hello",
		name:     "Go",
	}
	g.greet() //calling greet func on the g struct i Made above
	fmt.Println(g.name)

}

type greeter struct {
	greeting string
	name     string
}

//This below looks like a function, but its a method declaration
//A method is a func executing in a known context
//(g greeter) gives our func access to the greeter type. The greet method gets a copy of the greeter object and given the name g in the context
//of this method

func (g *greeter) greet() { //greeter here is a value type, no pointer. We get a full copy of the struct. Chaning g here won't change the above g
	fmt.Println(g.greeting, g.name) //Go can smartly implicitly deference the pointers.
	g.name = "Rohan"
	fmt.Println(g.name) //operating on copy of greeter obj
	//As expected, copying a huge struct is not good. So we can make our receiver a pointer receiver. g *greeter. LOOK HERE
}
