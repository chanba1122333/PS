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
	var sum int
	for i:=0; i<5; i++ {
		fmt.Fscan(reader, &n)
		sum += n
	}
	fmt.Fprintln(writer, sum)
	
}