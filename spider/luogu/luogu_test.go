package luogu

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
func luoGuTest(t *testing.T) {

	tp := "luoGu"
	fc1 := ScrapeAll
	var uid string
	var lgInt map[string]int

	//输入个例
	uid = model.TestLuoGuIdLYF
	lgInt = map[string]int{
		"luoGu_Basic_Problem_Number":       195,
		"luoGu_Elevated_Problem_Number":    368,
		"luoGu_Hard_Problem_Number":        1069,
		"luoGu_Person_Pass_Problem_Number": 1743,
		"luoGu_Person_Ranting":             837,
		"luoGu_Simple_Problem_Number":      43,
		"luoGu_UnKnow_Problem_Number":      68,
	}

	//开始测试
	checkIntError(t, uid, tp, fc1, lgInt)

}

func TestLg(t *testing.T) {
	luoGuTest(t)
}
