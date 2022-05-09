package luogu

import (
	_ "XCPCer_board/config"
	_ "XCPCer_board/dao"
	"XCPCer_board/model"
	"fmt"
	"testing"
)

///////////////////用户结构体//////////////////////
type userExample struct {
	msg map[string]int
}

var userSet = map[string]userExample{
	model.TestLuoGuIdLYF: userExample{
		msg: map[string]int{
			basicProblem:      30,
			elevatedProblem:   19,
			hardProblem:       6,
			passProblemNumber: 55,
			ranting:           94572,
			simpleProblem:     0,
			unKnowProblem:     0,
		},
	},
}
var subSet = map[string]string{
	model.TestLuoGuIdLYF: "map[P1009:{[NOIP1998 普及组] 阶乘之和 2} P1011:{[NOIP1998 提高组] 车站 2} P1020:{[NOIP1999 普及组] 导弹拦截 3} " +
		"P1048:{[NOIP2005 普及组] 采药 2} P1060:{[NOIP2006 普及组] 开心的金明 2} P1062:{[NOIP2006 普及组] 数列 2} " +
		"P1090:{[NOIP2004 提高组] 合并果子 / [USACO06NOV] Fence Repair G 3} P1091:{[NOIP2004 提高组] 合唱队形 3}" +
		" P1107:{[BJWC2008]雷涛的小猫 4} P1119:{灾后重建 4} P1135:{奇怪的电梯 3} P1144:{最短路计数 4} P1164:{小A点菜 2}" +
		" P1181:{数列分段 Section I 2} P1208:{[USACO1.3]混合牛奶 Mixing Milk 2} P1223:{排队接水 2} P1247:{取火柴游戏 4} " +
		"P1280:{尼克的任务 4} P1288:{取数游戏II 3} P1290:{欧几里德的游戏 3} P1326:{足球 2} P1339:{[USACO09OCT]Heat Wave G 3}" +
		" P1356:{数列的整除性 4} P1372:{又是毕业季I 2} P1443:{马的遍历 3} P1507:{NASA的食物计划 2} P1582:{倒水 4}" +
		" P1629:{邮递员送信 3} P1757:{通天之分组背包 2} P1776:{宝物筛选 4} P1833:{樱花 4} P1902:{刺杀大使 4}" +
		" P1962:{斐波那契数列 3} P2085:{最小函数值 4} P2196:{[NOIP1996 提高组] 挖地雷 3} P2197:{【模板】nim 游戏 4}" +
		" P2257:{YY的GCD 5} P2613:{【模板】有理数取余 4} P2953:{[USACO09OPEN]Cow Digit Game S 4} P3197:{[HNOI2008]越狱 3}" +
		" P3371:{【模板】单源最短路径（弱化版） 3} P3803:{【模板】多项式乘法（FFT） 5} P3812:{【模板】线性基 5} " +
		"P3868:{[TJOI2009] 猜数字 5} P3912:{素数个数 3} P4018:{Roy&October之取石子 3} P4239:{任意模数多项式乘法逆 6} " +
		"P4720:{【模板】扩展卢卡斯定理/exLucas 6} P4723:{【模板】常系数齐次线性递推 7} P4777:{【模板】扩展中国剩余定理（EXCRT） 6}" +
		" P4860:{Roy&October之取石子II 3} P5325:{【模板】Min_25筛 7} P5491:{【模板】二次剩余 6} P5520:{[yLOI2019] 青原樱 3}" +
		" UVA10214:{树林里的树 Trees in a Wood. 5}]",
}

func checkSub(t *testing.T, uid string, ans string) {
	ret, err := ScrapeSub(uid)
	if err != nil {
		t.Errorf("Error of  get sub : %v\n", err)
	}
	stq := fmt.Sprintf("%v", ret)
	if stq != ans {
		t.Errorf("Error of sub msg,Get: -------------\n%v\n"+
			"-------------\n But ans is:-------------\n%v\n-------------\n", stq, ans)
	}
}

/////////////////////////////////////
//////////  主测试函数  ///////////////
/////////////////////////////////////

func UserMsgTest(t *testing.T) {

	//test ScrapeUser
	for uid, correctMsg := range userSet {
		//get msg
		funcRet, err := ScrapeUser(uid)
		if err != nil {
			t.Errorf("Error of %v in all msg: %v", "luoGu", err)
		}
		//check map
		for k, v := range correctMsg.msg {
			if _, r := funcRet[k]; r == false || v != funcRet[k] {
				t.Errorf("Error of %v in all msg\n ret= %v  \nbut the ans is %v", "luoGu", funcRet, correctMsg.msg)
			}
		}
	}

}

func SubmissionTest(t *testing.T) {
	checkSub(t, model.TestLuoGuIdLYF, "map[P1009:{[NOIP1998 普及组] 阶乘之和 2} P1011:{[NOIP1998 提高组] 车站 2} P1020:{[NOIP1999 普及组] 导弹拦截 3} "+
		"P1048:{[NOIP2005 普及组] 采药 2} P1060:{[NOIP2006 普及组] 开心的金明 2} P1062:{[NOIP2006 普及组] 数列 2} "+
		"P1090:{[NOIP2004 提高组] 合并果子 / [USACO06NOV] Fence Repair G 3} P1091:{[NOIP2004 提高组] 合唱队形 3}"+
		" P1107:{[BJWC2008]雷涛的小猫 4} P1119:{灾后重建 4} P1135:{奇怪的电梯 3} P1144:{最短路计数 4} P1164:{小A点菜 2}"+
		" P1181:{数列分段 Section I 2} P1208:{[USACO1.3]混合牛奶 Mixing Milk 2} P1223:{排队接水 2} P1247:{取火柴游戏 4} "+
		"P1280:{尼克的任务 4} P1288:{取数游戏II 3} P1290:{欧几里德的游戏 3} P1326:{足球 2} P1339:{[USACO09OCT]Heat Wave G 3}"+
		" P1356:{数列的整除性 4} P1372:{又是毕业季I 2} P1443:{马的遍历 3} P1507:{NASA的食物计划 2} P1582:{倒水 4}"+
		" P1629:{邮递员送信 3} P1757:{通天之分组背包 2} P1776:{宝物筛选 4} P1833:{樱花 4} P1902:{刺杀大使 4}"+
		" P1962:{斐波那契数列 3} P2085:{最小函数值 4} P2196:{[NOIP1996 提高组] 挖地雷 3} P2197:{【模板】nim 游戏 4}"+
		" P2257:{YY的GCD 5} P2613:{【模板】有理数取余 4} P2953:{[USACO09OPEN]Cow Digit Game S 4} P3197:{[HNOI2008]越狱 3}"+
		" P3371:{【模板】单源最短路径（弱化版） 3} P3803:{【模板】多项式乘法（FFT） 5} P3812:{【模板】线性基 5} "+
		"P3868:{[TJOI2009] 猜数字 5} P3912:{素数个数 3} P4018:{Roy&October之取石子 3} P4239:{任意模数多项式乘法逆 6} "+
		"P4720:{【模板】扩展卢卡斯定理/exLucas 6} P4723:{【模板】常系数齐次线性递推 7} P4777:{【模板】扩展中国剩余定理（EXCRT） 6}"+
		" P4860:{Roy&October之取石子II 3} P5325:{【模板】Min_25筛 7} P5491:{【模板】二次剩余 6} P5520:{[yLOI2019] 青原樱 3}"+
		" UVA10214:{树林里的树 Trees in a Wood. 5}]")
}
func TestLg(t *testing.T) {
	//UserMsgTest(t)
	SubmissionTest(t)
}
