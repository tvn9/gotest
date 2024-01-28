package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// print a welcome message
	intro()

	// create a channel to indicate when the user want to quit
	doneChan := make(chan bool)

	// start a goroutine to read user input and run program
	go readUserInput(doneChan)

	// block until the doneChan gets a value
	<-doneChan

	// close the channel
	close(doneChan)

	// say goodbye
	fmt.Println("Goodbye.")

}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	// read user input
	scanner.Scan()

	// check user input
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole number!", false
	}
	_, msg := isPrime(numToCheck)
	return msg, false

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

func intro() {
	fmt.Println("Is it Prime?")
	fmt.Println("------------")
	fmt.Println("Enter a whole number, to test if it is a prime, or q to quite.")
	prompt()
}

func prompt() {
	fmt.Print("-> ")
}
func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not a prime!\n", n)
	}

	// negative number is not prime
	if n < 0 {
		return false, "Negative numbers are not prime!"
	}

	// use the modulus operator repeatedly to see if we have a prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			// not a prime number
			return false, fmt.Sprintf("%d is not a prime number!\n", n)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!\n", n)
}
