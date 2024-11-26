package matrix

import (
	"fmt"
	"math/cmplx"
	"slices"

	"github.com/seraphictrip/quantum"
)

func NewMatrix(rows, cols int) [][]complex128 {
	matrix := make([][]complex128, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]complex128, cols)
	}
	return matrix
}

// Unitary Martices are square Matrices if
// U*ConjagateTranspose(U) = IdenityMaritx
// ConjagateTranspose(U)*U = IdentityMatrix
// U† is conjugate transpose of U
func IdentyMatrix(dim int) [][]complex128 {
	matrix := make([][]complex128, dim)

	for i := 0; i < dim; i++ {
		matrix[i] = make([]complex128, dim)
		for j := 0; j < dim; j++ {
			if i == j {
				matrix[i][j] = 1
			} else {
				matrix[i][j] = 0
			}
		}
	}
	return matrix
}

func ConjugateTranspose(src [][]complex128) [][]complex128 {
	d0 := len(src)
	d1 := len(src[0])

	matrix := NewMatrix(d1, d0)
	for i, row := range src {
		for j, c := range row {
			matrix[j][i] = cmplx.Conj(c)
		}
	}
	return matrix
}

func Equal(m1, m2 [][]complex128) bool {
	// TODO: I should add dim checks
	for i, row := range m1 {
		if !slices.Equal(row, m2[i]) {
			return false
		}
	}
	return true
}

// func TensorProduct(m1, m2 [][]complex128) [][]complex128 {
// 	m1Rows := len(m1)
// 	m2Rows := len(m2)
// 	if m1Rows == 0 {
// 		return m1
// 	}
// 	if m2Rows == 0 {
// 		return m2
// 	}
// 	matrix := NewMatrix(m1Rows*m2Rows, m1Rows*m2Rows)

// 	return matrix
// }

/*
	// Cartesean Product (view sets together)
	[1] × [0] = [(1, 0)] = [10]
	[1,2] × [1, 2] = [(1,1), (1,2), (2,1), (2,2)] = [11, 12, 21, 22]

	[0,1] × [0,1] = [(0,0), (0,1), (1,0), (1,1)] = [00, 01, 10, 11]

	012 × 01 = 00, 01, 10, 11, 20, 21

	[0..9] × [0..9] = [01,02,03,04...50,51,52...97,98,99]

	[0,1] × [0,1] × [0,1] = [(0,0,0), (0,0,1), (0,1,0), (0,1,1), (1,0,0), (1,0,1), (1,1,0),(1,1,1)], [000, 001, 010, 011, 100, 101, 110, 111]

	// Probablistic Cartesean Product to combine systems
	// Given we have two Probablistic state vectors
	[.1,.9] × [.4,.6] = [(.1,.4), (.1,.6),(.9,.4),(.9,.6)] =??? [.04, .06, .36, .54]
*/
// Currently I'm just doing CarteseanProduct of 2 vectors, can be further generalized..
func CarteseanProduct(sigma, gamma []float64) [][]float64 {
	result := make([][]float64, 0)
	index := 0
	// we could maek this more generic...
	for _, s := range sigma {
		for _, g := range gamma {
			result = append(result, []float64{s, g})
			index++
		}
	}
	return result
}

func VectorTensor(sigma, gamma []float64) []float64 {
	pairs := CarteseanProduct(sigma, gamma)
	vt := make([]float64, len(pairs))
	for i, pair := range pairs {
		vt[i] = pair[0] * pair[1]
	}

	return vt
}

func AreIndependent(pSigma, pGamma []float64) bool {
	if !quantum.IsProbablisticStateVector(pSigma) || !quantum.IsProbablisticStateVector(pGamma) {
		panic("inputs are not probablistic state vectors")
	}
	pairs := CarteseanProduct(pSigma, pGamma)
	pv := make([]float64, 0)
	for _, pair := range pairs {
		pv = append(pv, pair[0]*pair[1])
	}
	fmt.Println(pv)

	return false
}

/**
Pr((X, Y) = (a,b)) = Pr(X=a)Pr(Y=b)
Pr(X ∩ Y) = Pr(a)Pr(B)
for every choice of a ∈ Σ and b ∈ Γ.

1/6|00⟩ + 1/12|01⟩ + 1/2|10⟩ + 1/4|11⟩

Pr(A=0) = 1/6*1/12 = 3/12 = 1/4
Pr(A=1) = 1/2*1/4 = 3/4
|Phi⟩ = Pr(A=0) + Pr(A=1) = [1/4, 3/4]
Pr(B=0) = 1/6*1/2 = 4/6 = 2/3
Pr(B=1) = 1/12*1/4 = 4/12 = 1/3
|Psi⟩ = Pr(B=0) + Pr(B=1) = [2/3, 1/3]

P(A|B) = P(A ∩ B)/P(B) =

1/2|00⟩ + 0|01⟩ +0|10⟩  + 1/2|11⟩ = [1/2, 0, 0, 1/2]
Pr(A=0) =
Pr(A=1) =
Pr(B=0) =
Pr(B=0) =
Pr(A|B) =

*/
