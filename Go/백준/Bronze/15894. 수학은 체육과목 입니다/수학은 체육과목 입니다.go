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

	var x int
	fmt.Fscan(reader, &x)
	fmt.Fprintln(writer, 4*x)
}
