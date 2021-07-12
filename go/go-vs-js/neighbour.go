package main

import "errors"

func multiply(x int, y int) (int, error) {
	if y <= 0 {
		return 0, errors.New("can't multiply by zero")
	}
	return x * y, nil
}
