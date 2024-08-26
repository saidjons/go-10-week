package main

import "fmt"

func incrementScore(s *int) {
    newScore := *s + 10
    *s = newScore
}

func main() {
    score := 20
    incrementScore(&score)
    fmt.Println("The score is", score) // Prints: "The score is 20"
}