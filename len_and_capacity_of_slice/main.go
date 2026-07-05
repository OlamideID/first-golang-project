package main

import "fmt"

func main() {
	scores := make([]int, 0, 5)
	fmt.Println(scores, len(scores), cap(scores))

	scores = append(scores, 200, 3000, 5000, 90000, 29101010, 2020129)
	fmt.Println(scores)
	fmt.Println(scores, len(scores), cap(scores))
}
