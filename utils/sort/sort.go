package sort

func BubbleSort(els []int) {
	working := true
	for working {
		working = false
		for i := 0; i < len(els)-1; i++ {
			if els[i] < els[i+1] {
				working = true
				els[i], els[i+1] = els[i+1], els[i]
			}
		}
	}
}
