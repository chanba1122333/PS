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

	var n, sum int
	var s string

	fmt.Fscan(reader, &n)
	fmt.Fscan(reader, &s)
	for i := 0; i < n; i++ {
		sum += int(s[i] - '0')
	}

	fmt.Fprintln(writer, sum)

}
