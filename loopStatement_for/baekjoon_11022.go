package loopStatement_for

import (
	"bufio"
	"fmt"
	"os"
)

func Bj11022()  {

	var a, b, length int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &length)

	for i := 0; i < length; i++ {
		fmt.Fscanln(reader, &a, &b)
		fmt.Fprintf(writer, "Case #%d: %d + %d = %d\n", i+1, a, b, a+b)
	}

	writer.Flush()
}
