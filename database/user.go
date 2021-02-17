package database

import (
	"WinterHomework/utilities"
	"fmt"
)

//获取密码。
func GetPassword(focus string,detail string) (*utilities.LoginCheck, error) {
	var tem utilities.LoginCheck
	pre:=fmt.Sprintf("select id,password,md5salt from users_information where %s =%s",focus,detail)
	rows,err:=DB.Query(pre)
	defer rows.Close()
	if err != nil{
		utilities.LogError("GetPassword Error",err)
		return nil, fmt.Errorf("未知错误#gp0016")
	}
	if rows.Next(){
		rows.Scan(&tem.Id,&tem.Password,&tem.Md5salt)
	}
	return &tem, nil
}

//录入注册信息。
func LogRegisterData(name string, telephone int64, email string, password string, MD5salt int64) error {
	pre:=fmt.Sprintf("insert users_information(name,telephone,email,password,md5salt) value('%s',%v,'%s','%s',%v)",
		name,telephone,email,password,MD5salt)
	stmt,err:=DB.Prepare(pre)
	defer stmt.Close()
	if err != nil {
		utilities.LogError("Prepare LogRegisterData Error",err)
		return fmt.Errorf("未知错误#lrd0032")
	}

	_,err=stmt.Exec()
	if err != nil {
		utilities.LogError("Insert LogRegisterData Error",err)
		return fmt.Errorf("未知错误#lgd0038")
	}

	return nil
}

//获取某个用户的某个信息。
//返回：数据（如果查找到），是否查询到目标信息，错误（如果遇到）。
func GetUserSingleInformation(require string, focus string, detail string) (string, bool, error) {
	pre:=fmt.Sprintf("select %s from users_information where %s =%s",require,focus,detail)
	rows,err:=DB.Query(pre)
	defer rows.Close()
	if err != nil {
		utilities.LogError("GetUserSingleInformation Error",err)
		return "", false, fmt.Errorf("未知错误#gusi0052")
	}
	var tem string
	if rows.Next(){			//如果查询到数据
		rows.Scan(&tem)
		return tem, true, nil
	}else {					//如果没有查询到数据
		return "", false, nil
	}
}

//更新用户信息。
func ChangeUserInformation(name string,focus string,detail string) error {
	pre:=fmt.Sprintf("update users_information set %s = %s where name = '%s' ",focus,detail,name)
	stmt,err:=DB.Prepare(pre)
	if err!=nil{
		utilities.LogError("ChangeUserInformation Error",err)
		return fmt.Errorf("未知错误#cui0069")
	}
	stmt.Exec()
	return nil
}