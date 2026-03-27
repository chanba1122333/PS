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
	max := n
	index := 1

	for i := 2; i <= 9; i++ {
		fmt.Fscan(reader, &n)
		if n > max {
			max = n
			index = i
		}
	}
	fmt.Fprintln(writer, max)
	fmt.Fprintln(writer, index)

}
