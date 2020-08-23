package main

import (
	"fmt"
	"reflect" //We use Go's reflection package to handle tags
)

type Animal struct {
	//Let's say we need this for a web application form. The max length of the name must be 100 characters. We use a tag as so:
	//Format: back ticks as the delimeters of the tag. and then we have a space delimited key value pair
	Name string `required max:"100"` //means Name field is required and max length is 100 char. We now have 2 space-delimited sub tags, required & max:100

	Origin string
}

func main() {
	t := reflect.TypeOf(Animal{})     //Need to get type of object I'm working with. So I pass in an Animal object. Animal{} initializes an empty Animal object
	field, _ := t.FieldByName("Name") //grabs a field from the Animal type, in this case grabbing Name field.
	fmt.Println(field.Tag)            //gets tag property of the Name field
	//This tag is just a random string. Your code logic later has to understand this tag and ensure that the Name field is filled out
	//and the tag requirements are met. Tags just provide a string of text, sm else has to figure out what to do w it.

}
