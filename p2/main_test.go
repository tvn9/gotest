// Improve isPrime test
package main

import (
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
			"prime", 7, true, "7 is a prime number!\n",
		},
		{
			"not prime", 8, false, "8 is not a prime number because it is divisible by 2!\n",
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
