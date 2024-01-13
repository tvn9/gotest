// Improve isPrime test - with table test cases
package main

import (
	"fmt"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{
			"zero", 0, false, fmt.Sprintf("%d is not prime by definition!\n", 0),
		},
		{
			"one", 1, false, fmt.Sprintf("%d is not prime by definition!\n", 1),
		},
		{
			"negative number", -21, false, fmt.Sprintf("%d negative numbers are not prime, by definition!\n", -21),
		},
		{
			"prime", 7, true, fmt.Sprintf("%d is a prime number!\n", 7),
		},
		{
			"not prime", 8, false, fmt.Sprintf("%d is not a prime number because it is divisible by 2!\n", 8),
		},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got %t\n", e.name, result)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got %t\n", e.name, result)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s\n", e.name, e.msg, msg)
		}
	}
}
