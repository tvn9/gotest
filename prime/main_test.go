package main

import (
	"fmt"
	"testing"
)

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
