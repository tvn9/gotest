package main

import (
	"fmt"
	"testing"
)

func Test_isPrime(t *testing.T) {
	fmt.Println("TEST - Not Prime")
	result, msg := isPrime(0)

	fmt.Printf("%t - %s\n", result, msg)

	if result {
		t.Errorf("with %d, expected true but receive %t\n", 0, result)
	}

	if msg != "0 is not prime by definition!\n" {
		t.Error("wrong message returned:", msg)
	}

	fmt.Println("TEST - Is Prime")
	result, msg = isPrime(3)

	fmt.Printf("%t - %s\n", result, msg)

	if !result {
		t.Errorf("with %d, expected false but receive %t\n", 3, result)
	}

	if msg != "3 is a prime number!\n" {
		t.Error("wrong message returned:", msg)
	}
}
