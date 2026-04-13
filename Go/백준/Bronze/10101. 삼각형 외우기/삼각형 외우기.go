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

	var a,b,c int
	fmt.Fscan(reader, &a,&b,&c)
	if a + b + c != 180{
		fmt.Fprintln(writer, "Error")
	}else {
		if a == 60 && a==b && b==c{
			fmt.Fprintln(writer, "Equilateral")
		}else if a == b && b!=c || a==c && a!=b || b==c && a!=b{
			fmt.Fprintln(writer, "Isosceles")
		}else {
			fmt.Fprintln(writer, "Scalene")
		}
	}
}