package main

//
// import (
// 	"fmt"
// 	"testing"
// )
//
// // Sample correct results
// // New(8, 20), SGF string "iu"
// // New(13, 5), "nf"
//
// // func main() {
// //
// //	fmt.Println()
// //	fmt.Println("*** Point Translation 02 Test: ")
// // }
//
// func TestPointToSGFTranslate(t01 *testing.T) {
// 	type Point struct {
// 		x int64
// 		y int64
// 	}
// 	//
// 	// type SGF struct {
// 	//	sgfX string
// 	//	sgfY string
// 	// }
//
// 	testToSGFCases := []struct {
// 		desc string
// 		in   Point
// 		want string
// 	}{
// 		{
// 			desc: "Point => SGF",
// 			in: Point{
// 				x: 8,
// 				y: 16,
// 			},
// 			want: "iq",
// 		},
// 		{
// 			desc: "Point => SGF",
// 			in: Point{
// 				x: 12,
// 				y: 5,
// 			},
// 			want: "mf",
// 		},
// 		{
// 			desc: "Point => SGF",
// 			in: Point{
// 				x: 33,
// 				y: 8,
// 			},
// 			want: "Hi",
// 		},
// 		{
// 			desc: "Point => SGF",
// 			in: Point{
// 				x: 40,
// 				y: 51,
// 			},
// 			want: "OZ",
// 		},
// 	}
//
// 	fmt.Println()
// 	for _, tc := range testToSGFCases {
// 		t01.Run(tc.desc, func(t01 *testing.T) {
// 			toSGFOut := New(tc.in.x, tc.in.y).ToSGF()
// 			// For testing purposes
// 			// fmt.Printf("%v.ToSGF() = %v, but wanted %v\n", tc.in,
// 			// 	toSGFOut, tc.want)
// 			if t01 == nil || toSGFOut != tc.want {
// 				t01.Errorf("%q.ToSGF() = %q, but wanted %q", tc.in,
// 					toSGFOut, tc.want)
// 			}
// 		})
// 	}
//
// 	testToPointCases := []struct {
// 		desc string
// 		in   string
// 		want *Point
// 	}{
// 		{
// 			desc: "SGF => Point",
// 			in:   "iq",
// 			want: &Point{
// 				x: 8,
// 				y: 16,
// 			},
// 		},
// 		{
// 			desc: "SGF => Point",
// 			in:   "mf",
// 			want: &Point{
// 				x: 12,
// 				y: 5,
// 			},
// 		},
// 		{
// 			desc: "SGF => Point",
// 			in:   "Hi",
// 			want: &Point{
// 				x: 33,
// 				y: 8,
// 			},
// 		},
// 		{
// 			desc: "SGF => Point",
// 			in:   "OZ",
// 			want: &Point{
// 				x: 40,
// 				y: 51,
// 			},
// 		},
// 	}
//
// 	fmt.Println()
// 	for _, tc := range testToPointCases {
// 		t01.Run(tc.desc, func(t01 *testing.T) {
// 			toPointOut := NewFromSGF(tc.in)
// 			pointX := toPointOut.x
// 			pointY := toPointOut.y
// 			// For testing purposes
// 			// fmt.Printf("%v.toPointOut.x = %v, but wanted %v", tc.in,
// 			//	pointX, tc.want.x)
// 			// The test condition
// 			if t01 == nil || pointX != tc.want.x {
// 				t01.Errorf("%q.toPointOut.x = %q, but wanted %q", tc.in,
// 					pointX, tc.want.x)
// 			}
// 			// For testing purposes
// 			// fmt.Printf("%v.toPointOut.y = %v, but wanted %v", tc.in,
// 			//	pointY, tc.want.y)
// 			// The test condition
// 			if t01 == nil || pointY != tc.want.y {
// 				t01.Errorf("%q.toPointOut.y = %q, but wanted %q", tc.in,
// 					pointY, tc.want.y)
// 			}
//
// 		})
// 	}
// }
