package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var s,p string
	fmt.Fscanln(reader, &s)
	fmt.Fscanln(reader, &p)
	if strings.Contains(s, p) {
    	fmt.Fprintln(writer, 1)
	} else {
    	fmt.Fprintln(writer, 0)
	}
	
}