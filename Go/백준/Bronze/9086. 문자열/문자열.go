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

	var t int
	var s string
	fmt.Fscan(reader, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(reader, &s)
		fmt.Fprintf(writer, "%c%c\n", s[0], s[len(s)-1])
	}

}
