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
	var memo [16]int
	
	memo[0] = 2
	fmt.Fscan(reader, &n)

	for i := 1; i <= n; i++ {
    	memo[i] = memo[i-1] + (memo[i-1] - 1)
	}
	result := memo[n] * memo[n]
	
	fmt.Fprint(writer, result)

}
