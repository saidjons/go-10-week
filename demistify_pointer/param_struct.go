package main

import "fmt"

type player struct {
    name  string
    score int
}

// Make the parameter a pointer to a player struct.
func incrementScore(p *player) {
    p.score += 10
}

func main() {
    // Initialize a player struct and assign it to the variable p1.
    p1 := player{name: "Alice", score: 20}
    // Pass a pointer to p1 to incrementScore().
    incrementScore(&p1)
    fmt.Printf("The score for %s is %d", p1.name, p1.score) // Prints: "The score for Alice is 30"
}
// https://www.alexedwards.net/blog/demystifying-function-parameters-in-go