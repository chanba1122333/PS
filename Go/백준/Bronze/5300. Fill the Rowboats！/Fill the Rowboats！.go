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
	for i := 1; i <= n; i++ {
		fmt.Fprint(writer, i , " ")
		if i % 6 == 0 || i == n {
			fmt.Fprint(writer, "Go!" , " ")
		}
	}
	
}