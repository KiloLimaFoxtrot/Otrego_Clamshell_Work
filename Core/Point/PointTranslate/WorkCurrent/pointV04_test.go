package main

// import (
// 	"fmt"
// 	"testing"
// )
//
// // The previous version of this is pointV04_test.go
//
// func TestTranslate(t *testing.T) {
//
// 	fmt.Println()
// 	fmt.Println("[*** pointV04 Point Translation Tests ***]")
//
// 	// Test Point Struct
// 	type PointStrct struct {
// 		x int64
// 		y int64
// 	}
// 	// Test SGF Point
// 	var SGFCrct01 string
// 	// Constant value to convert ascii to letter and int
// 	const aValue = int64('a')
// 	// Test *Point
// 	var x1 int64 = 8
// 	var y1 int64 = 20
// 	Pnt01 := New(x1, y1)
// 	// Test Control Point, is non-pointer type so can locally verify
// 	PntCrct01 := PointStrct{
// 		x: Pnt01.x,
// 		y: Pnt01.y,
// 	}
// 	// Convert *Point to non-pointer type for local verification
// 	PntTest01 := PointStrct{
// 		x: Pnt01.x,
// 		y: Pnt01.y,
// 	}
//
// 	// *** To Test *Point creation, to verify our *Point for below tests
// 	if PntTest01 != PntCrct01 {
// 		fmt.Println()
// 		t.Errorf("Point creation fail.\n"+
// 			"got: %v \n"+
// 			"want: %v \n", Pnt01, PntCrct01)
// 	}
//
// 	// *** To Test ToSGF method
//
// 	// Test SGF
// 	SGF01 := Pnt01.ToSGF()
// 	// Test control SGF, as a string "xy", with the above *Point values
// 	SGFCrct01 = string(rune((PntCrct01.x)+aValue)) + string(rune((
// 		PntCrct01.y)+aValue))
// 	// To interject a different value, comment out the above
// 	/*SGFCrct01 = "ju"*/
//
// 	if SGF01 != SGFCrct01 {
// 		fmt.Println()
// 		t.Errorf("ToSGF translation fail.\n"+
// 			"got: %v \n"+
// 			"want: %v \n", SGF01, SGFCrct01)
// 	}
//
// 	// *** To Test *Point method, with the above values
// 	// Test *Point
// 	Pnt02 := NewFromSGF(SGF01)
// 	// To interject a different value, comment out the above
// 	/*Pnt02 := PointStrct{
// 		x: 1,
// 		y: 24,
// 	}*/
// 	// Convert *Point to non-pointer type for local verification
// 	PntTest02 := PointStrct{
// 		x: Pnt02.x,
// 		y: Pnt02.y,
// 	}
//
// 	if PntTest02 != PntCrct01 {
// 		fmt.Println()
// 		t.Errorf("Point creation fail.\n"+
// 			"got: %v \n"+
// 			"want: %v \n", Pnt02, PntCrct01)
// 	}
// }
//
// // Correct results
// // PntCrct01 := New(8, 20)
// // PntCrct02 := New(13, 5)
//
// /*
// [*** pointV04 Point Translation Tests ***]
//
// intPnt01:  &{8 20}
//
// SGFPnt01:  iu
//
// intPnt02:  &{8 20}
//
// */
//
// /*
// intPnt01:  &{13 5}
//
// SGFPnt01:  nf
//
// intPnt02:  &{13 5}
//
// */
