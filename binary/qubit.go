package binary

import (
	"math"
	"math/cmplx"
)

// 1/math.Sqrt(2)
var HC = complex(1/math.Sqrt2, 0)

// -1/math.Sqrt(2)
var HCN = -HC

var PLUS_STATE = []complex128{HC, HC}

var MINUS_STATE = []complex128{HC, HCN}

var PauliIdentity = [][]complex128{
	{1, 0},
	{0, 1},
}

var PauliX = [][]complex128{
	{0, 1},
	{1, 0},
}

var PauliY = [][]complex128{
	{0, 0 + -1i},
	{0 + 1i, 0},
}

var PauliZ = [][]complex128{
	{1, 0},
	{0, -1},
}

var PhaseFlip = PauliZ

var Hadamard = [][]complex128{
	{HC, HC},
	{HC, HCN},
}

// TODO: I am doing EulersFormula wrong apparently, revisit...
func EF(theta float64) complex128 {
	return cmplx.Pow(math.E, 1i*(complex(theta, 0)))
}
func EulersFormula(theta float64) complex128 {
	cos := cmplx.Cos(complex(theta, 0))
	sin := cmplx.Sin(complex(theta, 0))
	return cos + 1i*sin
}

// TODO: Come back and look at Phase Operations, I'm not doing this righ
// I'm not doing Euler's formula right some how
func PhaseOp(theta float64) [][]complex128 {
	return [][]complex128{
		{1, 0},
		{0, EulersFormula(theta)},
	}
}

// The S gate is also known as the phase gate or the Z90 gate,
// because it represents a 90-degree rotation around the z-axis
// This should be PhaseOp(math.Pi/2), but I apparenntly can't get EulersFormula right
// S should equal T**2
var S = [][]complex128{
	{1, 0},
	{0, 0 + 1i},
}

// The conjugate transpose of the S gate, the S†
var Sdagger = [][]complex128{
	{1, 0},
	{0, 0 - 1i},
}

// The T gate
// should be PhasOp(math.Pi/4)
var T = [][]complex128{
	{1, 0},
	{0, (1 + 1i) / math.Sqrt2},
}

// The conjugate transpose of the T gate is T†
// come back to this as well...
var Tdaggger = [][]complex128{
	{1, 0},
	{0, (1 - 1i) / math.Sqrt2},
}

// H|0> = [[HC, HC], [HC, HCN]]*[1, 0] = [HC, HC] = |+>
var HadamardKet0 = [][]complex128{
	{HC},
	{HC},
}
var HKet0 = HadamardKet0
var KetPlus = HadamardKet0

// H|1> = [[HC, HC], [HC, HCN]]*[0, 1] = [HC, HCN] = |->
var HadamardKet1 = [][]complex128{
	{HC},
	{HCN},
}
var HKet1 = HadamardKet0
var KetMinus = HadamardKet1

var Ket0 = [][]complex128{
	{1},
	{0},
}

// H|+> = [[HC, HC], [HC, HCN]]*[HC, HC] = [1, 0] = |0>
var HadamardKetPlus = Ket0
var HKetPlus = HadamardKetPlus

var Ket1 = [][]complex128{
	{0},
	{1},
}

// H|-> = [[HC, HC], [HC, HCN]]*[HC, HCN] = [0, 1] = |1>
var HadamardKetMinus = Ket1
