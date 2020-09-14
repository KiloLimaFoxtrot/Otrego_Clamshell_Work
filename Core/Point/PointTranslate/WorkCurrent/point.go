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

// PointSGFSlce is a slice/Array translation reference between integer
// -Point (index position) and string SGF-Point (byte/char) values

// var PointSGFSlce = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h',
// 	'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u',
// 	'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
// 	'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U',
// 	'V', 'W', 'X', 'Y'}

var PointToSgfMap = map[int64]rune{
	1: 'a', 2: 'b', 3: 'c', 4: 'd', 5: 'e', 6: 'f', 7: 'g',
	8: 'h', 9: 'i', 10: 'j', 11: 'k', 12: 'l', 13: 'm', 14: 'n',
	15: 'o', 16: 'p', 17: 'q', 18: 'r', 19: 's', 20: 't', 21: 'u',
	22: 'v', 23: 'w', 24: 'x', 25: 'y', 26: 'z', 27: 'A', 28: 'B',
	29: 'C', 30: 'D', 31: 'E', 32: 'F', 33: 'G', 34: 'H', 35: 'I',
	36: 'J', 37: 'K', 38: 'L', 39: 'M', 40: 'N', 41: 'O', 42: 'P',
	43: 'Q', 44: 'R', 45: 'S', 46: 'T', 47: 'U', 48: 'V', 49: 'W',
	50: 'X', 51: 'Y', 52: 'Z',
}

var SgfToPointMap = map[rune]int{
	'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8,
	'i': 9, 'j': 10, 'k': 11, 'l': 12, 'm': 13, 'n': 14, 'o': 15,
	'p': 16, 'q': 17, 'r': 18, 's': 19, 't': 20, 'u': 21, 'v': 22,
	'w': 23, 'x': 24, 'y': 25, 'z': 26, 'A': 27, 'B': 28, 'C': 29,
	'D': 30, 'E': 31, 'F': 32, 'G': 33, 'H': 34, 'I': 35, 'J': 36,
	'K': 37, 'L': 38, 'M': 39, 'N': 40, 'O': 41, 'P': 42, 'Q': 43,
	'R': 44, 'S': 45, 'T': 46, 'U': 47, 'V': 48, 'W': 49, 'X': 50,
	'Y': 51, 'Z': 52,
}

/*
var SgfToPointMap = map[rune]int{
	'a': 1,
	'b': 2,
	'c': 3,
	'd': 4,
	'e': 5,
	'f': 6,
	'g': 7,
	'h': 8,
	'i': 9,
	'j': 10,
	'k': 11,
	'l': 12,
	'm': 13,
	'n': 14,
	'o': 15,
	'p': 16,
	'q': 17,
	'r': 18,
	's': 19,
	't': 20,
	'u': 21,
	'v': 22,
	'w': 23,
	'x': 24,
	'y': 25,
	'z': 26,
	'A': 27,
	'B': 28,
	'C': 29,
	'D': 30,
	'E': 31,
	'F': 32,
	'G': 33,
	'H': 34,
	'I': 35,
	'J': 36,
	'K': 37,
	'L': 38,
	'M': 39,
	'N': 40,
	'O': 41,
	'P': 42,
	'Q': 43,
	'R': 44,
	'S': 45,
	'T': 46,
	'U': 47,
	'V': 48,
	'W': 49,
	'X': 50,
	'Y': 51,
	'Z': 52,
}
*/

// New creates a new immutable Point.
func New(x, y int64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// X returns the x-value.
func (pt *Point) X() int64 { return pt.x }

// Y returns the y-value.
func (pt *Point) Y() int64 { return pt.y }

// ToSGF converts a pointer-type (immutable) *Point
// to an SGF Point (two letter string). The returned value is 1-indexed.
func (pt *Point) ToSGF() string {
	sgfOut := ""
	if (pt.X() <= 52) && (pt.Y() <= 52) {
		sgfX := string(PointToSgfMap[pt.X()])
		sgfY := string(PointToSgfMap[pt.Y()])
		sgfOut = sgfX + sgfY
	} else {
		panic("Error: *Point entries must not be nil for either" +
			" x y entries and have 0 <= x, y <= 51 values. ")
	}
	return sgfOut
}

// NewFromSGF converts an SGF point (
// two letter string, 0-indexed) to a pointer-type (immutable) *Point.
func NewFromSGF(sgfPt string) *Point {
	var intX int
	var intY int
	if (sgfPt != "") && (len(sgfPt) == 2) {
		intX = SgfToPointMap[rune(sgfPt[0])]
		intY = SgfToPointMap[rune(sgfPt[1])]

		// intX = bytes.Index(PointSGFSlce, []byte(string(sgfPt[0])))
		// intY = bytes.Index(PointSGFSlce, []byte(string(sgfPt[1])))
	} else {
		panic("Error: SGF string entries must not be empty and must" +
			" be of length = 2 byte/byte.  ")
	}
	return New(int64(intX), int64(intY))

}

// Testing functions for this build ??
func main() {
	fmt.Println()
	fmt.Println("*** Point Build v01: ")

	for pt, sgf := range PointToSgfMap {
		fmt.Printf("pt: %v, sgf: %q\n", pt, sgf)
	}

	for sgf, pt := range SgfToPointMap {
		fmt.Printf("sgf: %q, pt: %v\n", sgf, pt)
	}

	TestPointSGFBuild()

}

func TestPointSGFBuild() {
	fmt.Println()
	fmt.Println("TestPointSGFBuild() ")

	// Sample test run 01
	fmt.Println()
	fmt.Println("Sample test run 01: ")
	// 1.
	pnt01 := New(36, 50) // SGF string "iu"
	fmt.Println("pnt01: ", pnt01)
	// 2.
	sgf01 := pnt01.ToSGF()
	fmt.Println("sgf01: ", sgf01)
	// 3.
	pnt02 := NewFromSGF(sgf01) // SGF string "iu"
	fmt.Println("pnt02: ", pnt02)

	// Sample test run 02
	fmt.Println()
	fmt.Println("Sample test run 02: ")
	// 1.
	pnt03 := New(8, 47) // SGF string "iu"
	fmt.Println("pnt03: ", pnt03)
	// 2.
	sgf02 := pnt03.ToSGF()
	fmt.Println("sgf02: ", sgf02)
	// 3.
	pnt04 := NewFromSGF(sgf02) // SGF string "iu"
	fmt.Println("pnt04: ", pnt04)

}
