// Improve p1 - isPrime Test
package main

import "fmt"

func main() {
	n := 1

	_, msg := isPrime(n)

	fmt.Println(msg)
}

func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime by definition
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime by definition!\n", n)
	}

	// negative number are not prime
	if n < 0 {
		return false, fmt.Sprintf("%d negative numbers are not prime, by definition!\n", n)
	}

	// use the modulus operator repeatedly to use if we have a prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			// not a prime number
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d!\n", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!\n", n)

}
