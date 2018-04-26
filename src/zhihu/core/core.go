package core

import (
	"log"
	"fmt"
	"zhihu/xgboost"
)

type XBooster struct {
	ModelPath      string
	BoosterHandler xgboost.BoosterHandle
}

func NewXBooster(modelPath string) *XBooster {
	booster, _ := xgboost.XGBoosterCreateNil()
	booster.SetParam("seed", "0")
	err := booster.LoadModel(modelPath)
	if err != nil {
		log.Fatalln(err)
	}
	return &XBooster{
		ModelPath:      modelPath,
		BoosterHandler: *booster,
	}
}

func (self *XBooster) Predict(data xgboost.DMatrix) {
	testHandle, err := xgboost.XGDMatrixCreateFromMat(data, -1)
	result, err := self.BoosterHandler.Predict(testHandle, 0, 0)
	if err != nil {
		log.Fatalln(err)
	}
	for i, v := range result {
		fmt.Printf("prediction[%d]=%.2f\n", i, v)
	}
	testHandle.Free()

}

func (self *XBooster) Free() {
	self.BoosterHandler.Free()
}
