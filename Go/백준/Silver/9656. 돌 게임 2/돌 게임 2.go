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

	var n int
	fmt.Fscanln(reader, &n)
	if n%2 == 0 {
		fmt.Fprintln(writer, "SK")
	} else {
		fmt.Fprintln(writer, "CY")
	}
}

