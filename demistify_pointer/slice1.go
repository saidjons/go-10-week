package main

import "fmt"

func addBonus(s []int) {
    cs := append(s, s...)

    // cs := append([]int(nil), s...)
    for i := range cs {
        cs[i] += 50
    }
}
func addScores(s *[]int, values ...int) {
    *s = append(*s, values...)
}


func main() {
    scores := []int{10, 20, 30}
    addBonus(scores)
    fmt.Println(scores) // Prints: [60 70 80]
}