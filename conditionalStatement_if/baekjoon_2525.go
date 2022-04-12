package conditionalStatement_if

import "fmt"

func Bj2525()  {

	var a, b, c int

	fmt.Scanln(&a, &b)
	fmt.Scan(&c)

	b = b + c

	if b >= 60 {
		a = a + b/60
		b = b % 60
		if a > 23 {
			a = a - 24

			fmt.Print(a, b)
		} else {
			fmt.Print(a, b)
		}
	} else {
		fmt.Print(a, b)
	}
}
