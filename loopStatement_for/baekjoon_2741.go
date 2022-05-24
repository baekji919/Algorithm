package loopStatement_for

import (
	"bufio"
	"fmt"
	"os"
)

func Bj2741()  {

	var a int

	writer := bufio.NewWriter(os.Stdout)

	fmt.Scan(&a)

	for i := 1; i <= a; i++ {
		fmt.Fprintln(writer, i)
	}

	writer.Flush()
}
