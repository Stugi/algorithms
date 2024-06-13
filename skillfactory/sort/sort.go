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

// быстрая сортировка.
func quickSort(arr []int) []int {
	// ваш код здесь
	left := 0
	right := len(arr) - 1

	return quick(arr, left, right)
}

func quick(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quick(arr, low, p-1)
		arr = quick(arr, p+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]

	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]

	return arr, i
}
