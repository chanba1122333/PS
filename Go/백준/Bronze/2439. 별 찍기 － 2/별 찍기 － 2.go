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
	fmt.Fscan(reader, &n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Fprint(writer, " ")
		}
		for k := 1; k <= i; k++ {
			fmt.Fprint(writer, "*")
		}
		fmt.Fprintln(writer)
	}
}
