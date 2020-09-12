package main

//import "bytes"
//
///*
//Keep in mind, from the Google Developers Guide/ What to look for in a
//code review...
//
//A particular type of complexity is over-engineering,
//where developers have made the code more generic than it needs to be,
//or added functionality that isn't presently needed by the system.
//Reviewers should be especially vigilant about over-engineering.
//
//*** Encourage developers to solve the problem they know needs to be
//solved now, not the problem that the developer speculates might need
//to be solved in the future.
//
//The future problem should be solved once it arrives and you can see
//its actual shape and requirements in the physical universe.
//*/
//
//// ToDo:
//// 1. Rewrite the below translation methods to work with a slice
//// /array of lower and then upper case of the alphabet,
//// which will work as a translation dictionary/mapping from the int
//// point value (slice index value) to the SGF point value (
//// either a lower or uppercase letter)
//
//// Point is a basic point. Although simple, the member variables are kept
//// private to ensure that Point remains immutable.
//type Point struct {
//	x int64
//	y int64
//}
//
//type SGF struct {
//	sgfX string
//	sgfY string
//}
//
//// pointSGFSlce provides a mapping or reference to translate between
////SGF and Point integer(index position) values.
////The translation methods below will use pointSGFSlce zero-indexed with
////standard/natural sorting
//
////var PointSGFSlce = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h',
////	'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u',
////	'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
////	'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U',
////	'V', 'W', 'X', 'Y', 'Z'}
//
//const aValue = int64('a')
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
//	sgfX := string(rune((pt.X()) + aValue))
//	sgfY := string(rune((pt.Y()) + aValue))
//	return sgfX + sgfY
//}
//
//// NewFromSGF converts an SGF point (
//// two letter string, 0-indexed) to a pointer-type (immutable) *Point.
//func NewFromSGF(sgfPt string) *Point {
//	PointSGFSlce := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h',
//		'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u',
//		'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
//		'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U',
//		'V', 'W', 'X', 'Y', 'Z'}
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
