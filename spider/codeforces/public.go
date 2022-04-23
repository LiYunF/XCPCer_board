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
	//key

	// 个人总过题数
	codeforcesPracticePassAmountKey = "CodeForces_Practice_PassAmount"
	//个人最后一月过题数
	codeforcesLastMonthPracticePassAmount = "CodeForces_Last_Month_Practice_PassAmount"
	// 个人rating
	codeforcesMainRatingKey = "CodeForces_Main_Rating"
	// 个人历史最高rating
	codeforcesMaxMainRatingKey = "CodeForces_Main_Max_Rating"
	//当前rating所对应的等级（红名、紫名...)
	codeforcesMainRatingNameKey = "CodeForces_Main_Rating_Name"

	// CF finder关键词
	// 个人总过题数
	codeforcesPracticePassAmountKeyWord = "all"
	//个人最后一月过题数
	codeforcesLastMonthPracticePassAmountKeyWord = "month"
	// 个人rating
	codeforcesMainRatingKeyWord = "rating"
	// 个人历史最高rating
	codeforcesMaxMainRatingKeyWord = "max"
)

//获取cf个人总过题数
func cfPracticePassAmountHandler(doc *goquery.Selection) string {

	retStr := doc.Find(fmt.Sprintf("div[style=\"position: relative;\"] #pageContent ._UserActivityFrame_frame"+
		" .roundbox.userActivityRoundBox ._UserActivityFrame_footer"+
		" ._UserActivityFrame_countersRow ._UserActivityFrame_counter:contains(solved):contains(%v)"+
		" ._UserActivityFrame_counterValue", codeforcesPracticePassAmountKeyWord)).First().Text()
	return strings.Split(retStr, " ")[0]
	//"1000 problems" -> "1000"
}

//获取cf个人上个月总过题数
func cfPracticePassLastMonthAmountHandler(doc *goquery.Selection) string {

	retStr := doc.Find(fmt.Sprintf("div[style=\"position: relative;\"] #pageContent ._UserActivityFrame_frame"+
		" .roundbox.userActivityRoundBox ._UserActivityFrame_footer"+
		" ._UserActivityFrame_countersRow ._UserActivityFrame_counter:contains(solved):contains(%v)"+
		" ._UserActivityFrame_counterValue", codeforcesLastMonthPracticePassAmountKeyWord)).First().Text()
	return strings.Split(retStr, " ")[0]
}

func cfMainRatingHandler(doc *goquery.Selection) string {

	retStr := doc.Find(fmt.Sprintf("div[style=\"position: relative;\"] #pageContent "+
		" .userbox .info ul li:contains(%v)"+
		" span[style]", //有style的<span>是我们想要的
		codeforcesMainRatingKeyWord)).First().Text()
	return strings.Split(retStr, " ")[0]
}
func cfMaxMainRatingHandler(doc *goquery.Selection) string {

	retStr := doc.Find(fmt.Sprintf("div[style=\"position: relative;\"] #pageContent "+
		" .userbox .info ul li:contains(%v)"+
		" .smaller span+span", //选择他的邻居
		codeforcesMaxMainRatingKeyWord)).First().Text()
	return strings.Split(retStr, " ")[0]
}
func codeforcesMainRatingNameHandler(doc *goquery.Selection) string {

	retStr := doc.Find(fmt.Sprintf("div[style=\"position: relative;\"] #pageContent " +
		".userbox .info .main-info .user-rank  span")).First().Text()
	return retStr
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
	return "https://codeforces.com/profile/" + uid
}
