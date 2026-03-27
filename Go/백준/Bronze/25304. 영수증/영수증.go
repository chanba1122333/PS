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

	var x, n, a, b int
	fmt.Fscan(reader, &x, &n)
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a, &b)
		mul := a * b
		sum += mul
	}
	if x == sum {
		fmt.Fprintln(writer, "Yes")
	} else {
		fmt.Fprintln(writer, "No")
	}
}
