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

	var n, k int
	fmt.Fscanln(reader, &n, &k)

	removed := make([]bool, n+1) 
	count := 0

	for i := 2; i <= n; i++ {
		if removed[i] {
        continue
    	}

		for j := i; j <= n; j += i {
			if !removed[j] {
            removed[j] = true
            count++

				if count == k {
                fmt.Fprintln(writer, j)
            	}
        	}
    	}
	}
}
