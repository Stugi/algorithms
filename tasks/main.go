package main

import "fmt"

func main() {
	fmt.Println(findMaxNegative([]int{-1, -4, 1, 2, 4}))
	fmt.Println(findMostOftenRepeated([]int{1, 2, 1, 2, 2, 2, 3, 0, 4, 3, 2}))
	fmt.Println(findMostOftenRepeatedOptimized([]int{1, 2, 1, 2, 2, 2, 3, 0, 4, 3, 2}))
	fmt.Println(trimNegative([]int{1, -2, 1, 2, 2, -2, 3, 0, 4, 3, 2}))
	fmt.Println(trimLessAverage([]int{1, -32, 1, 2, 100, 22, 3, 0, 4, 3, 2}))
}

func findMaxNegative(array []int) (element int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("couldn find max element in blank slice")
	}
	// maxNegativeElement := 0
	// for _, elem := range array {
	maxNegativeElement := array[0]
	for _, elem := range array[1:] {
		if elem < maxNegativeElement {
			maxNegativeElement = elem
		}
	}

	if maxNegativeElement == 0 {
		return 0, fmt.Errorf("array havnt minus element")
	}

	return maxNegativeElement, nil
}

func findMostOftenRepeated(array []int) (mostOften int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found repeated numbers in empty slice")
	}

	var maxIndex, maxCount = 0, 0
	for i, number := range array {
		currentCount := 0
		for _, numberToCompare := range array {
			if number == numberToCompare {
				currentCount++
			}
		}

		if currentCount > maxCount {
			maxIndex = i
			maxCount = currentCount
		}
	}

	return array[maxIndex], nil
}

func findMostOftenRepeatedOptimized(array []int) (mostOften int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found repeated numbers in empty slice")
	}

	var maxIndex, maxCount = 0, 0
	for i, number := range array {
		currentCount := 0
		for _, numberToCompare := range array[i:] {
			if number == numberToCompare {
				currentCount++
			}
		}

		if currentCount > maxCount {
			maxIndex = i
			maxCount = currentCount
		}
	}

	return array[maxIndex], nil
}

func findMostOftenRepeatedWithMap(array []int) (mostOften int, err error) {
	// todo
	return 0, nil
}

func trimNegative(array []int) (returnArray []int, err error) {
	for _, elem := range array {
		if elem >= 0 {
			returnArray = append(returnArray, elem)
		}
	}

	return returnArray, nil
}

func trimLessAverage(array []int) (returnArray []int, err error) {
	if len(array) == 0 {
		return array, nil
	}
	var sum, average int
	for _, el := range array {
		sum += el
	}
	average = sum / len(array)

	for _, el := range array {
		if el >= average {
			returnArray = append(returnArray, el)
		}
	}

	return returnArray, nil
}
