package main

import "fmt"

func divide(a int, b int) (jay int, lee int) {
	jay = a / b
	lee = a + b

	return
}

func main() {
	q, r := divide(10, 2)

	fmt.Println(q, r)
}
