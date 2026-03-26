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

	var a,b int
	fmt.Fscan(reader, &a,&b)
	three := b%10
	fmt.Fprintln(writer,a*three)
	four := (b-three)/10
	rfour := four%10
	fmt.Fprintln(writer,a*rfour)
	five := (four-rfour)/10
	fmt.Fprintln(writer,a*five)
	fmt.Fprintln(writer,a*b)
}