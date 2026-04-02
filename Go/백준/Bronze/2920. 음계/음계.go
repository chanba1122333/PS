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
	
	if scanner.Scan() {
    s := scanner.Text()
    if s == "1 2 3 4 5 6 7 8" {
        fmt.Fprintln(writer, "ascending")
    } else if s == "8 7 6 5 4 3 2 1" {
        fmt.Fprintln(writer, "descending")
    } else {
        fmt.Fprintln(writer, "mixed")
        }
    }
}