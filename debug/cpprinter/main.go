package main

import (
	"github.com/seraphictrip/quantum/debug"
	"github.com/seraphictrip/quantum/matrix"
)

func main() {
	debug.PrintCarteseanProduct(matrix.CarteseanProduct([]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
}
