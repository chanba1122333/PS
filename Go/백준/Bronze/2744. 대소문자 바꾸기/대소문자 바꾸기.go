package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var s string
	fmt.Fscan(reader, &s)
	for _, r := range s {
		if unicode.IsUpper(r) {
			fmt.Fprintf(writer, "%c", unicode.ToLower(r))
		} else {
			fmt.Fprintf(writer, "%c", unicode.ToUpper(r))
		}
	}

}
