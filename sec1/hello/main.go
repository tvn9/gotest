package main

import (
	"fmt"
	"mastergo/sec1/records"
)

func sayBye(n string) {
	fmt.Printf("Good bye %s!\n", n)
}

func main() {
	name := "Chris"
	fmt.Println("Hello", name+"!")

	sayBye(name)

	// create a record
	person := records.Person{}

	person.FirstName = "Thanh"
	person.LastName = "Nguyen"
	person.Address.Addr = "12345 Smart St."
	person.Address.City = "Tokyo"

	fmt.Println(person)
}
