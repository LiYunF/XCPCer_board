package spider

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
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
	codeforcesMainMaxRatingKey = "CodeForces_Main_Max_Rating"
)

// CF finder关键词
const (
	// 个人总过题数
	codeforcesPracticePassAmountKeyWord = "all"

	//个人最后一月过题数
	codeforcesLastMonthPracticePassAmountKeyWord = "month"

	// 个人rating
	codeforcesMainRatingKeyWord = "CodeForces_Main_Rating"

	// 个人历史最高rating
	codeforcesMainMaxRatingKeyWord = "CodeForces_Main_Rating"
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

//GetCFContestPassAmount 获取Codeforces信息
func GetCFContestPassAmount(ctx context.Context, codeForcesId string) ([]*goQueryFinderReturn, error) {
	return getCFContestPersonalPracticePage(ctx, codeForcesId)
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
	//&goQueryFinder{
	//	findKey:     codeforcesMainRatingKey,
	//	findHandler: codeforcesMainRatingHandler,
	//},
}

func codeforcesPracticePassAmountHandler(doc *goquery.Document) string {

	retStr := doc.Find(fmt.Sprint("#body div[style=\"position: relative;\"] #pageContent ._UserActivityFrame_frame" +
		" .roundbox.userActivityRoundBox ._UserActivityFrame_footer" +
		" ._UserActivityFrame_countersRow ._UserActivityFrame_counter:contains(solved):contains(" +
		codeforcesPracticePassAmountKeyWord + ") ._UserActivityFrame_counterValue")).First().Text()
	return retStr[:len(retStr)-9]
	//"1000 problems" -> "1000"
}
func codeforcesPracticePassLastMonthAmountHandler(doc *goquery.Document) string {

	retStr := doc.Find(fmt.Sprint("#body div[style=\"position: relative;\"] #pageContent ._UserActivityFrame_frame" +
		" .roundbox.userActivityRoundBox ._UserActivityFrame_footer" +
		" ._UserActivityFrame_countersRow ._UserActivityFrame_counter:contains(solved):contains(" +
		codeforcesLastMonthPracticePassAmountKeyWord + ") ._UserActivityFrame_counterValue")).First().Text()
	return retStr[:len(retStr)-9]
	//"1000 problems" -> "1000"
}
func codeforcesMainRatingHandler(doc *goquery.Document) string {

	retStr := doc.Find(fmt.Sprint("#body div[style=\"position: relative;\"] #pageContent ._UserActivityFrame_frame" +
		" .roundbox.userActivityRoundBox ._UserActivityFrame_footer" +
		" ._UserActivityFrame_countersRow ._UserActivityFrame_counter:contains(solved):contains(" +
		codeforcesMainRatingKeyWord + ") ._UserActivityFrame_counterValue")).First().Text()
	return retStr[:len(retStr)-9]
	//"1000 problems" -> "1000"
}
