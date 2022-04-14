package luogu

import (
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"strconv"
)

const (
	//过题数
	luoGuPersonPassProblemNumber = "luoGu_Person_Pass_Problem_Number"
	//排名
	luoGuPersonRanting = "luoGu_Person_Ranting"
	//简单题个数
	luoGuSimpleProblemNumber = "luoGu_Simple_Problem_Number"
	//基础题个数
	luoGuBasicProblemNumber = "luoGu_Basic_Problem_Number"
	//提高题个数
	luoGuElevatedProblemNumber = "luoGu_Elevated_Problem_Number"
	//困难题个数
	luoGuHardProblemNumber = "luoGu_Hard_Problem_Number"
	//未知题个数
	luoGuUnKnowProblemNumber = "luoGu_UnKnow_Problem_Number"
)

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
