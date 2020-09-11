package main

import (
	"fmt"
)

// Game point format is x-axis string, y-axis int
type sgfPointStrct01 struct {
	x string
	y int
}

// Or, if the game point format is x-axis int, y-axis int
type intPointStrct01 struct {
	x, y int
}

// SGF Format is two lowercase letters
// type sgfPointStruct02 struct {
// 	xColVal, yRowVal string
// }

func main() {
	
	fmt.Println()
	fmt.Println("Go Game Translate SGF Point99 01")
	/*Signatures:
	1。 func translateToSGFPt01(pt Point99) string { .... }
	2。 func translateToPoint02(sgfPt string) Point99 {}
	*/
	fmt.Println()
	// Tester Method
	testMethods03()
}

/*
### Point99 to SGF Method
*/
// If game point format is x-axis string, y-axis int
func transPointToSGF03(ptIN sgfPointStrct01) string {
	sgfPtY := string(rune((ptIN.y) + 97))
	// SGF format of column x string then row y string
	return ptIN.x + sgfPtY
}

// If game point format is x-axis int, y-axis int
func transPointToSGF04(ptIN intPointStrct01) string {
	sgfPtX := string(rune((ptIN.x) + 97))
	sgfPtY := string(rune((ptIN.y) + 97))
	return sgfPtX + sgfPtY
}

/*
### SGF to Point99
*/
// If game point format is x-axis string, y-axis int
func transSGFToPoint03(sgfPt string) sgfPointStrct01 {
	sgfPtX := string(sgfPt[0])
	sgfPtY := sgfPt[1]
	return sgfPointStrct01{
		x: sgfPtX,
		y: int(sgfPtY) - 97,
	}
}

// If game point format is x-axis int, y-axis int
func transSGFToPoint04(sgfPt string) intPointStrct01 {
	sgfPtX := sgfPt[0]
	sgfPtY := sgfPt[1]
	return intPointStrct01{
		x: int(sgfPtX) - 97,
		y: int(sgfPtY) - 97,
	}
}

/*
### Tester Method
*/
func testMethods03() {
	fmt.Println("Results: ")
	
	// Test if game point format is x-axis string, y-axis int
	gamePoint01 := sgfPointStrct01{
		x: "q",
		y: 23,
	}
	fmt.Println()
	fmt.Printf("trnsPntToSGF02 gamePoint01: %v, to SGF: %v\n",
		gamePoint01, transPointToSGF03(gamePoint01))
	
	// Test if game point format is x-axis int, y-axis int
	gamePoint02 := intPointStrct01{
		x: 16,
		y: 23,
	}
	fmt.Println()
	fmt.Printf("trnsPntToSGF04 gamePoint02: %v, to SGF: %v\n",
		gamePoint02, transPointToSGF04(gamePoint02))
	
	// Test if game point format is x-axis string, y-axis int
	sgfPoint03 := "qx"
	fmt.Println()
	fmt.Printf("trnsSGFToPoint03 sgfPoint03: %v, to Point99: %v\n",
		sgfPoint03, transSGFToPoint03(sgfPoint03))
	
	// Test if game point format is x-axis int, y-axis int
	fmt.Println()
	fmt.Printf("trnsSGFToPoint04 sgfPoint03: %v, to Point99: %v\n",
		sgfPoint03, transSGFToPoint04(sgfPoint03))
	
	fmt.Println()
	fmt.Println("Otrego Go! ")
	
	fmt.Println()
	fmt.Println("GitHub Go! ")
	
	fmt.Println()
	fmt.Println("IntelliJ Go! ")
	
}
