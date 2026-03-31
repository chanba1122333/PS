package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, b, c int
	fmt.Fscan(reader, &a, &b, &c)
	fmt.Fprintln(writer, a+b-c)
	newa := strconv.Itoa(a)
	newb := strconv.Itoa(b)
	newnum := newa + newb
	num, _ := strconv.Atoi(newnum)
	fmt.Fprintln(writer, num-c)

}
