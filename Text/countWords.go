package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// todo: read from a file
func main() {
	fmt.Println("Input some random strings:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	fmt.Println("Total words:", len(strings.Split(s, " ")))
}
