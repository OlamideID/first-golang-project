package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// go doesn't use exceptions for normal failures
	//funcs return errors as non return values
	//val, err := something()
	//if err != nil {handle error}

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	input := "30"
	level, err := parseLevel(input)
	if err != nil {
		return err
	}

	fmt.Println("selected level is", level)
	return nil
}

func parseLevel(s string) (int, error) {
	// (value, err)
	// nil error -> success
	// not nil -> failure

	//
	n, error := strconv.Atoi(s)
	if error != nil {
		return 0, fmt.Errorf("Label must be a number")
	}

	if n < 1 || n > 5 {
		return 0, fmt.Errorf("must be a number")
	}

	return n, nil
}
