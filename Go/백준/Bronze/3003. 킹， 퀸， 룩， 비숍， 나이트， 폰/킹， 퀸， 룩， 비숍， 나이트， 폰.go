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

	var piece int
	chess := [6]int{1, 1, 2, 2, 2, 8}
	for i := 0; i < 6; i++ {
		fmt.Fscan(reader, &piece)
		fmt.Fprint(writer, chess[i]-piece, " ")
	}

}
