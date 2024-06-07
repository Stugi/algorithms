package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

// func main() {
// 	// ar := []int{1, 2, 3, 0, -3}
// 	ar := make([]int, 50)
// 	for i := range ar {
// 		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
// 	}

// 	// selectionSort(ar)
// 	fmt.Println(ar, len(ar))

// 	// selectionSortByMax(ar)
// 	fmt.Println(ar)

// 	selectionSortDoubleWay(ar)
// 	fmt.Println(ar)
// }

func selectionSort(ar []int) {
	// ваш код здесь
	for i := 0; i < len(ar); i++ {
		minIndex := i
		for j := i + 1; j < len(ar); j++ {
			if ar[minIndex] > ar[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			ar[i], ar[minIndex] = ar[minIndex], ar[i]
		}
	}
}

func selectionSortByMax(ar []int) {
	// Реализуйте сортировку выбором selectionSortByMax, работающую «справа налево»
	//(поиск максимальных элементов и перемещение их в конец).
	for i := len(ar) - 1; i > 0; i-- {
		maxIndex := i
		for j := 0; j < maxIndex; j++ {
			if ar[maxIndex] < ar[j] {
				maxIndex = j
			}
		}
		if maxIndex != i {
			ar[i], ar[maxIndex] = ar[maxIndex], ar[i]
		}
	}
}

func selectionSortDoubleWay(ar []int) {
	// Реализуйте двунаправленную сортировку выбором.
	// Это действительно сложная задача.
	for i := 0; i < len(ar)/2; i++ {
		minIndex := i
		maxIndex := i
		for j := i + 1; j < len(ar)-i; j++ {
			if ar[minIndex] > ar[j] {
				minIndex = j
			}
			if ar[maxIndex] < ar[j] {
				maxIndex = j
			}
		}
		if i == maxIndex {
			maxIndex = minIndex
		}
		ar[i], ar[minIndex] = ar[minIndex], ar[i]
		ar[len(ar)-1-i], ar[maxIndex] = ar[maxIndex], ar[len(ar)-1-i]

	}
}
