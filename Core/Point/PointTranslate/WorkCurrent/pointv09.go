package main

//
// import "fmt"
//
// // Point is a basic point. Although simple, the member variables are kept
// // private to ensure that Point remains immutable.
// type Point struct {
// 	x int64
// 	y int64
// }
//
// // PointSGFSlce is a slice/Array translation reference between integer
// // -Point (index position) and string SGF-Point (byte/char) values
//
// // var PointSGFSlce = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h',
// // 	'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u',
// // 	'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
// // 	'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U',
// // 	'V', 'W', 'X', 'Y'}
//
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
// /*
// var pointToSgfRef = map[int64]rune{
// 	1: 'a', 2: 'b', 3: 'c', 4: 'd', 5: 'e', 6: 'f', 7: 'g',
// 	8: 'h', 9: 'i', 10: 'j', 11: 'k', 12: 'l', 13: 'm', 14: 'n',
// 	15: 'o', 16: 'p', 17: 'q', 18: 'r', 19: 's', 20: 't', 21: 'u',
// 	22: 'v', 23: 'w', 24: 'x', 25: 'y', 26: 'z', 27: 'A', 28: 'B',
// 	29: 'C', 30: 'D', 31: 'E', 32: 'F', 33: 'G', 34: 'H', 35: 'I',
// 	36: 'J', 37: 'K', 38: 'L', 39: 'M', 40: 'N', 41: 'O', 42: 'P',
// 	43: 'Q', 44: 'R', 45: 'S', 46: 'T', 47: 'U', 48: 'V', 49: 'W',
// 	50: 'X', 51: 'Y', 52: 'Z',
// }
// */
//
// var sgfToPointRef = map[rune]int64{
// 	'a': 0, 'b': 1, 'c': 2, 'd': 3, 'e': 4, 'f': 5, 'g': 6, 'h': 7,
// 	'i': 8, 'j': 9, 'k': 10, 'l': 11, 'm': 12, 'n': 13, 'o': 14,
// 	'p': 15, 'q': 16, 'r': 17, 's': 18, 't': 19, 'u': 20, 'v': 21,
// 	'w': 22, 'x': 23, 'y': 24, 'z': 25, 'A': 26, 'B': 27, 'C': 28,
// 	'D': 29, 'E': 30, 'F': 31, 'G': 32, 'H': 33, 'I': 34, 'J': 35,
// 	'K': 36, 'L': 37, 'M': 38, 'N': 39, 'O': 40, 'P': 41, 'Q': 42,
// 	'R': 43, 'S': 44, 'T': 45, 'U': 46, 'V': 47, 'W': 48, 'X': 49,
// 	'Y': 50, 'Z': 51,
// }
//
// /*
// var sgfToPointRef = map[rune]int64{
// 	'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8,
// 	'i': 9, 'j': 10, 'k': 11, 'l': 12, 'm': 13, 'n': 14, 'o': 15,
// 	'p': 16, 'q': 17, 'r': 18, 's': 19, 't': 20, 'u': 21, 'v': 22,
// 	'w': 23, 'x': 24, 'y': 25, 'z': 26, 'A': 27, 'B': 28, 'C': 29,
// 	'D': 30, 'E': 31, 'F': 32, 'G': 33, 'H': 34, 'I': 35, 'J': 36,
// 	'K': 37, 'L': 38, 'M': 39, 'N': 40, 'O': 41, 'P': 42, 'Q': 43,
// 	'R': 44, 'S': 45, 'T': 46, 'U': 47, 'V': 48, 'W': 49, 'X': 50,
// 	'Y': 51, 'Z': 52,
// }
// */
// /*
// var sgfToPointRef = map[rune]int{
// 	'a': 1,
// 	'b': 2,
// 	'c': 3,
// 	'd': 4,
// 	'e': 5,
// 	'f': 6,
// 	'g': 7,
// 	'h': 8,
// 	'i': 9,
// 	'j': 10,
// 	'k': 11,
// 	'l': 12,
// 	'm': 13,
// 	'n': 14,
// 	'o': 15,
// 	'p': 16,
// 	'q': 17,
// 	'r': 18,
// 	's': 19,
// 	't': 20,
// 	'u': 21,
// 	'v': 22,
// 	'w': 23,
// 	'x': 24,
// 	'y': 25,
// 	'z': 26,
// 	'A': 27,
// 	'B': 28,
// 	'C': 29,
// 	'D': 30,
// 	'E': 31,
// 	'F': 32,
// 	'G': 33,
// 	'H': 34,
// 	'I': 35,
// 	'J': 36,
// 	'K': 37,
// 	'L': 38,
// 	'M': 39,
// 	'N': 40,
// 	'O': 41,
// 	'P': 42,
// 	'Q': 43,
// 	'R': 44,
// 	'S': 45,
// 	'T': 46,
// 	'U': 47,
// 	'V': 48,
// 	'W': 49,
// 	'X': 50,
// 	'Y': 51,
// 	'Z': 52,
// }
// */
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
// // ToSGF converts a pointer-type (immutable) *Point
// // to an SGF Point (two letter string). The returned value is 0-indexed.
// func (pt *Point) ToSGF() string {
// 	sgfOut := ""
// 	if (pt.X() <= 51) && (pt.Y() <= 51) {
// 		sgfX := string(pointToSgfRef[pt.X()])
// 		sgfY := string(pointToSgfRef[pt.Y()])
// 		sgfOut = sgfX + sgfY
// 	} else {
// 		sgfOut = "--"
// 	}
// 	return sgfOut
// }
//
// // NewFromSGF converts an SGF point (
// // two letter string, 0-indexed) to a pointer-type (immutable) *Point.
// func NewFromSGF(sgfPt string) *Point {
// 	var intX int64
// 	var intY int64
// 	if (sgfPt != "") && (sgfPt != "--") && (len(sgfPt) == 2) {
// 		intX = sgfToPointRef[rune(sgfPt[0])]
// 		intY = sgfToPointRef[rune(sgfPt[1])]
// 	} else {
// 		intX = 99
// 		intY = 99
// 	}
// 	return New(intX, intY)
//
// }
//
// // Testing functions for this build ??
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
// 	sgf01 := pnt01.ToSGF()
// 	fmt.Println("sgf01: ", sgf01)
// 	// 3.
// 	pnt02 := NewFromSGF(sgf01) // SGF string "iu"
// 	fmt.Println("pnt02: ", pnt02)
//
// 	// Sample test run 02
// 	fmt.Println()
// 	fmt.Println("Sample test run 02: ")
// 	// 1.
// 	pnt03 := New(36, 52) // SGF string "iu"
// 	fmt.Println("pnt03: ", pnt03)
// 	// 2.
// 	sgf02 := pnt03.ToSGF()
// 	fmt.Println("sgf02: ", sgf02)
// 	// 3.
// 	pnt04 := NewFromSGF(sgf02) // SGF string "iu"
// 	fmt.Println("pnt04: ", pnt04)
//
// }
