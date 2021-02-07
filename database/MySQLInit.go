package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var(
	DB *sql.DB
)

func OpenMysql(){
	mysql:="bluesky:135246Cjw@tcp(rm-bp14fk5x3q4byb6a2so.mysql.rds.aliyuncs.com:3306)"
	db,err:=sql.Open("mysql",mysql+"/winter_homework?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		fmt.Println(err)
		panic("数据库连接错误")
	}
	db.SetMaxOpenConns(1000)
	db.SetConnMaxIdleTime(100)
	DB=db
}