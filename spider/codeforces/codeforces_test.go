package codeforces

import (
	"XCPCer_board/model"
	"testing"
)

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

//测试codeforces
func cfTest(t *testing.T) {

	//基础设置
	tp := "codeforces"
	fc1 := ScrapeAll
	fc2 := ScrapeStr
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

func TestCF(t *testing.T) {
	cfTest(t)
}
