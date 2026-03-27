package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t, a, b int
	fmt.Fscan(reader, &t)

	for i := 1; i <= t; i++ {
		fmt.Fscan(reader, &a, &b)
		fmt.Fprintf(writer, "Case #%d: %d\n", i, a+b)
	}
}
