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

	var n,m int64
	fmt.Fscan(reader, &n,&m)
	ans := n - m
	if ans < 0 {
		ans = -ans
	}
	fmt.Fprintln(writer,ans)
	
}