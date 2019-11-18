package main

import "fmt"

func main() {
	// define map
	// emails := make(map[string]string)
	
	// assign key value
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"

	// decalre map and add kv
	        // key type // value type
	emails := map[string]string{"Bob":"bob@gmail.com", "Sharon":"sharon@gmail.com"}


	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))
	fmt.Println(emails)

	// delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}