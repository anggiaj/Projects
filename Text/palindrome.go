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

}
