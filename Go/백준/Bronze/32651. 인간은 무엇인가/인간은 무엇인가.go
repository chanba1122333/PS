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
	if n<=100000 && n%2024 == 0 {
		fmt.Fprintln(writer, "Yes")
	}else{
		fmt.Fprintln(writer, "No")
	}
	
}