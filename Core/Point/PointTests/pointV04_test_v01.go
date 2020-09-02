package main
/*
import (
	"fmt"
	"testing"
)

// func main() {
//
// }

func TestNew(t *testing.T) {
	
	fmt.Println()
	fmt.Println("[*** pointV04 Point Translation Tests ***]")
	
	// Test Structs
	type PointStrct struct {
		x int64
		y int64
	}
	
	type SGFStrct struct {
		sgfX string
		sgfY string
	}
	
	const aValue = int64('a')
	
	// Correct results
	// PntCrct01 := New(8, 20)
	// PntCrct02 := New(13, 5)
	
	// Point Creation Test
	var x1 int64 = 8
	var y1 int64 = 20
	
	// Test Control Point
	PntCrct01 := PointStrct{
		x: x1,
		y: y1,
	}
	Pnt01 := New(int64(x), int64(y))
	
	fmt.Println()
	if Pnt01 != PntCrct01 {
		fmt.Println()
		t.Errorf("Point creation fail.\n"+
			"got: %d \n"+
			"want: &d \n", Pnt01, PntCrct01)
	}
}
	/*
	// Test ToSGF method, a method on the Point object
	// Test Control SGF, uses the // Test Control Point values
	SGFCrct01 := SGF{
		string(rune((PntCrct01.x) + aValue)),
		string(rune((PntCrct01.y) + aValue)),
	}
	
	SGF01 := Pnt01.ToSGF()
	fmt.Println()
	if SGF01 != SGFCrct01 {
		fmt.Println()
		t.Errorf("SGF creation fail. \n"+
			"got: %d\n"+
			"want: &d \n", SGF01, SGFCrct01)
	}
	
	Pnt02 := NewFromSGF(SGF01)
	fmt.Println()
	if Pnt02 != PntCrct01 {
		fmt.Println()
		t.Errorf("Point creation fail.\n"+
			"got: %d \n"+
			"want: &d \n", Pnt02, PntCrct01)
	}
	
}*/

/*
[*** pointV04 Point Translation Tests ***]

intPnt01:  &{8 20}

SGFPnt01:  iu

intPnt02:  &{8 20}

*/

/*
intPnt01:  &{13 5}

SGFPnt01:  nf

intPnt02:  &{13 5}

*/
