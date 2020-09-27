package main

//
// import (
// 	"fmt"
// )
//
// // Point is a basic point. Although simple, the member variables are kept
// // private to ensure that Point remains immutable.
// type Point struct {
// 	x int64
// 	y int64
// }
//
// // pointToSgfRef is a translation reference between int64 Point
// // and string SGF-Point (rune) values
// var pointToSgfRef = map[int64]rune{
// 	0: 'a', 1: 'b', 2: 'c', 3: 'd', 4: 'e', 5: 'f', 6: 'g',
// 	7: 'h', 8: 'i', 9: 'j', 10: 'k', 11: 'l', 12: 'm', 13: 'n',
// 	14: 'o', 15: 'p', 16: 'q', 17: 'r', 18: 's', 19: 't', 20: 'u',
// 	21: 'v', 22: 'w', 23: 'x', 24: 'y', 25: 'z', 26: 'A', 27: 'B',
// 	28: 'C', 29: 'D', 30: 'E', 31: 'F', 32: 'G', 33: 'H', 34: 'I',
// 	35: 'J', 36: 'K', 37: 'L', 38: 'M', 39: 'N', 40: 'O', 41: 'P',
// 	42: 'Q', 43: 'R', 44: 'S', 45: 'T', 46: 'U', 47: 'V', 48: 'W',
// 	49: 'X', 50: 'Y', 51: 'Z',
// }
//
// // sgfToPointRef is an anonymous function that runs and inverts the
// // pointToSgfRef to create the reverse translation reference
// var sgfToPointRef = func(mapIn map[int64]rune) map[rune]int64 {
// 	mapElem := make(map[rune]int64)
// 	for key, val := range mapIn {
// 		mapElem[val] = key
// 	}
// 	return mapElem
// }(pointToSgfRef)
//
// // New creates a new immutable Point.
// func New(x, y int64) *Point {
// 	return &Point{
// 		x: x,
// 		y: y,
// 	}
// }
//
// // X returns the x-value.
// func (pt *Point) X() int64 { return pt.x }
//
// // Y returns the y-value.
// func (pt *Point) Y() int64 { return pt.y }
//
// // ToSGF converts a pointer-type (immutable) *Point to an SGF Point (
// // two letter string).
// // The returned value is 0-indexed.
// func (pt *Point) ToSGF() (string, error) {
// 	var sgfOut string
// 	var err01 = fmt.Errorf("")
// 	if (0 <= pt.X()) && (pt.X() <= 51) && (0 <= pt.Y()) && (pt.Y() <= 51) {
// 		sgfX := string(pointToSgfRef[pt.X()])
// 		sgfY := string(pointToSgfRef[pt.Y()])
// 		sgfOut = sgfX + sgfY
// 		err01 = nil
// 	} else {
// 		sgfOut = ""
// 		err01 = fmt.Errorf("*Point int64 x and y value entries must" +
// 			" be greater than or equal to 0, " +
// 			"and less than or equal to 51. ")
// 	}
// 	return sgfOut, err01
// }
//
// // NewFromSGF converts an SGF point (
// // two letter string) to a pointer-type (immutable) *Point (
// // (struct with two rune/char values).
// func NewFromSGF(sgfPt string) (*Point, error) {
// 	var intX int64
// 	var intY int64
// 	var err02 = fmt.Errorf("")
// 	if (sgfPt != "") && (len(sgfPt) == 2) {
// 		intX = sgfToPointRef[rune(sgfPt[0])]
// 		intY = sgfToPointRef[rune(sgfPt[1])]
// 		err02 = nil
// 	} else {
// 		intX = 99
// 		intY = 99
// 		err02 = fmt.Errorf("SGF string x and y value entries must non" +
// 			"-empty and of length 2 (runes/chars). ")
// 	}
// 	return New(intX, intY), err02
//
// }
//
// // *** DO NOT INCLUDE WITH PULL REQUEST OR OTHER REVIEWS !!! ***
// // *** FOR TESTING THIS BUILD VERSION ONLY !!! ***
// func main() {
// 	fmt.Println()
// 	fmt.Println("*** Point Build v01: ")
//
// 	for pt, sgf := range pointToSgfRef {
// 		fmt.Printf("pt: %v, sgf: %q\n", pt, sgf)
// 	}
//
// 	for sgf, pt := range sgfToPointRef {
// 		fmt.Printf("sgf: %q, pt: %v\n", sgf, pt)
// 	}
//
// 	TestPointSGFBuild()
//
// }
//
// func TestPointSGFBuild() {
// 	fmt.Println()
// 	fmt.Println("TestPointSGFBuild() ")
//
// 	// Sample test run 01
// 	fmt.Println()
// 	fmt.Println("Sample test run 01: ")
// 	// 1.
// 	pnt01 := New(36, 51) // SGF string "iu"
// 	fmt.Println("pnt01: ", pnt01)
// 	// 2.
// 	sgf01, err01 := pnt01.ToSGF()
// 	fmt.Printf("sgf01: %v - err01: %v\n", sgf01, err01)
// 	// 3.
// 	pnt02, err02 := NewFromSGF(sgf01) // SGF string "iu"
// 	fmt.Printf("pnt02: %v - err02: %v\n", pnt02, err02)
// 	// fmt.Println("pnt02: ", pnt02)
//
// 	// Sample test run 02
// 	fmt.Println()
// 	fmt.Println("Sample test run 02: ")
// 	// 1.
// 	pnt03 := New(36, 51) // SGF string "iu"
// 	fmt.Println("pnt03: ", pnt03)
// 	// 2.
// 	sgf02, err03 := pnt03.ToSGF()
// 	fmt.Printf("sgf02: %v - err03: %v\n", sgf02, err03)
// 	// fmt.Println("sgf02: ", sgf02)
// 	// 3.
// 	pnt04, err04 := NewFromSGF(sgf02) // SGF string "iu"
// 	fmt.Printf("pnt04: %v - err04: %v\n", pnt04, err04)
// 	// fmt.Println("pnt04: ", pnt04)
//
// }
