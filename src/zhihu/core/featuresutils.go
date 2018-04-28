package core

import (
	"reflect"
	"math"
)

func NewFeatureListNil() *FeatureList {
	return &FeatureList{}
}

// SetAdZoneId 广告位id 实时传递 8号位=8 30号位=2
func (self *FeatureList) SetAdZoneId(value int) {
	adZoneId := NewFeature("AdZoneId", OnehotValue, 30, 0, 0, 0, value)
	self.AdZoneId = *adZoneId
}

// SetCTR 广告历史CTR 实时传递click view imp数
func (self *FeatureList) SetCTR(clickNum, viewNum, impNum int) {
	showNum := 0
	if self.AdZoneId.Value.(int) == int(APP_FEED) {
		showNum = viewNum
	} else {
		showNum = impNum
	}
	smoothParam_alpha := 0.1
	smoothParam_beta := 10
	historyCtr := (float64(clickNum) + smoothParam_alpha) / float64(showNum+smoothParam_beta)
	ctr := NewFeature("CTR", ContinueValue, 0, 0, 0, 0, historyCtr)
	self.CTR = *ctr
}

// SetHasTarget 是否定向实时传递 0/1
func (self *FeatureList) SetHasTarget(value int) {
	hasTarget := NewFeature("HasTarget", ContinueValue, 0, 0, 0, 0, value)
	self.HasTarget = *hasTarget
}

// SetTerminalFeature 终端 1:DEFAULT、2:IOS 3:ANDROID、4:MAC_OS_X 5:WINDOWS
func (self *FeatureList) SetTerminalFeature(value int) {
	terminalFeature := NewFeature("TerminalFeature", OnehotValue, 5, 0, 0, 0, value)
	self.TerminalFeature = *terminalFeature
}

// SetAvgTopicInterestScore 用户话题和广告话题求交，结果的平均分数 参数map key 为topicid,value为权重
func (self *FeatureList) SetAvgTopicInterestScore(adtopicIds, usertopicids map[int]float64) {
	sum := 0.0
	count := 0.0
	value := 0.0
	for key := range adtopicIds {
		if val, ok := usertopicids[key]; ok {
			sum += val
			count += 1
		}
	}
	if count != 0.0 {
		value = sum / count
	}

	avgTopicInterestScore := NewFeature("AvgTopicInterestScore", ContinueValue, 0, 0, 0, 0, value)
	self.AvgTopicInterestScore = *avgTopicInterestScore
}

// SetMaxTopicInterestScore 用户话题和广告话题求交，结果的最大值 参数map key 为topicid,value为权重
func (self *FeatureList) SetMaxTopicInterestScore(adtopicIds, usertopicids map[int]float64) {
	value := 0.0
	for key := range adtopicIds {
		if val, ok := usertopicids[key]; ok {
			value = math.Max(value, val)
		}
	}
	maxTopicInterestScore := NewFeature("MaxTopicInterestScore", ContinueValue, 0, 0, 0, 0, value)
	self.MaxTopicInterestScore = *maxTopicInterestScore
}

// SetShowNumFeature 分广告位用户看过广告的次数（当天数据加周数据）from userActionJedis read  weekUserCrShow
// zoneJedisPools zoneDayUserCrShow value=zoneDayUserCrShow+weekUserCrShow
func (self *FeatureList) SetShowNumFeature(value int) {
	showNumFeature := NewFeature("ShowNumFeature", ContinueValue, 0, 0, 0, 0, value)
	self.ShowNumFeature = *showNumFeature
}

// SetDisplayHours 展现小时 实时获取 0-23
func (self *FeatureList) SetDisplayHours(value int) {
	displayHours := NewFeature("DisplayHours", OnehotValue, 24, 0, 0, 0, value)
	self.DisplayHours = *displayHours
}

// SetClickFeature 分广告位用户点广告的次数（当天数据加周数据）zoneDayUserCrClick + weekUserCrClick
func (self *FeatureList) SetClickFeature(value int) {
	clickFeature := NewFeature("ClickFeature", ContinueValue, 0, 0, 0, 0, value)
	self.ClickFeature = *clickFeature
}

// SetDayUserCrShow from zoneJedisPools read zoneDayUserCrShow
func (self *FeatureList) SetDayUserCrShow(value int) {
	dayUserCrShow := NewFeature("DayUserCrShow", ContinueValue, 0, 0, 0, 0, value)
	self.DayUserCrShow = *dayUserCrShow
}

// SetDayUserCrClick from zoneJedisPools read zoneDayUserCrClick
func (self *FeatureList) SetDayUserCrClick(value interface{}) {
	dayUserCrClick := NewFeature("DayUserCrClick", ContinueValue, 0, 0, 0, 0, value)
	self.DayUserCrClick = *dayUserCrClick
}
func (self *FeatureList) SetDayUserShow(value int) {
	dayUserShow := NewFeature("DayUserShow", ContinueValue, 0, 0, 0, 0, value)
	self.DayUserShow = *dayUserShow
}
func (self *FeatureList) SetDayUserClick(value interface{}) {
	dayUserClick := NewFeature("DayUserClick", ContinueValue, 0, 0, 0, 0, value)
	self.DayUserClick = *dayUserClick
}
func (self *FeatureList) SetDayUserAdShow(value interface{}) {
	dayUserAdShow := NewFeature("DayUserAdShow", ContinueValue, 0, 0, 0, 0, value)
	self.DayUserAdShow = *dayUserAdShow
}
func (self *FeatureList) SetDayUserAdClick(value interface{}) {
	dayUserAdClick := NewFeature("DayUserAdClick", ContinueValue, 0, 0, 0, 0, value)
	self.DayUserAdClick = *dayUserAdClick
}
func (self *FeatureList) SetUserInterests(value interface{}) {
	userInterests := NewFeature("UserInterests", OnehotValue, 19, 0, 0, 0, value)
	self.UserInterests = *userInterests
}
func (self *FeatureList) SetIndustry(value interface{}) {
	industry := NewFeature("Industry", OnehotValue, 19, 0, 0, 0, value)
	self.Industry = *industry
}
func (self *FeatureList) SetUserIndustrySim(value interface{}) {
	userIndustrySim := NewFeature("UserIndustrySim", ContinueValue, 0, 0, 0, 0, value)
	self.UserIndustrySim = *userIndustrySim
}
func (self *FeatureList) SetAnsIndustry(value interface{}) {
	ansIndustry := NewFeature("AnsIndustry", OnehotValue, 19, 0, 0, 0, value)
	self.UserIndustrySim = *ansIndustry
}
func (self *FeatureList) SetUserAdidSim(value interface{}) {
	userAdidSim := NewFeature("UserAdidSim", ContinueValue, 0, 0, 0, 0, value)
	self.UserAdidSim = *userAdidSim
}
func (self *FeatureList) SetAnsAdid(value interface{}) {
	ansAdid := NewFeature("AnsAdid", ContinueValue, 0, 0, 0, 0, value)
	self.AnsAdid = *ansAdid
}
func (self *FeatureList) SetClickIndustry(value interface{}) {
	clickIndustry := NewFeature("ClickIndustry", ContinueValue, 0, 0, 0, 0, value)
	self.ClickIndustry = *clickIndustry
}
func (self *FeatureList) SetDisplayWeeks(value interface{}) {
	displayWeeks := NewFeature("DisplayWeeks", OnehotValue, 8, 0, 0, 0, value)
	self.DisplayWeeks = *displayWeeks
}
func (self *FeatureList) SetTemplateFeature(value interface{}) {
	templateFeature := NewFeature("TemplateFeature", OnehotValue, 100, 0, 0, 0, value)
	self.TemplateFeature = *templateFeature
}
func (self *FeatureList) SetSimsMultiple(value interface{}) {
	simsMultiple := NewFeature("SimsMultiple", VectorValue, 0, 8, 0, 0, value)
	self.SimsMultiple = *simsMultiple
}
func (self *FeatureList) SetRecallSim(value interface{}) {
	recallSim := NewFeature("RecallSim", ContinueValue, 0, 0, 0, 0, value)
	self.RecallSim = *recallSim
}
func (self *FeatureList) SetSimIsEmpty(value interface{}) {
	simIsEmpty := NewFeature("SimIsEmpty", VectorValue, 0, 3, 0, 0, value)
	self.SimIsEmpty = *simIsEmpty
}
func (self *FeatureList) SetClusterCtr(value interface{}) {
	clusterCtr := NewFeature("ClusterCtr", ContinueValue, 0, 0, 0, 0, value)
	self.ClusterCtr = *clusterCtr
}
func (self *FeatureList) SetUserVectorByAls(value interface{}) {
	userVectorByAls := NewFeature("UserVectorByAls", VectorValue, 0, 16, 0, 0, value)
	self.UserVectorByAls = *userVectorByAls
}
func (self *FeatureList) SetUserVectorIsEmpty(value interface{}) {
	userVectorIsEmpty := NewFeature("UserVectorIsEmpty", ContinueValue, 0, 0, 0, 0, value)
	self.UserVectorIsEmpty = *userVectorIsEmpty
}
func (self *FeatureList) SetCreativeBrandSim(value interface{}) {
	creativeBrandSim := NewFeature("CreativeBrandSim", VectorValue, 0, 3, 0, 0, value)
	self.CreativeBrandSim = *creativeBrandSim
}
func (self *FeatureList) SetNegativeKeyWordSim(value interface{}) {
	negativeKeyWordSim := NewFeature("NegativeKeyWordSim", ContinueValue, 0, 0, 0, 0, value)
	self.NegativeKeyWordSim = *negativeKeyWordSim
}
func (self *FeatureList) SetPostiveKeyWordSim(value interface{}) {
	postiveKeyWordSim := NewFeature("PostiveKeyWordSim", ContinueValue, 0, 0, 0, 0, value)
	self.PostiveKeyWordSim = *postiveKeyWordSim
}
func (self *FeatureList) SetNegativeUserAdTopicSim(value interface{}) {
	negativeUserAdTopicSim := NewFeature("NegativeUserAdTopicSim", ContinueValue, 0, 0, 0, 0, value)
	self.NegativeUserAdTopicSim = *negativeUserAdTopicSim
}
func (self *FeatureList) SetPostiveUserAdTopicSim(value interface{}) {
	postiveUserAdTopicSim := NewFeature("PostiveUserAdTopicSim", ContinueValue, 0, 0, 0, 0, value)
	self.PostiveUserAdTopicSim = *postiveUserAdTopicSim
}
func (self *FeatureList) SetBrushnum(value interface{}) {
	brushnum := NewFeature("Brushnum", ContinueValue, 0, 0, 0, 0, value)
	self.Brushnum = *brushnum
}
func (self *FeatureList) SetCtrPro(value interface{}) {
	ctrPro := NewFeature("CtrPro", VectorValue, 0, 4, 0, 0, value)
	self.CtrPro = *ctrPro
}
func (self *FeatureList) SetViewClickNum(value interface{}) {
	viewClickNum := NewFeature("ViewClickNum", VectorValue, 0, 4, 0, 0, value)
	self.ViewClickNum = *viewClickNum
}
func (self *FeatureList) SetPageHomeType(value interface{}) {
	pageHomeType := NewFeature("PageHomeType", OnehotValue, 8, 0, 0, 0, value)
	self.PageHomeType = *pageHomeType
}
func (self *FeatureList) SetUserAdTagOneHotVector(value interface{}) {
	userAdTagOneHotVector := NewFeature("UserAdTagOneHotVector", OnehotVectorValue, 0, 0, 0, 69, value)
	self.UserAdTagOneHotVector = *userAdTagOneHotVector
}

type FeatureList struct {
	//1-30
	AdZoneId Feature
	//31,32
	CTR       Feature
	HasTarget Feature
	//33-37
	TerminalFeature Feature
	//38,39
	AvgTopicInterestScore Feature
	MaxTopicInterestScore Feature
	//40
	ShowNumFeature Feature
	//41-64
	DisplayHours Feature
	//65
	ClickFeature Feature
	//66,67
	DayUserCrShow  Feature
	DayUserCrClick Feature
	//68,69
	DayUserShow  Feature
	DayUserClick Feature
	//70,71
	DayUserAdShow  Feature
	DayUserAdClick Feature
	//72-90
	UserInterests Feature
	//91-109 客户行业
	Industry Feature
	// 110 用户兴趣和客户行业的相似度
	UserIndustrySim Feature
	//answer 行业向量 111 -129
	AnsIndustry Feature
	// uid - adid 相似度 130
	UserAdidSim Feature
	// awid - adid 相似度 131
	AnsAdid Feature
	// clickIndustry 132
	ClickIndustry Feature

	// 周特征 133 - 140
	DisplayWeeks Feature
	// 模板特征 141-240
	TemplateFeature Feature

	// 241 - 248
	SimsMultiple Feature
	// 249
	RecallSim Feature
	// 250-252
	SimIsEmpty Feature
	// 253
	ClusterCtr Feature
	// 254-269
	UserVectorByAls Feature
	// 270
	UserVectorIsEmpty Feature
	// 271-273
	CreativeBrandSim Feature
	//274
	NegativeKeyWordSim Feature
	//275
	PostiveKeyWordSim Feature
	//276
	NegativeUserAdTopicSim Feature
	//277
	PostiveUserAdTopicSim Feature
	// 278
	Brushnum Feature
	// 279-282 ctr
	CtrPro Feature
	// 283-286
	ViewClickNum Feature
	// 287-294
	PageHomeType Feature
	// 295-363
	UserAdTagOneHotVector Feature
}

func NewFeatureList() *FeatureList {
	//0 9:1 31:0.007049 32:1 35:1 40:0 56:1 65:0 66:0 67:0 68:3 69:0 70:0 71:0 130:0.024490 131:0 132:0 134:1 148:1 241:0 242:0.031287 243:0.060775 2
	//44:0 245:0 246:0 247:0.058366 248:0 250:384 251:0 252:128 253:0 270:0 271:0 272:0 273:0 274:0.285462 275:0.036747 276:0 277:0.408344 278:3 279:
	//0.005441 280:0.005441 281:0 282:0 283:0 284:0 285:0 286:0 295:0 299:1 305:1 309:1 353:1

	adZoneId := NewFeature("AdZoneId", OnehotValue, 30, 0, 0, 0, 9)
	ctr := NewFeature("CTR", ContinueValue, 0, 0, 0, 0, 0.007049)
	hasTarget := NewFeature("HasTarget", ContinueValue, 0, 0, 0, 0, 1)
	terminalFeature := NewFeature("TerminalFeature", OnehotValue, 5, 0, 0, 0, 3)
	avgTopicInterestScore := NewFeature("AvgTopicInterestScore", ContinueValue, 0, 0, 0, 0, 0)
	maxTopicInterestScore := NewFeature("MaxTopicInterestScore", ContinueValue, 0, 0, 0, 0, 0)
	showNumFeature := NewFeature("ShowNumFeature", ContinueValue, 0, 0, 0, 0, 0)
	displayHours := NewFeature("DisplayHours", OnehotValue, 24, 0, 0, 0, 16)
	clickFeature := NewFeature("ClickFeature", ContinueValue, 0, 0, 0, 0, 0)
	dayUserCrShow := NewFeature("DayUserCrShow", ContinueValue, 0, 0, 0, 0, 0)
	dayUserCrClick := NewFeature("DayUserCrClick", ContinueValue, 0, 0, 0, 0, 0)
	dayUserShow := NewFeature("DayUserShow", ContinueValue, 0, 0, 0, 0, 3)
	dayUserClick := NewFeature("DayUserClick", ContinueValue, 0, 0, 0, 0, 0)
	dayUserAdShow := NewFeature("DayUserAdShow", ContinueValue, 0, 0, 0, 0, 0)
	dayUserAdClick := NewFeature("DayUserAdClick", ContinueValue, 0, 0, 0, 0, 0)
	userInterests := NewFeature("UserInterests", OnehotValue, 19, 0, 0, 0, nil)
	industry := NewFeature("Industry", OnehotValue, 19, 0, 0, 0, nil)
	userIndustrySim := NewFeature("UserIndustrySim", ContinueValue, 0, 0, 0, 0, nil)
	ansIndustry := NewFeature("AnsIndustry", OnehotValue, 19, 0, 0, 0, nil)
	userAdidSim := NewFeature("UserAdidSim", ContinueValue, 0, 0, 0, 0, 0.024490)
	ansAdid := NewFeature("AnsAdid", ContinueValue, 0, 0, 0, 0, 0)
	clickIndustry := NewFeature("ClickIndustry", ContinueValue, 0, 0, 0, 0, 0)
	displayWeeks := NewFeature("DisplayWeeks", OnehotValue, 8, 0, 0, 0, 2)
	templateFeature := NewFeature("TemplateFeature", OnehotValue, 100, 0, 0, 0, 8)
	simsMultiple := NewFeature("SimsMultiple", VectorValue, 0, 8, 0, 0, []float32{0, 0.031287, 0.060775, 0, 0, 0, 0.058366, 0})
	recallSim := NewFeature("RecallSim", ContinueValue, 0, 0, 0, 0, 0)
	simIsEmpty := NewFeature("SimIsEmpty", VectorValue, 0, 3, 0, 0, []float32{384, 0, 128})
	clusterCtr := NewFeature("ClusterCtr", ContinueValue, 0, 0, 0, 0, 0)
	userVectorByAls := NewFeature("UserVectorByAls", VectorValue, 0, 16, 0, 0, nil)
	userVectorIsEmpty := NewFeature("UserVectorIsEmpty", ContinueValue, 0, 0, 0, 0, 0)
	creativeBrandSim := NewFeature("CreativeBrandSim", VectorValue, 0, 3, 0, 0, []float32{0, 0, 0})
	negativeKeyWordSim := NewFeature("NegativeKeyWordSim", ContinueValue, 0, 0, 0, 0, 0.285462)
	postiveKeyWordSim := NewFeature("PostiveKeyWordSim", ContinueValue, 0, 0, 0, 0, 0.036747)
	negativeUserAdTopicSim := NewFeature("NegativeUserAdTopicSim", ContinueValue, 0, 0, 0, 0, 0)
	postiveUserAdTopicSim := NewFeature("PostiveUserAdTopicSim", ContinueValue, 0, 0, 0, 0, 0.408344)
	brushnum := NewFeature("Brushnum", ContinueValue, 0, 0, 0, 0, 3)
	ctrPro := NewFeature("CtrPro", VectorValue, 0, 4, 0, 0, []float32{0.005441, 0.005441, 0, 0})
	viewClickNum := NewFeature("ViewClickNum", VectorValue, 0, 4, 0, 0, []float32{0, 0, 0, 0})
	pageHomeType := NewFeature("PageHomeType", OnehotValue, 8, 0, 0, 0, nil)
	userAdTagOneHotVector := NewFeature("UserAdTagOneHotVector", OnehotVectorValue, 0, 0, 0, 69, []float32{1, 5, 11, 15, 59})

	featureList := &FeatureList{
		AdZoneId:               *adZoneId,
		CTR:                    *ctr,
		HasTarget:              *hasTarget,
		TerminalFeature:        *terminalFeature,
		AvgTopicInterestScore:  *avgTopicInterestScore,
		MaxTopicInterestScore:  *maxTopicInterestScore,
		ShowNumFeature:         *showNumFeature,
		DisplayHours:           *displayHours,
		ClickFeature:           *clickFeature,
		DayUserCrShow:          *dayUserCrShow,
		DayUserCrClick:         *dayUserCrClick,
		DayUserShow:            *dayUserShow,
		DayUserClick:           *dayUserClick,
		DayUserAdShow:          *dayUserAdShow,
		DayUserAdClick:         *dayUserAdClick,
		UserInterests:          *userInterests,
		Industry:               *industry,
		UserIndustrySim:        *userIndustrySim,
		AnsIndustry:            *ansIndustry,
		UserAdidSim:            *userAdidSim,
		AnsAdid:                *ansAdid,
		ClickIndustry:          *clickIndustry,
		DisplayWeeks:           *displayWeeks,
		TemplateFeature:        *templateFeature,
		SimsMultiple:           *simsMultiple,
		RecallSim:              *recallSim,
		SimIsEmpty:             *simIsEmpty,
		ClusterCtr:             *clusterCtr,
		UserVectorByAls:        *userVectorByAls,
		UserVectorIsEmpty:      *userVectorIsEmpty,
		CreativeBrandSim:       *creativeBrandSim,
		NegativeKeyWordSim:     *negativeKeyWordSim,
		PostiveKeyWordSim:      *postiveKeyWordSim,
		NegativeUserAdTopicSim: *negativeUserAdTopicSim,
		PostiveUserAdTopicSim:  *postiveUserAdTopicSim,
		Brushnum:               *brushnum,
		CtrPro:                 *ctrPro,
		ViewClickNum:           *viewClickNum,
		PageHomeType:           *pageHomeType,
		UserAdTagOneHotVector:  *userAdTagOneHotVector,
	}
	return featureList
}

func (self *FeatureList) Transform2MapFeature() map[int]float32 {
	sparseFeature := make(map[int]float32)
	t := reflect.TypeOf(*self)
	v := reflect.ValueOf(*self)

	currentIndex := 0

	for k := 0; k < t.NumField(); k++ {
		feature := v.Field(k).Interface().(Feature)

		if feature.FType == VectorValue {
			switch feature.Value.(type) {
			case []float32:
				for index, val := range feature.Value.([]float32) {
					sparseFeature[index+currentIndex+1] = val
				}
				currentIndex += feature.VectorInfo
			default:
				currentIndex += feature.VectorInfo
			}

		} else if feature.FType == OnehotValue {
			switch feature.Value.(type) {
			case float32:
				index := feature.Value.(float32)
				sparseFeature[int(index)+currentIndex] = 1.0
				currentIndex += feature.CategoricalInfo
			case float64:
				index := feature.Value.(float64)
				sparseFeature[int(index)+currentIndex] = 1.0
				currentIndex += feature.CategoricalInfo
			case int:
				index := feature.Value.(int)
				sparseFeature[index+currentIndex] = 1.0
				currentIndex += feature.CategoricalInfo
			default:
				currentIndex += feature.CategoricalInfo
			}
		} else if feature.FType == ContinueValue {
			switch feature.Value.(type) {
			case float32:
				currentIndex += 1
				sparseFeature[currentIndex] = feature.Value.(float32)
			case float64:
				currentIndex += 1
				sparseFeature[currentIndex] = float32(feature.Value.(float64))
			case int:
				currentIndex += 1
				sparseFeature[currentIndex] = float32(feature.Value.(int))
			default:
				currentIndex += 1
			}
		} else if feature.FType == OnehotVectorValue {
			switch feature.Value.(type) {
			case []float32:
				for _, val := range feature.Value.([]float32) {
					sparseFeature[int(val)+currentIndex] = 1.0
				}
				currentIndex += feature.OnehotVectorInfo
			default:
				currentIndex += feature.OnehotVectorInfo
			}
		}
	}
	return sparseFeature

}
