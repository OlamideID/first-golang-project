package main

import "fmt"

func main() {
	var marks [3]int

	marks[0] = 10
	marks[1] = 20
	marks[2] = 30

	fmt.Println(marks)

	res := [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(res))
}
