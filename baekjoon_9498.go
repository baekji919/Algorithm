package main

import "fmt"

func Bj9498()  {

	var a int

	fmt.Scan(&a)
	var score string

	if a <= 100 && a >= 90 {
		score = "A"
	} else if a <= 89 && a >= 80 {
		score = "B"
	} else if a <= 79 && a >= 70 {
		score = "C"
	} else if a <= 69 && a >= 60 {
		score = "D"
	} else {
		score = "F"
	}

	fmt.Print(score)
}
