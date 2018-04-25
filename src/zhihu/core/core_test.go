package core

import "testing"

func TestXBooster_Predict(t *testing.T) {
	booster := new(XBooster)
	booster.ModelPath = "../resources/models/model_20180412_Pos8.bin"
	booster.LoadModel()
	booster.Predict()

}
