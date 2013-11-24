package main

import (
	"bufio"
	"fmt"
	"os"

//	"strings"
)

func main() {
	fmt.Println("Input some random strings:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		if scanner.Text()[0] == '\n' {
			break
		}
		count++
	}
	fmt.Println(count)
}
