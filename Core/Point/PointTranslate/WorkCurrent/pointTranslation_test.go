package main

import (
	"fmt"
	"testing"
)

// Sample correct results
// New(8, 20), SGF string "iu"
// New(13, 5), nf

func TestTranslate(t *testing.T) {

	fmt.Println("[*** Go Game Point Translation Tests ***]")

	// *** Test 'Fields' *** //
	var x1 int64 = 8          // User chosen x values for testing
	var y1 int64 = 20         // User chosen y values for testing
	const aValue = int64('a') // Ascii int string conversion value

	// ** Create test pointer-type *Point with point.go New() method
	Pnt01 := New(x1, y1)
	// Create non-pointer-type copy of Pnt01,
	// and a test reference copy, for this unit test
	type PointStrct struct { // non-pointer-type struct
		x int64
		y int64
	}
	PntToTest01 := PointStrct{ // Non-pointer Pnt01 copy to test
		x: Pnt01.x,
		y: Pnt01.y,
	}
	PntRefrnce01 := PointStrct{ // Non-pointer Pnt01 reference copy
		x: Pnt01.x,
		y: Pnt01.y,
	}

	// ** Create test SGF Point with point.go ToSGF() method
	SGF01 := Pnt01.ToSGF()
	// Create a test reference copy of SGF01, for this unit test
	var SGFCrct01 string // string variable to hold SGF values
	SGFCrct01 = string(rune((PntToTest01.x)+aValue)) + string(rune((PntToTest01.y)+aValue))
	// Interject different user values with below, and comment out above
	// SGFCrct01 = "ju"

	// ** Create test *Point from SGF01, which should equal Pnt01
	Pnt02 := NewFromSGF(SGF01)
	// Create non-pointer-type copy of Pnt01, for this unit test.
	// Do not need reference copy as Pnt02 should equal Pnt01
	PntTest02 := PointStrct{ // Non-pointer Pnt02 copy to test
		x: Pnt02.x,
		y: Pnt02.y,
	}
	// Interject different user values with below, and comment out above
	// Pnt02 := PointStrct{
	// 	x: 1,
	// 	y: 24,
	// }

	// *** Test Statements *** //

	// *** To Test Pnt01 creation, using non-pointer copies
	if PntToTest01 != PntRefrnce01 {
		fmt.Println()
		t.Errorf("Point creation fail.\n"+
			"got: %v \n"+
			"want: %v \n", Pnt01, PntRefrnce01)
	} else {
		fmt.Println("Point creation PASS. ")
	}

	// *** To Test SGF Point creation
	if SGF01 != SGFCrct01 {
		fmt.Println()
		t.Errorf("To SGF translation fail.\n"+
			"got: %v \n"+
			"want: %v \n", SGF01, SGFCrct01)
	} else {
		fmt.Println("To SGF translation PASS. ")
	}

	// *** To Test Pnt02 creation, using non-pointer copies
	if PntTest02 != PntRefrnce01 {
		fmt.Println()
		t.Errorf("To Point translation fail.\n"+
			"got: %v \n"+
			"want: %v \n", Pnt02, PntRefrnce01)
	} else {
		fmt.Println("To Point translation PASS. ")
	}
}
