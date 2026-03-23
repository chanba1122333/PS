package main

import (
    "bufio"
    "fmt"
    "os"
)


func bigger(a, b int) (result string) {
    if a > b {
        result = "Yes"
    } else {
        result = "No"
    }
    return 
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()

    for {
        var a, b int
        _, err := fmt.Fscanln(reader, &a, &b)
        if err != nil {
            break
        }

        if a == 0 && b == 0 {
            break
        }

        res := bigger(a, b)
        fmt.Fprintln(writer, res)
    }
}