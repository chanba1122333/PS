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

	var a, b, c int
	fmt.Fscan(reader, &a, &b, &c)

	if a == b && b == c {
		fmt.Fprintln(writer, 10000+a*1000)
	} else if a == b || a == c {
		fmt.Fprintln(writer, 1000+a*100)
	} else if b == c {
		fmt.Fprintln(writer, 1000+b*100)
	} else {
		maxVal := a
		if b > maxVal {
			maxVal = b
		}
		if c > maxVal {
			maxVal = c
		}
		fmt.Fprintln(writer, maxVal*100)
	}
}
