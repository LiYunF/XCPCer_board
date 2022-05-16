package nowcoder

import "fmt"

// @Author: Feng
// @Date: 2022/5/16 17:48

func getRatingKey(uid string) string {
	return fmt.Sprintf("%v_%v", ratingKey, uid)
}

func getRankingKey(uid string) string {
	return fmt.Sprintf("%v_%v", rankingKey, uid)
}

func getContestAmountKey(uid string) string {
	return fmt.Sprintf("%v_%v", contestAmountKey, uid)
}

func getPassAmountKey(uid string) string {
	return fmt.Sprintf("%v_%v", passAmountKey, uid)
}
