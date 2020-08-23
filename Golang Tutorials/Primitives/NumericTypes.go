package main

import "fmt"

func main() {
	//every language can choose to define its own size of Int type
	//guaranteed to be at least 32 bits

	//We can have 8,16,32,64-bit integers
	//n := 42
	//fmt.Printf

	//creating unsigned integer:

	//var n uint16 = 42
	//fmt.Printf("%v, %T\n", n, n)

	// BIT OPERATORS:

	/*a := 10             //1010
	b := 3              //0011
	fmt.Println(a & b)  //0010
	fmt.Println(a | b)  //1011
	fmt.Println(a ^ b)  //XOR 1001
	fmt.Println(a &^ b) //AND NOT     0100  */

	// BIT SHIFTING

	/*a := 8              //1000
	fmt.Println(a << 3) //bit left shifts a 3 places, left shifting is just mutliplying a by itself 2^3 * 2^3 = 2^6 = 64 (adding to exponent w base 2 )
	fmt.Println(a >> 3) //2^3 / 2^3 = 2^0 = 1 */

	//FLOATING POINTS

	//float32 and float64 available

	//Declaration:

	/*n := 3.14   //declaration
	n = 13.7e72 //reassignment
	n = 2.1E14

	fmt.Printf("%v, %T", n, n)

	//can also do
	//var n float32 = 3.14 */

	// COMPLEX TYPE

	//have complex64 and complex128

	var n complex64 = 1 + 2i //complex64 composed of 2 float32s
	fmt.Printf("%v, %T\n", n, n)

	//To get real part, use real(n) or imag(n) for imaginary part
	//can also use complex func as so: var n complex128 = complex(5, 12) takes real part and imag part

}
