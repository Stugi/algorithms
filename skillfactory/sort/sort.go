package sort

func bubbleSort(ar []int) {
	isNotSorted := true
	for i := 1; i < len(ar); i++ {
		isNotSorted = true
		for j := 1; j < len(ar); j++ {
			if ar[j-1] > ar[j] {
				ar[j-1], ar[j] = ar[j], ar[j-1]
				isNotSorted = false
			}
		}
		if isNotSorted {
			return
		}
	}
}
