package main

import (
	"fmt"
)

func main() {

	// array
	array := [5]int{1, 2, 3, 4, 5}

	fmt.Println(array)
	// slice
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 4, 5, 6}
	fmt.Println(slice)
	slice = append(slice, 10)
	fmt.Println(slice)

	// with golang, you can add more than 1 return values
	plus, min := add(5, 10)
	fmt.Println(plus, min)

	// error handling in go
	data, err := multiply(2, -2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

}

func add(x int, y int) (int, int) {
	return x + y, x - y
}
