package utils

import (
	"testing"
	"fmt"
)

func TestCosineSimilarity(t *testing.T) {
	A := []float64{1, 2, 3}
	B := []float64{2, 3, 4}
	res, _ := CosineSimilarity(A, B)
	fmt.Println(res)
	A = []float64{0}
	B = []float64{0}
	res, _ = CosineSimilarity(A, B)
	fmt.Println(res)

}
