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

    var a, b, v int
    fmt.Fscan(reader, &a, &b, &v)
    
    day := (v-b-1)/(a-b) + 1

    fmt.Fprintln(writer, day)
}