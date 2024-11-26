package quantum

import (
	"math"
	"math/cmplx"
)

func IsQuantumStateVector(v []complex128) bool {
	// The sum of the absolute values squared of the entries of a quantum state vector is
	norm := EuclideanNorm(v)
	return InTolerance(norm, .999999, 1.00009)
}

func InTolerance(val, low, high float64) bool {
	if low == high {
		return val == low
	}
	return val > low && val < high
}

func UniformQuantumStateVector[T any](orderedSet []T) []complex128 {
	popSize := len(orderedSet)
	v := make([]complex128, popSize)
	for i := range orderedSet {
		v[i] = complex(1./math.Sqrt(float64(popSize)), 0)
	}

	return v
}

func AbsoluteSquare(i complex128) float64 {
	return math.Pow(cmplx.Abs(i), 2)
}

// sum of the absolute values squared
func EuclideanNorm(v []complex128) float64 {
	norm := 0.0
	for _, i := range v {
		norm += AbsoluteSquare(i)
	}
	return norm
}

func UniformProbablisticStateVector(orderedSet []any) []float64 {
	popSize := len(orderedSet)
	v := make([]float64, popSize)
	for i := range orderedSet {
		v[i] = 1. / float64(popSize)
	}

	return v
}

func IsProbablisticStateVector(v []float64) bool {
	// the sum of all entries == 1
	var sum float64 = 0
	for _, val := range v {
		if val < 0 {
			return false
		}
		sum += val
	}

	return sum == 1
}

func ToProbablisticStateVector(v []complex128) []float64 {
	if !IsQuantumStateVector(v) {
		panic("no compute")
	}
	ps := make([]float64, len(v))
	for i, num := range v {
		ps[i] = AbsoluteSquare(num)
	}

	return ps
}
