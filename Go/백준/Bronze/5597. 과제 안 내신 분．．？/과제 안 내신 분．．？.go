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

	submitted := make([]bool, 31)

    var n int
    for i := 0; i < 28; i++ {
        fmt.Fscan(reader, &n)
        submitted[n] = true
    }

    for i := 1; i <= 30; i++ {
        if !submitted[i] {
            fmt.Fprintln(writer, i)
        }
    }
}