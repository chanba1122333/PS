package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"unicode"
)

func sum(s string) int {
	sum := 0
	for _, r := range s {
		if unicode.IsDigit(r) {
			sum += int(r - '0')
		}
	}
	return sum
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
    fmt.Fscan(reader, &n)
	serial := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &serial[i])
	}
	sort.Slice(serial, func(i, j int) bool {
		a, b := serial[i], serial[j]
		if len(a) != len(b) {
			return len(a) < len(b)
		}
		suma, sumb := sum(a), sum(b)
		if suma != sumb {
			return suma < sumb
		}
		return a < b
	})

	for _, s := range serial {
		fmt.Println(s)
	}

}
