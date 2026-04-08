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

    var n int64
    fmt.Fscan(reader, &n)

    line := int64(1)
    count := 1

    if n == 1 {
        fmt.Fprintln(writer, 1)
        return
    }

    for i := 1; ; i++ {
        line += int64(6 * i)
        count++  

        if n <= line {
            break 
        }
    }
    
    fmt.Fprintln(writer, count)
}