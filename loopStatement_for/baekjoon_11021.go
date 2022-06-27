package loopStatement_for

import "fmt"

func Bj11021()  {

	var a, b, c int

	fmt.Scan(&a)

	var sum []int
	for i := 0; i < a; i++ {
		fmt.Scanln(&b, &c)
		sum = append(sum, b+c)
	}

	for i := 0; i < a; i++ {
		fmt.Printf("Case #%d: %d\n", i+1, sum[i])
	}
}
