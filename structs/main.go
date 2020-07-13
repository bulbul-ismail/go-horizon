package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// ismail := person{"John", "Doe"}
	john := person{firstName: "John", lastName: "Doe"}
	var doe person
	fmt.Println(john, doe)
	fmt.Printf("%+v", doe)

	doe.firstName = "john"
	doe.lastName = "doe"
	fmt.Printf("%+v", doe)
}
