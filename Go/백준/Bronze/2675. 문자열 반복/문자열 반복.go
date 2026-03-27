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

	var r, t int
	var s string
	fmt.Fscan(reader, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(reader, &r, &s)

		for j := 0; j < len(s); j++ {
			for k := 0; k < r; k++ {
				fmt.Fprint(writer, string(s[j]))
			}
		}
		fmt.Fprintln(writer)
	}

}
