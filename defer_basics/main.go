package main

import (
	"errors"
	"fmt"
)

func main() {
	// defer resp.body.close()

	// fmt.Println("Case 1: success")
	// if err := doWork(true); err != nil {
	// 	fmt.Println("error:", err)
	// }

	fmt.Println("Case 1: failure")
	if err := doWork(false); err != nil {
		fmt.Println("error:", err)
	}
}

func doWork(success bool) error {
	// resource related
	// start message -> resource acquired
	// cleanup message -> resource released

	fmt.Println("start : resource acquired")
	// defer will guaerantee this runs at the end of the function
	// return success
	// return error
	defer fmt.Println("cleanup : resource released")
	if !success {
		return errors.New("wahala")
	}

	fmt.Println("something is happening")
	fmt.Println("done")

	return nil
}
