package main

import "zhihu/core"

func main() {
	booster := new(core.XBooster)
	booster.ModelPath = "./resources/models/model_20180412_Pos8.bin"
	//booster.LoadModel()
	booster.Predict()

}
