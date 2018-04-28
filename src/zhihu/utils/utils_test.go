package utils

import (
	"testing"
	"fmt"
	"math"
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

	aa := make(map[int]float64)
	bb := make(map[int]float64)

	aa[1] = 0.1
	aa[2] = 0.2
	bb[1] = 0.3
	bb[2] = 0.4
	bb[3] = 0.5

	value := 0.0
	for key := range bb {
		//fmt.Println(aa[key])
		if val, ok := aa[key]; ok {
			fmt.Println(val)
			value = math.Max(value, val)
		}
	}
	fmt.Println("*************")
	fmt.Println(value)

}
