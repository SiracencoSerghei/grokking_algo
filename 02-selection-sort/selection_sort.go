package main

import "fmt"

func selectionSort(arr []int) []int {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		smallestIndex := i

		// Find smallest element in remaining array
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[smallestIndex] {
				smallestIndex = j
			}
		}

		// Swap
		arr[i], arr[smallestIndex] = arr[smallestIndex], arr[i]

		fmt.Printf("Step %d: %v\n", i+1, arr)
	}

	return arr
}

func main() {
	numbers := []int{5, 3, 10, 2, 6}

	fmt.Println("Before:", numbers)

	sorted := selectionSort(numbers)

	fmt.Println("After: ", sorted)
}