package database

import(
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var(
	DB *sql.DB
	LogFile *log.Logger
)

func OpenMysql(){
	mysql:="bluesky:135246Cjw@tcp(rm-bp14fk5x3q4byb6a2so.mysql.rds.aliyuncs.com:3306)"
	db,err:=sql.Open("mysql",mysql+"/winter_homework?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		fmt.Println(err)
		panic("数据库连接错误")
	}
	DB=db
}

func LogInit(){
	File,err:=os.OpenFile("logFile.md" ,os.O_RDWR, os.ModePerm)
	if err!=nil{
		fmt.Println(err)
		panic("打开日志失败")
	}
	logFile := log.New(File,"[GIN-debug]",log.Ldate | log.Ltime)
	LogFile =logFile
}

func LogError(detail string,err error){
	detail="["+detail+"]"
	log.Printf(detail,err,"\n")

	LogFile.Println(detail,err)
}