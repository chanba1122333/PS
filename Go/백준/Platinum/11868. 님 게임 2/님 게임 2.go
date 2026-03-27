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

	var n, p, result int
	fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p)
		result ^= p
	}

	if result != 0 {
		fmt.Fprintln(writer, "koosaga")
	} else {
		fmt.Fprintln(writer, "cubelover")
	}

}
