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
    maxVal := -1 
    maxRow, maxCol := 0, 0

    for i := 1; i <= 9; i++ {
        for j := 1; j <= 9; j++ {
            fmt.Fscan(reader, &n)

            if n > maxVal {
                maxVal = n
                maxRow = i
                maxCol = j
            }
        }
    }

    fmt.Fprintln(writer, maxVal)
    fmt.Fprintln(writer, maxRow, maxCol)
}