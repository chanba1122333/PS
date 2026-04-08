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

    for i := 0; i < n; i++ {
        var m int
        fmt.Fscan(reader, &m)

        countQ := m / 25
        m %= 25 

        countD := m / 10
        m %= 10

        countN := m / 5
        m %= 5

        countP := m / 1
        m %= 1

        fmt.Fprintln(writer, countQ, countD, countN, countP)
    }
}