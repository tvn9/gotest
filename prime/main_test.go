package main

import (
	"fmt"
	"testing"
)

func TestTable_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!\n"},
		{"not prime", 8, false, "8 is not a prime number!\n"},
		{"zero", 0, false, "0 is not a prime!\n"},
		{"one", 1, false, "1 is not a prime!\n"},
		{"negative", -5, false, "Negative numbers are not prime!"},
	}

	for _, e := range primeTests {
		res, msg := isPrime(e.testNum)
		if e.expected && !res {
			t.Errorf("%s: expected true but got false", e.name)
		} else {
			fmt.Println("PASS")
		}
		if !e.expected && res {
			t.Errorf("%s: expected false but got true", e.name)
		} else {
			fmt.Println("PASS")
		}
		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func Test_isPrime(t *testing.T) {
	n := 0
	res, msg := isPrime(n)

	if res {
		t.Errorf("expected false, got %v\n", res)
	}

	if msg != fmt.Sprintf("%d is not a prime!\n", n) {
		t.Error("wrong message returned:", msg)
	}

	n = 7
	res, msg = isPrime(n)
	if !res {
		t.Errorf("test value %d, expected %v, but got %v instead\n", n, false, res)
	}
	testMsg := fmt.Sprintf("%d is a prime number!\n", n)
	if msg != testMsg {
		t.Errorf("expected %s, but got %s instead\n", testMsg, msg)
	}
}
