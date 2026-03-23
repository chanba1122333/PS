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
    fmt.Fscanln(reader, &n)

    var res int64 = 1

    for i := 1; i <= n; i++ {
        res *= int64(i)
    }

    fmt.Fprintln(writer, res)
}