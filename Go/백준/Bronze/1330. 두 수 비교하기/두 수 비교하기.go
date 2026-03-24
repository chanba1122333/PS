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
	fmt.Fscanln(reader, &a, &b)
	if a>b {
		fmt.Fprintln(writer, ">")
	} else if a<b {
		fmt.Fprintln(writer, "<")
	} else {
		fmt.Fprintln(writer, "==")
	}
	
}