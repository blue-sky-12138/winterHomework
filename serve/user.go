package serve

import (
	"WinterHomework/model"
	"WinterHomework/utilities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//用户登录
func Login(context *gin.Context){
	var (
		login utilities.LoginCheck
		resp utilities.Resp
	)
	err := context.ShouldBind(&login)
	if err != nil{
		utilities.LogError("GetLoginInformation Error",err)
		resp.Code = 20003
		resp.Message = "未知错误"
		context.JSON(200,resp)
		return
	}

	check, err := model.GetUserPassword(login.Name)
	if err != nil {
		resp.Code = 20003
		resp.Message = err.Error()
		context.JSON(200,resp)
		return
	}

	if check.Md5salt == 0{		//检查用户名是否存在
		resp.Code =	 20001
		resp.Message = "该用户不存在"
		context.JSON(200,resp)
	}else if utilities.Cryptography(login.Password,check.Md5salt) != check.Password {	//检查密码是否正确
		resp.Code =	 20002
		resp.Message = "密码错误"
		context.JSON(200,resp)
	}else{
		cookie := &http.Cookie{
			Name:     	"user",
			Value:  	strconv.FormatInt(check.Id,10),
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
		tem utilities.RegisterInformation		//用于接收传输的数据
		resp utilities.Resp						//响应体
		ok bool									//用于临时存储函数的bool值
	)
	err := context.ShouldBind(&tem)
	if err != nil{
		utilities.LogError("GetRegisterInformation Error",err)
		resp.Code = 30003
		resp.Message = err.Error()
		context.JSON(200,resp)
		return
	}

	//检查邮箱的正确性、是否存在
	if tem.Email != ""{
		if !utilities.EmailCheck(tem.Email){
			resp.Code = 30002
			resp.Message = "邮箱错误"
			context.JSON(200,resp)
			return
		}else {
			ok, err = model.CheckRegisterOrNot("email",tem.Email)
			if err != nil {
				resp.Code = 30003
				resp.Message = err.Error()
				context.JSON(200,resp)
				return
			}else if ok{				//如果查找到数据
				resp.Code = 30002
				resp.Message = "该邮箱已被注册"
				context.JSON(200,resp)
				return
			}
		}
	}

	//检查手机号的正确性、是否存在
	if tem.Telephone != 0 {
		if !utilities.TelephoneCheck(tem.Telephone){
			resp.Code = 30002
			resp.Message = "手机号错误"
			context.JSON(200,resp)
			return
		}else {
			ok, err = model.CheckRegisterOrNot("telephone",strconv.FormatInt(tem.Telephone, 10))
			if err != nil {
				resp.Code = 30003
				resp.Message = err.Error()
				context.JSON(200,resp)
				return
			}else if ok{				//如果查找到数据
				resp.Code = 30002
				resp.Message = "该手机号已被注册"
				context.JSON(200,resp)
				return
			}
		}
	}

	//检查用户名的正确性、是否存在
	if tem.Name!= ""{
		if !utilities.UserNameCheck(tem.Name){
			resp.Code = 30002
			resp.Message = "用户名不符合要求"
			context.JSON(200,resp)
			return
		}else {
			ok, err = model.CheckRegisterOrNot("name", tem.Name)
			if err != nil {
				resp.Code = 30003
				resp.Message = err.Error()
				context.JSON(200,resp)
				return
			}else if ok{				//如果查找到数据
				resp.Code = 30002
				resp.Message = "该用户名已被注册"
				context.JSON(200,resp)
				return
			}
		}
	}

	//检查密码规范性
	if !utilities.PasswordCheck(tem.Password){
		resp.Code = 30002
		resp.Message = "密码不规范"
		context.JSON(200,resp)
		return
	}

	//md5加盐加密
	password,MD5salt := utilities.CryptographyNow(tem.Password)

	if model.PostRegisterInformation(tem.Name,tem.Telephone,tem.Email,password,MD5salt) == nil {
		resp.Code = 300
		resp.Message = "注册成功"
		context.JSON(200,resp)
	}else {
		resp.Code = 30003
		resp.Message = "未知错误"
		context.JSON(200,resp)
	}
}

func Update(context *gin.Context) {
	var (
		resp utilities.Resp
		err error
	)

	userId, _ := strconv.ParseInt(context.Query("user_id"),10,64)	//用户id
	operateType, _ := strconv.Atoi(context.Query("type"))						//操作类型
	content := context.Query("content")										//更改成什么内容

	//简单排除类型错误
	if operateType > 6 || operateType < 1 {
		resp.Code = 12001
		resp.Message = "类型不合法"
		context.JSON(200,resp)
		return
	}

	if operateType == 1 {			//更新手机号
		newTele, _ := strconv.ParseInt(content,10,64)
		if utilities.TelephoneCheck(newTele) {
			err = model.ChangeUserTelephone(userId,newTele)
		}
	}else if operateType == 2 {		//更新邮箱
		if utilities.EmailCheck(content) {
			err = model.ChangeUserEmail(userId,content)
		}
	}else if operateType == 3 {		//更新昵称
		if utilities.UserNameCheck(content) {
			err = model.ChangeUserNickname(userId,content)
		}
	}else if operateType == 4 {		//更新签名
		err = model.ChangeUserSignature(userId,content)
	}

	//操作错误集中检测
	if err != nil {
		resp.Code = 12002
		resp.Message = "未知错误"
		context.JSON(200,resp)
		return
	}

	resp.Code = 1200
	resp.Message = "更新成功"
	context.JSON(200,resp)
}