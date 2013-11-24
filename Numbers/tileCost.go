package main

import (
	"fmt"
)

func main() {
	var w, h, c float64

	fmt.Println("What's the cost per sq. feet? ")
	fmt.Scanf("%G", &c)
	fmt.Println("What's the width of the floor? ")
	fmt.Scanf("%G", &w)
	fmt.Println("What's the height of the floor? ")
	fmt.Scanf("%G", &h)

	fmt.Printf("The total cost is $%G for %G square feet\n", w*h*c, w*h)

}
