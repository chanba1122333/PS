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

	remainders := make(map[int]bool)

	for i := 0; i < 10; i++ {
    var num int
    fmt.Fscan(reader, &num)
    r := num % 42
    remainders[r] = true
	}
	fmt.Fprintln(writer, len(remainders))
}