package main

import "fmt"

func main() {
	results := []string{"Akinola", "Idowu"}

	fmt.Println(results, results[0], results[len(results)-1])

	results[1] = "Olamide"
	fmt.Println(results)

	var nums []int
	nums = append(nums, 10)
	nums = append(nums, 20, 30)
	fmt.Println(nums)
}
