package loopStatement_for

import "fmt"

func Bj2438()  {

	var a int

	fmt.Scan(&a)

	for i := 0; i < a; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Printf("\n")
	}
}
