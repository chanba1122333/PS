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

	var t, k, n int
	var dp [15][15]int
	for j := 1; j < 15; j++ {
		dp[0][j] = j
	}
	for i := 1; i < 15; i++ {
		for j := 1; j < 15; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}

	fmt.Fscan(reader, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(reader, &k, &n)
		fmt.Fprintln(writer, dp[k][n])
	}

}
