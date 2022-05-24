package loopStatement_for

import (
	"bufio"
	"fmt"
	"os"
)

func Bj2742()  {

	var a int

	writer := bufio.NewWriter(os.Stdout)

	fmt.Scan(&a)

	for i := a; i >= 1; i-- {
		fmt.Fprintln(writer, i)
	}

	writer.Flush()
}
