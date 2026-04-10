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

	var x, y, w, h int
	fmt.Fscan(reader, &x, &y, &w, &h)

	d1 := x
	d2 := y
	d3 := w - x
	d4 := h - y

	ans := min(min(d1, d2), min(d3, d4))

	fmt.Fprintln(writer, ans)
}
