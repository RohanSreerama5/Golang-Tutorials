package main

import (
	"fmt"
)

type Doctor struct { //Any other src file in this package cud access this struct. TO use outside the package (export). capitalize field names
	number     int
	actorName  string
	companions []string //comapnions equal to a slice of strings
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions)
	fmt.Println(aDoctor.companions[2])
}

//Can also use positional syntax to define struct contents as so:

/*aDoctor := Doctor{
	3,
	"Jon Pertwee",
	[]string{
		"Liz Shaw",
		"Jo Grant",
		"Sarah Jane Smith",
	},
}*/

//This way allows u to not have to write out what each thing is : (NOT RECOMMENEDED)

//always use field names when defining struct contents; allows u to skip defining fields too if you don't want to define a field
