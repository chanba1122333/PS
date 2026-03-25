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

	var c,n int
	fmt.Fscan(reader, &c)
	for i:=0; i<c; i++ {
		cnt := 0
		fmt.Fscan(reader,&n)
		for j:=1; j<=n; j++ {
			if n%j == 0 {
				cnt++
			}
		}
		fmt.Fprintln(writer,n,cnt)
	}
	
}