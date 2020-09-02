package main

import (
	"fmt"
	"testing"
)

// func main() {
//
// }

func TestSum(t *testing.T) {
	xVal := 5
	yVal := 5
	crrctRslt := xVal + yVal
	total := Sum(xVal, yVal)
	// Test method
	if total != crrctRslt {
		fmt.Println()
		t.Errorf("Sum was incorrect, \n"+
			"got: %d \n"+
			"want: %d \n", total, crrctRslt)
	}
}
