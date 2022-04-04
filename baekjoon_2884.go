package main

import "fmt"

func Bj2884(a int, b int)  {

	b = b - 45

	if b < 0 {
		if a == 0 {
			a = 24 -1
			b = 60 + b
			fmt.Print(a, b)
		} else {
			a = a - 1
			b = 60 + b
			fmt.Print(a, b)
		}
	} else {
		fmt.Print(a, b)
	}
}
