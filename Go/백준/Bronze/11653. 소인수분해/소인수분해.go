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

	var n int
	fmt.Fscanln(reader, &n)

	for i := 2; i*i <= n; i++ {
    	for n%i == 0 {
        fmt.Println(i)
        n = n/i
    	}
	}
	if n > 1 {
    fmt.Println(n)
	}
}