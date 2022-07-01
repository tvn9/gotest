package sort

func BubbleSortDes(els []int) {
	run := true
	for run {
		run = false
		for i := 0; i < len(els)-1; i++ {
			if els[i] < els[i+1] {
				run = true
				els[i], els[i+1] = els[i+1], els[i]
			}
		}
	}
}

func BubbleSortAsc(els []int) {
	run := true
	for run {
		run = false
		for i := 0; i < len(els)-1; i++ {
			if els[i] > els[i+1] {
				run = true
				els[i], els[i+1] = els[i+1], els[i]
			}
		}
	}
}
