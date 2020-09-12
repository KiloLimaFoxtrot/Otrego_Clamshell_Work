package main

//import (
//	"bytes"
//	"fmt"
//)
//
///*
//Keep in mind...
//solve the problem they know needs to be solved now,
//not the problem that the developer speculates might need to be solved
//in the future.
//*/
//
//// Sample correct results
//// New(8, 20), SGF string "iu"
//// New(13, 5), "nf"
//
//// Point is a basic point. Although simple, the member variables are kept
//// private to ensure that Point remains immutable.
//type Point struct {
//	x int64
//	y int64
//}
//
//const aValue = int64('a')
//
//var SGFstr = ""
//
//type SGF struct {
//	sgfX string
//	sgfY string
//}
//
//var PointSGFSlce = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h',
//	'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u',
//	'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
//	'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U',
//	'V', 'W', 'X', 'Y', 'Z'}
//
//// New creates a new immutable Point.
//func New(x, y int64) *Point {
//	return &Point{
//		x: x,
//		y: y,
//	}
//}
//
//// X returns the x-value.
//func (p *Point) X() int64 { return p.x }
//
//// Y returns the y-value.
//func (p *Point) Y() int64 { return p.y }
//
//// ToSGF() converts a pointer-type (immutable) *Point
//// to an SGF Point (two letter string). The returned value is 0-indexed.
//func (pt *Point) ToSGF() string {
//	sgfX := string(PointSGFSlce[pt.x])
//	sgfY := string(PointSGFSlce[pt.y])
//
//	return sgfX + sgfY
//}
//
//// NewFromSGF converts an SGF point (
//// two letter string, 0-indexed) to a pointer-type (immutable) *Point.
//func NewFromSGF(sgfPt string) *Point {
//	//PointSGFSlce := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h',
//	//	'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u',
//	//	'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
//	//	'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U',
//	//	'V', 'W', 'X', 'Y', 'Z'}
//
//	//x := string(sgfPt[0])
//	//y := string(sgfPt[1])
//	//xIndx := bytes.Index(PointSGFSlce, []byte(x))
//	//yIndx := bytes.Index(PointSGFSlce, []byte(y))
//	// condensed below
//	xIndx := bytes.Index(PointSGFSlce, []byte(string(sgfPt[0])))
//	yIndx := bytes.Index(PointSGFSlce, []byte(string(sgfPt[1])))
//
//	//x := int64(sgfPt[0]) - aValue
//	//y := int64(sgfPt[1]) - aValue
//	return New(int64(xIndx), int64(yIndx))
//}
//
//// Testing functions for this build ??
//func main() {
//
//	fmt.Println()
//	fmt.Println("*** Point Build v01: ")
//}
//
//func TestSort(sgfPnt []byte) {
//	//type (
//	//	//Possibly for testing...
//	//	PointSGFStct = []struct {
//	//		sgfX byte
//	//		sgfY byte
//	//	}
//	//)
//	//{
//	//	{'a', 'b'},
//	//	{'y', 'z'},
//	//
//	//}
//}
