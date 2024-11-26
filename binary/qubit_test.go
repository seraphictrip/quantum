package binary_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/seraphictrip/quantum/binary"
)

var PhaseOpTests = []struct {
}{}

func TestPhaseOp(t *testing.T) {
	// var a = (1 + 1i) / math.Sqrt2
	fmt.Println(binary.PhaseOp(math.Pi / 4))
	fmt.Println(binary.HC)
	//	for i, e := range PhaseOpTests {
	//
	// // 		e := e
	// // 		t.Run(strconv.Itoa(i), func(t *testing.T){
	// // // code
	// // 		})
	//
	//	}
}
