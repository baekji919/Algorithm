package loopStatement_for

import "fmt"

func Bj8393()  {

	var a int

	fmt.Scan(&a)

	sum := 0
	for i := 1; i <= a; i++ {
		sum = sum + i
	}
	fmt.Print(sum)
}
