package main

import "fmt"

func main() {
	// Create a dummy array of names
	names := [5]string{"Alice", "Bob", "Charlie", "David", "Eve"}

	// Create another array with only 3 elements to copy into
	var copiedNames [3]string

	// Use the copy function to copy the first 3 elements
	copy(copiedNames[:], names[1:])

	// Print both arrays
	fmt.Println("Original names:", names)
	fmt.Println("Copied names:", copiedNames)
}
