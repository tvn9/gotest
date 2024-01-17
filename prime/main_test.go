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
}
