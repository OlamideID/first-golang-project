package main

import "fmt"

func main() {
	items := 3
	pricePerItem := 4

	if total := items * pricePerItem; total >= 100 {
		fmt.Println("Eligible")
	} else {
		fmt.Println("ineligible")
	}

}
