package main

import "fmt"
func PrintAnything(value interface{}) {
    fmt.Println(value)
}

func main() {
    PrintAnything("Hello, World!")
    PrintAnything(42)
    PrintAnything(3.14)
}
