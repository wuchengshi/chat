package service

import (
	"chat/util"
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lunny/godbc"
)

var (
	showSQL  = true
	maxCon   = 10
	NONERROR = "noerror" //没有错误
)

var DbEngin *xorm.Engine

func init() {
	err := errors.New(NONERROR)

	if util.Db == "sqlserver" {
		dsName := fmt.Sprintf("driver={SQL Server};Server=%s;Database=%s;uid=%s;pwd=%s;", util.DbHost, util.DbName, util.DbUser, util.DbPassWord)
		DbEngin, err = xorm.NewEngine("odbc", dsName)
		if nil != err && NONERROR != err.Error() {
			log.Fatal(err.Error())
		}
		if err := DbEngin.Ping(); err != nil {
			fmt.Println(err)
			return
		}
		//是否显示SQL语句
		DbEngin.ShowSQL(showSQL)
		//数据库最大打开的连接数
		DbEngin.SetMaxOpenConns(maxCon)
		if err != nil {
			fmt.Println(err)
		}
	} else if util.Db == "mysql" {
		dsName := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8", util.DbUser, util.DbPassWord, util.DbHost, util.DbPort, util.DbName)

		DbEngin, err = xorm.NewEngine("mysql", dsName)
		if nil != err && NONERROR != err.Error() {
			log.Fatal(err.Error())
		}
		//是否显示SQL语句
		DbEngin.ShowSQL(showSQL)
		//数据库最大打开的连接数
		DbEngin.SetMaxOpenConns(maxCon)

		//自动
		//DbEngin.Sync2(new(model.AdminUser), new(model.Contact), new(model.Record), new(model.Community))
	}

	fmt.Println("init data base ok")
}
