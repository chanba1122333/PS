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

	var a, b, c int64
	fmt.Fscan(reader, &a, &b, &c)
	result := a + b + c
	fmt.Fprintln(writer, result)
	
}