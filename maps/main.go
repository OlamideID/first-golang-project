package main

import "fmt"

func main() {
	//map[keyType]valueType
	ages := map[string]int{
		"Olamide": 24,
		"Gabriel": 29,
	}

	fmt.Println(ages["Olamide"])

	var scores map[string]int
	scores = make(map[string]int)
	scores["jagaban"] = 90

	fmt.Println(scores)

	users := map[string]string{
		"u1": "Raul",
		"u2": "Sesko",
		"u3": "Bruno",
	}
	fmt.Println(users["u3"])

	delete(users, "u1")
	fmt.Println(users)

}
