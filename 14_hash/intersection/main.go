package main

import "fmt"

func main() {
	var size1, size2 int
	fmt.Println("Enter first array size:")
	_, err := fmt.Scanf("%d", &size1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println("Enter second array size:")
	_, err = fmt.Scanf("%d", &size2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	var arr1 = make([]string, size1)
	var arr2 = make([]string, size2)
	fmt.Println("Enter first array:")
	for i := 0; i < size1; i++ {
		_, err = fmt.Scanf("%v", &arr1[i])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
	fmt.Println("Enter second array:")
	for j := 0; j < size2; j++ {
		_, err = fmt.Scanf("%v", &arr2[j])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}

	fmt.Println("Intersection: ", intersect(arr1, arr2))
}

func intersect(arr1, arr2 []string) []string {
	var result []string
	for _, val := range arr1 {
		for _, val2 := range arr2 {
			if val == val2 {
				result = append(result, val)
			}
		}
	}
	return result
}
