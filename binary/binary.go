package binary

import "fmt"

// This probably isn't what you expected... ha, I need notes

// a | f(a)
// 0 | 0
// 1 | 0
func Const0(alpha complex128) complex128 {
	return 0
}

// a | f(a)
// 0 | 0
// 1 | 1
func Identity(alpha complex128) complex128 {
	return alpha
}

// a | f(a)
// 0 | 1
// 1 | 0
func Not(alpha complex128) complex128 {
	if alpha == 0 {
		return 1
	}
	return 0
}

// a | f(a)
// 0 | 1
// 1 | 1
func Const1(alpha complex128) complex128 {
	return alpha
}

/*
	M|a} = |(f(a)}
*/

var Const0Matrix = NewMatrix([][]complex128{
	{1, 1},
	{0, 0},
})

var IdentityMatrix = NewMatrix([][]complex128{
	{1, 0},
	{0, 1},
})

var NotMatrix = NewMatrix([][]complex128{
	{0, 1},
	{1, 0},
})

var Const1Matrix = NewMatrix([][]complex128{
	{0, 0},
	{1, 1},
})

type Matrix struct {
	data [][]complex128
}

func (m Matrix) Data() [][]complex128 {
	return m.data
}

func (m Matrix) IsNil() bool {
	return len(m.data) == 0
}

// {-3, 0, 3, 2 }  	[2]	[-3*2+ 0*-3+3*4+ 2*1]	[4]
// {1, 7, -1, 9}   	[-3] [1*2+7*-3+-1*4+9*1]	[-32]
//
//	[4]
//	[1]
func (m Matrix) VectorProduct(v []complex128) []complex128 {
	if m.IsNil() {
		return nil
	}
	if len(v) != len(m.data[0]) {
		panic(fmt.Errorf("undefined"))
	}
	data := make([]complex128, len(m.data))
	for i := range m.data {
		for j := range m.data[i] {
			data[i] += m.data[i][j] * v[j]
		}
	}

	return data
}

func NewMatrix(data [][]complex128) Matrix {
	return Matrix{
		data: data,
	}
}
