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

    var n, m int
    fmt.Fscan(reader, &n, &m)
    matrixA := make([][]int, n)
    for i := 0; i < n; i++ {
        matrixA[i] = make([]int, m)
        for j := 0; j < m; j++ {
            fmt.Fscan(reader, &matrixA[i][j])
        }
    }

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            var valB int
            fmt.Fscan(reader, &valB)
            
            fmt.Fprintf(writer, "%d ", matrixA[i][j] + valB)
        }
        fmt.Fprintln(writer) 
    }
}