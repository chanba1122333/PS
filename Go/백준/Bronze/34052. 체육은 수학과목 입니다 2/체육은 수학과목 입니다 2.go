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

	var a,b,c,d int
	fmt.Fscan(reader, &a, &b, &c, &d)
	if (a+b+c+d) <= 1500 {
		fmt.Fprintln(writer, "Yes")
	}else {
		fmt.Fprintln(writer, "No")
	}
	
}