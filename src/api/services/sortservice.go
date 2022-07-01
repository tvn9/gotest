package services

import "github.com/tvn9/gotest/src/api/utils/sort"

func SortDes(els []int) {
	sort.BubbleSortDes(els)
}

func SortAsc(els []int) {
	sort.BubbleSortAsc(els)
}
