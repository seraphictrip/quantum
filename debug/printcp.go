package debug

import "fmt"

func PrintCarteseanProduct(cp [][]float64) {
	fmt.Println("[][]float64{")
	for _, pair := range cp {
		fmt.Printf("{%v, %v},\n", pair[0], pair[1])
	}
	fmt.Println("}")
}
