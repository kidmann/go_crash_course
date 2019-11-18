package main

import (
	"fmt"
	"strconv"
)

// Person struct (how to define)
type Person struct {
	// firstName string
	// lastName string
	// city string
	// gender string
	// age int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith",
		city: "Boston", gender: "f", age: 25}

	// Alternative
	person2 := Person{"Bob", "Johnson", "New York", "m", 30}

	// fmt.Println(person1)

	// Get single field
	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)

	// Pointer receivers - for methods that change something
	person1.hasBirthday()
	person1.hasBirthday()
	person1.getMarried("Williams")
	person2.getMarried("Thomson")

	// Value receivers - for methods that do calculation
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())

}
