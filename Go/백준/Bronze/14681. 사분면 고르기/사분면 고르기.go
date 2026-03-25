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

	var x,y int

	fmt.Fscan(reader, &x, &y)

	switch {
	case x > 0 && y > 0:
    	fmt.Fprintln(writer,"1")
	case x < 0 && y > 0:
    	fmt.Fprintln(writer,"2")
	case x < 0 && y < 0:
    	fmt.Fprintln(writer,"3")
	case x > 0 && y < 0:
		fmt.Fprintln(writer,"4")
	}
	
}