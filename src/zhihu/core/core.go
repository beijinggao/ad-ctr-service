package core

import (
	"log"
	"fmt"
	"zhihu/xgboost"
	"sort"
)

type XBooster struct {
	ModelPath string
}

func (self *XBooster) Predict() {
	featureList := NewFeatureList()
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

	var test xgboost.DMatrix
	test = append(test, dataRow)
	testHandle, err := xgboost.XGDMatrixCreateFromMat(test, -1)
	// create the booster and load some parameters
	booster, err := xgboost.XGBoosterCreate([]*xgboost.DMatrixHandle{testHandle})
	booster.LoadModel(self.ModelPath)

	if err != nil {
		log.Fatalln(err)
	}
	result, err := booster.Predict(testHandle, 0, 0)
	if err != nil {
		log.Fatalln(err)
	}
	for i, v := range result {
		fmt.Printf("prediction[%d]=%.2f\n", i, v)
	}

}
