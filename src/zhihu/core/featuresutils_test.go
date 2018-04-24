package core

import "testing"

func TestFeatureList_Transform2MapFeature(t *testing.T) {
	featureList:=NewFeatureList()
	featureList.Transform2MapFeature(featureList)
}
