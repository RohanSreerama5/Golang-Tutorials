package main

import (
	"fmt"
)

func main() {
	//Declaring func signature for divide func that takes 2 float64s and returns float64 and err object
	var divide func(float64, float64) (float64, error) //this func takes float64s amd returns a float64 and err object
	divide = func(a, b float64) (float64, error) {     //initializing var to an anonymous function that takes a and b that are float64s and returns those
		if b == 0.0 {
			return 0.0, fmt.Errorf("Cannot divide by zero")
		} else {
			return a / b, nil
		}
	}
	d, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

}
