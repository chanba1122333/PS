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
	con := int64(1)
	for i:=int64(1); i<=n; i++ {
		con *= i
	}
	fmt.Fprintln(writer,con)

}