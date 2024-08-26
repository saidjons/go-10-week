package main

import "fmt"

func IdentifyType(value interface{}) {
    switch v := value.(type) {
    case string:
        fmt.Println("String:", v)
    case int:
        fmt.Println("Integer:", v)
    case float64:
        fmt.Println("Float64:", v)
    default:
        fmt.Println("Unknown type")
    }
}

func main() {
    IdentifyType("Hello")
    IdentifyType(123)
    IdentifyType(3.14)
}
