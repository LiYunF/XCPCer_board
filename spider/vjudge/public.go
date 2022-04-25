package vjudge

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"strconv"
)

//---------------------------------------------------------------------//
// vj //
//---------------------------------------------------------------------//
// vj  Key
const (
	//24小时前的过题数
	vjPersonLast24HoursPassNumber = "vj_Person_Last_24_Hours_Pass_Number"
	//7天前的过题数
	vjPersonLast7DaysPassNumber = "vj_Person_Last_7_Days_Pass_Number"
	//30天前的过题数
	vjPersonLast30DaysPassNumber = "vj_Person_Last_30_Days_Pass_Number"
	//总过题数
	vjPersonPassNumber = "vj_Person_Pass_Number"
)

//获取24小时前的过题数
func vjPersonLast24HoursPassNumberHandler(doc *goquery.Selection) string {
	retStr := doc.Find(fmt.Sprintf(".container a[title=\"New solved in last 24 hours\"]")).First().Text()
	return retStr
}

//获取7天前的过题数
func vjPersonLast7DaysPassNumberHandler(doc *goquery.Selection) string {
	retStr := doc.Find(fmt.Sprintf(".container a[title=\"New solved in last 7 days\"]")).First().Text()
	return retStr
} //获取一个月前的过题数
func vjPersonLast30DaysPassNumberHandler(doc *goquery.Selection) string {
	retStr := doc.Find(fmt.Sprintf(".container a[title=\"New solved in last 30 days\"]")).First().Text()
	return retStr
} //获取总的过题数
func vjPersonPassNumberHandler(doc *goquery.Selection) string {
	retStr := doc.Find(fmt.Sprintf(".container a[title=\"Overall solved\"]")).First().Text()
	return retStr
}

//---------------------------------------------------------------------//
// 部分共用函数 //
//---------------------------------------------------------------------//
//转化int
func strToInt(doc *goquery.Selection, f func(doc *goquery.Selection) string) int {
	ret := f(doc)
	num, err := strconv.Atoi(ret)
	if err != nil {
		log.Errorf("VJ strToInt get err:%v\tand the return is %v:", num, err)
		return -1
	}
	return num
}
func getPersonPage(uid string) string {
	return "https://vjudge.net/user/" + uid
}
