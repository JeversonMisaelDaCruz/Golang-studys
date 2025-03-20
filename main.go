package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	var person Person
	fmt.Println("Insert your name:")
	fmt.Scanln(&person.Name)
	fmt.Println("Insert your age:")
	fmt.Scanln(&person.Age)
	fmt.Println("Your name is", person.Name, "and your age is", person.Age)
	if person.Age >= 18 {
		fmt.Println("You can drive")
	} else {
		fmt.Println("You can't drive")
	}
}
