package sort

import (
	"math/rand"
	"testing"
)

func BenchmarkBubbleSort(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})
}

func generateSlice(max, size int) []int {
	ar := make([]int, size)
	for i := range ar {
		ar[i] = rand.Intn(max*2) - max
	}

	return ar
}

func TestSort_Buble(t *testing.T) {
	testCases := []struct {
		input []int
	}{
		{input: []int{5, 3, 5, 23, 6}},
		{input: []int{5, 123}},
		{input: []int{5, 1, 2, 3}},
		{input: []int{-35, 3, -5, 23, 6}},
	}

	ascendingOrder := func(in []int) bool {
		prev := in[0]

		for _, el := range in[1:] {
			if el < prev {
				return false
			}
			prev = el
		}
		return true
	}

	for _, tt := range testCases {

	}
}
