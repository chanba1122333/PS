package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var T int
	fmt.Fscan(reader, &T)

	for i := 0; i < T; i++ {
		var n int
		fmt.Fscan(reader, &n)

		var maxS string
		maxL := math.MinInt32

		for j := 0; j < n; j++ {
			var s string
			var l int
			fmt.Fscan(reader, &s, &l)

			if l > maxL {
				maxL = l
				maxS = s
			}

		}
		
		fmt.Fprintln(writer, maxS)

	}

}
