package main

// *** This is the work file for the point translation methods add pull
// request 01

/*
// ### ### START For testing...
*/

/*
func main() {
	fmt.Println()
	fmt.Println("pointTranslateINTSGF: ")
	
	p01 := Point{
		x: 7,
		y: 19,
	}
	
	// New Immutable int *Point p02
	p02Int := New(p01.X(), p01.Y())
	fmt.Println()
	fmt.Println("Imm p02Int: ", p02Int)
	
	// New Immutable int *Point p03
	p03Int := Point{
		x: 7,
		y: 19,
	}
	fmt.Println()
	fmt.Println("Reg p03Int: ", p03Int)
	
	// INT TO SGF
	// p02Sgf point from p02 int *Point
	p02Sgf := translateToSGFPt01(p02Int)
	fmt.Println()
	fmt.Println("SGF p02Sgf: ", p02Sgf)
	
	// p02Sgf point from p02 int Point
	p03Sgf := translateToSGFPt02(p03Int)
	fmt.Println()
	fmt.Println("SGF p03Sgf: ", p03Sgf)
	
	// SGF TO INT
	// p03int point from p02Sgf point
	p04Int := translateToINTPt01(p02Sgf)
	fmt.Println()
	fmt.Println("INT p04Int: ", p04Int)
	
	// p03int point from p02Sgf point
	p05Int := translateToINTPt01(p03Sgf)
	fmt.Println()
	fmt.Println("INT p05Int: ", p05Int)
	
}

// Package point is a basic package for points.

// Point is a basic point. Although simple, the member variables are kept
// private to ensure that Point remains immutable.

type Point struct {
	x int64
	y int64
}

// New creates a new immutable Point.
func New(x, y int64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// X returns the x-value.
func (p *Point) X() int64 {
	return p.x
}

// Y returns the y-value.
func (p *Point) Y() int64 {
	return p.y
}
*/

/*
// ### ### END For testing...
*/

/*
// Three translation methods, one is an alternate

// 1. '*Point' immutable int point, to SGF(String of two letters) Point
func translateToSGFPt01(pt *Point) string {
	sgfPtX := string(rune((pt.X()) + 97))
	sgfPtY := string(rune((pt.Y()) + 97))
	// two lower case letters
	// return sgfPtX + sgfPtY
	return strings.ToLower(sgfPtX) + strings.ToLower(sgfPtY)
}

// Alternate, in case the incoming point is not immutable
// 2. 'Point' mutable int point, to SGF(String of two letters) Point
func translateToSGFPt02(pt Point) string {
	sgfPtX := string(rune((pt.X()) + 97))
	sgfPtY := string(rune((pt.Y()) + 97))
	// two lower case letters
	// return sgfPtX + sgfPtY
	return strings.ToLower(sgfPtX) + strings.ToLower(sgfPtY)
}

// 3. 'p' SGF(two letter string) Point, to '*Point' immutable Int Point
func translateToINTPt01(sgfPt string) *Point {
	sgfPtX := sgfPt[0]
	sgfPtY := sgfPt[1]
	x := int(sgfPtX) - 97
	y := int(sgfPtY) - 97
	return New(int64(x), int64(y))
}*/