// This app allows user to enter USD bills from 1USD to 100USD and get total USD amount
// This calclation is simular to the bank adding system, the bank allways sort the bill 
// into kinds before they insert into the cound machine, then after they have the total
// number of bills of each kind then they will enter the number to the calculator.
package main

import (
	"fmt"
	"mastergo/utility"
)

func main() {
	// Create the USDBills object
	usd := utility.USDBills{}

	// AddBills adding USD bills by entering number of each type of bill
	// start from 1USD bill, 5USD, 10USD, 20USD, 50USD, and 100USD
	usdTotal := usd.AddBills(1, 1, 1, 1, 1, 1, 1)

	fmt.Printf("Total amount: $%d\n", usdTotal)
}
