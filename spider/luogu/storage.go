package luogu

import (
	"XCPCer_board/db/mysql"
	"fmt"
	log "github.com/sirupsen/logrus"
)

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
	//get struct
	user, ok := MapToStruct(res)
	if !ok {
		log.Errorf("Map of " + uid + "is incomplete\n")
	}
	user.Uid = uid
	//insert data
	_, err = mysql.Db.Exec(sqlStr, user.Uid, user.PassProblemNumber, user.Ranting,
		user.SimpleProblem, user.BasicProblem, user.ElevatedProblem, user.HardProblem,
		user.UnKnowProblem)
	if err != nil {
		log.Errorf("insert failed, err:%v\n", err)
		return err
	}

	return nil
}

//QuerySql 查询
func QuerySql(uid string) error {
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
		log.Errorf("scan failed, err: %v\n", err)
		return err
	}

	//log.Printf("data of :%v\n "+
	//	"PassProblemNumber:%v\n"+
	//	"Ranting:%v\n"+
	//	"SimpleProblemPass:%v\n"+
	//	"BasicProblemPass:%v\n"+
	//	"ElevatedProblemPass:%v\n"+
	//	"HardProblemPass:%v\n"+
	//	"UnKnowProblemPass:%v\n", uid, user.PassProblemNumber, user.Ranting,
	//	user.SimpleProblem, user.BasicProblem, user.ElevatedProblem, user.HardProblem,
	//	user.UnKnowProblem)

	return err
}

//UpdateSql 更新
func UpdateSql(age int, id int) {
	sqlStr := "update lg set problem_number=?, ranting=?, simple_problem_number=?," +
		"base_problem_number=?, elevated_problem_number=?, hard_problem_number=?, " +
		"unKnow_problem_number=? where uid=?"
	ret, err := mysql.Db.Exec(sqlStr, age, id)
	if err != nil {
		fmt.Printf("update failed, err: %v\n", err)
	}
	n, err := ret.RowsAffected() // 操作收影响的行
	if err != nil {
		fmt.Printf("get RowsAffected failed: %v\n", err)
		return
	}
	fmt.Printf("update success, affected rows: %d\n", n)

}
