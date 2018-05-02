package core

import (
	"reflect"
	"math"
	"zhihu/utils"
	"sort"
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
func (self *FeatureList) SetDayUserCrClick(value int) {
	dayUserCrClick := NewFeature("DayUserCrClick", ContinueValue, 0, 0, 0, 0, value)
	self.DayUserCrClick = *dayUserCrClick
}

// SetDayUserShow from allredis read dayUserShowAll
func (self *FeatureList) SetDayUserShow(value int) {
	dayUserShow := NewFeature("DayUserShow", ContinueValue, 0, 0, 0, 0, value)
	self.DayUserShow = *dayUserShow
}

// SetDayUserClick from allredis read dayUserClickAll
func (self *FeatureList) SetDayUserClick(value int) {
	dayUserClick := NewFeature("DayUserClick", ContinueValue, 0, 0, 0, 0, value)
	self.DayUserClick = *dayUserClick
}

// SetDayUserAdShow key = s"ad-effect:ad-uad:${day}:${user.userId}:${adId}" get views
func (self *FeatureList) SetDayUserAdShow(value int) {
	dayUserAdShow := NewFeature("DayUserAdShow", ContinueValue, 0, 0, 0, 0, value)
	self.DayUserAdShow = *dayUserAdShow
}

// SetDayUserAdClick key = s"ad-effect:ad-uad:${day}:${user.userId}:${adId}" get clicks
func (self *FeatureList) SetDayUserAdClick(value int) {
	dayUserAdClick := NewFeature("DayUserAdClick", ContinueValue, 0, 0, 0, 0, value)
	self.DayUserAdClick = *dayUserAdClick
}

// SetUserInterests set nil
func (self *FeatureList) SetUserInterests(value interface{}) {
	userInterests := NewFeature("UserInterests", OnehotValue, 19, 0, 0, 0, value)
	self.UserInterests = *userInterests
}

// SetIndustry set nil
func (self *FeatureList) SetIndustry(value interface{}) {
	industry := NewFeature("Industry", OnehotValue, 19, 0, 0, 0, value)
	self.Industry = *industry
}

// SetUserIndustrySim set nil
func (self *FeatureList) SetUserIndustrySim(value interface{}) {
	userIndustrySim := NewFeature("UserIndustrySim", ContinueValue, 0, 0, 0, 0, value)
	self.UserIndustrySim = *userIndustrySim
}

// SetAnsIndustry set nil
func (self *FeatureList) SetAnsIndustry(value interface{}) {
	ansIndustry := NewFeature("AnsIndustry", OnehotValue, 19, 0, 0, 0, value)
	self.UserIndustrySim = *ansIndustry
}

// SetUserAdidSim from ansJedis read advector and uservecotr
func (self *FeatureList) SetUserAdidSim(adVector, userVecotr []float64) {
	cosine, err := utils.CosineSimilarity(adVector, userVecotr)
	if (err != nil) {
		cosine = 0.0
	}
	userAdidSim := NewFeature("UserAdidSim", ContinueValue, 0, 0, 0, 0, cosine)
	self.UserAdidSim = *userAdidSim
}

// SetAnsAdid from ansJedis read advector and ansVector
func (self *FeatureList) SetAnsAdid(adVector, ansVector []float64) {
	cosine, err := utils.CosineSimilarity(adVector, ansVector)
	if (err != nil) {
		cosine = 0.0
	}
	ansAdid := NewFeature("AnsAdid", ContinueValue, 0, 0, 0, 0, cosine)
	self.AnsAdid = *ansAdid
}

// SetClickIndustry 用户在过去一周不同行业上的点击数量 从ctrJedis读取用户点击日志， 从mysql中读取adId和行业，
func (self *FeatureList) SetClickIndustry(value int) {
	clickIndustry := NewFeature("ClickIndustry", ContinueValue, 0, 0, 0, 0, value)
	self.ClickIndustry = *clickIndustry
}

// SetDisplayWeeks 实时获取星期几 DAY_OF_WEEK
func (self *FeatureList) SetDisplayWeeks(value int) {
	displayWeeks := NewFeature("DisplayWeeks", OnehotValue, 8, 0, 0, 0, value)
	self.DisplayWeeks = *displayWeeks
}

// SetTemplateFeature 模版id 根据adid从myql中获取
func (self *FeatureList) SetTemplateFeature(value int) {
	templateFeature := NewFeature("TemplateFeature", OnehotValue, 100, 0, 0, 0, value)
	self.TemplateFeature = *templateFeature
}

//SimsMultiple " 广告（128 * 1 关键词）和用户（128）的相似度
// 广告（128 * 1 TITLE）和用户（128）的相似度
// 广告（128 * 1 DESC ）和用户（128）的相似度
//广告（128 * 1 关键词）和回答（128）的相似度
//广告（128 * 1 TITLE）和回答（128）的相似度
//广告（128 * 1 DESC）和回答（128）的相似度
//广告创意（128 * 1）和用户（128）的相似度
//广告创意（128 * 1）和回答（128）的相似度
//from ansJedis read advector
func (self *FeatureList) SetSimsMultiple(value interface{}) {
	simsMultiple := NewFeature("SimsMultiple", VectorValue, 0, 8, 0, 0, value)
	self.SimsMultiple = *simsMultiple
}

// SetRecallSim 问题和创意的相似度	直接用引擎的数据sim
func (self *FeatureList) SetRecallSim(value float64) {
	recallSim := NewFeature("RecallSim", ContinueValue, 0, 0, 0, 0, value)
	self.RecallSim = *recallSim
}

// SetSimIsEmpty 从SimsMultiple中扩展特征：广告向量是否为空 回答向量是否为空 用户向量是否为空(0/1) 例：[0,0,1]
func (self *FeatureList) SetSimIsEmpty(value []float32) {
	simIsEmpty := NewFeature("SimIsEmpty", VectorValue, 0, 3, 0, 0, value)
	self.SimIsEmpty = *simIsEmpty
}

// SetClusterCtr 从文件中读取("strategy.topicCluster.path") key=maxTopic + "&" + newIndustry.toString
func (self *FeatureList) SetClusterCtr(value float64) {
	clusterCtr := NewFeature("ClusterCtr", ContinueValue, 0, 0, 0, 0, value)
	self.ClusterCtr = *clusterCtr
}

// SetUserVectorByAls read from userkeywordJedis s"usrvtals:${user}"
func (self *FeatureList) SetUserVectorByAls(value []float32) {
	userVectorByAls := NewFeature("UserVectorByAls", VectorValue, 0, 16, 0, 0, value)
	self.UserVectorByAls = *userVectorByAls
}

// SetUserVectorIsEmpty UserVectorByAls 扩展特征是否为空（0/1）
func (self *FeatureList) SetUserVectorIsEmpty(value interface{}) {
	userVectorIsEmpty := NewFeature("UserVectorIsEmpty", ContinueValue, 0, 0, 0, 0, value)
	self.UserVectorIsEmpty = *userVectorIsEmpty
}

//SetCreativeBrandSim from cidJedis read cidBrandVector key=s"cvtbd:${cid}" 同SetSimsMultiple中uservecotr、ansvector计算相似度
func (self *FeatureList) SetCreativeBrandSim(cidBrandVector, userVecotr, ansVector []float64) {
	cid_user_cosine, err := utils.CosineSimilarity(cidBrandVector, userVecotr)
	if (err != nil) {
		cid_user_cosine = 0.0
	}
	cid_ans_cosine, err := utils.CosineSimilarity(cidBrandVector, ansVector)
	if (err != nil) {
		cid_ans_cosine = 0.0
	}

	cidBrandVectorisEmpty := 0
	if (len(cidBrandVector) > 0) {
		cidBrandVectorisEmpty = 1
	}
	value := []float32{float32(cid_user_cosine), float32(cid_ans_cosine), float32(cidBrandVectorisEmpty)}
	creativeBrandSim := NewFeature("CreativeBrandSim", VectorValue, 0, 3, 0, 0, value)
	self.CreativeBrandSim = *creativeBrandSim
}

// SetNegativeKeyWordSim from userAdkeywordJedis read user negative vector and ad vector
func (self *FeatureList) SetNegativeKeyWordSim(user_negative_vector, ad_vecotr []float64) {
	user_negative_cosine, err := utils.CosineSimilarity(ad_vecotr, user_negative_vector)
	if (err != nil) {
		user_negative_cosine = 0.0
	}
	negativeKeyWordSim := NewFeature("NegativeKeyWordSim", ContinueValue, 0, 0, 0, 0, user_negative_cosine)
	self.NegativeKeyWordSim = *negativeKeyWordSim
}

// SetPostiveKeyWordSim from userAdkeywordJedis read user postive vector and ad vector
func (self *FeatureList) SetPostiveKeyWordSim(user_postive_vector, ad_vecotr []float64) {
	user_postive_cosine, err := utils.CosineSimilarity(ad_vecotr, user_postive_vector)
	if (err != nil) {
		user_postive_cosine = 0.0
	}
	postiveKeyWordSim := NewFeature("PostiveKeyWordSim", ContinueValue, 0, 0, 0, 0, user_postive_cosine)
	self.PostiveKeyWordSim = *postiveKeyWordSim
}

// SetNegativeUserAdTopicSim from userAdkeywordJedis read user negative vector and ad vector
func (self *FeatureList) SetNegativeUserAdTopicSim(adVector, negative_userTopicVector []float64) {
	negative_userTopic_cosine, err := utils.CosineSimilarity(adVector, negative_userTopicVector)
	if (err != nil) {
		negative_userTopic_cosine = 0.0
	}
	negativeUserAdTopicSim := NewFeature("NegativeUserAdTopicSim", ContinueValue, 0, 0, 0, 0, negative_userTopic_cosine)
	self.NegativeUserAdTopicSim = *negativeUserAdTopicSim
}
func (self *FeatureList) SetPostiveUserAdTopicSim(adVector, postive_userTopicVector []float64) {
	postive_userTopic_cosine, err := utils.CosineSimilarity(adVector, postive_userTopicVector)
	if (err != nil) {
		postive_userTopic_cosine = 0.0
	}
	postiveUserAdTopicSim := NewFeature("PostiveUserAdTopicSim", ContinueValue, 0, 0, 0, 0, postive_userTopic_cosine)
	self.PostiveUserAdTopicSim = *postiveUserAdTopicSim
}

// SetBrushnum 刷数 实时获取
func (self *FeatureList) SetBrushnum(value int) {
	brushnum := NewFeature("Brushnum", ContinueValue, 0, 0, 0, 0, value)
	self.Brushnum = *brushnum
}

// SetCtrPro from proJedis and proUserJedis and aderUserJedis read views and clicks
func (self *FeatureList) SetCtrPro(proDayShow, proDayClick, aderDayShow, aderDayClick, proUserDayShow, proUserDayClick, aderUserDayShow, aderUserDayClick int) {
	ctrarr := []float32{float32(proDayClick / (proDayShow + 1.0)),
		float32(aderDayClick / (aderDayShow + 1.0)), float32(proUserDayClick / (proUserDayShow + 1.0)), float32(aderUserDayClick / (aderUserDayShow + 1.0))}
	ctrPro := NewFeature("CtrPro", VectorValue, 0, 4, 0, 0, ctrarr)
	self.CtrPro = *ctrPro
}

// SetViewClickNum proJedis and proUserJedis and aderUserJedis read and clicks
func (self *FeatureList) SetViewClickNum(proDayClick, aderDayClick, proUserDayClick, aderUserDayClick int) {
	value := []float32{float32(proDayClick), float32(aderDayClick), float32(proUserDayClick), float32(aderUserDayClick)}
	viewClickNum := NewFeature("ViewClickNum", VectorValue, 0, 4, 0, 0, value)
	self.ViewClickNum = *viewClickNum
}

// SetPageHomeType RECOMMEND =0, FOLLOW =1,DEFAULT = 2
func (self *FeatureList) SetPageHomeType(value int) {
	pageHomeType := NewFeature("PageHomeType", OnehotValue, 8, 0, 0, 0, 2)
	self.PageHomeType = *pageHomeType
}

// SetUserAdTagOneHotVector  from userAdkeywordJedis read user and ad tagid
func (self *FeatureList) SetUserAdTagOneHotVector(usertag_vecotr, adtag_vecotr []float64) {
	resArray := []float64{}
	if (len(usertag_vecotr) > 0 && len(adtag_vecotr) > 0) {
		if (usertag_vecotr[0] == adtag_vecotr[0]) {
			resArray = append(resArray, 1)
		}
	}
	if (len(adtag_vecotr) > 0) {
		for _, element := range adtag_vecotr[1:] {
			resArray = append(resArray, 1+element)
		}
	}
	if (len(usertag_vecotr) > 0) {
		for _, element := range usertag_vecotr[1:] {
			resArray = append(resArray, 35+element)
		}
	}
	sort.Float64s(resArray)

	userAdTagOneHotVector := NewFeature("UserAdTagOneHotVector", OnehotVectorValue, 0, 0, 0, 69, resArray)
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
			case []float64:
				for _, val := range feature.Value.([]float64) {
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
