package main

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
// type SGF struct {
// 	x string
// 	y string
// }
//
// const aValue = int64('a')
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
// // Two Go-game board position/point translation methods
// // 1. '*Point' immutable int point, to SGF(String of two letters) point
// func (pt *Point) ToSGF() string {
// 	sgfX := string(rune((pt.X()) + aValue))
// 	sgfY := string(rune((pt.Y()) + aValue))
// 	// two lower case letters
// 	return sgfX + sgfY
// }
//
// // 3. 'p' SGF(two letter string) Point, to '*Point' immutable Int Point
// func NewFromSGF(sgfPt string) *Point {
// 	sgfX := sgfPt[0]
// 	sgfY := sgfPt[1]
// 	x := int64(sgfX) - aValue
// 	y := int64(sgfY) - aValue
// 	return New(int64(x), int64(y))
// }
//
// // For Testing the above translation methods
// func main() {
//
// 	fmt.Println()
// 	fmt.Println("[*** pointV04 Point Translation Tests ***]")
//
// 	intPnt01 := New(13, 5)
// 	fmt.Println()
// 	fmt.Println("intPnt01: ", intPnt01)
//
// 	// This test the ToSGF method, which is a method on the Point object
// 	SGFPnt01 := intPnt01.ToSGF()
// 	fmt.Println()
// 	fmt.Println("SGFPnt01: ", SGFPnt01)
//
// 	intPnt02 := NewFromSGF(SGFPnt01)
// 	fmt.Println()
// 	fmt.Println("intPnt02: ", intPnt02)
//
// }
