package luogu

import (
	"XCPCer_board/db/mysql"
	"XCPCer_board/scraper"
	"strconv"
)

func ScrapeAll(uid string) (map[string]int, error) {
	// 请求所有并合并所有
	res, err := scraper.MergeAllResults[string, int](
		GetStrMsg(uid),
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func InsertSql(uid string) error {
	res, err := scraper.MergeAllResults[string, int](
		GetStrMsg(uid),
	)
	if err != nil {
		return err
	}
	var mp map[string]int
	mp = res
	n := len(mp)
	sqlStr := "insert into " + "lg " + "(user_luogu_id, "
	keyList := [10]string{}
	valList := [10]any{}
	valList[0], err = strconv.Atoi(uid)
	if err != nil {
		return err
	}
	cnt := 0
	for key, val := range res {
		keyList[cnt] = key
		valList[cnt] = val
		cnt++
	}
	for i, v := range keyList {
		sqlStr += v
		if i < n-1 {
			sqlStr += ", "
		}
	}
	sqlStr += ")values (?,"
	for i := 0; i < n; i++ {
		sqlStr += "?"
		if i < n-1 {
			sqlStr += ","
		}
	}
	sqlStr += ")"

	//fmt.Printf(sqlStr)
	//") values (?,?)"
	ret, err := mysql.MySql.Exec(sqlStr, valList[0:n+1]...)
	if err != nil {
		//fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	_, err = ret.LastInsertId() // 新插入数据的id
	if err != nil {
		//fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return err
	}
	//fmt.Printf("insert success, the id is %d.\n", theID)

	return nil
}
