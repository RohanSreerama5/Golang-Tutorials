package main //packages are how code is organized into sub libraries inside of Go

import "fmt"

func main() {
	fmt.Println("Hello Go!")
}

//go run to temporarily compile and run, rlly fast, will also compile in any 3rd party libraries like fmt

//go build will use the actual package path, not the go file path
//if it finds a main package with a main() it will compile it into an .exe

//go install expects to be pointed to a pacakge that has an entry point and installs that into bin folder,
// so the .exe file will be in bin folder, instead of the main overall directory
