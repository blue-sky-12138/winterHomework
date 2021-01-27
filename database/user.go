package database

import "fmt"

type UserInformation struct {
	UserName string `json:"name"`

}
type CheckPassword struct {
	Password string
	Md5salt  int64
}

//这是用email登录的快捷方式
func GetPasswordFromEmail(email string) *CheckPassword{
	return GetPassword("email","'"+email+"'")
}
//这是用手机号登录的快捷方式
func GetPasswordFromTelephone(telephone string) *CheckPassword{
	return GetPassword("telephone",telephone)
}
//这是用用户名登录的快捷方式
func GetPasswordFromName(name string) *CheckPassword{
	return GetPassword("name","'"+name+"'")
}
//获取密码
func GetPassword(focus string,detail string) *CheckPassword {
	var tem CheckPassword

	pre:=fmt.Sprintf("select password,md5salt from users_information where %s =%s",focus,detail)
	stmt,err:=DB.Query(pre)
	defer stmt.Close()
	if err != nil{
		LogError("GetPassword Error",err)
	}
	if stmt.Next(){
		stmt.Scan(&tem.Password,&tem.Md5salt)
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
		LogError("Prepare LogRegisterData Error",err)
		return false
	}

	_,err=stmt.Exec()
	if err != nil {
		LogError("Insert LogRegisterData Error",err)
		return false
	}

	return true
}

//获取某个用户的某个信息
func GetUserSingleInformation(require string, focus string, detail string) string {
	pre:=fmt.Sprintf("select %s from users_information where %s =%s",require,focus,detail)
	stmt,err:=DB.Query(pre)
	defer stmt.Close()
	if err != nil {
		LogError("GetUserSingleInformation Error",err)
	}
	var tem string
	if stmt.Next(){
		stmt.Scan(&tem)
	}
	return tem
}