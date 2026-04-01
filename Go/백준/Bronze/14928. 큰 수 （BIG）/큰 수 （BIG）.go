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

	var n string
	const num = 20000303
	fmt.Fscan(reader, &n)
	remainder := 0
    for i := 0; i < len(n); i++ {
        digit := int(n[i] - '0')
        remainder = (remainder*10 + digit) % num
    }
	fmt.Fprintln(writer, remainder)
	
}