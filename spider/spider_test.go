package spider

import (
	"XCPCer_board/model"
	"XCPCer_board/spider/codeforces"
	"XCPCer_board/spider/luogu"
	"XCPCer_board/spider/nowcoder"
	"XCPCer_board/spider/vjudge"
	"testing"
)

type cfIntMsg struct {
	uid string
}

//检查函数，若不一致返回1
func isIntMsgDifferent(funcRet, ans map[string]int) bool {
	for k, v := range ans {
		if _, r := funcRet[k]; r == false || v != funcRet[k] {
			return true
		}
	}
	return false
}
func isStringMsgDifferent(funcRet, ans map[string]string) bool {
	for k, v := range ans {
		if _, r := funcRet[k]; r == false || v != funcRet[k] {
			return true
		}
	}
	return false
}

//判断并输出错误
func checkIntError(t *testing.T, uid string, tp string, all func(uid string) (map[string]int, error),
	acInt map[string]int) {
	if ret, err := all(uid); isIntMsgDifferent(ret, acInt) {
		if err != nil {
			t.Errorf("Error of %v in all msg: %v", tp, err)
		}
		t.Errorf("Error of %v in all msg\n ret= %v  \nbut the ans is %v", tp, ret, acInt)
	}
}
func checkStrError(t *testing.T, uid string, tp string, all func(uid string) (map[string]string, error),
	acInt map[string]string) {
	if ret, err := all(uid); isStringMsgDifferent(ret, acInt) {
		if err != nil {
			t.Errorf("Error of %v in all msg: %v", tp, err)
		}
		t.Errorf("Error of %v in str msg\n ret= %v  \nbut the ans is %v", tp, ret, acInt)
	}
}

//////////////////////////////////////
/////		测试四大模块     //////////
/////////////////////////////////////

//测试codeforces
func cfTest(t *testing.T) {

	//基础设置
	tp := "codeforces"
	fc1 := codeforces.ScrapeAll
	fc2 := codeforces.ScrapeStr
	var uid string
	var cfInt map[string]int
	var cfStr map[string]string

	//个例赋值
	uid = model.TestCodeForcesIdLYF
	cfInt = map[string]int{
		"CodeForces_Last_Month_Practice_PassAmount": 0,
		"CodeForces_Main_Max_Rating":                1837,
		"CodeForces_Main_Rating":                    1742,
		"CodeForces_Practice_PassAmount":            350,
	}
	cfStr = map[string]string{
		"CodeForces_Main_Rating_Name": "Expert ",
	}

	//开始测试
	checkIntError(t, uid, tp, fc1, cfInt)
	checkStrError(t, uid, tp, fc2, cfStr)

}
func luoGuTest(t *testing.T) {

	tp := "luoGu"
	fc1 := luogu.ScrapeAll
	var uid string
	var lgInt map[string]int

	//输入个例
	uid = model.TestLuoGuIdLYF
	lgInt = map[string]int{
		"luoGu_Basic_Problem_Number":       195,
		"luoGu_Elevated_Problem_Number":    368,
		"luoGu_Hard_Problem_Number":        1069,
		"luoGu_Person_Pass_Problem_Number": 1743,
		"luoGu_Person_Ranting":             796,
		"luoGu_Simple_Problem_Number":      43,
		"luoGu_UnKnow_Problem_Number":      68,
	}

	//开始测试
	checkIntError(t, uid, tp, fc1, lgInt)

}
func nowCoderTest(t *testing.T) {

	tp := "nowCoder"
	fc1 := nowcoder.ScrapeAll
	var uid string
	var ncInt map[string]int

	//输入个例
	uid = model.TestNowCoderIdLYF
	ncInt = map[string]int{
		"NowCoder_Main_AttendContestAmount": 23,
		"NowCoder_Main_Rating":              -1,
		"NowCoder_Main_RatingRanking":       -1,
		"NowCoder_Practice_PassAmount":      39,
	}
	//开始测试
	checkIntError(t, uid, tp, fc1, ncInt)

}
func vjTest(t *testing.T) {

	tp := "vJudge"
	fc1 := vjudge.ScrapeAll
	var uid string
	var vjInt map[string]int

	uid = model.TestVJIdLYF
	vjInt = map[string]int{
		"vj_Person_Last_24_Hours_Pass_Number": 0,
		"vj_Person_Last_30_Days_Pass_Number":  0,
		"vj_Person_Last_7_Days_Pass_Number":   0,
		"vj_Person_Pass_Number":               30,
	}
	//开始测试
	checkIntError(t, uid, tp, fc1, vjInt)

}

func TestMul(t *testing.T) {
	cfTest(t)
	luoGuTest(t)
	vjTest(t)

}
