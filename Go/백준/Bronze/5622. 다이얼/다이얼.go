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
	fmt.Fscan(reader, &s)

	save := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 65, 66, 67:
			save += 3
		case 68, 69, 70:
			save += 4
		case 71, 72, 73:
			save += 5
		case 74, 75, 76:
			save += 6
		case 77, 78, 79:
		    save += 7
		case 80, 81, 82, 83:
			save += 8
		case 84, 85, 86:
			save += 9
		case 87, 88, 89, 90:
			save += 10
		}
	}

	fmt.Fprintln(writer, save)
}