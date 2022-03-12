package spider

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"strings"
)

//---------------------------------------------------------------------//
// codeforces个人信息 //
//---------------------------------------------------------------------//
// CF findler Key
const (
	numberOfKey = 4
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
)

// CF finder关键词
const (
	// 个人总过题数
	codeforcesPracticePassAmountKeyWord = "all"

	//个人最后一月过题数
	codeforcesLastMonthPracticePassAmountKeyWord = "month"

	// 个人rating
	codeforcesMainRatingKeyWord = "rating"

	// 个人历史最高rating
	codeforcesMaxMainRatingKeyWord = "max"

	//当前rating所对应的等级
	//不需要finder关键词
)

//---------------------------------------------------------------------//
// 个人基础信息获取
//---------------------------------------------------------------------//

//getCodeForcesContestProfileBaseUrl 获取codeForces个人主页URL
func getCodeForcesContestProfileBaseUrl(codeForcesId string) string {
	return "https://codeforces.com/profile/" + codeForcesId
}

//getCodeForcesContestProfilePracticeUrl 获取codeForces个人提交页面URL
func getCodeForcesContestProfileSubmissionsUrl(codeForcesId string) string {
	return "https://codeforces.com/submissions/" + codeForcesId
}

//getCFContestPersonalPracticePage Codeforces获取个人主页信息
func getCFContestPersonalPracticePage(ctx context.Context, codeForcesId string) ([]*goQueryFinderReturn, error) {

	goQueryFinderRets, err := doHTTPGetAndGoQuery(ctx, getCodeForcesContestProfileBaseUrl(codeForcesId),
		codeforcesPersonalMainPageFinderList...,
	)

	if err != nil {
		log.Errorf("getCFContestPersonalPracticePage doHTTPGetAndGoQuery Error err = %v", err)
		return nil, err
	}

	return goQueryFinderRets, nil
}

//codeforcesPersonalMainPageFinderList 需要抓取的个人主页信息finder列表
var codeforcesPersonalMainPageFinderList = []*goQueryFinder{
	&goQueryFinder{ //获取过题数
		findKey:     codeforcesPracticePassAmountKey,
		findHandler: codeforcesPracticePassAmountHandler,
	},
	&goQueryFinder{ //最后一月过题数
		findKey:     codeforcesLastMonthPracticePassAmount,
		findHandler: codeforcesPracticePassLastMonthAmountHandler,
	},
	&goQueryFinder{ //rating
		findKey:     codeforcesMainRatingKey,
		findHandler: codeforcesMainRatingHandler,
	},
	&goQueryFinder{ //最高rating
		findKey:     codeforcesMaxMainRatingKey,
		findHandler: codeforcesMaxMainRatingHandler,
	},
	&goQueryFinder{ //个人rating对应的等级
		findKey:     codeforcesMainRatingNameKey,
		findHandler: codeforcesMainRatingNameHandler,
	},
}

func codeforcesPracticePassAmountHandler(doc *goquery.Document) string {

	retStr := doc.Find(fmt.Sprintf("#body div[style=\"position: relative;\"] #pageContent ._UserActivityFrame_frame"+
		" .roundbox.userActivityRoundBox ._UserActivityFrame_footer"+
		" ._UserActivityFrame_countersRow ._UserActivityFrame_counter:contains(solved):contains(%v)"+
		" ._UserActivityFrame_counterValue", codeforcesPracticePassAmountKeyWord)).First().Text()
	return strings.Split(retStr, " ")[0]
	//"1000 problems" -> "1000"
}
func codeforcesPracticePassLastMonthAmountHandler(doc *goquery.Document) string {

	retStr := doc.Find(fmt.Sprintf("#body div[style=\"position: relative;\"] #pageContent ._UserActivityFrame_frame"+
		" .roundbox.userActivityRoundBox ._UserActivityFrame_footer"+
		" ._UserActivityFrame_countersRow ._UserActivityFrame_counter:contains(solved):contains(%v)"+
		" ._UserActivityFrame_counterValue", codeforcesLastMonthPracticePassAmountKeyWord)).First().Text()
	return strings.Split(retStr, " ")[0]
}
func codeforcesMainRatingHandler(doc *goquery.Document) string {

	retStr := doc.Find(fmt.Sprintf("#body div[style=\"position: relative;\"] #pageContent "+
		"div[style=\"padding:1em 1em 0 1em;\"] .userbox .info ul li:contains(%v)"+
		" span[style]", //有style的<span>是我们想要的
		codeforcesMainRatingKeyWord)).First().Text()
	return strings.Split(retStr, " ")[0]
}
func codeforcesMaxMainRatingHandler(doc *goquery.Document) string {

	retStr := doc.Find(fmt.Sprintf("#body div[style=\"position: relative;\"] #pageContent "+
		"div[style=\"padding:1em 1em 0 1em;\"] .userbox .info ul li:contains(%v)"+
		" .smaller span+span", //选择他的邻居
		codeforcesMaxMainRatingKeyWord)).First().Text()
	return strings.Split(retStr, " ")[0]
}
func codeforcesMainRatingNameHandler(doc *goquery.Document) string {

	retStr := doc.Find(fmt.Sprintf("#body div[style=\"position: relative;\"] #pageContent " +
		"div[style=\"padding:1em 1em 0 1em;\"] .userbox .info .main-info .user-rank  span")).First().Text()
	return retStr
}

//GetCFContestPassAmount 获取Codeforces信息
func GetCFContestPassAmount(ctx context.Context, codeForcesId string) ([]*goQueryFinderReturn, error) {
	return getCFContestPersonalPracticePage(ctx, codeForcesId)
}
