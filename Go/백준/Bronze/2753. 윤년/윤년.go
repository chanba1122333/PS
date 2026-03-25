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

	var year int

	fmt.Fscan(reader, &year)

	switch {
	case (year%4==0 && year%100!=0) || year%400 == 0:
    	fmt.Fprintln(writer,"1")
	default :
    	fmt.Fprintln(writer,"0")
	}
	
}