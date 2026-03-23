package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	
    fmt.Fprintln(writer, "문제의 정답")
	
}