// Just some example of go package
package utility

type usd int
type USDBills struct {
	cent1, cent5, cent10, cent50, usd1, usd2, usd5, usd10, usd20, usd50, usd100 usd
}

// just return what ever string received
func Say(s string) string {
	return s
}

// Add - adds two numbers of type int and return result of type int
func Add(n1, n2 int) int {
	return n1 + n2
}

// multiply - multiplys two numbers of type int and return result of type int
func Multiply(n1, n2 int) int {
	return n1 * n2
}

// Devide - devides two numbers of type int and return result of they int
func Devide(n1, n2 int) int {
	return n1 / n2
}

// Add money
func (db *USDBills) AddDollarCent(d1, c1, d2, c2 usd) (dollars, cents usd) {
	db.cent1 = 1
	db.usd1 = db.cent1 * 100

	d1 = d1 * db.usd1
	c1 = c1 * db.cent1
	c1 = c1 + d1

	d2 = d2 * db.usd1
	c2 = c2 * db.cent1
	c2 = c2 + d2
	cents = c1 + c2
	dollars = cents / db.usd1
	cents = cents % db.usd1
	return dollars, cents
}

// AddBills allow input of bill in kinds starting from 1USD, 5USD, 10USD, 20USD, 50USD, and 100USD
func (db *USDBills) AddBills(oneUSD, twoUSD, fiveUSD, tenUSD, twentyUSD, fityUSD, hundredUSD usd) (usdTotal usd) {

	db.usd1 = 1
	db.usd2 = 2
	db.usd5 = 5
	db.usd10 = 10
	db.usd20 = 20
	db.usd50 = 50
	db.usd100 = 100
	usd1Sub := oneUSD * db.usd1
	usd2Sub := twoUSD * db.usd2
	usd5Sub := fiveUSD * db.usd5
	usd10Sub := tenUSD * db.usd10
	usd20Sub := twentyUSD * db.usd20
	usd50Sub := fityUSD * db.usd50
	usd100Sub := hundredUSD * db.usd100

	usdTotal = usd1Sub + usd2Sub + usd5Sub + usd10Sub + usd20Sub + usd50Sub + usd100Sub
	return usdTotal
}

// Adding Coins
