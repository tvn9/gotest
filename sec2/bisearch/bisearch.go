// Binary search
package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func biSearch(ar []int, tg int, lo int, hi int) int {
	if hi < lo {
		return -1
	}
	mid := int((lo + hi) / 2)
	if ar[mid] > tg {
		return biSearch(ar, tg, lo, mid)
	} else if ar[mid] < tg {
		return biSearch(ar, tg, mid+1, hi)
	} else {
		return mid
	}
}

func iterBiSearch(ar []int, tg int, lo int, hi int) int {
	startIndex := lo
	endIndex := hi
	var mid int
	for startIndex < endIndex {
		mid = int((startIndex + endIndex) / 2)
		if ar[mid] > tg {
			endIndex = mid
		} else if ar[mid] < tg {
			startIndex = mid
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please enter a number for binary search")
		return
	}

	target, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("%s is not a number\n", args[0])
		return
	}

	nums := []int{21, 24, 31, 41, 51, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	// for binary search first we must sort the nums slice
	sort.Ints(nums)

	fmt.Println(nums)

	// This step is not part of the binary search argorithms
	for _, v := range nums {
		if target == v {
			fmt.Printf("Your search number is %d\n", target)

			find := biSearch(nums, target, 0, len(nums)-1)
			fmt.Printf("biSearch func found %d at index %d\n", target, find)

			find1 := iterBiSearch(nums, target, 0, len(nums)-1)
			fmt.Printf("iterBiSearch func found %d at index %d\n", target, find1)
			return
		} else {
			continue
		}
	}
	fmt.Printf("%d is not in the list.\n", target)
}
