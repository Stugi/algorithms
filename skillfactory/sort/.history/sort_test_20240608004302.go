package sort

import (
	"math/rand"
	"testing"
)

// запустить бенчмарки go test -bench ./...
// запустить тесты go test ./...

// bubblesort.
func BenchmarkBubbleSort(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		// показывает сколько раз алоцирована память
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})
}

// selectionSort
func BenchmarkSelectionSort(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		// показывает сколько раз алоцирована память ReportAllocs
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})
}

// insertionSort
func BenchmarkInsertionSort(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		// показывает сколько раз алоцирована память
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})
}

// mergeSort
func BenchmarkMergeSort(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		// показывает сколько раз алоцирована память
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			mergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			mergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			mergeSort(ar)
			b.StopTimer()
		}
	})
}

// quickSort
func BenchmarkQuickSort(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		// показывает сколько раз алоцирована память
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})
}

func generateSlice(max, size int) []int {
	ar := make([]int, size)
	for i := range ar {
		// чтобы генерировать числа от -max до max
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
		bubbleSort(tt.input)
		if !ascendingOrder(tt.input) {
			t.Error("order is not valid")
		}
	}
}

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
