package loopStatement_for

import (
	"bufio"
	"fmt"
	"os"
)

func Bj15552()  {

	var a, b, length int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &length)

	for i := 0; i < length; i++ {
		fmt.Fscanln(reader, &a, &b)
		fmt.Fprintln(writer, a+b)
	}

	writer.Flush()
}