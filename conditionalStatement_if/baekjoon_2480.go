package conditionalStatement_if

import (
	"fmt"
)

func Bj2480()  {

	var a, b, c, total int

	fmt.Scan(&a, &b, &c)

	if a == b && b == c {
		total = 10000 + a * 1000
	} else if a == b || b == c || c == a {
		if a == b {
			total = 1000 + a * 100
		} else if b == c {
			total = 1000 + b * 100
		} else {
			total = 1000 + c * 100
		}
	} else {
		arr := [3]int{a, b, c}
		max := arr[0]
		for i := range arr {
			if max < arr [i] {
				max = arr[i]
			}
		}
		total = max * 100
	}

	fmt.Print(total)
}
