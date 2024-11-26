package binary

import (
	"slices"
	"strconv"
	"testing"
)

var VectorProductTests = []struct {
	m        Matrix
	v        []complex128
	expected []complex128
}{
	{
		m:        Matrix{},
		v:        []complex128{},
		expected: []complex128{},
	},
	{
		m: NewMatrix([][]complex128{
			{1},
		}),
		v:        []complex128{128},
		expected: []complex128{128},
	},
	{
		m: NewMatrix([][]complex128{
			{1, 0},
			{0, 1},
		}),
		v:        []complex128{128, 256},
		expected: []complex128{128, 256},
	},
	{
		m:        IdentityMatrix,
		v:        []complex128{128, 256},
		expected: []complex128{128, 256},
	},
	{
		m: NewMatrix([][]complex128{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		}),
		v:        []complex128{128, 256, -1},
		expected: []complex128{128, 256, -1},
	},
	{
		m:        NotMatrix,
		v:        []complex128{0, 1},
		expected: []complex128{1, 0},
	},
	{
		m: NewMatrix([][]complex128{
			{-3, 0, 3, 2},
			{1, 7, -1, 9},
		}),
		v:        []complex128{2, -3, 4, -1},
		expected: []complex128{4, -32},
	},
	{
		m:        Const0Matrix,
		v:        []complex128{0, 1},
		expected: []complex128{1, 0},
	},
	{
		m:        Const1Matrix,
		v:        []complex128{0, 1},
		expected: []complex128{0, 1},
	},
	{
		// I get lost here, I think because I put 0 in I should expect 1
		// as that is definition,  but this is actually .5(IdentityMatrix) + .5(Const0)?
		m: NewMatrix([][]complex128{
			{1, .5},
			{0, .5},
		}),
		v:        []complex128{0, 1},
		expected: []complex128{.5, .5},
	},
	{
		m: NewMatrix([][]complex128{
			{1, .25},
			{0, .75},
		}),
		v:        []complex128{0, 1},
		expected: []complex128{.25, .75},
	},
	{
		m: NewMatrix([][]complex128{
			{1, .5, .25},
			{0, .5, .75},
		}),
		v:        []complex128{0, 1, 0},
		expected: []complex128{.5, .5},
	},
	{
		m: NewMatrix([][]complex128{
			{1, .5, .75},
			{0, .5, .25},
		}),
		v:        []complex128{0, 1, 0},
		expected: []complex128{.5, .5},
	},
}

func TestVectorProduct(t *testing.T) {
	for i, e := range VectorProductTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			output := e.m.VectorProduct(e.v)

			if !slices.Equal(output, e.expected) {
				t.Fatalf("m.VectorProduct(%v) = %v, want %v", e.v, output, e.expected)
			}
		})
	}
}
