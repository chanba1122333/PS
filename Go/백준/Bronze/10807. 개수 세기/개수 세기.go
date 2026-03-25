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

	var n,v int
	fmt.Fscan(reader, &n)
	nums := make([]int,n)
	for i := 0; i < n; i++ {
        fmt.Fscan(reader, &nums[i])
    }
	fmt.Fscan(reader, &v)
	count := 0
	for _,val := range nums{
		if val == v {
			count++
		}
	}
	fmt.Fprintln(writer, count)
}