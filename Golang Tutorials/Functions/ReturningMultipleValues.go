package main

import (
	"fmt"
)

func main() {
	d, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return //print errr and exit main app
	}
	fmt.Println(d) //print asnwer
}

func divide(a, b float64) (float64, error) { //Panic isn't the best course of action bc it means a pgm can't rlly continue.
	//Instead return an error object letting the calling func know sm went wrong //so we rreturn an error object too
	if b == 0.0 { //guard that checks for error conditions CHECK FOR ERRORS FIRST
		return 0.0, fmt.Errorf("cannot divide by 0") //pass 0.0 bc can't divide
		//generate an error object using the Errorf func
	}
	return a / b, nil //pass nil for error value bc no error present
}
