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
	res:=a*b*c
    var counts [10]int

    for res > 0 {
        counts[res % 10]++
        res /= 10
    }

    for i := 0; i < 10; i++ {
        fmt.Fprintln(writer, counts[i])
    }

}