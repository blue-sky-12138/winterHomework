package controller

import (
	"WinterHomework/database"
	"WinterHomework/model"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

type LoginCheck struct {
	Name string 	`json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
type RegisterInformation struct {
	Name string		`json:"username" form:"username"`
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

	if tem.Email != ""{
		if !RegisterEmailCheck(tem.Email){
			context.JSON(200,gin.H{
				"code":		30012,
				"message":	"邮箱错误",
			})
			return
		}else if model.CheckRegisterOrNot("email",tem.Email) != ""{
			context.JSON(200,gin.H{
				"code":		30002,
				"message":	"该邮箱已被注册",
			})
			return
		}
	}else if tem.Telephone != 0 {
		if !RegisterTelephoneCheck(tem.Telephone){
			context.JSON(200,gin.H{
				"code":		30011,
				"message":	"手机号错误",
			})
			return
		}else if model.CheckRegisterOrNot("telephone",strconv.FormatInt(tem.Telephone, 10)) != "" {
			context.JSON(200,gin.H{
				"code":		30001,
				"message":	"该手机号已被注册",
			})
			return
		}
	}

	if tem.Name!=""{
		if !RegisterUserNameCheck(tem.Name){
			context.JSON(200,gin.H{
				"code":		30013,
				"message":	"用户名不符合要求",
			})
			return
		}else if model.CheckRegisterOrNot("name", tem.Name) != "" {
			context.JSON(200,gin.H{
				"code":		30003,
				"message":	"该用户名已被注册",
			})
			return
		}
	}


	if !RegisterPasswordCheck(tem.Password){
		context.JSON(200,gin.H{
			"code":		30014,
			"message":	"密码不规范",
		})
		return
	}

	password,MD5salt:=CryptographyNow(tem.Password)

	if model.PostRegisterInformation(tem.Name,tem.Telephone,tem.Email,password,MD5salt) {
		context.JSON(200,gin.H{
			"code":		300,
			"message":	"注册成功",
		})
	}else{
		context.JSON(200,gin.H{
			"code":		30015,
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


//注册手机号正确性检查
func RegisterTelephoneCheck(number int64) bool {
	reg:=regexp.MustCompile("(13|14|15|17|18|19)[0-9]{9}")
	if reg.MatchString(strconv.FormatInt(number,10)){
		return true
	}else{
		return false
	}
}
//注册用户名有效性检查
//只能包括数字、字母、汉字的组合，且不能只包含数字
func RegisterUserNameCheck(name string) bool {
	reg:=regexp.MustCompile("[^A-Za-z0-9\u4e00-\u9fa5]")
	if !reg.MatchString(name) {
		reg=regexp.MustCompile("[^0-9]")
		if reg.MatchString(name){
			return true
		}
		return false
	}else{
		return false
	}
}
//邮箱正确性检测
func RegisterEmailCheck(email string) bool {
	reg:=regexp.MustCompile("\\w[-\\w.+]*@([A-Za-z0-9][-A-Za-z0-9]+\\.)+[A-Za-z]{2,14}")
	if reg.MatchString(email){
		return true
	}else{
		return false
	}
}
//注册密码规范性检查
//最少8位,最多16位，只能包括数字、字母、下划线
func RegisterPasswordCheck(password string) bool {
	reg:=regexp.MustCompile("[^\\w]")
	if len(password) >= 8 && len(password) <= 16 && !reg.MatchString(password){
		return true
	}else{
		return false
	}
}