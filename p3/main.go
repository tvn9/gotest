// Improve p2 - main
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Print a welcome message
	intro()

	// create a  channel to indicate when the user want to quit
	doneChan := make(chan bool)

	// start a goroutine to read user input
	go readUserInput(doneChan)

	// block until the doneChan gets a value
	<-doneChan

	// close the channel
	close(doneChan)

	// say goodbye
	fmt.Println("Goodbye.")

}

func readUserInput(doneChan chan bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		res, done := checkNumbers(scanner)

		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	// read user input
	scanner.Scan()

	//  check to see if the use want to quit
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	// convert user input (string) to int
	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole number!\n", false
	}

	_, msg := isPrime(num)

	return msg, false
}

func intro() {
	fmt.Println("Is it Prime?")
	fmt.Println("-------------")
	fmt.Println("Enter a whole number, and we'll tell you if it is prime number or not. Enter q to quit.")
	prompt()
}

func prompt() {
	fmt.Print("-> ")
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
