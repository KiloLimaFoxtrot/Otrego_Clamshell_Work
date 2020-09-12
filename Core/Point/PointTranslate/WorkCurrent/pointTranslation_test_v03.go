package main

// Sample correct results
// New(8, 20), SGF string "iu"
// New(13, 5), "nf"

/*
// Notes from Discord texting with Josh 7 September 2020,
about the translation point test I am building for Otrego
[... missing notes from Discord chat with Josh]
4 Types of tests
1. PntPrefernc01 != PtToTest01
2. TestConverSGFToPoint
3. TestConvertPointToSGF
4. TestRoundTripConvert

Considerations
1. Have a single test method test a thing
2. Don't compare structs, compare fields
e.g. below
a.x == b.x
a.y == b.y

or
[Create an equals method, that returns a bool value]
func (p *Point) Equals(pt *Point) bool {
  return p.x == pt.x && p.y == pt.y
}

probably easier to do [test field values like above]
a.x == b.x
a.y == b.y
[Instead of the struct comparison]
if PntTest01 != PntCrct01 {
        fmt.Println()
        t.Errorf("Point creation fail.\n"+
            "got: %v \n"+
            "want: %v \n", Pnt01, PntCrct01)
    } else {
        fmt.Println("Point creation PASS. ")
    }
you won't need the PASS

*/
//
// func TestTranslate(t *testing.T) {
//
// 	fmt.Println("[*** Go Game Point Translation Tests ***]")
//
// 	// *** Test 'Fields' *** //
// 	var x1 int64 = 8          // User chosen x values for testing
// 	var y1 int64 = 20         // User chosen y values for testing
// 	const aValue = int64('a') // Ascii int string conversion value
//
// 	// ** Create test pointer-type *Point with pointV06.go New() method
// 	Pnt01 := New(x1, y1)
// 	// Create non-pointer-type copy of Pnt01,
// 	// and a test reference copy, for this unit test
// 	type PointStrct struct { // non-pointer-type struct
// 		x int64
// 		y int64
// 	}
// 	PntToTest01 := PointStrct{ // Non-pointer Pnt01 copy to test
// 		x: Pnt01.x,
// 		y: Pnt01.y,
// 	}
// 	PntRefrnce01 := PointStrct{ // Non-pointer Pnt01 reference copy
// 		x: Pnt01.x,
// 		y: Pnt01.y,
// 	}
//
// 	// ** Create test SGF Point with pointV06.go ToSGF() method
// 	SGF01 := Pnt01.ToSGF()
// 	// Create a test reference copy of SGF01, for this unit test
// 	var SGFCrct01 string // string variable to hold SGF values
// 	SGFCrct01 = string(rune((PntToTest01.x)+aValue)) + string(rune((PntToTest01.y)+aValue))
// 	// Interject different user values with below, and comment out above
// 	// SGFCrct01 = "ju"
//
// 	// ** Create test *Point from SGF01, which should equal Pnt01
// 	Pnt02 := NewFromSGF(SGF01)
// 	// Create non-pointer-type copy of Pnt01, for this unit test.
// 	// Do not need reference copy as Pnt02 should equal Pnt01
// 	PntTest02 := PointStrct{ // Non-pointer Pnt02 copy to test
// 		x: Pnt02.x,
// 		y: Pnt02.y,
// 	}
// 	// Interject different user values with below, and comment out above
// 	// Pnt02 := PointStrct{
// 	// 	x: 1,
// 	// 	y: 24,
// 	// }
//
// 	// *** Test Statements *** //
//
// 	// *** To Test Pnt01 creation, using non-pointer copies
// 	if PntToTest01 != PntRefrnce01 {
// 		fmt.Println()
// 		t.Errorf("Point creation fail.\n"+
// 			"got: %v \n"+
// 			"want: %v \n", Pnt01, PntRefrnce01)
// 	} else {
// 		fmt.Println("Point creation PASS. ")
// 	}
//
// 	// *** To Test SGF Point creation
// 	if SGF01 != SGFCrct01 {
// 		fmt.Println()
// 		t.Errorf("To SGF translation fail.\n"+
// 			"got: %v \n"+
// 			"want: %v \n", SGF01, SGFCrct01)
// 	} else {
// 		fmt.Println("To SGF translation PASS. ")
// 	}
//
// 	// *** To Test Pnt02 creation, using non-pointer copies
// 	if PntTest02 != PntRefrnce01 {
// 		fmt.Println()
// 		t.Errorf("To Point translation fail.\n"+
// 			"got: %v \n"+
// 			"want: %v \n", Pnt02, PntRefrnce01)
// 	} else {
// 		fmt.Println("To Point translation PASS. ")
// 	}
// }
