package main

import (
	"bufio"
	"fmt"
	"os"
)

func square(a int)(int) {
	return a * a
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a,b,c,d,e int
	fmt.Fscanln(reader, &a, &b, &c, &d, &e)
	res := square(a)+square(b)+square(c)+square(d)+square(e)
	fmt.Fprintln(writer, res%10)
	
}