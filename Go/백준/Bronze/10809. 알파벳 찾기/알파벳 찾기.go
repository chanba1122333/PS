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

	var s string

	abt := make([]int, 26)

	for i := 0; i < 26; i++ {
		abt[i] = -1
	}

	fmt.Fscan(reader, &s)

	for i := 0; i < len(s); i++ {

		idx := s[i] - 'a'

		if abt[idx] == -1 {
			abt[idx] = i
		}
	}

	for i := 0; i < 26; i++ {
		fmt.Fprint(writer, abt[i], " ")
	}

	fmt.Fprintln(writer)

}
