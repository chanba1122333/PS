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

   	for i := 0; i < len(s); i++ {
		fmt.Fprint(writer, string(s[i]))
		if ( (i+1) % 10 == 0 ) {
			fmt.Fprint(writer,"\n")
		}
	}
}