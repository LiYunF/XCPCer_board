package codeforces

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

//---------------------------------------------------------------------//
// codeforces个人信息 //
//---------------------------------------------------------------------//
// CF  Key
const (
	// key
	// 个人总过题数
	problemPassAmountKey = "codeForces_problem_pass_amount"
	// 个人最后一月过题数
	lastMonthPassAmount = "codeForces_last_month_problem_pass_amount"

	// CF finder关键词
	// 个人总过题数
	problemPassKeyWord = "all"
	//个人最后一月过题数
	lastMonthPassKeyWord = "month"
)

//problemPassAmountHandler 获取cf个人总过题数
func problemPassAmountHandler(doc *goquery.Selection) string {
	retStr := doc.Find(fmt.Sprintf("div[style=\"position: relative;\"] #pageContent ._UserActivityFrame_frame"+
		" .roundbox.userActivityRoundBox ._UserActivityFrame_footer"+
		" ._UserActivityFrame_countersRow ._UserActivityFrame_counter:contains(solved):contains(%v)"+
		" ._UserActivityFrame_counterValue", problemPassKeyWord)).First().Text()
	return strings.Split(retStr, " ")[0]
	//"1000 problems" -> "1000"
}

//lastMonthAmountHandler 获取cf个人上个月总过题数
func lastMonthAmountHandler(doc *goquery.Selection) string {
	retStr := doc.Find(fmt.Sprintf("div[style=\"position: relative;\"] #pageContent ._UserActivityFrame_frame"+
		" .roundbox.userActivityRoundBox ._UserActivityFrame_footer"+
		" ._UserActivityFrame_countersRow ._UserActivityFrame_counter:contains(solved):contains(%v)"+
		" ._UserActivityFrame_counterValue", lastMonthPassKeyWord)).First().Text()
	return strings.Split(retStr, " ")[0]
}

//---------------------------------------------------------------------//
// 部分共用函数 //
//---------------------------------------------------------------------//
//转化int
func strToInt(doc *goquery.Selection, f func(doc *goquery.Selection) string) int {
	ret := f(doc)
	num, err := strconv.Atoi(ret)
	if err != nil {
		log.Errorf("CF strToInt get err:%v\tand the return is %v:", num, err)
		return -1
	}
	return num
}

func getPersonPage(uid string) string {
	return "https://codeforces.com/profile/" + uid
}
func getPersonProblemPage(uid string) string {
	return "https://codeforces.com/submissions/" + uid
}
