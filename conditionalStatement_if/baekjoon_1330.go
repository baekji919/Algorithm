package conditionalStatement_if

import "fmt"

func Bj1330(a int, b int) {

	var math string
	if a > b {
		math = ">"
	} else if a < b {
		math = "<"
	} else {
		math = "=="
	}

	fmt.Printf(math)
}
