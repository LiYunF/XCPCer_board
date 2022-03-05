package spider

import (
	"XCPCer_board/util"
	"context"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"strconv"
)

const (
	PracticePassAmountKeyWord = "题已通过"

	MainRatingKeyWord = "Rating"
)

// 处理goquery匹配的回调函数
type goQueryFindHandler func(*goquery.Document) string

//getNowCoderContestProfileBaseUrl 获取牛客竞赛区个人主页URL
func getNowCoderContestProfileBaseUrl(nowCoderId string) string {
	return "https://ac.nowcoder.com/acm/contest/profile/" + nowCoderId
}

//getNowCoderContestProfilePracticeUrl 获取牛客竞赛区个人练习URL
func getNowCoderContestProfilePracticeUrl(nowCoderId string) string {
	return getNowCoderContestProfileBaseUrl(nowCoderId) + "/practice-coding"
}

// 由于需要手动关闭http client的body进行连接复用 使用回调函数的方式，保证可以关闭读响应的流
//doHTTPGetAndGoQuery 进行http请求和html解析
func doHTTPGetAndGoQuery(ctx context.Context, url string, findHandler goQueryFindHandler) (string, error) {

	// 请求阶段，并完成请求相应状态错误判断
	res, err := util.SendHTTPGet(ctx, url)
	if err != nil {
		log.Errorf("HTTP Get Error err = %v", err)
		return "", err
	}

	// 关闭io读，方便连接复用
	defer res.Body.Close()

	//解析html阶段
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Errorf("GoQuery Pharse HTML Error err = %v", err)
		return "", err
	}

	return findHandler(doc), nil
}

//---------------------------------------------------------------------//
// 个人练习信息获取
//---------------------------------------------------------------------//

//getNCContestPersonalPracticePage 牛客竞赛区获取个人练习页面信息
func getNCContestPersonalPracticePage(ctx context.Context, nowCoderId string, keyWord string) (int, error) {

	infoStr, err := doHTTPGetAndGoQuery(ctx, getNowCoderContestProfilePracticeUrl(nowCoderId),
		func(doc *goquery.Document) string {
			//解析个人状态行
			return doc.Find(".nk-container.acm-container .nk-container .nk-main.with-profile-menu.clearfix ." +
				"my-state-main .my-state-item:contains(" + keyWord + ") .state-num").Text()
		})
	if err != nil {
		log.Errorf("getNCContestPersonalPracticePage doHTTPGetAndGoQuery Error err = %v", err)
		return 0, err
	}

	return strconv.Atoi(infoStr)
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

	infoStr, err := doHTTPGetAndGoQuery(ctx, getNowCoderContestProfileBaseUrl(nowCoderId),
		func(doc *goquery.Document) string {
			//解析个人状态行
			return doc.Find(".nk-container.acm-container .nk-container .nk-main.with-profile-menu.clearfix " +
				".my-state-main .my-state-item:contains(" + keyWord + ") .state-num").Text()
		})
	if err != nil {
		log.Errorf("GetNCContestPersonalMainPage doHTTPGetAndGoQuery Error err = %v", err)
		return 0, err
	}

	return strconv.Atoi(infoStr)
}

//GetNCContestRating 获取牛客竞赛区rating
func GetNCContestRating(ctx context.Context, nowCoderId string) (int, error) {
	return getNCContestPersonalMainPage(ctx, nowCoderId, MainRatingKeyWord)
}


