//IMPORTANT VERY COOL STUFF

package main

import (
	"fmt"
)

const ( //These are 8 access roles you cna have at a company
	isAdmin          = 1 << iota //00000001
	isHeadquarters               //2
	canSeeFinancials             //4 00000100

	canSeeAfrica //8 = 2^3
	canSeeAsia
	canSeeEurope // 00100000
	canSeeNorthAmerica
	canSeeSouthAmerica

	//Each one of these constants is going to occupy one location in a byte. A byte is 8 bits and we have 8 constants (8 locations)
)

func main() {

	//defining the roles in a single byte
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope //When you OR these, your resultant bit pattern (byte) will have these flags as 1 (true)
	fmt.Printf("%b\n", roles)                                  //prints 100101 or 00100101

	//We've encoded 8 access roles for a user into a single byte of data (very efficient)

	//Now to check if this user is an admin?

	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin) //This applies a bit mask, so that we can compare it to isAdmin constant
	//and check if i have admin access

	//roles are for my user. I'm ANDing that with isAdmin to check if I actually have admin access
	//The bit mask is just isolating the bit  I need (the admin bit )

	fmt.Printf("Is HQ? %v", isHeadquarters&roles == isHeadquarters) //false bc don;'t have hq true access
}
