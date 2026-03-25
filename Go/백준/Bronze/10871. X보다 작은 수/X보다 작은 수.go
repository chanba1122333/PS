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

	var n,x,num int
	fmt.Fscan(reader, &n,&x)
	for i:=0; i<n; i++ {
		fmt.Fscan(reader, &num)
		if num<x {
			fmt.Fprint(writer,num," ")
		}
	}
}