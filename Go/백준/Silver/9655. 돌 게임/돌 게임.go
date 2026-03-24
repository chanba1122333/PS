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
	if n%2 == 0 {
		fmt.Fprintln(writer, "CY")
	} else {
		fmt.Fprintln(writer, "SK")
	}
}

/*
1 상
2 dp1 창
3 상
4 dp1 dp3 창
5 dp2 dp4 상
6 dp5 dp3 창
7 dp4 dp6 상
8 dp5 dp7 창
9 dp6 dp8 상
10 dp7 dp9 창

홀수 상 짝수 창 */