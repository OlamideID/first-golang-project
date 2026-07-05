package main

import "fmt"

func main() {
	score := 10
	fmt.Println("before", score)

	addScore(&score)
	fmt.Println("after", score)
}

func addScore(score *int) {
	*score = *score + 5
}
