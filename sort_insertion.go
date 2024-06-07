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
// 		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
// 	}

// 	insertionSort(ar)

// 	fmt.Println(ar)
// }

func insertionSort(arr []int) {
	// for i := 1; i < len(ar); i++ {
	// 	index := i
	// 	for j := i - 1; j >= 0; j-- {
	// 		if ar[j] < ar[i] {
	// 			index = j
	// 		} else {
	// 			continue
	// 		}
	// 	}
	// 	if index != i {
	// 		ar[i], ar[index] = ar[index], ar[i]
	// 	}
	// }
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}

}
