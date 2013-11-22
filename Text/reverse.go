package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Input some random strings:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := []rune(scanner.Text())

	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}

	fmt.Println(string(s))
}
