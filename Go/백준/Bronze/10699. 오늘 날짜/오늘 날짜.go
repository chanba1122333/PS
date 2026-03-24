package main

import (
    "bufio"
    "fmt"
    "os"
    "time"
)

func main() {
    //reader := bufio.NewReader(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()
    now := time.Now()
    result := now.Format("2006-01-02")
    fmt.Fprintln(writer, result)
}