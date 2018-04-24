package core

const (
	_                 FeatureType = iota
	ContinueValue
	OnehotValue
	VectorValue
	XgbNode
	OnehotVectorValue
)

type FeatureType int

type Feature struct {
	Name             string
	FType            FeatureType
	CategoricalInfo  int
	VectorInfo       int // 是否是向量 0 不是 向量，其他的代表向量长度，严格非负
	OrgId            int
	OnehotVectorInfo int
}

func NewFeature(name string, ftype FeatureType, categoricalInfo int, vectorInfo int, orgId int, onehotVectorInfo int) *Feature {
	return &Feature{
		Name:             name,
		FType:            ftype,
		CategoricalInfo:  categoricalInfo,
		VectorInfo:       vectorInfo,
		OrgId:            orgId,
		OnehotVectorInfo: onehotVectorInfo,
	}

}

/*type AdZoneId struct {
	Feature
}

func NewAdZoneId() *AdZoneId {
	feature := NewFeature("AdZoneId", OnehotValue, 16, 0, 0, 0)
	return &AdZoneId{
		Feature: *feature,
	}
}*/
