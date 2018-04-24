package core

import (
	"reflect"
	"fmt"
)

type FeatureList struct {
	//1-30
	AdZoneId Feature
}

func NewFeatureList() *FeatureList {
	adZoneId := NewFeature("AdZoneId", OnehotValue, 16, 0, 0, 0)
	featureList := &FeatureList{
		AdZoneId: *adZoneId,
	}
	return featureList
}

func (self *FeatureList) Transform2MapFeature(featureList *FeatureList) {
	//sparseFeature:=make(map[int]float32)

	t := reflect.TypeOf(*self)
	v := reflect.ValueOf(*self)

	for k := 0; k < t.NumField(); k++ {
		feature := v.Field(k)

		fmt.Println(feature.Interface().(Feature).Name)
		fmt.Println(feature.Interface().(Feature).CategoricalInfo)
		fmt.Println(feature.Interface().(Feature).FType)
	}

	/*val sparseFeature = new java.util.HashMap[java.lang.Integer, java.lang.Double]()
	var currentIndex = 1
	(categoryInfos.indices zip featureList).foreach(
		x => {
	if (x._2.getType == featureType.VectorValue.id){
	x._2.computeVector(matchedItem) match {
	case None =>
	currentIndex += vectorInfos(x._1)
	case Some(i) =>
	if (vectorInfos(x._1) == i.length && i.length>=1){
	i.indices.foreach(r =>
	sparseFeature.put(r + currentIndex, i(r))
	)
	}
	currentIndex += vectorInfos(x._1)
	}
	}
	else if (x._2.getType ==featureType.OnehotValue.id){
	val OnehotTmp = x._2.compute(matchedItem)
	if (OnehotTmp.nonEmpty){
	val hottedIdx = OnehotTmp.getOrElse(0.0).toInt
	if (0<=hottedIdx && hottedIdx<categoryInfos(x._1)){
	sparseFeature.put(hottedIdx + currentIndex, 1.0)
	}
	}
	currentIndex += categoryInfos(x._1)
	}
	else if (x._2.getType ==featureType.ContinueValue.id){
	if (x._2.compute(matchedItem).nonEmpty){
	sparseFeature.put(currentIndex,x._2.compute(matchedItem).get)
	}
	currentIndex += 1
	}
	else if (x._2.getType ==featureType.XgbNode.id){
	if (x._2.computeVector(matchedItem).nonEmpty){
	val NodeIndx = x._2.computeVector(matchedItem).getOrElse(Array.empty[Double])
	var tmpcurrentIndex = currentIndex
	NodeIndx.foreach(
	r => {
	sparseFeature.put(r.toInt + tmpcurrentIndex, 1.0)
	tmpcurrentIndex = tmpcurrentIndex + treeLeafSize
	}
	)
	currentIndex = currentIndex + treeLeafSize*treeNum
	}
	else{
	currentIndex = currentIndex + treeLeafSize*treeNum
	}
	}
	else if (x._2.getType ==featureType.OnehotVectorValue.id){
	if (x._2.computeVector(matchedItem).nonEmpty){
	val NodeIndx = x._2.computeVector(matchedItem).getOrElse(Array.empty[Double])
	NodeIndx.distinct.foreach(
	r => {
	sparseFeature.put(r.toInt + currentIndex, 1.0)
	}
	)
	currentIndex = currentIndex + x._2.getOnehotVectorInfo
	}
	else{
	sparseFeature.put( currentIndex, 1.0)
	currentIndex = currentIndex + x._2.getOnehotVectorInfo
	}
	}
	}
)
	sparseFeature*/

}
