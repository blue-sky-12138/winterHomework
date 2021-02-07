package model

import (
	"WinterHomework/database"
	"WinterHomework/utilities"
	"regexp"
	"strconv"
	"strings"
)

//获取用户密码
func GetUserPassword(detail string) *utilities.LoginCheck {
	if strings.Contains(detail,"@"){		//判断是否有@，有即为邮箱登录
		return database.GetPassword("email","'"+detail+"'")
	}else{
		reg:=regexp.MustCompile("[^0-9]")	//判断是否为纯数字，是即为手机号登录
		if reg.MatchString(detail){
			return database.GetPassword("name","'"+detail+"'")
		}else{
			return database.GetPassword("telephone",detail)
		}
	}
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

//更新生日
func ChangeUserBirthday(name string,detail int64) bool {
	return database.ChangeUserInformation(name,"birthday",strconv.FormatInt(detail,10))
}
//更新签名
func ChangeUserSignature(name string, detail string) bool {
	return database.ChangeUserInformation(name,"signature","'"+detail+"'")
}
//更新性别
func ChangeUserGender(name string,detail int) bool {
	return database.ChangeUserInformation(name,"gender",strconv.Itoa(detail))
}
//更新昵称
func ChangeUserNickname(name string, detail string) bool {
	return database.ChangeUserInformation(name,"nickname","'"+detail+"'")
}
//更新邮箱
func ChangeUserEmail(name string,detail string) bool {
	return database.ChangeUserInformation(name,"email","'"+detail+"'")
}
//更新手机号
func ChangeUserTelephone(name string,detail int64) bool {
	return database.ChangeUserInformation(name,"telephone",strconv.FormatInt(detail,10))
}