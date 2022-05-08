package mysql

import (
	"XCPCer_board/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

var MySql *sql.DB


//初始化数据库的函数
func init() {
	msg := config.Config.Database.MysqlConf

	password := msg.Password
	host := msg.Host
	dbname := msg.Name
	user := msg.Username
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v", user, password, host, dbname)

	var err error
	Db, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Errorf("Open Sql Error: %v", err)
		return err
	}
	//fmt.Println("成功打开MYSQL")
	MySql.SetMaxOpenConns(20)
	MySql.SetMaxIdleConns(10)
	return nil
}
