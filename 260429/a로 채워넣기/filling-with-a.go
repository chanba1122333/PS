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

	var s string
	fmt.Fscan(reader, &s)

	arr := []byte(s)

	arr[1] = 'a'
	arr[len(arr)-2] = 'a'

	fmt.Fprintln(writer, string(arr))
}