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

	var h,i,a,r,c int
	fmt.Fscanln(reader, &h, &i, &a, &r, &c)

	fmt.Fprintln(writer, (h*i)-(a*r*c))
	
}