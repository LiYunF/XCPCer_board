package luogu

import (
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"strconv"
)

const (
	//过题数
	luoGuPersonPassProblemNumber = "lg_problem_number"
	//排名
	luoGuPersonRanting = "lg_ranting"
	//简单题个数
	luoGuSimpleProblemNumber = "simple_problem_number"
	//基础题个数
	luoGuBasicProblemNumber = "base_problem_number"
	//提高题个数
	luoGuElevatedProblemNumber = "elevated_problem_number"
	//困难题个数
	luoGuHardProblemNumber = "hard_problem_number"
	//未知题个数
	luoGuUnKnowProblemNumber = "unKnow_problem_number"
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
