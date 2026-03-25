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

	var s string
	var t int

	fmt.Fscan(reader, &s)
	fmt.Fscan(reader, &t)
	fmt.Fprintf(writer, "%c\n", s[t-1])
}
