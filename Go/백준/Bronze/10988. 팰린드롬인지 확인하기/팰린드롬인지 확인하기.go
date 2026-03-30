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

	result := 1
	
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i]{
			result = 0
			break
		}
	}
	fmt.Fprintln(writer,result)
}