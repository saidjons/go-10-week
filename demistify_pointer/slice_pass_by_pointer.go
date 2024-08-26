package main

import "fmt"

func addScores(s *[]int, values ...int) {
    *s = append(*s, values...)
}

func main() {
    scores := []int{10, 20, 30}
    addScores(&scores, 40, 50, 60)
    fmt.Println(scores) // Prints: [10 20 30 40 50 60]
}
