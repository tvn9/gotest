package records

type Person struct {
	FirstName string
	LastName  string
	Address   Address
}

type Address struct {
	Addr    string
	Suit    string
	City    string
	State   string
	ZipCode string
	Country string
}
