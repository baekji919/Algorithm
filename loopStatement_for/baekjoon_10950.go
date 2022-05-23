package loopStatement_for

import "fmt"

func Bj10950()  {

	var a, A, B int
	var total []int

	fmt.Scan(&a)

	for i := 0; i < a; i++ {
		fmt.Scanln(&A, &B)
		total = append(total, A+B)
	}

	for i := 0; i < a; i++ {
		fmt.Println(total[i])
	}
}
