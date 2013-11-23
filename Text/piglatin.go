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
	s := strings.ToLower(scanner.Text())
	vowels := "aeiuo"
	pig := "ay"

	if strings.IndexAny(s, vowels) == 0 {
		fmt.Printf("%s%s\n", s, pig)
	} else {
		fmt.Printf("%s%s%s\n", s[1:], s[0:1], pig)
	}
}
