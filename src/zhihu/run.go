package main

import (
	"zhihu/core"
	"fmt"
	"sort"
	"zhihu/xgboost"
)

func main() {
	modelPath := "./resources/models/model_20180412_Pos8.bin"

	var inference xgboost.DMatrix

	booster := core.NewXBooster(modelPath)
	featureList := core.NewFeatureList()
	res := featureList.Transform2MapFeature(featureList)
	fmt.Println(res)
	sortedKeys := make([]int, 0, len(res))
	for k := range res {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Ints(sortedKeys)
	var dataRow []float32
	for _, k := range sortedKeys {
		fmt.Printf("%d:%v ", k, res[k])
		dataRow = append(dataRow, res[k])
	}
	inference = append(inference, dataRow)

	booster.Predict(inference)
	booster.Free()

}
