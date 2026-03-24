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
    var s string

    fmt.Fscan(reader, &n)

    for i := 0; i < n; i++ {
        fmt.Fscan(reader, &s)

        if len(s) >= 6 && len(s) <= 9 {
            fmt.Fprintln(writer, "yes")
        } else {
            fmt.Fprintln(writer, "no")
        }
    }
}