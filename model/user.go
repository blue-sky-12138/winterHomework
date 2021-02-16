package model

import (
	"WinterHomework/database"
	"WinterHomework/utilities"
	"regexp"
	"strconv"
	"strings"
)

//获取用户密码。
func GetUserPassword(detail string) (*utilities.LoginCheck, error) {
	if strings.Contains(detail,"@"){		//判断是否有@，有即为邮箱登录
		data, err := database.GetPassword("email","'"+detail+"'")
		if err == nil {
			return data, nil
		}else {
			return nil, err
		}
	}else{
		reg:=regexp.MustCompile("[^0-9]")	//判断是否为纯数字，是即为手机号登录
		if reg.MatchString(detail){
			data, err := database.GetPassword("name","'"+detail+"'")
			if err == nil {
				return data, nil
			}else {
				return nil, err
			}
		}else {
			data, err := database.GetPassword("telephone", detail)
			if err == nil {
				return data, nil
			} else {
				return nil, err
			}
		}
	}
}

//检查是否已被注册。
//返回：是否查询到目标数据，错误（如果有）。
func CheckRegisterOrNot(focus string,detail string) (bool, error) {
	if focus=="telephone"{
		_, ok, err := database.GetUserSingleInformation(focus,focus,detail)
		if err == nil {
			return ok, nil
		}else {
			return ok, err
		}
	}else {
		_, ok, err := database.GetUserSingleInformation(focus,focus,"'"+detail+"'")
		if err == nil {
			return ok, nil
		}else {
			return ok, err
		}
	}
}

//注册。
func PostRegisterInformation(name string,telephone int64,email string,password string,MD5salt int64) error {
	return database.LogRegisterData(name,telephone,email,password,MD5salt)
}

//更新生日。
func ChangeUserBirthday(name string,detail int64) error {
	return database.ChangeUserInformation(name,"birthday",strconv.FormatInt(detail,10))
}
//更新签名。
func ChangeUserSignature(name string, detail string) error {
	return database.ChangeUserInformation(name,"signature","'"+detail+"'")
}
//更新性别。
func ChangeUserGender(name string,detail int) error {
	return database.ChangeUserInformation(name,"gender",strconv.Itoa(detail))
}
//更新昵称。
func ChangeUserNickname(name string, detail string) error {
	return database.ChangeUserInformation(name,"nickname","'"+detail+"'")
}
//更新邮箱。
func ChangeUserEmail(name string,detail string) error {
	return database.ChangeUserInformation(name,"email","'"+detail+"'")
}
//更新手机号。
func ChangeUserTelephone(name string,detail int64) error {
	return database.ChangeUserInformation(name,"telephone",strconv.FormatInt(detail,10))
}