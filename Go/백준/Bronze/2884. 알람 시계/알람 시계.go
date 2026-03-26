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

	var a, b int
	fmt.Fscan(reader, &a, &b)
	if b >= 45 {
		fmt.Fprintln(writer, a, b-45)
	} else if a == 0 {
		fmt.Fprintln(writer, "23", b+15)
	} else {
		fmt.Fprintln(writer, a-1, b+15)
	}

}
