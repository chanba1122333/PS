package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	
	s, _ := reader.ReadString('\n')

	words := strings.Fields(s)

	fmt.Println(len(words))
}