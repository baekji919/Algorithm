package loopStatement_for

import "fmt"

func Bj2739()  {

	var a int

	fmt.Scan(&a)

	for i := 1; i < 10; i++ {
		fmt.Println(a, "*", i, "=", a*i)
	}
}
