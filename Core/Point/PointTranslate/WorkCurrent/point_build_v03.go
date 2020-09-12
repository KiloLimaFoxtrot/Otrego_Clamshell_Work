package main

//
// import (
// 	"bytes"
// )
//
// // Point is a basic point. Although simple, the member variables are kept
// // private to ensure that Point remains immutable.
// type Point struct {
// 	x int64
// 	y int64
// }
//
// // PointSGFSlce is a slice/Array that serves as the translation key
// // between integer-Point (index position) and string SGF-Point (
// // byte/char) values at the index position
// var PointSGFSlce = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h',
// 	'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u',
// 	'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
// 	'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U',
// 	'V', 'W', 'X', 'Y', 'Z'}
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
// func (p *Point) X() int64 { return p.x }
//
// // Y returns the y-value.
// func (p *Point) Y() int64 { return p.y }
//
// // ToSGF() converts a pointer-type (immutable) *Point
// // to an SGF Point (two letter string). The returned value is 0-indexed.
// func (pt *Point) ToSGF() string {
// 	sgfX := string(PointSGFSlce[pt.X()])
// 	sgfY := string(PointSGFSlce[pt.Y()])
//
// 	return sgfX + sgfY
// }
//
// // NewFromSGF converts an SGF point (
// // two letter string, 0-indexed) to a pointer-type (immutable) *Point.
// func NewFromSGF(sgfPt string) *Point {
// 	// x := string(sgfPt[0])
// 	// y := string(sgfPt[1])
// 	// xIndx := bytes.Index(PointSGFSlce, []byte(x))
// 	// yIndx := bytes.Index(PointSGFSlce, []byte(y))
// 	// condensed below
// 	xIndx := bytes.Index(PointSGFSlce, []byte(string(sgfPt[0])))
// 	yIndx := bytes.Index(PointSGFSlce, []byte(string(sgfPt[1])))
//
// 	// x := int64(sgfPt[0]) - aValue
// 	// y := int64(sgfPt[1]) - aValue
// 	return New(int64(xIndx), int64(yIndx))
// }

// // Testing functions for this build ??
// func main() {
// 	fmt.Println()
// 	fmt.Println("*** Point Build v01: ")
//
// 	for i := 0; i < len(PointSGFSlce); i++ {
// 		fmt.Printf("indx: %v, val: %v\n", i, string(PointSGFSlce[i]))
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
// 	pnt01 := New(8, 20) // SGF string "iu"
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
// 	pnt03 := New(8, 32) // SGF string "iu"
// 	fmt.Println("pnt03: ", pnt03)
// 	// 2.
// 	sgf02 := pnt03.ToSGF()
// 	fmt.Println("sgf02: ", sgf02)
// 	// 3.
// 	pnt04 := NewFromSGF(sgf02) // SGF string "iu"
// 	fmt.Println("pnt04: ", pnt04)
//
// }
