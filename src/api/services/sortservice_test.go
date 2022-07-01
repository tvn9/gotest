package services

import (
	"fmt"
	"testing"
)

func TestSortAsc(t *testing.T) {
	elements := []int{9, 8, 3, 4, 5, 7, 6, 0, 1, 2}
	fmt.Println("Testing Bubble Sort:")
	fmt.Println(elements)
	SortAsc(elements)

	// Validation
	if elements[0] != 0 {
		t.Error("first element should be 0")
	}

	if elements[len(elements)-1] != 9 {
		t.Error("last elements should be 9")
	}
	fmt.Println("bubble sort in ascending order")
	fmt.Println(elements)
}

func TestSortDes(t *testing.T) {
	elements := []int{9, 8, 3, 4, 5, 7, 6, 0, 1, 2}
	fmt.Println("Testing Bubble Sort:")
	fmt.Println(elements)
	SortDes(elements)

	// Validation
	if elements[0] != 9 {
		t.Error("first element should be 9")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("last elements should be 0")
	}
	fmt.Println("bubble sort in descending order")
	fmt.Println(elements)
}
