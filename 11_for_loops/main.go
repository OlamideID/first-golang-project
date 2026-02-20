// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	items := []string{"Apple", "Banana", "Cherry", "Plantain", "coke", "Sabs"}

// 	for i := 0; i <= 5; i++ {
// 		if i < len(items) {
// 			fmt.Println("Index:", i, "Value:", strings.ToUpper(items[i]))
// 		} else {
// 			fmt.Println("Index:", i, "Out of range")
// 		}
// 	}
// 	pie()
// }

// func pie() {
// 	items := []string{"Apple", "Banana", "Cherry", "Plantain", "coke", "Sabs"}

// 	for i, value := range items {
// 		if i < 3 {
// 			fmt.Println("Index:", i, "Value:", value)
// 		} else {
// 			fmt.Println("Index:", i, "Value:", strings.ToUpper(value))
// 		}
// 	}
// }

package main

import "fmt"

func main() {
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	N := 10
	sum := 10

	for i := 1; i <= N; i++ {
		sum = sum + i
	}

	fmt.Println(sum)
}
