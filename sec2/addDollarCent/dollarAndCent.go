// Calculating money example
package main

import (
	"fmt"
	"mastergo/utility"
)

func main() {
	// Create a money object
	usd := utility.USDBills{}

	// Adding dollars and cents of (group A and Group B).
	// Enter (dollar, cent, dollar, cents)
	dollars, cents := usd.AddDollarCent(10, 50, 4, 50)
	fmt.Printf("dollars: %d cents: %d\n", dollars, cents)
}
