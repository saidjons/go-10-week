package main

import "fmt"

func main() {
    // Create a dummy array of names
    names := [5]string{"Alice", "Bob", "Charlie", "David", "Eve"}

    // Create a slice to hold the copied elements without restrictions on capacity
    // copiedNames := append([]string{}, names[:]...)
    // copiedNames := [...]string{"Penn", "Teller"}
    copiedNames := names[1:]

    // Print both the original array and the slice
    fmt.Println("Original names:", names)
    fmt.Println("Copied names:", copiedNames)
}
