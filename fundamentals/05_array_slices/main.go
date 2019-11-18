package main

import "fmt"

func main() {
	// // Arrays (must be a fixed length and name your types)
	// var fruitArray [2]string

	// // Assign values
	// fruitArray[0] = "Apple"
	// fruitArray[1] = "Orange"

	// fmt.Println(fruitArray)
	// fmt.Println(fruitArray[1])

	// Declare and assign
	fruitArr := [2]string{"Apple", "Orange"}

	fmt.Println(fruitArr)

	// Slices (Don't need to declare length)
	fruitSlice := []string{"Apple", "Orange", "Pear"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitArr[1:2])

}
