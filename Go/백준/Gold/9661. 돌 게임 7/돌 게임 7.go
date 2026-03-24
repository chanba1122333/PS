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
	if n%5 == 0 || n%5 == 2 {
	fmt.Fprintln(writer, "CY")
	} else {
		fmt.Fprintln(writer, "SK")
	}
}