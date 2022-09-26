package main

import "fmt"

func main() {
	a := struct {
		message string
	}{"hello"}

	b := &a

	// mutate
	// note b.message is short for (*b).message
	b.message = "goodbye"
	fmt.Println(a.message == b.message) // prints "true"

	// reassign
	*b = struct {
		message string
	}{"galaxy"}
	fmt.Println(a.message == b.message) // prints "true"
}
