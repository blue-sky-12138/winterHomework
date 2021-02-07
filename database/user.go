package database

import (
	"WinterHomework/utilities"
	"fmt"
)

//获取密码
func GetPassword(focus string,detail string) *utilities.LoginCheck {
	var tem utilities.LoginCheck
	pre:=fmt.Sprintf("select password,md5salt from users_information where %s =%s",focus,detail)
	rows,err:=DB.Query(pre)
	defer rows.Close()
	if err != nil{
		utilities.LogError("GetPassword Error",err)
	}
	if rows.Next(){
		rows.Scan(&tem.Password,&tem.Md5salt)
	}
	return &tem
}

//录入注册信息
func LogRegisterData(name string, telephone int64, email string, password string, MD5salt int64) bool {
	pre:=fmt.Sprintf("insert users_information(name,telephone,email,password,md5salt) value('%s',%v,'%s','%s',%v)",
		name,telephone,email,password,MD5salt)
	stmt,err:=DB.Prepare(pre)
	defer stmt.Close()
	if err != nil {
		utilities.LogError("Prepare LogRegisterData Error",err)
		return false
	}

	_,err=stmt.Exec()
	if err != nil {
		utilities.LogError("Insert LogRegisterData Error",err)
		return false
	}

	return true
}

//获取某个用户的某个信息
func GetUserSingleInformation(require string, focus string, detail string) string {
	pre:=fmt.Sprintf("select %s from users_information where %s =%s",require,focus,detail)
	rows,err:=DB.Query(pre)
	defer rows.Close()
	if err != nil {
		utilities.LogError("GetUserSingleInformation Error",err)
	}
	var tem string
	if rows.Next(){
		rows.Scan(&tem)
	}
	return tem
}

//更新用户信息
func ChangeUserInformation(name string,focus string,detail string) bool{
	pre:=fmt.Sprintf("update users_information set %s = %s where name ='%s'",focus,detail,name)
	stmt,err:=DB.Prepare(pre)
	if err!=nil{
		utilities.LogError("ChangeUserInformation Error",err)
		return false
	}
	stmt.Exec()
	return true
}