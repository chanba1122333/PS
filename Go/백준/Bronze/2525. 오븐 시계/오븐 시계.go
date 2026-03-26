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

	var a, b, c int
	fmt.Fscan(reader, &a, &b)
	fmt.Fscan(reader, &c)
	ah := (b + c) / 60
	am := (b + c) % 60
	if b+c >= 60 && a+ah < 24 {
		fmt.Fprintln(writer, a+ah, am)
	} else if a+ah >= 24 && b+c >= 60 {
		fmt.Fprintln(writer, a+ah-24, am)
	} else {
		fmt.Fprintln(writer, a, b+c)
	}

}
