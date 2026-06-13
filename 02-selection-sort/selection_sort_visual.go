package main

import (
	"fmt"
)

type Task struct {
	name     string
	duration int
}

// Selection sort по часу виконання задачі
func selectionSort(tasks []Task) []Task {
	n := len(tasks)

	for i := 0; i < n-1; i++ {
		minIndex := i

		for j := i + 1; j < n; j++ {
			if tasks[j].duration < tasks[minIndex].duration {
				minIndex = j
			}
		}

		tasks[i], tasks[minIndex] = tasks[minIndex], tasks[i]

		fmt.Printf("Step %d:\n", i+1)
		printTasks(tasks)
		fmt.Println("-------------------")
	}

	return tasks
}

func printTasks(tasks []Task) {
	for _, t := range tasks {
		fmt.Printf("%s (%dh)\n", t.name, t.duration)
	}
	fmt.Println()
}

func main() {
	tasks := []Task{
		{"Build API", 8},
		{"Fix bugs", 3},
		{"Write docs", 2},
		{"Deploy service", 5},
		{"Code review", 1},
	}

	fmt.Println("Before sorting:\n")
	printTasks(tasks)

	sorted := selectionSort(tasks)

	fmt.Println("After sorting:\n")
	printTasks(sorted)
}