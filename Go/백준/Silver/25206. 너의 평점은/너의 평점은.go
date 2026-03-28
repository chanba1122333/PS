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

	var sj, sc string
	var n float64
	av := float64(0)
	count := float64(0)
	for i := 0; i < 20; i++ {
		fmt.Fscan(reader, &sj, &n, &sc)
		if sc != "P" {
			count += n
			switch sc {
			case "A+":
				av += n * 4.5
			case "A0":
				av += n * 4.0
			case "B+":
				av += n * 3.5
			case "B0":
				av += n * 3.0
			case "C+":
				av += n * 2.5
			case "C0":
				av += n * 2.0
			case "D+":
				av += n * 1.5
			case "D0":
				av += n * 1.0
			case "F":
				av += n * 0.0
			}
		}

	}
	if count != 0 {
		fmt.Fprintf(writer, "%.6f\n", av/count)
	} else {
		fmt.Fprintln(writer, "0.000000")
	}
}
