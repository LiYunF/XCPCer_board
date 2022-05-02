package luogu

import (
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"strconv"
)

const (
	//过题数
	passProblemNumber = "problem_number"
	//排名
	ranting = "ranting"
	//简单题个数
	simpleProblem = "simple_problem_number"
	//基础题个数
	basicProblem = "base_problem_number"
	//提高题个数
	elevatedProblem = "elevated_problem_number"
	//困难题个数
	hardProblem = "hard_problem_number"
	//未知题个数
	unKnowProblem = "unKnow_problem_number"
)

type UserMsg struct {
	Uid               string
	PassProblemNumber int
	Ranting           int
	SimpleProblem     int
	BasicProblem      int
	ElevatedProblem   int
	HardProblem       int
	UnKnowProblem     int
}

func strToInt(doc *goquery.Selection, f func(doc *goquery.Selection) string) int {
	ret := f(doc)
	num, err := strconv.Atoi(ret)
	if err != nil {
		log.Errorf("luogu strToInt get err:%v\tand the return is %v:", num, err)
		return -1
	}
	return num
}
func getPersonPage(uid string) string {
	return "https://www.luogu.com.cn/user/" + uid
}
func getPersonPractice(uid string) string {
	return getPersonPage(uid) + "#practice"
}

/////////////结构体对外暴露函数///////////////

//StructToMap 结构体转Map
func StructToMap(user UserMsg) (map[string]int, string) {
	var mp map[string]int
	mp[passProblemNumber] = user.PassProblemNumber
	mp[ranting] = user.Ranting
	mp[simpleProblem] = user.SimpleProblem
	mp[basicProblem] = user.BasicProblem
	mp[elevatedProblem] = user.ElevatedProblem
	mp[hardProblem] = user.HardProblem
	mp[unKnowProblem] = user.UnKnowProblem
	return mp, user.Uid
}

//MapToStruct Map转结构体, 返回的bool=1为正常，0为map里没有该值
func MapToStruct(mp map[string]int) (UserMsg, bool) {

	var user UserMsg
	var ok bool

	if user.PassProblemNumber, ok = mp[passProblemNumber]; !ok {
		return user, ok
	}
	if user.Ranting, ok = mp[ranting]; !ok {
		return user, ok
	}
	if user.SimpleProblem, ok = mp[simpleProblem]; !ok {
		return user, ok
	}
	if user.BasicProblem, ok = mp[basicProblem]; !ok {
		return user, ok
	}
	if user.ElevatedProblem, ok = mp[elevatedProblem]; !ok {
		return user, ok
	}
	if user.HardProblem, ok = mp[hardProblem]; !ok {
		return user, ok
	}
	if user.UnKnowProblem, ok = mp[unKnowProblem]; !ok {
		return user, ok
	}
	return user, true
}
