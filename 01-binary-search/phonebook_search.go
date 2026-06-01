package main

import "fmt"

type Contact struct {
	Name  string
	Phone string
}

func binarySearch(phonebook []Contact, target string) (string, bool) {
	low := 0
	high := len(phonebook) - 1

	for low <= high {

		// overflow-safe midpoint
		mid := low + (high-low)/2

		guess := phonebook[mid]

		switch {
		case guess.Name == target:
			fmt.Printf(" guess.Name == target: %s == %s\n", guess.Name, target)
			fmt.Println("case1")
			return guess.Phone, true

		case guess.Name > target:
			high = mid - 1
			fmt.Printf(" guess.Name > target: %s > %s\n", guess.Name, target)
			fmt.Println("case2")

		default:
			low = mid + 1
			fmt.Printf(" guess.Name < target: %s < %s\n", guess.Name, target)
			fmt.Println("case3")
		}
	}

	return "", false
}

func main() {

	// MUST be sorted
	phonebook := []Contact{
		{"Alice", "555-123-4567"},
		{"Andrew", "555-987-6543"},
		{"Bob", "555-555-5555"},
		{"Charlie", "555-111-2222"},
		{"Jane", "987-654-3210"},
		{"John", "123-456-7890"},
		{"Maria", "777-888-9999"},
		{"Sergio", "444-555-6666"},
		{"Zoe", "111-222-3333"},
	}

	name := "Pietro"

	phone, found := binarySearch(phonebook, name)

	if found {
		fmt.Printf("Found: %s -> %s\n", name, phone)
	} else {
		fmt.Printf("%s not found\n", name)
	}
}

