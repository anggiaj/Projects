package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var h string
	fmt.Print("Enter a host address: ")
	fmt.Scanf("%s", &h)
	out, err := exec.Command("Whois", h).Output()
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(out))
}
