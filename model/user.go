package model

import (
	"WinterHomework/database"
	"strings"
)

//获取用户密码
func GetUserPassword(detail string)*database.CheckPassword{
	var tem *database.CheckPassword
	if strings.Contains(detail,"@"){
		tem=database.GetPasswordFromEmail(detail)
	}else {
		tem=database.GetPasswordFromTelephone(detail)
	}
	return tem
}

//检查是否已被注册
func CheckRegisterOrNot(detail string) string {
	if strings.Contains(detail,"@"){
		return database.GetUserSingleInformation("email","email","'"+detail+"'")
	}else {
		return database.GetUserSingleInformation("telephone","telephone",detail)
	}
}

//注册
func PostRegisterInformation(name string,telephone int64,email string,password string,MD5salt int64) bool {
	return database.LogRegisterData(name,telephone,email,password,MD5salt)
}