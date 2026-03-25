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

	var score int
    
	fmt.Fscanln(reader, &score)

	switch {
	case score >= 90:
    	fmt.Fprintln(writer,"A")
	case score >= 80:
    	fmt.Fprintln(writer,"B")
	case score >= 70:
    	fmt.Fprintln(writer,"C")
	case score >= 60:
		fmt.Fprintln(writer,"D")
	default:
    	fmt.Fprintln(writer,"F")
	}
	
}