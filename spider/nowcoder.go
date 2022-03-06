package spider

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"strconv"
)

//---------------------------------------------------------------------//
// 牛客个人信息爬取 //
//---------------------------------------------------------------------//
// 牛客findler Key
const (

	// 个人练习页面
	PracticePassAmountKey = "NowCoder_Practice_PassAmount"

	// 个人主页
	MainRatingKey = "NowCoder_Main_Rating"
)

// 牛客finder关键词
const (
	// 个人练习页面
	PracticePassAmountKeyWord = "题已通过"

	// 个人主页
	MainRatingKeyWord = "Rating"
)

//getNowCoderContestProfileBaseUrl 获取牛客竞赛区个人主页URL
func getNowCoderContestProfileBaseUrl(nowCoderId string) string {
	return "https://ac.nowcoder.com/acm/contest/profile/" + nowCoderId
}

//getNowCoderContestProfilePracticeUrl 获取牛客竞赛区个人练习URL
func getNowCoderContestProfilePracticeUrl(nowCoderId string) string {
	return getNowCoderContestProfileBaseUrl(nowCoderId) + "/practice-coding"
}

//---------------------------------------------------------------------//
// 个人练习信息获取
//---------------------------------------------------------------------//

//getNCContestPersonalPracticePage 牛客竞赛区获取个人练习页面信息
func getNCContestPersonalPracticePage(ctx context.Context, nowCoderId string, keyWord string) (int, error) {

	goQueryFinderRets, err := doHTTPGetAndGoQuery(ctx, getNowCoderContestProfilePracticeUrl(nowCoderId),
		goQueryFinder{
			findKey: PracticePassAmountKey,
			findHandler: func(doc *goquery.Document) string {
				//解析个人状态行
				return doc.Find(".nk-container.acm-container .nk-container .nk-main.with-profile-menu.clearfix ." +
					"my-state-main .my-state-item:contains(" + keyWord + ") .state-num").Text()
			}})
	if err != nil {
		log.Errorf("getNCContestPersonalPracticePage doHTTPGetAndGoQuery Error err = %v", err)
		return 0, err
	}

	return strconv.Atoi(goQueryFinderRets[0].value)
}

//GetNCContestPassAmount 获取牛客竞赛区过题数
func GetNCContestPassAmount(ctx context.Context, nowCoderId string) (int, error) {
	return getNCContestPersonalPracticePage(ctx, nowCoderId, PracticePassAmountKeyWord)
}

//---------------------------------------------------------------------//
// 个人主页信息获取
//---------------------------------------------------------------------//

//getNCContestPersonalMainPage 牛客竞赛区获取个人主页信息
func getNCContestPersonalMainPage(ctx context.Context, nowCoderId string, keyWord string) (int, error) {

	goQueryFinderRets, err := doHTTPGetAndGoQuery(ctx, getNowCoderContestProfileBaseUrl(nowCoderId),
		goQueryFinder{
			findKey: MainRatingKey,
			findHandler: func(doc *goquery.Document) string {
				//解析个人状态行
				return doc.Find(".nk-container.acm-container .nk-container .nk-main.with-profile-menu.clearfix " +
					".my-state-main .my-state-item:contains(" + keyWord + ") .state-num").Text()
			}})
	if err != nil {
		log.Errorf("GetNCContestPersonalMainPage doHTTPGetAndGoQuery Error err = %v", err)
		return 0, err
	}

	return strconv.Atoi(goQueryFinderRets[0].value)
}

//GetNCContestRating 获取牛客竞赛区rating
func GetNCContestRating(ctx context.Context, nowCoderId string) (int, error) {
	return getNCContestPersonalMainPage(ctx, nowCoderId, MainRatingKeyWord)
}
