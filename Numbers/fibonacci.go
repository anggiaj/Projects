package main

import (
	"fmt"
)

func main() {
	var n int
	series := []int{0, 1}

	fmt.Print("How many numbers do you need? ")
	fmt.Scanf("%d", &n)

	for i := 2; i < n; i++ {
		series = append(series, series[i-2]+series[i-1])
	}
	fmt.Println(series)
}
