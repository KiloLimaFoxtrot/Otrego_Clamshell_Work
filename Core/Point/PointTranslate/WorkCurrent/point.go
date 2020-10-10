package main

import (
	"fmt"
)

// Point is a basic point. Although simple, the member variables are kept
// private to ensure that Point remains immutable.
type Point struct {
	x int64
	y int64
}

// pointToSgfRef is a translation reference between int64 Point
// and string SGF-Point (rune) values
var pointToSgfRef = map[int64]rune{
	0: 'a', 1: 'b', 2: 'c', 3: 'd', 4: 'e', 5: 'f', 6: 'g',
	7: 'h', 8: 'i', 9: 'j', 10: 'k', 11: 'l', 12: 'm', 13: 'n',
	14: 'o', 15: 'p', 16: 'q', 17: 'r', 18: 's', 19: 't', 20: 'u',
	21: 'v', 22: 'w', 23: 'x', 24: 'y', 25: 'z', 26: 'A', 27: 'B',
	28: 'C', 29: 'D', 30: 'E', 31: 'F', 32: 'G', 33: 'H', 34: 'I',
	35: 'J', 36: 'K', 37: 'L', 38: 'M', 39: 'N', 40: 'O', 41: 'P',
	42: 'Q', 43: 'R', 44: 'S', 45: 'T', 46: 'U', 47: 'V', 48: 'W',
	49: 'X', 50: 'Y', 51: 'Z',
}

// sgfToPointRef is an anonymous function that runs and inverts the
// pointToSgfRef to create the reverse translation reference
var sgfToPointRef = func(in map[int64]rune) map[rune]int64 {
	elem := make(map[rune]int64)
	for key, val := range in {
		elem[val] = key
	}
	return elem
}(pointToSgfRef)

// New creates a new immutable Point.
func New(x, y int64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// X() method receives a pointer argument of type *Point,
// and returns the x-value.
func (pt *Point) X() int64 { return pt.x }

// Y() method receives a pointer argument of type *Point,
// and returns the y-value.
func (pt *Point) Y() int64 { return pt.y }

// ToSGF() method receives a pointer argument of type *Point,
// and converts the pointer-type (immutable) *Point
// to an SGF Point (two letter string).
func (pt *Point) ToSGF() (string, error) {
	if (pt.X() < 0) || (51 < pt.X()) || (pt.Y() < 0) || 51 < (pt.Y()) {
		return "", fmt.Errorf(
			"*Point int64 x and y value entries must be greater than" +
				" or equal to 0, " +
				"and less than or equal to 51. ")
	}
	return string(pointToSgfRef[pt.X()]) + string(
		pointToSgfRef[pt.Y()]), nil

}

// String() method receives a pointer argument of type Point,
// and performs to represent and print a Point;
// useful for debugging and test purposes - ilmanzo contributing
func (pt Point) String() string {
	return fmt.Sprintf("{%d,%d}", pt.x, pt.y)
}

// NewFromSGF function receives a pointer argument of type string,
// converts an SGF point (two letter string) to a pointer point of
// *Point (an immutable type struct with two rune/char values).
func NewFromSGF(sgfPt string) (*Point, error) {
	if (sgfPt == "") || (len(sgfPt) != 2) {
		return nil, fmt.Errorf(
			"SGF string x and y value entries must be non-empty and" +
				" of length 2 (runes/chars). ")
	}
	return New(sgfToPointRef[rune(sgfPt[0])],
		sgfToPointRef[rune(sgfPt[1])]), nil

}

// *** DO NOT INCLUDE WITH PULL REQUEST OR OTHER REVIEWS !!! ***
// *** FOR TESTING THIS BUILD VERSION ONLY !!! ***
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
