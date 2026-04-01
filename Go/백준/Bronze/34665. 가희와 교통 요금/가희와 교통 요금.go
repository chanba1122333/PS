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

	var a,b string
	fmt.Fscan(reader, &a, &b)
	if a == b {
		fmt.Fprintln(writer, "0")
	}else {
		fmt.Fprintln(writer, "1550")
	}
}