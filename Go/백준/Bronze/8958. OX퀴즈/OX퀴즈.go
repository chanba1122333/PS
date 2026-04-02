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

	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(reader, &s)

		sum := 0
		current := 0

		for j := 0; j < len(s); j++ {
			if s[j] == 'O' {
				current++
				sum += current
			} else {
				current = 0
			}
		}
		fmt.Fprintln(writer, sum)
	}
}