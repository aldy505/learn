package main

import "fmt"

func main() {
	// What is pointer?
	// A pointer holds the memory address of a value.

	var name string
	fmt.Println("Memory address of name variable:", &name)

	// Set the value of name:
	*&name = "Timmy"
	// Above is the same as name = "Timmy"
	fmt.Println("*&name =", name)

	thomas := &name
	*thomas = "Thomas"
	fmt.Println("Set name to thomas:", thomas)
	fmt.Println("Memory address for thomas", thomas)

	// Check if thomas and name have the same memory address
	if thomas == &name {
		fmt.Println("Thomas and Name have the same memory address")
	} else {
		fmt.Println("Thomas and Name don't have the same memory address")
	}

	// If they are the same, how about we call name once again?
	fmt.Println("Name is:", name)
}
