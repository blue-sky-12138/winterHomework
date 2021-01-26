package controller

import (
	"WinterHomework/database"
	"WinterHomework/model"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
	"time"
)

type LoginCheck struct {
	Name string 	`json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}
type RegisterInformation struct {
	Name string		`json:"name" form:"name"`
	Password string	`json:"password" form:"password"`
	Telephone int64	`json:"telephone" form:"telephone"`
	Email string	`json:"email" form:"email"`
}
type Resp struct {
	Code int 			`json:"code"`
	Message string 		`json:"message"`
	Data interface{} 	`json:"data"`
}

//用户登录
func PostLogin(context *gin.Context){
	var login LoginCheck
	err:=context.ShouldBind(&login)
	if err!=nil{
		database.LogError("GetLoginInformation Error",err)
	}

	check:=model.GetUserPassword(login.Name)
	if check.Md5salt == 0{
		context.JSON(200,gin.H{
			"code":		20001,
			"message":	"该用户不存在",
		})
	}else if Cryptography(login.Password,check.Md5salt) != check.Password{
		context.JSON(200,gin.H{
			"code":		20002,
			"message":	"密码错误",
		})
	}else{
		cookie := &http.Cookie{
			Name:     	"user",
			Value:  	login.Name,
			MaxAge:	 	100000,
			Path:     	"/",
			Domain: 	"localhost",
			Secure: 	false,
			HttpOnly: 	true,
		}
		http.SetCookie(context.Writer,cookie)

		context.JSON(200, gin.H{
			"code":    	200,
			"message": 	"登陆成功",
		})
	}
}

//用户注册
func Register(context *gin.Context){
	var tem RegisterInformation
	err:=context.ShouldBind(&tem)
	if err!=nil{
		database.LogError("GetRegisterInformation Error",err)
	}

	if tem.Email != "" && model.CheckRegisterOrNot(tem.Email) != ""{
		if tem.Telephone != 0 && model.CheckRegisterOrNot(strconv.FormatInt(tem.Telephone,10)) != ""{
			context.JSON(200,gin.H{
				"code":		30003,
				"message":	"手机号和邮箱都已被注册",
			})
		}else{
			context.JSON(200,gin.H{
				"code":		30002,
				"message":	"邮箱已被注册",
			})
		}
	}else if tem.Telephone != 0 && model.CheckRegisterOrNot(strconv.FormatInt(tem.Telephone, 10)) != "" {
		context.JSON(200,gin.H{
			"code":		30001,
			"message":	"手机号已被注册",
		})
	}

	password,MD5salt:=CryptographyNow(tem.Password)
	if model.PostRegisterInformation(tem.Name,tem.Telephone,tem.Email,password,MD5salt) {
		context.JSON(200,gin.H{
			"code":		300,
			"message":	"注册成功",
		})
	}else{
		context.JSON(200,gin.H{
			"code":		30004,
			"message":	"未知错误",
		})
	}
}

//MD5加密的快捷方式
func CryptographyNow(Data string) (string,int64) {
	Md5salt:=time.Now().Unix()
	return Cryptography(Data,Md5salt),Md5salt
}
//MD5加密
func Cryptography(Data string,Md5salt int64)string{
	has:=md5.New()
	io.WriteString(has,Data+strconv.FormatInt(Md5salt,10))
	tem:= has.Sum(nil)
	Result:=hex.EncodeToString(tem)
	return Result
}