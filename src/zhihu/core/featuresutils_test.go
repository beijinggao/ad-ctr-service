package core

import (
	"testing"
	"fmt"
	"sort"
)

func TestFeatureList_Transform2MapFeature(t *testing.T) {
	featureList := NewFeatureList()
	res := featureList.Transform2MapFeature()
	fmt.Println(res)
	sortedKeys := make([]int, 0, len(res))
	for k := range res {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Ints(sortedKeys)
	for _, k := range sortedKeys {
		fmt.Printf("%d:%v ", k, res[k])
	}
}
