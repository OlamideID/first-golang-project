package main

import "fmt"

// struct groups related fields into one type

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func main() {
	u1 := User{ID: 1, Name: "Olamide", Email: "iakinola926@gmail.com", Age: 100}
	fmt.Println(u1)
	fmt.Println(u1.ID, u1.Email, u1.Name, u1.Age)

	u1.Age = 200
	fmt.Println(u1)

	u2 := User{
		Name:  "Akinola",
		Email: "Idowuakinola27@gmail.com",
	}

	fmt.Println("partial", u2)

}
