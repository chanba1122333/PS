package main

import (
	"bufio"
	"fmt"
	"os"
)

func win(a, b string) bool {
	return (a == "R" && b == "S") ||
		(a == "S" && b == "P") ||
		(a == "P" && b == "R")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var ml, mr, tl, tr string
	fmt.Fscan(reader, &ml, &mr, &tl, &tr)

	if tl == tr && (win(ml, tl) || win(mr, tl)) {
		fmt.Fprintln(writer, "MS")
	} else if ml == mr && (win(tl, ml) || win(tr, ml)) {
		fmt.Fprintln(writer, "TK")
	} else {
		fmt.Fprintln(writer, "?")
	}
}