package mysql

import (
	"XCPCer_board/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

var MySql *sql.DB

//InitDB 初始化数据库的函数
func InitDB() (err error) {
	msg := config.Config.Database.Msg
	password := msg.Password
	host := msg.Host
	dbname := msg.Name
	user := msg.Username
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v", user, password, host, dbname)
	//dsn := "user:password@tcp(127.0.0.1:3306)/dbname"

	MySql, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Errorf("Open Sql Error: %v", err)
		return err
	}
	//fmt.Println("成功打开MYSQL")
	MySql.SetMaxOpenConns(20)
	MySql.SetMaxIdleConns(10)
	return nil
}
