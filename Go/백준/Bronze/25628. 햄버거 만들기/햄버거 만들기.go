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

	var a,b int
	fmt.Fscan(reader, &a,&b)
	a = a/2
	count:=0
	for a>=1 && b>=1 {
		a=a-1
		b=b-1
		count++
	} 
	fmt.Fprintln(writer, count)
}