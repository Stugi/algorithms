package main

import "fmt"

func main() {
	// m1 := []int{0, 1, 2, 3, 4, 5}
	// m2 := []int{9, 7, 4, 1, 3, 5}
	// m1 := []int{0}
	// m2 := []int{}
	// m1 := []int{1, 1}
	// m2 := []int{3, 2, 1}
	m1 := []int{5, 15, 2, 13, 7, 16, 10, 2}
	m2 := []int{1, 9, 7, 4, 6, 2, 1, 13, 22, -3, 12, 76}
	bubbleSort(m1)
	fmt.Println(m1, checkSliceIsSorted(m1))
	bubbleSort(m2)
	fmt.Println(m2, checkSliceIsSorted(m2))
}

func checkSliceIsSorted(a []int) bool {
	for j := 1; j < len(a); j++ {
		if a[j] < a[j-1] {
			return false
		}
	}
	return true
}

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
