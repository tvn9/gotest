package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSortNoError(t *testing.T) {
	// Init
	elements := []int{9, 8, 3, 4, 5, 7, 0, 1, 2}

	// Execution
	BubbleSort(elements)

	// Validation
	if elements[0] != 9 {
		t.Error("first element should be 9")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("last elements should be 0")
	}
	fmt.Println(elements)
}
