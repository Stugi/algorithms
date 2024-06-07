package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

// func main() {
// 	ar := make([]int, 50)
// 	for i := range ar {
// 		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
// 	}

// 	bubbleSort(ar)

// 	fmt.Println(ar)

// 	bubbleSortReversed(ar)

// 	fmt.Println(ar)
// }

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

func bubbleSortReversed(ar []int) {
	for i := 0; i < len(ar); i++ {
		for j := 1; j < len(ar); j++ {
			if ar[j-1] < ar[j] {
				ar[j-1], ar[j] = ar[j], ar[j-1]
			}
		}
	}
}

func bubbleSortRecursive(ar []int) {

}
