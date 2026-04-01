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

	for {
		var s string
		fmt.Fscan(reader, &s)
		if s == "0" {
			break
		}

		num := 0
		for i := 0; i < len(s); i++ {
			new := int(s[i] - '0')
			fac := 1 
			for j := 1; j <= len(s)-i; j++ {
				fac *= j
			}
			num += new * fac 
		}
		fmt.Fprintln(writer, num)
	}
}