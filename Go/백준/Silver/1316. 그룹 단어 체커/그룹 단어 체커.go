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
	var s string
	count := 0

	fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {

		fmt.Fscan(reader, &s)

		check := make([]bool, 26)
		var prev byte
		included := true

		for j := 0; j < len(s); j++ {
			curr := s[j]

			if curr != prev {
				index := curr - 'a'

				if check[index] {
					included = false
					break
				}
				check[index] = true
				prev = curr
			}
		}

		if included == true {
			count++
		}
	}
	fmt.Fprintln(writer, count)
}
