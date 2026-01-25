package main

import (
	"fmt"
	"strings"
)

func main() {
	firstName := "Akinola"
	lastName := "Idowu"
	fullName := firstName + " " + lastName

	fmt.Print(fullName)

	fmt.Println(strings.ToUpper(fullName))
}
