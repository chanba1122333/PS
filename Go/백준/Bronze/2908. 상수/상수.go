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
	fia := a / 100
	sea := (a - (fia * 100)) / 10
	tha := a % 10
	reala := 100*tha + 10*sea + fia

	fib := b / 100
	seb := (b - (fib * 100)) / 10
	thb := b % 10
	realb := 100*thb + 10*seb + fib

	if reala > realb {
		fmt.Fprintln(writer, reala)
	} else {
		fmt.Fprintln(writer, realb)
	}

}
