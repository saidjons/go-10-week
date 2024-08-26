package main

import "fmt"

func main() {
    // Create a dummy array of names
    names := [5]string{"Alice", "Bob", "Charlie", "David", "Eve"}

    // Create another array to copy into
    var copiedNames [5]string

    // Copy the names array into copiedNames
    copiedNames = names

    // Print both arrays
    fmt.Println("Original names:", names)
    fmt.Println("Copied names:", copiedNames)
}
