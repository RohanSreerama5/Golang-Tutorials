package main

//NOTE: GO does not support traditional OOP things like inheritance. Go uses a similar principle called composition (embeddimg)
import (
	"fmt"
)

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal   //We say Bird has animal like characterisitics by embedding Animal struct like so
	SpeedKPH float32
	CanFly   bool //WE say Bird has-a Animal. It has animal-like characteristics, so we embed Animal. But Bird is still an independent struct. It's not
	//an is-a relation like in inheritance. Invalid to say Bird is-a Animal. This is a composition relation. Cannnot use bird
	//and animal interchangeably. Need interfaces to use data interchangeably
}

func main() {
	/*b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false */

	//can also declare like this:

	b := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly:   false,
	}
	fmt.Println(b.Name)
}
