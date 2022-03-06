package spider

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)

//---------------------------------------------------------------------//
// 牛客竞赛区个人信息爬取 //
//---------------------------------------------------------------------//
// 牛客finder存储Key
const (
	// 个人练习页面
	practicePassAmountKey = "NowCoder_Practice_PassAmount"

	// 个人主页
	mainRatingKey              = "NowCoder_Main_Rating"
	mainRatingRatingKey        = "NowCoder_Main_RatingRanking"
	mainAttendContestAmountKey = "NowCoder_Main_AttendContestAmount"
)

// 牛客finder关键词
const (
	// 个人练习页面
	practicePassAmountKeyWord = "题已通过"

	// 个人主页
	mainRatingKeyWord              = "Rating"
	mainRatingRankingKeyWord       = "Rating排名"
	mainAttendContestAmountKeyWord = "次比赛"
)

//getNowCoderContestProfileBaseUrl 获取牛客竞赛区个人主页URL
func getNowCoderContestProfileBaseUrl(nowCoderId string) string {
	return "https://ac.nowcoder.com/acm/contest/profile/" + nowCoderId
}

//getNowCoderContestProfilePracticeUrl 获取牛客竞赛区个人练习URL
func getNowCoderContestProfilePracticeUrl(nowCoderId string) string {
	return getNowCoderContestProfileBaseUrl(nowCoderId) + "/practice-coding"
}

//getNowCoderContestBaseFindRule 获取牛客竞赛区基础的
func getNowCoderContestBaseFindRule(keyWord string) string {
	return fmt.Sprintf(".nk-container.acm-container .nk-container .nk-main.with-profile-menu.clearfix "+
		".my-state-main .my-state-item:contains(%v) .state-num", keyWord)
}

//---------------------------------------------------------------------//
// 个人练习信息获取
//---------------------------------------------------------------------//

//personalPracticePageFinderList 需要抓取的个人练习页面finder列表
var personalPracticePageFinderList = []*goQueryFinder{
	&goQueryFinder{
		findKey:     practicePassAmountKey,
		findHandler: passAmountHandler,
	},
}

//passAmountHandler 获取竞赛区题目通过数量handler
func passAmountHandler(doc *goquery.Document) string {
	return doc.Find(getNowCoderContestBaseFindRule(practicePassAmountKeyWord)).First().Text()
}

//getNCContestPersonalPracticePage 牛客竞赛区获取个人练习页面信息
func getNCContestPersonalPracticePage(ctx context.Context, nowCoderId string) ([]*goQueryFinderReturn, error) {

	goQueryFinderRets, err := doHTTPGetAndGoQuery(ctx, getNowCoderContestProfilePracticeUrl(nowCoderId),
		personalPracticePageFinderList...)
	if err != nil {
		log.Errorf("getNCContestPersonalPracticePage doHTTPGetAndGoQuery Error err = %v", err)
		return nil, err
	}

	return goQueryFinderRets, nil
}

//---------------------------------------------------------------------//
// 个人主页信息获取
//---------------------------------------------------------------------//

//personalMainPageFinderList 需要抓取的个人主页信息finder列表
var personalMainPageFinderList = []*goQueryFinder{
	&goQueryFinder{
		findKey:     mainRatingKey,
		findHandler: ratingHandler,
	},
	&goQueryFinder{
		findKey:     mainRatingRatingKey,
		findHandler: ratingRankingHandler,
	},
	&goQueryFinder{
		findKey:     mainAttendContestAmountKey,
		findHandler: attendContestAmountHandler,
	},
}

//getNCContestPersonalMainPage 牛客竞赛区获取个人主页信息
func getNCContestPersonalMainPage(ctx context.Context, nowCoderId string, keyWord string) ([]*goQueryFinderReturn, error) {

	goQueryFinderRets, err := doHTTPGetAndGoQuery(ctx, getNowCoderContestProfileBaseUrl(nowCoderId),
		personalMainPageFinderList...)
	if err != nil {
		log.Errorf("GetNCContestPersonalMainPage doHTTPGetAndGoQuery Error err = %v", err)
		return nil, err
	}

	return goQueryFinderRets, nil
}

//ratingHandler 获取竞赛区Rating handler (需要的条件多一个比较特殊)
func ratingHandler(doc *goquery.Document) string {
	return doc.Find(fmt.Sprintf(".nk-container.acm-container .nk-container .nk-main.with-profile-menu.clearfix "+
		".my-state-main .my-state-item:contains(%v) .state-num.rate-score5", mainRatingKeyWord)).First().Text()
}

//ratingRankingHandler 获取竞赛区Rating排名handler
func ratingRankingHandler(doc *goquery.Document) string {
	return doc.Find(getNowCoderContestBaseFindRule(mainRatingRankingKeyWord)).First().Text()
}

//attendContestAmountHandler 获取竞赛区Rating排名handler
func attendContestAmountHandler(doc *goquery.Document) string {
	return doc.Find(getNowCoderContestBaseFindRule(mainAttendContestAmountKeyWord)).First().Text()
}
