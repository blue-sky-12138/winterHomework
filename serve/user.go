package serve

import (
	"WinterHomework/model"
	"WinterHomework/utilities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//用户登录
func PostLogin(context *gin.Context){
	var (
		login utilities.LoginCheck
		resp utilities.Resp
	)
	err:=context.ShouldBind(&login)
	if err!=nil{
		utilities.LogError("GetLoginInformation Error",err)
	}

	check:=model.GetUserPassword(login.Name)
	if check.Md5salt == 0{		//检查用户名是否存在
		resp.Code =	 20001
		resp.Message = "该用户不存在"
		context.JSON(200,resp)
	}else if utilities.Cryptography(login.Password,check.Md5salt) != check.Password{	//检查密码是否正确
		resp.Code =	 20002
		resp.Message = "密码错误"
		context.JSON(200,resp)
	}else{
		cookie := &http.Cookie{
			Name:     	"user",
			Value:  	login.Name,
			MaxAge:	 	100000,
			Path:     	"/",
			Secure: 	false,
			HttpOnly: 	true,
		}
		http.SetCookie(context.Writer,cookie)

		resp.Code = 200
		resp.Message = "登陆成功"
		context.JSON(200, resp)
	}
}

//用户注册
func Register(context *gin.Context){
	var (
		tem utilities.RegisterInformation
		resp utilities.Resp
	)
	err:=context.ShouldBind(&tem)
	if err!=nil{
		utilities.LogError("GetRegisterInformation Error",err)
	}

	//检查邮箱的正确性、是否存在
	if tem.Email != ""{
		if !utilities.RegisterEmailCheck(tem.Email){
			resp.Code = 30012
			resp.Message = "邮箱错误"
			context.JSON(200,resp)
			return
		}else if model.CheckRegisterOrNot("email",tem.Email) != ""{
			resp.Code = 30002
			resp.Message = "该邮箱已被注册"
			context.JSON(200,resp)
			return
		}
	}

	//检查手机号的正确性、是否存在
	if tem.Telephone != 0 {
		if !utilities.RegisterTelephoneCheck(tem.Telephone){
			resp.Code = 30011
			resp.Message = "手机号错误"
			context.JSON(200,resp)
			return
		}else if model.CheckRegisterOrNot("telephone",strconv.FormatInt(tem.Telephone, 10)) != "" {
			resp.Code = 30001
			resp.Message = "该手机号已被注册"
			context.JSON(200,resp)
			return
		}
	}

	//检查用户名的正确性、是否存在
	if tem.Name!= ""{
		if !utilities.RegisterUserNameCheck(tem.Name){
			resp.Code = 30013
			resp.Message = "用户名不符合要求"
			context.JSON(200,resp)
			return
		}else if model.CheckRegisterOrNot("name", tem.Name) != "" {
			resp.Code = 30003
			resp.Message = "该用户名已被注册"
			context.JSON(200,resp)
			return
		}
	}

	//检查密码规范性
	if !utilities.RegisterPasswordCheck(tem.Password){
		resp.Code = 30014
		resp.Message = "密码不规范"
		context.JSON(200,resp)
		return
	}

	//md5加盐加密
	password,MD5salt:=utilities.CryptographyNow(tem.Password)

	if model.PostRegisterInformation(tem.Name,tem.Telephone,tem.Email,password,MD5salt) {
		resp.Code = 300
		resp.Message = "注册成功"
		context.JSON(200,resp)
	}else{
		resp.Code = 30015
		resp.Message = "未知错误"
		context.JSON(200,resp)
	}
}