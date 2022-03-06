package spider

import (
	"context"
	log "github.com/sirupsen/logrus"
	"strconv"
)

// 确认请求情况的暂存文件，写了单测就可以删了

//GetNCContestRating 牛客竞赛区获取个人Rating
func GetNCContestRating(ctx context.Context, nowCoderId string) (int, error) {

	goQueryFinderRets, err := doHTTPGetAndGoQuery(ctx, getNowCoderContestProfileBaseUrl(nowCoderId),
		&goQueryFinder{
			findKey:     mainRatingKey,
			findHandler: ratingHandler,
		})
	if err != nil {
		log.Errorf("GetNCContestPersonalMainPage doHTTPGetAndGoQuery Error err = %v", err)
		return 0, err
	}

	return strconv.Atoi(goQueryFinderRets[0].value)
}

//GetNCContestPassAmount 获取牛客竞赛区过题数
func GetNCContestPassAmount(ctx context.Context, nowCoderId string) (int, error) {
	goQueryFinderRets, err := doHTTPGetAndGoQuery(ctx, getNowCoderContestProfilePracticeUrl(nowCoderId),
		&goQueryFinder{
			findKey:     practicePassAmountKey,
			findHandler: passAmountHandler,
		})
	if err != nil {
		log.Errorf("getNCContestPersonalPracticePage doHTTPGetAndGoQuery Error err = %v", err)
		return 0, err
	}

	return strconv.Atoi(goQueryFinderRets[0].value)
}