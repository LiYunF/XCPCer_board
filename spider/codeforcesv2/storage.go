package codeforcesv2

import (
	"XCPCer_board/db/mysql"
	_ "XCPCer_board/db/mysql"
	_ "github.com/go-sql-driver/mysql"
)

//数据库增加单人过题数
func insertTable(mp map[string]Submission, useName string) error {
	db := mysql.Db

	sql, _ := db.Prepare("insert into cf " +
		"(`Id`,`user_id`,`problem_index`," +
		"`contest_id`,`rating`,`problem_name`)" +
		"value(?,?,?,?,?,?)")

	for _, i := range mp {
		sql.Exec(i.Id, useName, i.Problem.Index, i.ContestId, i.Problem.Rating, i.Problem.Name)
	}
	return nil
}
