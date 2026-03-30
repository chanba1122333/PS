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

    var s string
    fmt.Fscan(reader, &s)
    abt := make([]int, 26)

    for i := 0; i < len(s); i++ {
        curr := s[i]
        
        if curr >= 'a' && curr <= 'z' {
            abt[curr-'a']++
        } else {
            abt[curr-'A']++
        }
    }

    maxCount := -1
    var resultIndex int
    isDuplicate := false

    for i := 0; i < 26; i++ {
        if abt[i] > maxCount {
            maxCount = abt[i]
            resultIndex = i
            isDuplicate = false
        } else if abt[i] == maxCount && maxCount != 0 {
            isDuplicate = true
        }
    }

    if isDuplicate {
        fmt.Fprintln(writer, "?")
    } else {
        fmt.Fprintf(writer, "%c\n", resultIndex+'A')
    }
}