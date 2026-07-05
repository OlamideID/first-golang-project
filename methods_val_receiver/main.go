package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{Name: "Olamide", Age: 24}
	fmt.Println(u.Intro())
}

func (u User) Intro() string {
	return fmt.Sprint("Hello, I am ", u.Name)
}
