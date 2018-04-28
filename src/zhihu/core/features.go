package core

const (
	_                 FeatureType = iota
	ContinueValue
	OnehotValue
	VectorValue
	XgbNode
	OnehotVectorValue
)

const (
	_                     AdZoneType = iota
	WEB_PAGE_TOP
	WEB_PATE_BOTTOM
	INNER_WEB_PAGE_TOP
	INNER_WEB_PAGE_BOTTOM
	WEB_FEED
	APP_FEED
	MOBILE_WEB_BANNER
)

type AdZoneType int

type FeatureType int

type Feature struct {
	Name             string
	FType            FeatureType
	CategoricalInfo  int
	VectorInfo       int // 是否是向量 0 不是 向量，其他的代表向量长度，严格非负
	OrgId            int
	OnehotVectorInfo int
	Value            interface{}
}

func NewFeature(name string, ftype FeatureType, categoricalInfo int, vectorInfo int, orgId int, onehotVectorInfo int, value interface{}) *Feature {
	return &Feature{
		Name:             name,
		FType:            ftype,
		CategoricalInfo:  categoricalInfo,
		VectorInfo:       vectorInfo,
		OrgId:            orgId,
		OnehotVectorInfo: onehotVectorInfo,
		Value:            value,
	}

}
