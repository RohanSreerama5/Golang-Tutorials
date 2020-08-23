package main

import (
	"fmt"
)

func main() {
	var w Writer = ConsoleWriter{} //var Writer equal to a Console Writer instance
	w.Write([]byte("Hello Go!"))
}

//Now structs are basically data containers. On the other hand,interfaces store behaviors instead of raw data
//Interfaces describe behaviors
type Writer interface { //interfaces are also a type //THIS IS THE INTERFACE
	//store method definitions
	Write([]byte) (int, error) //This accpets a slice of bytes and returns an int and err

}

//Now implement it

type ConsoleWriter struct{} //THIS IS THE CONCRETE TYPE THAT IMPLEMENTS THE BEHAVIORS THE INTERFACE DESCRIBES

func (cw ConsoleWriter) Write(data []byte) (int, error) { //method called Write. It accpets slic eof byes and returns int and error
	//this func also has access to a ConsoleWriter object
	n, err := fmt.Println(string(data)) //convert byte slice to string
	return n, err
} //This implicityl implements the Writer interface by creating a method on console writer that has signtaure of Writer interface
