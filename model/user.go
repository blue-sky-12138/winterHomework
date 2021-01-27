package model

import (
	"WinterHomework/database"
	"regexp"
	"strings"
)

//获取用户密码
func GetUserPassword(detail string)*database.CheckPassword{
	var tem *database.CheckPassword
	if strings.Contains(detail,"@"){
		tem=database.GetPasswordFromEmail(detail)
	}else{
		reg:=regexp.MustCompile("[^0-9]")
		if reg.MatchString(detail){
			tem=database.GetPasswordFromName(detail)
		}else{
			tem=database.GetPasswordFromTelephone(detail)
		}
	}
	return tem
}

//检查是否已被注册
func CheckRegisterOrNot(focus string,detail string) string {
	if focus=="telephone"{
		return database.GetUserSingleInformation(focus,focus,detail)
	}else {
		return database.GetUserSingleInformation(focus,focus,"'"+detail+"'")
	}
}

//注册
func PostRegisterInformation(name string,telephone int64,email string,password string,MD5salt int64) bool {
	return database.LogRegisterData(name,telephone,email,password,MD5salt)
}