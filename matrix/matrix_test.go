package matrix_test

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/seraphictrip/quantum/matrix"
)

var MatrixTests = []struct {
	m      [][]complex128
	result [][]complex128
}{
	{
		m: [][]complex128{
			{1, 0},
			{0, 1},
		},
		result: matrix.IdentyMatrix(2),
	},
	{
		m: [][]complex128{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		},
		result: matrix.IdentyMatrix(3),
	},
}

func TestMatrix(t *testing.T) {
	for i, e := range MatrixTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := matrix.ConjugateTranspose(e.m)
			if !matrix.Equal(result, e.result) {
				t.Fatal("nope")
			}
		})
	}
}

var CarteseanProductTests = []struct {
	sigma []float64
	gamma []float64

	expected [][]float64
}{
	{sigma: []float64{0}, gamma: []float64{1}, expected: [][]float64{{0, 1}}},
	{
		[]float64{0, 1, 2},
		[]float64{0, 1, 2},
		[][]float64{
			{0, 0}, {0, 1}, {0, 2},
			{1, 0}, {1, 1}, {1, 2},
			{2, 0}, {2, 1}, {2, 2},
		},
	},
	{
		[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		[][]float64{
			{0, 0},
			{0, 1},
			{0, 2},
			{0, 3},
			{0, 4},
			{0, 5},
			{0, 6},
			{0, 7},
			{0, 8},
			{0, 9},
			{1, 0},
			{1, 1},
			{1, 2},
			{1, 3},
			{1, 4},
			{1, 5},
			{1, 6},
			{1, 7},
			{1, 8},
			{1, 9},
			{2, 0},
			{2, 1},
			{2, 2},
			{2, 3},
			{2, 4},
			{2, 5},
			{2, 6},
			{2, 7},
			{2, 8},
			{2, 9},
			{3, 0},
			{3, 1},
			{3, 2},
			{3, 3},
			{3, 4},
			{3, 5},
			{3, 6},
			{3, 7},
			{3, 8},
			{3, 9},
			{4, 0},
			{4, 1},
			{4, 2},
			{4, 3},
			{4, 4},
			{4, 5},
			{4, 6},
			{4, 7},
			{4, 8},
			{4, 9},
			{5, 0},
			{5, 1},
			{5, 2},
			{5, 3},
			{5, 4},
			{5, 5},
			{5, 6},
			{5, 7},
			{5, 8},
			{5, 9},
			{6, 0},
			{6, 1},
			{6, 2},
			{6, 3},
			{6, 4},
			{6, 5},
			{6, 6},
			{6, 7},
			{6, 8},
			{6, 9},
			{7, 0},
			{7, 1},
			{7, 2},
			{7, 3},
			{7, 4},
			{7, 5},
			{7, 6},
			{7, 7},
			{7, 8},
			{7, 9},
			{8, 0},
			{8, 1},
			{8, 2},
			{8, 3},
			{8, 4},
			{8, 5},
			{8, 6},
			{8, 7},
			{8, 8},
			{8, 9},
			{9, 0},
			{9, 1},
			{9, 2},
			{9, 3},
			{9, 4},
			{9, 5},
			{9, 6},
			{9, 7},
			{9, 8},
			{9, 9},
		},
	},
}

func TestCarteseanProduct(t *testing.T) {
	for i, e := range CarteseanProductTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := matrix.CarteseanProduct(e.sigma, e.gamma)
			if len(result) != len(e.expected) {
				t.Fatalf("nope")
			}
			for i, pair := range result {
				if !slices.Equal(pair, e.expected[i]) {
					t.Fatalf("CarteseanProduct(%v, %v) = %v, want %v", e.sigma, e.gamma, result, e.expected)
				}
			}
		})
	}
}

var VectorTensorTests = []struct {
	sigma, gamma, expected []float64
}{
	// 1/6|00⟩ + 1/12|01⟩ + 1/2|10⟩ + 1/4|11⟩
	{[]float64{1 / 4., 3 / 4.}, []float64{2 / 3., 1 / 3.}, []float64{1 / 6., 1 / 12., 1 / 2., 1 / 4.}},
}

func TestVectorTensor(t *testing.T) {
	for i, e := range VectorTensorTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := matrix.VectorTensor(e.sigma, e.gamma)
			if !slices.Equal(result, e.expected) {
				t.Fatalf("got %v, want %v", result, e.expected)
			}
		})
	}
}

func StringCarteseanProduct(sigma, gamma []string) []string {
	result := make([]string, len(sigma)*len(gamma))
	i := 0
	for _, s := range sigma {
		for _, g := range gamma {
			result[i] = s + g
			i++
		}
	}
	return result
}
func StringPermutation(input string) []string {
	// progressively re-apply CarteseanProduct
	v := strings.Split(input, "")
	if len(v) == 0 {
		return nil
	}
	prefixes := v
	for i := 0; i < len(input)-1; i++ {
		prefixes = StringCarteseanProduct(prefixes, v)
	}
	if len(prefixes) != int(math.Pow(float64(len(v)), float64(len(v)))) {
		panic("you know nothing")
	}
	return prefixes
}

func Factorial(n int) (result int) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func swap(str string, x, y int) string {
	if x < 0 || x > len(str)-1 {
		panic("invalid x")
	}
	if y < 0 || y > len(str)-1 {
		panic("invalid y")
	}
	bytes := []byte(str)
	temp := str[x]
	bytes[x] = str[y]
	bytes[y] = temp

	return string(bytes)
}

// abc, acb, bac, bca
func permute(str string, l, r int, agg []string) []string {
	// fmt.Printf("permute(%q, %v, %v, %v)\n", str, l, r, agg)
	fmt.Println(str)
	if l == r {
		agg = append(agg, str)
		return agg
	}
	for i := l; i <= r; i++ {
		str = swap(str, l, i)
		agg = permute(str, l+1, r, agg)
		str = swap(str, l, i)
	}
	return agg
}

// ab => [ab, ba]
// abc => [abc, acb, bac, bca, cba, cab]
func StringPermutationNoRepeat(input string) []string {

	return permute(input, 0, len(input)-1, []string{})
}

var StringPermutationTests = []struct {
	input    string
	expected []string
}{
	{},
	{"ab", []string{"aa", "ab", "ba", "bb"}},
	{"abc", []string{
		"aaa", "aab", "aac",
		"aba", "abb", "abc",
		"aca", "acb", "acc",
		"baa", "bab", "bac",
		"bba", "bbb", "bbc",
		"bca", "bcb", "bcc",
		"caa", "cab", "cac",
		"cba", "cbb", "cbc",
		"cca", "ccb", "ccc",
	}},
}

func TestStringPermutation(t *testing.T) {
	for i, e := range StringPermutationTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := StringPermutation(e.input)
			if !slices.Equal(result, e.expected) {
				t.Fatalf("got %v, want %v", result, e.expected)
			}
		})
	}
}

var StringPermutationNoRepeatTests = []struct {
	input    string
	expected []string
}{
	// {},
	// {"ab", []string{"ab", "ba"}},
	// {"abc", []string{
	// 	"abc", "acb",
	// 	"bac", "bca",
	// 	"cba", "cab",
	// }},
	// {"1234", []string{
	// 	"1234", "1243", "1324", "1342", "1423", "1432", "2134", "2143", "2314", "2341", "2413", "2431", "3124", "3142", "3214", "3241", "3412", "3421", "4123", "4132", "4213", "4231", "4312", "4321",
	// }},
	{"1357", nil},
}

func TestStringPermutationNoRepeat(t *testing.T) {
	for i, e := range StringPermutationNoRepeatTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := StringPermutationNoRepeat(e.input)
			if !slices.Equal(result, e.expected) {
				t.Fatalf("got %v, want %v", result, e.expected)
			}
		})
	}
}

/*
1357

1357
1375
1527
1572
1753
1735

*/
