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

	minx, maxx := 10001, -10001
	miny, maxy := 10001, -10001

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)

		if a < minx {
			minx = a
		}
		if a > maxx {
			maxx = a
		}

		if b < miny {
			miny = b
		}
		if b > maxy {
			maxy = b
		}
	}

	w := maxx - minx
	h := maxy - miny

	fmt.Fprintln(writer, w*h)
}