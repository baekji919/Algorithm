package loopStatement_for

import "fmt"

func Bj2439()  {

	var a int

	fmt.Scan(&a)

	for i := 0; i < a; i++ {
		for j := a-1; j > i; j-- {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
