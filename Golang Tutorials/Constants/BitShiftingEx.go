package main

import (
	"fmt"
)

const (
	_  = iota             //ignore first value by assigning it to blank identifier
	KB = 1 << (10 * iota) //bit shift 1 by 10*iota = 10 (same as 1*2^10)
	MB                    //2^100 etc.
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fileSize := 4000000000
	fmt.Printf("%.2fGB".fileSize / GB) //%2f means expecting floating point of 2 decimal places
}
