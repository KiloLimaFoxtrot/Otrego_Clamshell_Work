package main

import (
	"fmt"
	"testing"
)

func TestPointTranslation(t01 *testing.T) {
	type Point struct {
		x int64
		y int64
	}

	// First test translation from integer-point to SGF-string-point
	testToSGFCases := []struct {
		desc string
		in   Point
		want string
	}{
		{
			desc: "Point => SGF",
			in: Point{
				x: 8,
				y: 16,
			},
			want: "iq",
		},
		{
			desc: "Point => SGF",
			in: Point{
				x: 12,
				y: 5,
			},
			want: "mf",
		},
		{
			desc: "Point => SGF",
			in: Point{
				x: 33,
				y: 8,
			},
			want: "Hi",
		},
		{
			desc: "Point => SGF",
			in: Point{
				x: 40,
				y: 51,
			},
			want: "OZ",
		},
	}

	fmt.Println()
	for _, tc := range testToSGFCases {
		t01.Run(tc.desc, func(t01 *testing.T) {
			toSGFOut := New(tc.in.x, tc.in.y).ToSGF()
			if t01 == nil || toSGFOut != tc.want {
				t01.Errorf("%q.ToSGF() = %q, but wanted %q", tc.in,
					toSGFOut, tc.want)
			}
		})
	}

	// Second test translation from SGF-string-point to integer-point
	testToPointCases := []struct {
		desc string
		in   string
		want *Point
	}{
		{
			desc: "SGF => Point",
			in:   "iq",
			want: &Point{
				x: 8,
				y: 16,
			},
		},
		{
			desc: "SGF => Point",
			in:   "mf",
			want: &Point{
				x: 12,
				y: 5,
			},
		},
		{
			desc: "SGF => Point",
			in:   "Hi",
			want: &Point{
				x: 33,
				y: 8,
			},
		},
		{
			desc: "SGF => Point",
			in:   "OZ",
			want: &Point{
				x: 40,
				y: 51,
			},
		},
	}

	fmt.Println()
	for _, tc := range testToPointCases {
		t01.Run(tc.desc, func(t01 *testing.T) {
			toPointOut := NewFromSGF(tc.in)
			// include the point.go *Point type X() Y() getters below
			pointX := toPointOut.X()
			pointY := toPointOut.Y()
			if t01 == nil || pointX != tc.want.x {
				t01.Errorf("%q.toPointOut.x = %q, but wanted %q", tc.in,
					pointX, tc.want.x)
			}
			if t01 == nil || pointY != tc.want.y {
				t01.Errorf("%q.toPointOut.y = %q, but wanted %q", tc.in,
					pointY, tc.want.y)
			}
		})
	}
}
