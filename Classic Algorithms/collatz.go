package main

import (
	"fmt"
)

func main() {
	var n int
	c := 0

	fmt.Println("Input a positive number: ")
	// do it faster with bufio
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		panic(err)
	}

	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
		c++
	}
	fmt.Println("Number of steps:", c)
}
