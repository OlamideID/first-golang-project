package main

import "fmt"

func main() {
	points := map[string]int{

		"a": 120,
		"b": 0,
	}
	fmt.Println("a", points["a"])
	fmt.Println("b", points["b"])
	fmt.Println("c", points["c"])

	valB, okB := points["b"]
	fmt.Println(valB, okB)

	valC, okC := points["c"]
	fmt.Println(valC, okC)

	if val, ok := points["b"]; ok {
		fmt.Println(val, "Sabi boy Siuuu")
	} else {
		fmt.Println("Hello")
	}

	prices := map[string]int{
		"abc": 500,
		"def": 300,
	}

	total := 0
	for item, price := range prices {
		fmt.Println(item, price)
		total = total + price
	}

	fmt.Println(total)

	for item := range prices {
		fmt.Println(item)
	}
}
