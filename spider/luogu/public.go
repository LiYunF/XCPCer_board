package luogu

import (
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

const (
	////////////////user//////////////////
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


const (
////////submission/////////
//题号

)

//获取网页函数
func getPersonPage(uid string) string {
	return "https://www.luogu.com.cn/user/" + uid
}
func getPersonPractice(uid string) string {
	return getPersonPage(uid) + "#practice"
}

//字符转int

func strToInt(doc *goquery.Selection, f func(doc *goquery.Selection) string) int {
	ret := f(doc)
	num, err := strconv.Atoi(ret)
	if err != nil {
		log.Errorf("luogu strToInt get err:%v\tand the return is %v:", num, err)
		return -1
	}
	return num
}


//unicode转中文
func unicodeToChinese(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil

func getPersonPage(uid string) string {
	return "https://www.luogu.com.cn/user/" + uid
}
func getPersonPractice(uid string) string {
	return getPersonPage(uid) + "#practice"

}
