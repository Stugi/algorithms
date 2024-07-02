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

	var map1, map2 = make(map[int]int, size1), make(map[int]int, size2)
	var key1, key2 int

	fmt.Println("Enter first array:")
	for i := 0; i < size1; i++ {
		_, err := fmt.Scanf("%v", &key1)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		map1[key1] += 1
	}
	fmt.Println("Enter second array:")
	for i := 0; i < size2; i++ {
		_, err := fmt.Scanf("%v", &key2)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		map2[key2] += 1
	}
	fmt.Println(map1, map2)
	fmt.Println("Intersection: ", intersect(map1, map2))
}

func intersect(map1, map2 map[int]int) []int {
	var result []int
	for key, _ := range map1 {
		_, ok := map2[key]
		if ok {
			result = append(result, key)
		}
	}
	return result
}
