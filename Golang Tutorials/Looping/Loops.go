package main

import "fmt"

func main() {
	/*for i := 0; i < 5; i++ {
		fmt.Println(i)
	}*/ //Here i is scoped only to for loop, can't use it outside for loop

	//To use 2 vars do this:

	/*for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}*/

	//Another way

	/*i := 0      //Here i is scoped to main func
	for i < 5 { //This makes it like a do-while or while loop
		fmt.Println(i)
		i++ //WE can also add the incrementer here
	}*/

	//Other special case: This is if you need to break out of a for loop based on some logical test and not just counter:

	/*i := 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break //Leaves entire for loop
		}
	}*/

	//Using a continue statement

	/*for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue //means exit this iteration of the loop and start back over
		}
		fmt.Println(i) //This prints only odd nums
	}*/

	//NESTED LOOPS:

	//How to break out of nested loop? Use a label:

Loop:

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop //means program will break out and go to the Loop label
			}
		}
	}
}
