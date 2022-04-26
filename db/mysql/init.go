package mysql

import (
	"XCPCer_board/config"
	"database/sql"
	log "github.com/sirupsen/logrus"
)

var MySql *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
	password:=config.Config
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname"
	MySql, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Errorf("Open Sql Error: %v", err)
		return err
	}

	return nil
}
