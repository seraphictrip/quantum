package quantum_test

import (
	"math"
	"slices"
	"strconv"
	"testing"

	"github.com/seraphictrip/quantum"
	"github.com/seraphictrip/quantum/binary"
)

var IsProbablisticStateVectorTests = []struct {
	v      []float64
	result bool
}{
	{
		result: false,
	},
	{
		v:      []float64{1},
		result: true,
	},
	{
		v:      []float64{.5, .5},
		result: true,
	},
	{
		v:      []float64{.5, .25, .25},
		result: true,
	},
	{
		v:      []float64{.5, .3, .1, .1},
		result: true,
	},
	{
		v:      []float64{.52, .48},
		result: true,
	},
	{
		v:      []float64{-1, 2},
		result: false,
	},
}

func TestIsProbablisticStateVector(t *testing.T) {
	for i, e := range IsProbablisticStateVectorTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := quantum.IsProbablisticStateVector(e.v)

			if e.result != result {
				t.Fatalf("IsProbablisticStateVector(%v) = %v, want %v", e.v, result, e.result)
			}
		})
	}
}

var EuclideanNormTests = []struct {
	v      []complex128
	result float64
}{
	// zero as expected
	{},
	{
		v:      []complex128{},
		result: 0,
	},
	{
		v:      []complex128{0, 1},
		result: 1,
	},
	{
		v:      []complex128{complex(1/math.Sqrt(2), 0), complex(1/math.Sqrt(2), 0)},
		result: 1,
	},
	{
		v:      []complex128{complex(1/math.Sqrt(2), 0), -complex(1/math.Sqrt(2), 0)},
		result: 1,
	},
	{
		v:      []complex128{(1 + 2i) / 3.0, complex(2.0/3.0, 0)},
		result: 1,
	},
}

func TestEuclideanNorm(t *testing.T) {
	for i, e := range EuclideanNormTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := quantum.EuclideanNorm(e.v)
			if result != e.result {
				t.Fatalf("EuclideanNorm(%v) = %v, want %v", e.v, result, e.result)
			}
		})
	}
}

func QuantumDie(n int) []complex128 {
	v := make([]complex128, n)
	for i := 0; i < n; i++ {
		v[i] = complex(1/math.Sqrt(float64(n)), 0)
	}
	return v
}

var IsQuantumStateVectorTests = []struct {
	v      []complex128
	result bool
}{
	{},
	{v: binary.PLUS_STATE, result: true},
	{v: binary.MINUS_STATE, result: true},
	{v: []complex128{(1 + 2i) / 3.0, complex(2.0/3.0, 0)}, result: true},
	{
		v:      []complex128{0, 1},
		result: true,
	},
	{v: []complex128{.5, 0 + 1i/2, binary.HC}, result: true},
	{v: []complex128{.5, 0, 0 + 1i/2, binary.HC}, result: true},
	{v: QuantumDie(6), result: true},
	{v: QuantumDie(20), result: true},
	{v: QuantumDie(100), result: true},
	{v: quantum.UniformQuantumStateVector([]int{1, 2, 3, 4, 5, 6}), result: true},
	{v: []complex128{.5, 0 + 1i/2, -0.5, -0 + 1i/2}, result: true},
}

func TestIsQuantumStateVector(t *testing.T) {
	for i, e := range IsQuantumStateVectorTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := quantum.IsQuantumStateVector(e.v)
			if result != e.result {
				t.Fatalf("IsQuantumStateVector(%v) = %v, want %v", e.v, result, e.result)
			}
		})
	}
}

var ToProbablisticStateVecortTests = []struct {
	v  []complex128
	ps []float64
}{
	{
		v:  binary.PLUS_STATE,
		ps: []float64{.5, .5},
	},
	{
		v:  binary.MINUS_STATE,
		ps: []float64{.5, .5},
	},
	{
		v:  []complex128{(1 + 2i) / 3, -(2.0 / 3)},
		ps: []float64{5.0 / 9, 4.0 / 9},
	},
}

func TestToProbablisticStateVector(t *testing.T) {
	for i, e := range ToProbablisticStateVecortTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ps := quantum.ToProbablisticStateVector(e.v)

			if !slices.Equal(ps, e.ps) {
				t.Fatalf("ToProbablisticStateVector(%v) = %v, want %v", e.v, ps, e.ps)
			}
		})
	}
}
