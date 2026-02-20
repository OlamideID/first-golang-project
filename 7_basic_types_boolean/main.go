package main

import "fmt"

func main() {
	isLoggedin := true
	isAdmin := false
	hasSubscription := true

	canOpenDash := isLoggedin && hasSubscription

	canDeletePost := isAdmin || (isLoggedin && hasSubscription)

	fmt.Println(canOpenDash, canDeletePost)

	age := 2
	isAdult := age > 18
	fmt.Println(isAdult)
}
