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
	found := false
    for i := 1; i <= 9; i++ {
        if n%i == 0 && (n/i >= 1 && n/i <= 9) {
            found = true
            break
        }
    }
    if found {
        fmt.Fprintln(writer, "1")
    } else {
        fmt.Fprintln(writer, "0")
    }
}