package main

import (
	"fmt"
)

func main() {
	xVal := 5
	yVal := 5
	fmt.Println()
	fmt.Printf("testMain: Sum(%v, %v) = %v \n", xVal, yVal,
		Sum(xVal, yVal))
	Sum(xVal, yVal)
}

func Sum(x int, y int) int {
	return x + y
}
