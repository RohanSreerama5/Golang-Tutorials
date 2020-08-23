package main

//2 remaining collection types available in go are maps and structs
import (
	"fmt"
)

func main() {

	statePopulations := make(map[string]int) //can also use make func to make map
	statePopulations = map[string]int{       //statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	//this map of state populations takes some key, the state name, and map that over to its respective population (int)
	//mapping one key type to one value type

	//Values can be of any type; however,
	//keys must be a type that can be tested for equality i.e. strings, pointers, structs, arrays, interfaces, bools, channels, ints, etc.
	//slices, maps, and other functions cannot be keys as they cannot be tested for equality. Array is a valid key type

	//adding value to map :

	statePopulations["Georgia"] = 10310371

	//deleting value from map:

	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)

	//NOTE: IF you interrogate a map for a key value that does not exist, console will print 0. Use 'ok' to ensure key actually exists in map as so:

	pop, ok1 := statePopulations["Oho"]  //key DNE, so ok will return false and 0 will be still printed.
	pop2, ok := statePopulations["Ohio"] //ths will print 11614373 true

	fmt.Println(pop, ok1) //This is a special syntax
	fmt.Println(pop2, ok)

	//IF you want to just check if a certain key exists: we can use the write only variable _ as so:

	_, ok3 := statePopulations["Rohan"]

	fmt.Println(ok3)

	//To find length of map :

	fmt.Println(len(statePopulations))

	//Note unlike a slice or array, a map does not return its values in the same order we provided them. Order is not guaranteed .

	//NOTE: Just liek slices, underlying data of map is passed by reference. Means manipulating one thing that points to a map alters the
	//entire orignal map too. When passed into functions, maps are passed by reference.

	//Below is an example:

	sp := statePopulations
	delete(sp, "Ohio")
	fmt.Println(sp)
	fmt.Println(statePopulations) //sp and statePopulations will print same thing cuz deleting Ohio deletes it in the original map, sp points
	// to the same underlying data that statePopulations points to.

}
