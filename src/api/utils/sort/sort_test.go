package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSortDes(t *testing.T) {
	// Init
	elements := []int{9, 8, 3, 4, 5, 7, 6, 0, 1, 2}

	// Execution
	BubbleSortDes(elements)

	// Validation
	if elements[0] != 9 {
		t.Error("first element should be 9")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("last elements should be 0")
	}
	fmt.Println("Testing bubble sort in descending order")
	fmt.Println(elements)
}

func TestBubbleSortAsc(t *testing.T) {
	// Init
	elements := []int{9, 8, 3, 4, 5, 7, 0, 1, 2, 6}

	// Execution
	BubbleSortAsc(elements)

	// Validation
	if elements[0] != 0 {
		t.Error("first element should be 0")
	}

	if elements[len(elements)-1] != 9 {
		t.Error("last elements should be 9")
	}
	fmt.Println("Testing bubble sort in ascending order")
	fmt.Println(elements)
}
