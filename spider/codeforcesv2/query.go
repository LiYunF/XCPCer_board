package codeforcesv2

import (
	"XCPCer_board/db/mysql"
)

/*
------------------------------
@Time : 5/3/2022 7:36 PM
@Author : Sariel Crescent
@File : query
@Software: GoLand
----------------------------
*/
type rate struct {
	sum       int
	countRate [50]int
}

var (
	name  string
	count int
	num   [50]int
)

//人与题目相关联写一个查询函数返回 账号&codeforces过题总数
func userSumNumber() (map[string]int, error) {
	db := mysql.Db
	str := "select `user_id` ,count(*) from codeforces group by `user_id` order by count(*);"
	qry, err := db.Query(str)
	if err != nil {
		return nil, err
	}
	var mp map[string]int
	for qry.Next() {
		if err := qry.Scan(&name, &count); err != nil {
			return nil, err
		}
		mp[name] = count
	}
	return mp, nil
}

//人与难度相关联再来写一个查询函数返回map
func userRatingProblem() (map[string]rate, error) {
	db := mysql.Db
	str := "select `user_id`,`rating` from codeforces"
	qry, err := db.Query(str)
	if err != nil {
		return nil, err
	}
	var mp map[string]rate
	for qry.Next() {
		if err := qry.Scan(&name, &count); err != nil {
			return nil, err
		}
		r := mp[name]
		r.countRate[count/100+1]++
		mp[name] = r
	}
	return mp, err
}
