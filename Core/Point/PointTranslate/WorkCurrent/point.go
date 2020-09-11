package main

// ToDo:
// 1. Rewrite the below translation methods to work with a slice
// /array of lower and then upper case of the alphabet,
// which will work as a translation dictionary/mapping from the int
// point value (slice index value) to the SGF point value (
// either a lower or uppercase letter)

// Point is a basic point. Although simple, the member variables are kept
// private to ensure that Point remains immutable.
type Point struct {
	x int64
	y int64
}

const aValue = int64('a')

// New creates a new immutable Point.
func New(x, y int64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// X returns the x-value.
func (p *Point) X() int64 { return p.x }

// Y returns the y-value.
func (p *Point) Y() int64 { return p.y }

// ToSGF() converts a pointer-type (immutable) *Point
// to an SGF Point (two letter string). The returned value is 0-indexed.
func (pt *Point) ToSGF() string {
	sgfX := string(rune((pt.X()) + aValue))
	sgfY := string(rune((pt.Y()) + aValue))
	return sgfX + sgfY
}

// NewFromSGF converts an SGF point (
// two letter string, 0-indexed) to a pointer-type (immutable) *Point.
func NewFromSGF(sgfPt string) *Point {
	x := int64(sgfPt[0]) - aValue
	y := int64(sgfPt[1]) - aValue
	return New(int64(x), int64(y))
}
