package luogu

import (
	"XCPCer_board/db/mysql"
	log "github.com/sirupsen/logrus"
)

/////////////用户结构体///////////////

type UserMsg struct {
	Uid               string
	PassProblemNumber int
	Ranting           int
	SimpleProblem     int
	BasicProblem      int
	ElevatedProblem   int
	HardProblem       int
	UnKnowProblem     int
}

//StructToMap 结构体转Map
func StructToMap(user UserMsg) map[string]int {
	mp := map[string]int{
		passProblemNumber: user.PassProblemNumber,
		ranting:           user.Ranting,
		simpleProblem:     user.SimpleProblem,
		basicProblem:      user.BasicProblem,
		elevatedProblem:   user.ElevatedProblem,
		hardProblem:       user.HardProblem,
		unKnowProblem:     user.UnKnowProblem,
	}
	return mp
}

//MapToStruct Map转结构体, 返回的bool=1为正常，0为map里没有该值
func MapToStruct(mp map[string]int) (UserMsg, bool) {

	var user UserMsg
	var ok bool
	if user.PassProblemNumber, ok = mp[passProblemNumber]; !ok {
		return user, ok
	}
	if user.Ranting, ok = mp[ranting]; !ok {
		return user, ok
	}
	if user.SimpleProblem, ok = mp[simpleProblem]; !ok {
		return user, ok
	}
	if user.BasicProblem, ok = mp[basicProblem]; !ok {
		return user, ok
	}
	if user.ElevatedProblem, ok = mp[elevatedProblem]; !ok {
		return user, ok
	}
	if user.HardProblem, ok = mp[hardProblem]; !ok {
		return user, ok
	}
	if user.UnKnowProblem, ok = mp[unKnowProblem]; !ok {
		return user, ok
	}
	return user, true
}

/////////////增删改查///////////////

//Storage for mock
type Storage interface {
	InsertSql(uid string) error
	QuerySql(uid string) (map[string]int, error)
	UpdateSql(uid string) error
	DeleteSql(uid string) error
}

//InsertSql 插入
func InsertSql(uid string) error {

	//get data
	res, err := ScrapeAll(uid)
	if err != nil {
		return err
	}
	//get str
	sqlStr := "insert into lg(uid, problem_number, ranting, simple_problem_number" +
		",base_problem_number,elevated_problem_number,hard_problem_number," +
		"unKnow_problem_number) values (?,?,?,?,?,?,?,?)"
	//check the len of map
	if len(res) != 7 {
		log.Errorf("the len of uid=%v 's map is not correct\n", uid)
	}
	//insert data
	_, err = mysql.Db.Exec(sqlStr, uid, res[passProblemNumber], res[ranting],
		res[simpleProblem], res[basicProblem], res[elevatedProblem], res[hardProblem],
		res[unKnowProblem])
	if err != nil {
		log.Errorf("database insert luogu uid=%v failed,  err:%v\n", uid, err)
		return err
	}

	return nil
}

//QuerySql 查询
func QuerySql(uid string) (map[string]int, error) {
	// 获取单条数据
	sqlStr := "select problem_number, ranting, simple_problem_number" +
		",base_problem_number,elevated_problem_number,hard_problem_number," +
		"unKnow_problem_number from lg where uid=?"
	var user UserMsg
	user.Uid = uid
	// 非常重要:确保QueryRow之后调用Scan方法,否则持有数据的连接不会被释放
	err := mysql.Db.QueryRow(sqlStr, uid).Scan(&user.PassProblemNumber, &user.Ranting,
		&user.SimpleProblem, &user.BasicProblem, &user.ElevatedProblem, &user.HardProblem,
		&user.UnKnowProblem)
	if err != nil {
		log.Errorf("database Query luogu uid=%v failed,  err:%v\n", uid, err)
		return nil, err
	}
	mp := StructToMap(user)
	return mp, nil
}

//UpdateSql 更新
func UpdateSql(uid string) error {
	res, err := ScrapeAll(uid)
	if err != nil {
		return err
	}
	//get str
	sqlStr := "update lg set problem_number=?, ranting=?, simple_problem_number=?," +
		"base_problem_number=?, elevated_problem_number=?, hard_problem_number=?, " +
		"unKnow_problem_number=? where uid=?"
	//check the len of map
	if len(res) != 7 {
		log.Errorf("the len of uid=%v 's map is not correct\n", uid)
	}
	//insert data
	_, err = mysql.Db.Exec(sqlStr, uid, res[passProblemNumber], res[ranting],
		res[simpleProblem], res[basicProblem], res[elevatedProblem], res[hardProblem],
		res[unKnowProblem])
	if err != nil {
		log.Errorf("database Update luogu uid=%v failed,  err:%v\n", uid, err)
		return err
	}
	return nil
}

//DeleteSql 删除
func DeleteSql(uid string) error {

	sqlStr := "delete from lg where uid=?"
	//get struct
	_, err := mysql.Db.Exec(sqlStr, uid)
	if err != nil {
		log.Errorf("database Delete luogu uid=%v failed,  err:%v\n", uid, err)
		return err
	}
	return nil
}
