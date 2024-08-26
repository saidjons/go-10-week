package main

import (
    "fmt"
    "sync"
)

func say(s string, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 5; i++ {
        fmt.Println(s)
    }
}

func main() {
    var wg sync.WaitGroup
    wg.Add(1) // Add a counter for the goroutine
	wg.Add(1)
    go say("world", &wg)
    go say("hello", &wg)
    
    wg.Wait() // Wait for all goroutines to finish
}
