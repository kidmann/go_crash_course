package main

import "fmt"

func main() {
	a := 5
	b := &a // assigning b to a pointer of a
	fmt.Println(a, b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b) // * represents a pointer 

	// Use * to read val from address
	fmt.Println(*b) 

	// Change val with pointer
	*b = 10
	fmt.Println(a)
	
	// everything in Go is passed by value 
}