package main

import (
	"bufio"
	"fmt"
	"os"
)

func odd(a, b int64) int64 {
	return ((a + b) * (a - b))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, b int64
	fmt.Fscan(reader, &a, &b)
	fmt.Fprintln(writer, odd(a, b))

}
