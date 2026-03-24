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

    var n float64
    fmt.Fscanln(reader, &n)

	result := int64(n) 

    fmt.Fprintln(writer, result)
}