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
// 	ar := make([]int, 50)
// 	for i := range ar {
// 		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
// 	}

// 	ar = mergeSort(ar)

// 	fmt.Println(ar)
// }

func mergeSort(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}
	first := ar[:len(ar)/2]
	second := ar[len(ar):]

	return merge(first, second)
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
