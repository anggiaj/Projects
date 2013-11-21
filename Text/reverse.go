package main

import (
	"fmt"
)

/*
func reverse(s string) string {
	b := []rune(s)
	for i := 0; i < len(b)/2; i++ {
		j := len(b) - i - 1
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func main() {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", reverse(input))


}
*/
func main() {
	var in string
	_, err := fmt.Scanln(&in)
	if err != nil {
		fmt.Println(err)
	}
	for i := len(in) - 1; i >= 0; i-- {
		fmt.Print(string(in[i]))
	}
}
