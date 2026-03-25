package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
   	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

    count := 0
    
    for scanner.Scan() {
        count++
    }

    fmt.Fprintln(writer, count)
}