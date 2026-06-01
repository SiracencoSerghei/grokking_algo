package main

import "fmt"

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}

	return smallestIndex
}

func selectionSort(arr []int) []int {
	var newArr []int

	for len(arr) > 0 {
		smallestIndex := findSmallest(arr)

		newArr = append(newArr, arr[smallestIndex])

		arr = append(
			arr[:smallestIndex],
			arr[smallestIndex+1:]...,
		)
	}

	return newArr
}

func main() {
	arr := []int{5, 3, 6, 2, 10}

	fmt.Println(selectionSort(arr))
}