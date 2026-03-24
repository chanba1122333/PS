package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	cat := `\    /\
 )  ( ')
(  /  )
 \(__)|`

	fmt.Fprintln(writer, cat)
}