package conditionalStatement_if

import "fmt"

func Bj2753()  {

	var a int

	fmt.Scan(&a)

	var result int

	if a % 4 == 0 {
		if a % 100 != 0 || a % 400 == 0{
			result = 1
		} else {
			result = 0
		}
	} else {
		result = 0
	}

	fmt.Print(result)
}
