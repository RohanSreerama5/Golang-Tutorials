package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 5; i++ {
		func(i int) {
			//msg := "Hello Go!" //msg only in this scope not in main func
			fmt.Println(i)
		}(i) //Anonymous Function   //Now that we pass in i, we are not reading from outer scope for i anymore. Its passed in

		//Although this seems to work, if the Anon func is executing asynchronously, (not synchronized w for loop) then the for loop counter may just
		//keep incrementing, and we may have odd behavior in the println statement. Best practice is to pass in / provide var to inner func
	}
}
