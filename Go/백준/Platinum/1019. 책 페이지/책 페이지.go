package main

import (
	"bufio"
	"fmt"
	"os"
)

var page [10]int64

func addcount(x,digit int64) {
	for x > 0 {
		d := x % 10
		page[d] += digit
		x /= 10
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int64
	fmt.Fscanln(reader, &n)

	var start int64 = 1
	var end int64 = n
	var digit int64 = 1

	for start <= end {
		for start%10 != 0 && start <= end {
			addcount(start, digit)
			start++
		}

		if start > end {
			break
		}

		for end%10 != 9 && start <= end {
			addcount(end, digit)
			end--
		}

		start /= 10
		end /= 10
		count := (end - start + 1)
		for i := 0; i < 10; i++ {
			page[i] += count * digit
		}

		digit *= 10
	}

	for i := 0; i < 10; i++ {
		fmt.Fprintln(writer, page[i])
	}
	
}