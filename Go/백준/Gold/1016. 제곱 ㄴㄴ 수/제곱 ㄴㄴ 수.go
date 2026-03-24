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

	var min, max int64
	fmt.Fscan(reader, &min, &max)

	rangeSize := max - min + 1
	sqrssnum := make([]bool, rangeSize)

	for i := int64(2); i*i <= max; i++ {
		square := i * i

		startMult := (min + square - 1) / square
		
		for num := startMult * square; num <= max; num += square {
			index := num - min
			if index >= 0 && index < rangeSize {
				sqrssnum[index] = true
			}
		}
	}

	count := 0

	for i := int64(0); i < rangeSize; i++ {
		if !sqrssnum[i] {
			count++
		}
	}

	fmt.Fprintln(writer, count)
}