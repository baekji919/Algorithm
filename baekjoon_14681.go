package main

import "fmt"

func Bj14681(a float64, b float64) {

	var result int

	if a > 0 {
		if b > 0 {
			result = 1
		} else {
			result = 4
		}
	} else {
		if b > 0 {
			result = 2
		} else {
			result = 3
		}
	}

	fmt.Print(result)
}
