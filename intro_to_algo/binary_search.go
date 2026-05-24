package main

import "fmt"

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low<=high {
		// mid := (low + high) / 2 in case of overflow
		mid := low + (high - low) / 2
		guess := arr[mid]
		if guess == target {
			return mid
		}
		if guess > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 5
	result := binarySearch(arr, target)
	fmt.Println(result)
}

// БД → B-Tree → O(log n)
// файлові системи → дерево директорій → O(log n)
// DNS → ієрархія доменів → O(log n)
// memory management → багаторівневі таблиці → O(log n)