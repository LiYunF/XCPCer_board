package codeforces

import (
	"XCPCer_board/spider/codeforcesv2"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
)

func createTableName() error {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)"+
		"/codeforces?charset=utf8&multiStatements=true")
	if err != nil {
		log.Fatal("fail to start mysql", err)
		return err
	}
	defer func() {
		if db != nil {
			db.Close()
		}
	}()
	Trans, err := db.Begin()
	if err != nil {
		log.Fatal("fail to start transaction", err)
		return err
	}
	fmt.Println(Trans)
	path := "createCodeforces.sql"
	sqlByte, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("open file failure", err)
		return err
	}
	sqlTable := string(sqlByte)
	_, err = db.Exec(sqlTable)
	if err != nil {
		log.Fatal("sql create err", err)
		Trans.Rollback()
		return err
	}
	return nil
}

func insertTable(mp map[string]codeforcesv2.Submission, useName string) error {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)"+
		"/codeforces?charset=utf8&multiStatements=true")

	if err != nil {
		return err
	}
	defer db.Close()

	sql, _ := db.Prepare("insert into cf (`Id`,`user_id`,`problem_index`,`contest_id`,`rating`,`problem_name`)value(?,?,?,?,?,?)")
	for _, i := range mp {
		sql.Exec(i.Id, useName, i.Problem.Index, i.ContestId, i.Problem.Rating, i.Problem.Name)
	}
	return nil
}
