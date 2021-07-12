package main

import "fmt"

func main() {
	array := make([]int, 5)
	fmt.Println(array)
	fmt.Println(cap(array))
	fmt.Println(len(array))

	complex := make([]complex64, 1)
	complex[0] = 197812.1232
	fmt.Println(complex[0])
	fmt.Println(imag(complex[0]))
	fmt.Println(real(complex[0]))

	array = append(array, 23, 34, 35)
	fmt.Println(array)
	var another []int
	numberCopied := copy(another, array)
	fmt.Println(numberCopied)
	fmt.Println(another)
}
