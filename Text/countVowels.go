package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Input some random strings:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	vowels := map[rune]int{'a': 0, 'e': 0, 'i': 0, 'o': 0, 'u': 0}

	s := strings.ToLower(scanner.Text())
	for _, c := range s {
		if _, ok := vowels[c]; ok {
			vowels[c]++
		}
	}

	t := 0
	for r, c := range vowels {
		fmt.Println(string(r), "=", c)
		t += c
	}
	fmt.Println("Total =", t)
}
