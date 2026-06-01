package main

import "fmt"

// алгоритм O(n)
func linearSearch(arr []int, target int) int {
	for i, v := range arr{
		if v == target {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 5
	result := linearSearch(arr, target)
	fmt.Println(result)
}